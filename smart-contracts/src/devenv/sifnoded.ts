import * as ChildProcess from "child_process"
import { ShellCommand } from "./devEnv"
import { GolangResults } from "./golangBuilder"
import * as path from "path"
import * as fs from "fs"
import YAML from "yaml"
import notifier from "node-notifier"
import { EbrelayerArguments } from "./ebrelayer"
import * as delay from "delay"

import {
  ExecFileSyncOptions,
  ExecFileSyncOptionsWithStringEncoding,
  ExecSyncOptionsWithStringEncoding,
  StdioOptions,
} from "child_process"
import { network } from "hardhat"
import { sleep } from "./devEnvUtilities"
import { log } from "console"

export const crossChainFeeBase: number = 1
export const crossChainLockFee: number = 1
export const crossChainBurnFee: number = 1
const ethereumCrossChainFeeToken: string =
  "sif5ebfaf95495ceb5a3efbd0b0c63150676ec71e023b1043c40bcaaf91c00e15b2"

export interface ValidatorValues {
  chain_id: string
  node_id: string
  ipv4_address: string
  moniker: string
  password: string
  address: string
  pub_key: string
  mnemonic: string
  validator_address: string
  validator_consensus_address: string
  is_seed: boolean
}
export interface EbRelayerAccount {
  name: string
  account: string
  homeDir: string
}
export interface SifnodedResults {
  validatorValues: ValidatorValues[]
  relayerAddresses: EbRelayerAccount[]
  witnessAddresses: EbRelayerAccount[]
  adminAddress: EbRelayerAccount
  process: ChildProcess.ChildProcess
  tcpurl: string
}

export async function waitForSifAccount(address: string, sifnoded: string) {
  for (;;) {
    try {
      console.log("Attempting to check account")
      ChildProcess.execSync(`${sifnoded} query account ${address}`, {
        encoding: "utf8",
      }).trim()
      console.log("Sifnoded is now running, continunig onwards")
      return
    } catch {
      await sleep(1000)
    }
  }
}

export class SifnodedRunner extends ShellCommand<SifnodedResults> {
  output: Promise<SifnodedResults>
  private outputResolve: any
  private sifnodedCommand: string
  private sifgenCommand: string

  constructor(
    readonly golangResults: GolangResults,
    readonly logfile = "/tmp/sifnode/sifnoded.log",
    readonly rpcPort = 9000,
    readonly nValidators = 1,
    readonly nRelayers = 1,
    readonly nWitnesses = 1,
    readonly chainId = "localnet",
    readonly networkConfigFile = "/tmp/sifnodedConfig.yml",
    readonly networkDir = "/tmp/sifnodedNetwork",
    readonly seedIpAddress = "10.10.1.1",
    readonly whitelistFile = "../test/integration/whitelisted-denoms.json"
  ) {
    super()
    this.sifnodedCommand = path.join(this.golangResults.goBin, "sifnoded")
    this.output = new Promise<SifnodedResults>((res, _) => {
      this.outputResolve = res
    })
    this.sifgenCommand = path.join(this.golangResults.goBin, "sifgen")
    this.sifnodedCommand = path.join(this.golangResults.goBin, "sifnoded")
  }

  cmd(): [string, string[]] {
    return ["sifgen", ["node"]]
  }

  async sifgenNetworkCreate(): Promise<SifnodedResults> {
    const sifgenArgs = [
      "network",
      "create",
      this.chainId,
      this.nValidators.toString(),
      this.networkDir,
      this.seedIpAddress,
      this.networkConfigFile,
      "--keyring-backend",
      "test",
      // Mint goes to validator
      "--mint-amount",
      "999999000000000000000000000rowan,1370000000000000000ibc/FEEDFACEFEEDFACEFEEDFACEFEEDFACEFEEDFACEFEEDFACEFEEDFACEFEEDFACE,999999000000000000000000000sif5ebfaf95495ceb5a3efbd0b0c63150676ec71e023b1043c40bcaaf91c00e15b2",
    ]

    await fs.promises.mkdir(this.networkDir, { recursive: true })

    const sifnodedLogFile = fs.openSync(this.logfile, "w")

    let stdioOptions: StdioOptions = ["ignore", sifnodedLogFile, sifnodedLogFile]

    const sifgenOutput = ChildProcess.execFileSync(this.sifgenCommand, sifgenArgs, {
      encoding: "utf8",
    })

    // Debug log
    // TODO: Add formal loglevel aware logging
    console.log("SifgenOutput", sifgenOutput)

    const file = fs.readFileSync(this.networkConfigFile, "utf8")
    const networkConfig: ValidatorValues[] = YAML.parse(file)
    let homeDir: string = ""

    // TODO: Extract this into function
    for (const validator of networkConfig) {
      const moniker = validator["moniker"]
      const mnemonic = validator["mnemonic"]
      const password = validator["password"]
      let chainDir: string = path.join(this.networkDir, "validators", this.chainId, moniker)

      homeDir = path.join(chainDir, ".sifnoded")
      await this.addValidatorKeyToTestKeyring(moniker, mnemonic)

      const valOperKey = this.readValoperKey(moniker, homeDir)
      await this.addGenesisValidator(chainDir, valOperKey)
      console.log(
        `Added genesis validator: ${JSON.stringify({ moniker, homeDir, chainDir, valOperKey })}`
      )

      // TODO This is not needed, plus it doesn't need a password
      const whitelistedValidator = ChildProcess.execSync(
        `${this.sifnodedCommand} keys show -a --bech val ${moniker} --keyring-backend test`,
        { encoding: "utf8", input: password }
      ).trim()
      console.log(`--bech val output: ${whitelistedValidator}`)
    }

    // Create an ADMIN account on sifnode with name sifnodeadmin
    const sifnodedAdminAddress: EbRelayerAccount = this.addAccount("sifnodeadmin", homeDir, true)
    // Create an account for each relayer as requested
    const relayerAddresses = Array.from({ length: this.nRelayers }, (_, relayer) =>
      this.addRelayerWitnessAccount(`relayer-${relayer}`, homeDir)
    )
    // Create an account for each witness as requested
    const witnessAddresses = Array.from({ length: this.nWitnesses }, (_, witness) =>
      this.addRelayerWitnessAccount(`witness-${witness}`, homeDir)
    )

    let sifnodedDaemonCmd = `${this.sifnodedCommand} start --log_level debug --log_format json --minimum-gas-prices 0.5rowan --rpc.laddr tcp://0.0.0.0:26657 --home ${homeDir}`

    console.log(`start sifnoded with: \n${sifnodedDaemonCmd}`)
    const sifnoded = ChildProcess.spawn(sifnodedDaemonCmd, { shell: true, stdio: stdioOptions })

    // Register tokens in the token registry
    // Must wait for sifnode to fully start first
    await waitForSifAccount(networkConfig[0].address, this.sifnodedCommand)
    const registryPath = path.resolve(__dirname, "./", "registry.json")


    // TODO This command looks like something is failing. The exit code is 0, but it prints this (x2):
    // {"height":"2","txhash":"E84F6BB8D3C74F8CFF9867C0322E6E1C4E3485BAD839CC145EF2F881D651ADD6","codespace":"sdk","code":7,"data":"","raw_log":"failed to execute message; message index: 0: unauthorised signer: invalid address","logs":[],"info":"","gas_wanted":"200000","gas_used":"57304","tx":null,"timestamp":""}
    const tmp = ChildProcess.execSync(
      `${this.sifnodedCommand} tx tokenregistry register-all ${registryPath} --home ${homeDir} --gas-prices 0.5rowan --gas-adjustment 1.5 --from ${sifnodedAdminAddress.name} --yes --keyring-backend test --chain-id ${this.chainId}`,
      { encoding: "utf8" }
    ).trim()
    console.log("Result from token registry: " + tmp)

    await this.setCrossChainFee(
      sifnodedAdminAddress,
      "31337",
      ethereumCrossChainFeeToken,
      String(crossChainFeeBase),
      String(crossChainLockFee),
      String(crossChainBurnFee),
      this.chainId
    )

    sifnoded.on("exit", (code) => {
      notifier.notify({
        title: "Sifnoded Notice",
        message: `Sifnoded has just exited with exit code: ${code}`,
      })
    })

    return {
      validatorValues: networkConfig,
      adminAddress: sifnodedAdminAddress,
      relayerAddresses: relayerAddresses,
      witnessAddresses: witnessAddresses,
      process: sifnoded,
      tcpurl: "tcp://0.0.0.0:26657",
    }
    //    return lastValueFrom(eventEmitterToObservable(sifnoded, "sifnoded"))
  }

  addAccount(name: string, homeDir: string, isAdmin: boolean): EbRelayerAccount {
    let accountAddCmd = `${this.sifnodedCommand} keys add ${name} --keyring-backend test --output json --home ${homeDir}`
    const accountJSON = ChildProcess.execSync(accountAddCmd, {
      encoding: "utf8",
      input: "yes\nyes",
    }).trim()
    const accountAddress = JSON.parse(accountJSON)["address"]

    // TODO: Homedir would contain value of last assignment. Might need to be fixed when we support more than 1 acc
    ChildProcess.execSync(
      `${this.sifnodedCommand} add-genesis-account ${accountAddress} 100000000000000000000rowan,20000000000000000000ceth --home ${homeDir}`,
      { encoding: "utf8" }
    ).trim()
    if (isAdmin === true) {
      ChildProcess.execSync(
        `${this.sifnodedCommand} set-genesis-oracle-admin ${accountAddress} --home ${homeDir}`,
        { encoding: "utf8" }
      ).trim()
    }

    ChildProcess.execSync(
      `${this.sifnodedCommand} set-genesis-whitelister-admin ${accountAddress} --home ${homeDir}`,
      { encoding: "utf8" }
    ).trim()

    return {
      account: accountAddress,
      name: name,
      homeDir,
    }
  }

  addRelayerWitnessAccount(name: string, homeDir: string): EbRelayerAccount {
    const adminAccount = this.addAccount(name, homeDir, false)
    // Whitelist Relayer/Witness Account
    const EVM_Network_Descriptor = 31337
    const Validator_Power = 100
    const bachAddress = this.readValoperKey(name, homeDir)
    ChildProcess.execSync(
      `${this.sifnodedCommand} set-gen-denom-whitelist ${this.whitelistFile} --home ${homeDir}`,
      { encoding: "utf8" }
    ).trim()
    ChildProcess.execSync(
      `${this.sifnodedCommand} add-genesis-validators ${EVM_Network_Descriptor} ${bachAddress} ${Validator_Power} --home ${homeDir}`,
      { encoding: "utf8" }
    ).trim()

    return adminAccount
  }

  async addValidatorKeyToTestKeyring(moniker: string, mnemonic: string) {
    const sifnodedArgs = ["keys", "add", moniker, "--keyring-backend", "test", "--recover"]

    console.log("Add Validator with mnemonics: ", mnemonic)

    let child = ChildProcess.execFileSync(this.sifnodedCommand, sifnodedArgs, {
      encoding: "utf8",
      shell: false,
      input: `${mnemonic}\n`,
    })
    console.log("Add Validator key to test ring output:", child)
  }

  // TODO: args Position
  readValoperKey(moniker: string, homeDir: string): string {
    return ChildProcess.execSync(
      `${this.sifnodedCommand} keys show -a --bech val ${moniker} --keyring-backend test --home ${homeDir}`,
      { encoding: "utf8" }
    ).trim()
  }

  // sifnoded add-genesis-validators $valoper --home $CHAINDIR/.sifnoded
  async addGenesisValidator(chainDir: string, valoper: string): Promise<string> {
    const sifgenArgs = [
      "add-genesis-validators",
      "1",
      valoper,
      "100",
      "--home",
      path.join(chainDir, ".sifnoded"),
    ]

    console.log(`Add genesis validator: ${JSON.stringify({ chainDir, valoper })}`)
    return ChildProcess.execFileSync(this.sifnodedCommand, sifgenArgs, { encoding: "utf8" })
  }

  // sifnoded tx ethbridge set-cross-chain-fee sif1f8sz5779td3y6xsq296k3wurflsdnfxmq5hudd 1 ceth 1 1 1
  // set-cross-chain-fee [cosmos-sender-address] [network-id] [cross-chain-fee] [fee-currency-gas] [minimum-lock-cost] [minimum-burn-cost]
  async setCrossChainFee(
    sifnodeAdminAccount: EbRelayerAccount,
    networkId: string,
    crossChainFee: string,
    feeCurrencyGas: string,
    minLockCost: string,
    minBurnCost: string,
    chainId: string
  ): Promise<string> {
    const sifgenArgs = [
      "tx",
      "ethbridge",
      "set-cross-chain-fee",
      sifnodeAdminAccount.account,
      networkId, // This is 31377 for HARDHAT
      crossChainFee,
      feeCurrencyGas,
      minLockCost,
      minBurnCost,
      "--home",
      sifnodeAdminAccount.homeDir,
      "--from",
      sifnodeAdminAccount.name,
      "--keyring-backend",
      "test",
      "--chain-id",
      chainId,
      "--gas-prices",
      "0.5rowan",
      "--gas-adjustment",
      "1.5",
      "-y",
    ]

    console.log(sifgenArgs) // TODO I'm getting this: 
    // sifnoded tx ethbridge set-cross-chain-fee sifnodeadmin 31337 sif5ebfaf95495ceb5a3efbd0b0c63150676ec71e023b1043c40bcaaf91c00e15b2 1 1 1 --from sifnodeadmin 
    //     --chain-id localnet --gas-prices 0.5rowan --gas-adjustment 1.5 -y --home /tmp/sifnodedNetwork/validators/localnet/quiet-sea/.sifnoded --keyring-backend test
    // Error: decoding bech32 failed: invalid index of 1

    return ChildProcess.execFileSync(this.sifnodedCommand, sifgenArgs, { encoding: "utf8" })
  }

  override async run(): Promise<void> {
    const output = await this.sifgenNetworkCreate()
    this.outputResolve(output)
  }

  override async results(): Promise<SifnodedResults> {
    return this.output
  }
}
