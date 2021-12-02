import json
import time
from command import Command
from common import *


class Sifnoded(Command):
    def __init__(self):
        self.binary = "sifnoded"
        # home = None
        # keyring_backend = None

    def sifnoded_init(self, moniker, chain_id):
        args = [self.binary, "init", moniker, "--chain-id", chain_id]
        res = self.execst(args)
        return json.loads(res[2])  # output is on stderr

    def sifnoded_keys_show(self, name, bech=None, keyring_backend=None, sifnoded_home=None):
        keyring_backend = keyring_backend or "test"
        args = ["keys", "show", name] + \
            (["--bech", bech] if bech else [])
        res = self.sifnoded_exec(args, keyring_backend=keyring_backend, sifnoded_home=sifnoded_home)
        return yaml_load(stdout(res))

    def sifnoded_get_val_address(self, moniker, sifnoded_home=None):
        expected = exactly_one(stdout_lines(self.sifnoded_exec(["keys", "show", "-a", "--bech", "val", moniker], keyring_backend="test", sifnoded_home=sifnoded_home)))
        result = exactly_one(self.sifnoded_keys_show(moniker, bech="val", keyring_backend="test", sifnoded_home=sifnoded_home))["address"]
        assert result == expected
        return result

    # Why we need several different functions here?
    # How "sifnoded keys add <name> --keyring-backend test" works:
    # If name does not exist yet, it creates it and returns a yaml
    # If name alredy exists, prompts for overwrite (y/n) on standard input, generates new address/pubkey/mnemonic
    # Directory used is xxx/keyring-test if "--home xxx" is specified, otherwise $HOME/.sifnoded/keyring-test

    def sifnoded_keys_add(self, moniker, mnemonic, sifnoded_home=None):
        stdin = [" ".join(mnemonic)]
        res = self.sifnoded_exec(["keys", "add", moniker, "--recover"], keyring_backend="test", sifnoded_home=sifnoded_home, stdin=stdin)
        account = exactly_one(yaml_load(stdout(res)))
        return account

    # Creates a new key in the keyring and returns its address ("sif1xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx").
    # Since this is a test keyring, we don't need to save the generated private key.
    # If we wanted to recreate it, we can capture the mnemonic from the message that is printed to stderr.
    def sifnoded_keys_add_1(self, moniker, sifnoded_home=None):
        res = self.sifnoded_exec(["keys", "add", moniker], keyring_backend="test", sifnoded_home=sifnoded_home, stdin=["y"])
        account = exactly_one(yaml_load(stdout(res)))
        unused_mnemonic = stderr(res).splitlines()[-1].split(" ")
        return account

    # # This is without "--recover", it will generate a new mnemonic.
    # # Using this is probably a bug.
    # # From peggy
    # # @TODO Passing mnemonic to stdin is useless, only "y/n" makes sense, probably could use sifnoded_keys_add_1
    # # See smart-contracts/src/devenv/sifnoded.ts:addValidatorKeysToTestKeyring
    # def sifnoded_keys_add_2(self, moniker, mnemonic):
    #     stdin = [" ".join(mnemonic)]
    #     res = self.sifnoded_exec(["keys", "add", moniker], keyring_backend="test", stdin=stdin)
    #     result = exactly_one(yaml_load(stdout(res)))
    #     # {"name": "<moniker>", "type": "local", "address": "sif1...", "pubkey": "sifpub1...", "mnemonic": "", "threshold": 0, "pubkeys": []}
    #     return result

    def sifnoded_keys_delete(self, name):
        self.execst(["sifnoded", "keys", "delete", name, "--keyring-backend", "test"], stdin=["y"], check_exit=False)

    def sifnoded_add_genesis_account(self, sifnodeadmin_addr, tokens, sifnoded_home=None):
        tokens_str = ",".join([sif_format_amount(amount, denom) for amount, denom in tokens])
        self.sifnoded_exec(["add-genesis-account", sifnodeadmin_addr, tokens_str], sifnoded_home=sifnoded_home)

    def sifnoded_add_genesis_validators(self, address):
        args = ["sifnoded", "add-genesis-validators", address]
        res = self.execst(args)
        return res

    # At the moment only on future/peggy2 branch, called from PeggyEnvironment
    def sifnoded_add_genesis_validators_peggy(self, evm_network_descriptor, valoper, validator_power, sifnoded_home=None):
        self.sifnoded_exec(["add-genesis-validators", str(evm_network_descriptor), valoper, str(validator_power)],
            sifnoded_home=sifnoded_home)

    def sifnoded_set_genesis_oracle_admin(self, address, sifnoded_home=None):
        self.sifnoded_exec(["set-genesis-oracle-admin", address], sifnoded_home=sifnoded_home)

    def sifnoded_set_genesis_whitelister_admin(self, address, sifnoded_home=None):
        self.sifnoded_exec(["set-genesis-whitelister-admin", address], sifnoded_home=sifnoded_home)

    def sifnoded_set_gen_denom_whitelist(self, denom_whitelist_file, sifnoded_home=None):
        self.sifnoded_exec(["set-gen-denom-whitelist", denom_whitelist_file], sifnoded_home=sifnoded_home)

    # At the moment only on future/peggy2 branch, called from PeggyEnvironment
    # This was split from init_common
    def sifnoded_peggy2_add_account(self, name, tokens, is_admin=False, sifnoded_home=None):
        # TODO Peggy2 devenv feed "yes\nyes" into standard input, we only have "y\n"
        account = self.sifnoded_keys_add_1(name, sifnoded_home=sifnoded_home)
        account_address = account["address"]

        self.sifnoded_add_genesis_account(account_address, tokens, sifnoded_home=sifnoded_home)
        if is_admin:
            self.sifnoded_set_genesis_oracle_admin(account_address, sifnoded_home=sifnoded_home)
        self.sifnoded_set_genesis_whitelister_admin(account_address, sifnoded_home)
        return account_address

    def sifnoded_peggy2_add_relayer_witness_account(self, name, tokens, evm_network_descriptor, validator_power, denom_whitelist_file, sifnoded_home=None):
        account_address = self.sifnoded_peggy2_add_account(name, tokens, sifnoded_home=sifnoded_home)  # Note: is_admin=False
        # Whitelist relayer/witness account
        valoper = self.sifnoded_get_val_address(name, sifnoded_home=sifnoded_home)
        self.sifnoded_set_gen_denom_whitelist(denom_whitelist_file, sifnoded_home=sifnoded_home)
        self.sifnoded_add_genesis_validators_peggy(evm_network_descriptor, valoper, validator_power, sifnoded_home=sifnoded_home)
        return account_address

    def sifnoded_tx_clp_create_pool(self, chain_id, keyring_backend, from_name, symbol, fees, native_amount, external_amount):
        args = [self.binary, "tx", "clp", "create-pool", "--chain-id={}".format(chain_id),
            "--keyring-backend={}".format(keyring_backend), "--from", from_name, "--symbol", symbol, "--fees",
            sif_format_amount(*fees), "--nativeAmount", str(native_amount), "--externalAmount", str(external_amount),
            "--yes"]
        res = self.execst(args)
        return yaml_load(stdout(res))

    def sifnoded_peggy2_token_registry_register_all(self, registry_path, gas_prices, gas_adjustment, from_account,
        chain_id, sifnoded_home=None):
        args = ["tx", "tokenregistry", "register-all", registry_path, "--gas-prices", sif_format_amount(*gas_prices),
            "--gas-adjustment", str(gas_adjustment), "--from", from_account, "--chain-id", chain_id, "--yes"]
        res = self.sifnoded_exec(args, keyring_backend="test", sifnoded_home=sifnoded_home)
        return [json.loads(x) for x in stdout(res).splitlines()]

    def sifnoded_peggy2_set_cross_chain_fee(self, admin_account_address, network_id, ethereum_cross_chain_fee_token,
        cross_chain_fee_base, cross_chain_lock_fee, cross_chain_burn_fee, admin_account_name, chain_id, gas_prices,
        gas_adjustment, sifnoded_home=None
    ):
        # Checked OK
        args = ["tx", "ethbridge", "set-cross-chain-fee", admin_account_address, str(network_id),
            ethereum_cross_chain_fee_token, str(cross_chain_fee_base), str(cross_chain_lock_fee),
            str(cross_chain_burn_fee), "--from", admin_account_name, "--chain-id", chain_id, "--gas-prices",
            sif_format_amount(*gas_prices), "--gas-adjustment", str(gas_adjustment), "-y"]
        res = self.sifnoded_exec(args, keyring_backend="test", sifnoded_home=sifnoded_home)
        return res

    def sifnoded_start(self, tcp_url=None, minimum_gas_prices=None, sifnoded_home=None, log_file=None,
        log_format_json=False
    ):
        args = [self.binary, "start"] + \
            (["--minimum-gas-prices", sif_format_amount(*minimum_gas_prices)] if minimum_gas_prices is not None else []) + \
            (["--rpc.laddr", tcp_url] if tcp_url else []) + \
            (["--log_level", "debug"] if log_format_json else []) + \
            (["--log_format", "json"] if log_format_json else []) + \
            (["--home", sifnoded_home] if sifnoded_home else [])
        return self.popen(args, log_file=log_file)

    def sifnoded_exec(self, args, sifnoded_home=None, keyring_backend=None, stdin=None, cwd=None):
        args = [self.binary] + args + \
            (["--home", sifnoded_home] if sifnoded_home else []) + \
            (["--keyring-backend", keyring_backend] if keyring_backend else [])
        res = self.execst(args, stdin=stdin, cwd=cwd)
        return res

    def sifnoded_get_status(self, host, port):
        url = "http://{}:{}/node_info".format(host, port)
        return json.loads(http_get(url).decode("UTF-8"))

    def tcp_probe_connect(self, host, port):
        res = self.execst(["nc", "-z", host, str(port)], check_exit=False)
        return res[0] == 0

    def wait_for_file(self, path):
        while not self.exists(path):
            time.sleep(1)

    def wait_for_sif_account_up(self, address, tcp_url=None):
        # TODO Deduplicate: this is also in run_ebrelayer()
        # netdef_json is path to file containing json_dump(netdef)
        # while not self.tcp_probe_connect("localhost", tendermint_port):
        #     time.sleep(1)
        # self.wait_for_sif_account(netdef_json, validator1_address)

        # Peggy2
        # How this works: by default, the command below will try to do a POST to http://localhost:26657.
        # So the port has to be up first, but this query will fail anyway if it is not.
        args = ["sifnoded", "query", "account", address] + \
            (["--node", tcp_url] if tcp_url else [])
        while True:
            try:
                self.execst(args)
                break
            except Exception as e:
                log.debug(f"Waiting for sif account {address}... ({repr(e)})")
                time.sleep(1)

class Sifgen:
    def __init__(self, cmd):
        self.cmd = cmd
        self.binary = "sifgen"

    # Reference: docker/localnet/sifnode/root/scripts/sifnode.sh (branch future/peggy2):
    # sifgen node create "$CHAINNET" "$MONIKER" "$MNEMONIC" --bind-ip-address "$BIND_IP_ADDRESS" --standalone --keyring-backend test
    def create_standalone(self, chainnet, moniker, mnemonic, bind_ip_address, keyring_backend=None):
        args = ["node", "create", chainnet, moniker, mnemonic, bind_ip_address]
        return self.sifgen_exec(args, keyring_backend=keyring_backend)

    def sifgen_exec(self, args, keyring_backend=None, cwd=None, env=None):
        args = [self.binary] + args + \
            (["--keyring-backend", keyring_backend] if keyring_backend else [])
        return self.cmd.execst(args, cwd=cwd, env=env)


class Ebrelayer:
    def __init__(self, cmd):
        self.cmd = cmd
        self.binary = "ebrelayer"

    def __deleteme__peggy2_init_relayer(
        self,
        network_descriptor,
        tendermint_node,
        web3_provider,
        bridge_registry_contract_address,
        validator_moniker,
        validator_mnemonic,
        chain_id,
        symbol_translator_file,
        ethereum_address,
        ethereum_private_key,
        keyring_backend=None,
        log_file=None,
        cwd=None,
    ):
        # Usage:
        #   ebrelayer init-relayer [networkDescriptor] [tendermintNode] [web3Provider] [bridgeRegistryContractAddress] [validatorMnemonic] [flags]
        #
        # Examples:
        # ebrelayer init-relayer 1 tcp://localhost:26657 ws://localhost:7545/ 0x30753E4A8aad7F8597332E813735Def5dD395028 mnemonic --chain-id=peggy
        return self.peggy2_run_ebrelayer("init-relayer", network_descriptor, tendermint_node, web3_provider,
            bridge_registry_contract_address, validator_mnemonic, chain_id=chain_id, node=tendermint_node,
            sign_with=validator_moniker, symbol_translator_file=symbol_translator_file,
            ethereum_address=ethereum_address, ethereum_private_key=ethereum_private_key,
            keyring_backend=keyring_backend, cwd=cwd, log_file=log_file)

    def __deleteme__peggy2_init_witness(
        self,
        network_descriptor,
        tendermint_node,
        web3_provider,
        bridge_registry_contract_address,
        validator_moniker,
        validator_mnemonic,
        chain_id,
        symbol_translator_file,
        ethereum_address,
        ethereum_private_key,
        relayerdb_path=None,
        keyring_backend=None,
        log_file=None,
        cwd=None,
    ):
        # Usage:
        #   ebrelayer init-witness [networkDescriptor] [tendermintNode] [web3Provider] [bridgeRegistryContractAddress] [validatorMnemonic] [flags]
        #
        # Examples:
        # ebrelayer init-witness 1 tcp://localhost:26657 ws://localhost:7545/ 0x30753E4A8aad7F8597332E813735Def5dD395028 mnemonic --chain-id=peggy
        extra_args = [] + \
            (["--relayerdb-path", relayerdb_path] if relayerdb_path else [])
        return self.peggy2_run_ebrelayer("init-witness", network_descriptor, tendermint_node, web3_provider,
            bridge_registry_contract_address, validator_mnemonic, chain_id=chain_id, node=tendermint_node,
            sign_with=validator_moniker, symbol_translator_file=symbol_translator_file,
            ethereum_address=ethereum_address, ethereum_private_key=ethereum_private_key,
            keyring_backend=keyring_backend, cwd=cwd, log_file=log_file,
            log_format="json", extra_args=extra_args)

    def peggy2_run_ebrelayer(self, init_what, network_descriptor, tendermint_node, web3_provider,
        bridge_registry_contract_address, validator_mnemonic, chain_id, node=None, keyring_backend=None,
        sign_with=None, symbol_translator_file=None, relayerdb_path=None, log_format=None, extra_args=None,
        ethereum_private_key=None, ethereum_address=None, home=None, log_file=None, cwd=None
    ):
        env = {}
        if ethereum_private_key:
            assert not ethereum_private_key.startswith("0x")
            env["ETHEREUM_PRIVATE_KEY"] = ethereum_private_key
        if ethereum_address:
            assert ethereum_address.startswith("0x")
            env["ETHEREUM_ADDRESS"] = ethereum_address
        env = env or None  # Avoid passing empty environment
        args = [
            self.binary,
            init_what,
            "--network-descriptor", str(network_descriptor),  # Network descriptor for the chain (31337)
            "--tendermint-node", tendermint_node,  # URL to tendermint node
            "--web3-provider", web3_provider,  # Ethereum web3 service address (ws://localhost:8545/)
            "--bridge-registry-contract-address", bridge_registry_contract_address,
            "--validator-mnemonic", validator_mnemonic,
            "--chain-id", chain_id  # chain ID of tendermint node (localnet)
        ] + \
            (extra_args if extra_args else []) + \
            (["--node", node] if node else []) + \
            (["--keyring-backend", keyring_backend] if keyring_backend else []) + \
            (["--from", sign_with] if sign_with else []) + \
            (["--relayerdb-path", relayerdb_path] if relayerdb_path else []) + \
            (["--home", home] if home else []) + \
            (["--symbol-translator-file", symbol_translator_file] if symbol_translator_file else []) + \
            (["--log_format", log_format] if log_format else [])
        return self.cmd.popen(args, env=env, cwd=cwd, log_file=log_file)

    # Legacy stuff - pre-peggy2
    # Called from IntegrationContext
    def init(self, tendermind_node, web3_provider, bridge_registry_contract_address, validator_moniker,
        validator_mnemonic, chain_id, ethereum_private_key=None, ethereum_address=None, gas=None, gas_prices=None,
        node=None, keyring_backend=None, sign_with=None, symbol_translator_file=None, relayerdb_path=None,
        cwd=None, log_file=None
    ):
        env = {}
        if ethereum_private_key:
            assert not ethereum_private_key.startswith("0x")
            env["ETHEREUM_PRIVATE_KEY"] = ethereum_private_key
        if ethereum_address:
            assert ethereum_address.startswith("0x")
            env["ETHEREUM_ADDRESS"] = ethereum_address
        env = env or None  # Avoid passing empty environment
        args = [self.binary, "init", tendermind_node, web3_provider, bridge_registry_contract_address,
            validator_moniker, " ".join(validator_mnemonic), "--chain-id={}".format(chain_id)] + \
            (["--gas", str(gas)] if gas is not None else []) + \
            (["--gas-prices", sif_format_amount(*gas_prices)] if gas_prices is not None else []) + \
            (["--node", node] if node is not None else []) + \
            (["--keyring-backend", keyring_backend] if keyring_backend is not None else []) + \
            (["--from", sign_with] if sign_with is not None else []) + \
            (["--symbol-translator-file", symbol_translator_file] if symbol_translator_file else []) + \
            (["--relayerdb-path", relayerdb_path] if relayerdb_path else [])
        return self.cmd.popen(args, env=env, cwd=cwd, log_file=log_file)
