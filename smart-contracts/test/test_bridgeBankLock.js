const Web3Utils = require("web3-utils");
const web3 = require("web3");
const BigNumber = web3.BigNumber;

const { ethers, upgrades } = require("hardhat");
const { use, expect } = require("chai");
const { solidity } = require("ethereum-waffle");
const { setup, batchAddTokensToEthWhitelist } = require("./helpers/testFixture");

require("chai")
  .use(require("chai-as-promised"))
  .use(require("chai-bignumber")(BigNumber))
  .should();

use(solidity);

const getBalance = async function (address) {
  return await network.provider.send("eth_getBalance", [address]);
}

describe("Test Bridge Bank", function () {
  let userOne;
  let userTwo;
  let userThree;
  let userFour;
  let accounts;
  let signerAccounts;
  let operator;
  let owner;
  const consensusThreshold = 75;
  let initialPowers;
  let initialValidators;
  let networkDescriptor;
  // track the state of the deployed contracts
  let state;

  before(async function () {
    accounts = await ethers.getSigners();

    signerAccounts = accounts.map((e) => { return e.address });

    operator = accounts[0];
    userOne = accounts[1];
    userTwo = accounts[2];
    userFour = accounts[3];
    userThree = accounts[7].address;

    owner = accounts[5];
    pauser = accounts[6];

    initialPowers = [25, 25, 25, 25];
    initialValidators = signerAccounts.slice(0, 4);

    networkDescriptor = 1;
  });

  beforeEach(async function () {
    state = await setup({
      initialValidators,
      initialPowers,
      operator,
      consensusThreshold,
      owner,
      user: userOne,
      recipient: userThree,
      pauser,
      networkDescriptor
    });
  });

  describe("BridgeBank", function () {
    it("should deploy the BridgeBank, correctly setting the operator", async function () {
      state.bridgeBank.should.exist;

      const bridgeBankOperator = await state.bridgeBank.operator();
      bridgeBankOperator.should.be.equal(operator.address);
    });

    it("should allow user to lock ERC20 tokens", async function () {
      // Add the token into white list
      await state.bridgeBank.connect(operator)
        .updateEthWhiteList(state.token1.address, true)
        .should.be.fulfilled;

      // Attempt to lock tokens
      await state.bridgeBank.connect(userOne).lock(
        state.sender,
        state.token1.address,
        state.amount
      );

      // Confirm that the user has been minted the correct token
      const afterUserBalance = Number(
        await state.token1.balanceOf(userOne.address)
      );
      afterUserBalance.should.be.bignumber.equal(state.amount);
    });

    it("should allow users to lock Ethereum in the bridge bank", async function () {
      const tx = await state.bridgeBank.connect(userOne).lock(
        state.sender,
        state.constants.zeroAddress,
        state.weiAmount, {
        value: state.weiAmount
      }
      ).should.be.fulfilled;
      await tx.wait();

      const contractBalanceWei = await getBalance(state.bridgeBank.address);
      const contractBalance = Web3Utils.fromWei(contractBalanceWei, "ether");

      contractBalance.should.be.bignumber.equal(
        Web3Utils.fromWei((+state.weiAmount).toString(), "ether")
      );
    });
  });

  describe("Multi Lock ERC20 Tokens", function () {
    it("should allow user to multi-lock ERC20 tokens", async function () {
      // Add the tokens into white list
      await batchAddTokensToEthWhitelist(state, [
        state.token1.address,
        state.token2.address,
        state.token3.address,
      ]);

      // Attempt to lock tokens
      await state.bridgeBank.connect(userOne).multiLock(
        [state.sender, state.sender, state.sender],
        [state.token1.address, state.token2.address, state.token3.address],
        [state.amount, state.amount, state.amount]
      );

      // Confirm that the user has been minted the correct token
      let afterUserBalance = Number(
        await state.token1.balanceOf(userOne.address)
      );
      afterUserBalance.should.be.bignumber.equal(state.amount);

      afterUserBalance = Number(
        await state.token2.balanceOf(userOne.address)
      );
      afterUserBalance.should.be.bignumber.equal(state.amount);

      afterUserBalance = Number(
        await state.token3.balanceOf(userOne.address)
      );
      afterUserBalance.should.be.bignumber.equal(state.amount);
    });

    it("should allow user to multi-lock ERC20 tokens with multiLockBurn method", async function () {
      // Add the tokens into white list
      await batchAddTokensToEthWhitelist(state, [
        state.token1.address,
        state.token2.address,
        state.token3.address,
      ]);
      
      // Attempt to lock tokens
      await state.bridgeBank.connect(userOne).multiLockBurn(
        [state.sender, state.sender, state.sender],
        [state.token1.address, state.token2.address, state.token3.address],
        [state.amount, state.amount, state.amount],
        [false, false, false]
      );

      // Confirm that the user has been minted the correct token
      let afterUserBalance = Number(
        await state.token1.balanceOf(userOne.address)
      );
      afterUserBalance.should.be.bignumber.equal(state.amount);

      afterUserBalance = Number(
        await state.token2.balanceOf(userOne.address)
      );
      afterUserBalance.should.be.bignumber.equal(state.amount);

      afterUserBalance = Number(
        await state.token3.balanceOf(userOne.address)
      );
      afterUserBalance.should.be.bignumber.equal(state.amount);
    });

    it("should not allow user to multi-burn ERC20 tokens that are not cosmos native assets", async function () {
      // Attempt to lock tokens
      await expect(
        state.bridgeBank.connect(userOne).multiLockBurn(
          [state.sender, state.sender, state.sender],
          [state.token1.address, state.token2.address, state.token3.address],
          [state.amount, state.amount, state.amount],
          [true, false, false]
        ),
      ).to.be.revertedWith("Token is not in Cosmos whitelist");
    });

    it("should allow user to multi-lock and burn ERC20 tokens and rowan with multiLockBurn method", async function () {
      // Add the tokens into white list
      await batchAddTokensToEthWhitelist(state, [
        state.token1.address,
        state.token2.address,
      ]);

      // approve bridgebank to spend rowan
      await state.rowan.connect(userOne).approve(
        state.bridgeBank.address,
        state.amount
      );

      // Lock & burn tokens
      const tx = await state.bridgeBank.connect(userOne).multiLockBurn(
        [state.sender, state.sender, state.sender],
        [state.token1.address, state.token2.address, state.rowan.address],
        [state.amount, state.amount, state.amount],
        [false, false, true]
      );

      await tx.wait();

      // Confirm that the user has the proper balance after the multiLockBurn
      let afterUserBalance = Number(
        await state.token1.balanceOf(userOne.address)
      );
      afterUserBalance.should.be.bignumber.equal(state.amount);

      afterUserBalance = Number(
        await state.token2.balanceOf(userOne.address)
      );
      afterUserBalance.should.be.bignumber.equal(state.amount);

      afterUserBalance = Number(
        await state.rowan.balanceOf(userOne.address)
      );
      afterUserBalance.should.be.bignumber.equal(state.amount);
    });

    it("should not allow user to multi-lock ERC20 tokens if one token is not fully approved", async function () {
      // Add the tokens into white list
      await batchAddTokensToEthWhitelist(state, [
        state.token1.address,
        state.token2.address,
        state.token3.address,
      ]);
      
      const tx = await state.token1.connect(userOne).approve(
        state.bridgeBank.address,
        0
      );
      const receipt = await tx.wait();

      // Attempt to lock tokens
      await expect(
        state.bridgeBank.connect(userOne).multiLock(
          [state.sender, state.sender, state.sender],
          [state.token1.address, state.token2.address, state.token3.address],
          [state.amount, state.amount, state.amount]
        )
      ).to.be.revertedWith("transfer amount exceeds allowance");

      // Confirm that user token balances have stayed the same
      let afterUserBalance = Number(
        await state.token1.balanceOf(userOne.address)
      );
      afterUserBalance.should.be.bignumber.equal(state.amount * 2);

      afterUserBalance = Number(
        await state.token2.balanceOf(userOne.address)
      );
      afterUserBalance.should.be.bignumber.equal(state.amount * 2);

      afterUserBalance = Number(
        await state.token3.balanceOf(userOne.address)
      );
      afterUserBalance.should.be.bignumber.equal(state.amount * 2);
    });

    it("should not allow user to multi-lock when parameters are malformed, not enough token amounts", async function () {
      // Attempt to lock tokens
      await expect(
        state.bridgeBank.connect(userOne).multiLock(
          [state.sender, state.sender, state.sender],
          [state.token1.address, state.token2.address, state.token3.address],
          [state.amount, state.amount]
        )
      ).to.be.revertedWith("M_P");
    });

    it("should not allow user to multi-lock when parameters are malformed, not enough token addresses", async function () {
      // Attempt to lock tokens
      await expect(
        state.bridgeBank.connect(userOne).multiLock(
          [state.sender, state.sender, state.sender],
          [state.token1.address, state.token2.address],
          [state.amount, state.amount, state.amount]
        )
      ).to.be.revertedWith("M_P");
    });

    it("should not allow user to multi-lock when parameters are malformed, not enough sif addresses", async function () {
      // Attempt to lock tokens
      await expect(
        state.bridgeBank.connect(userOne).multiLock(
          [state.sender, state.sender],
          [state.token1.address, state.token2.address, state.token3.address],
          [state.amount, state.amount, state.amount]
        )
      ).to.be.revertedWith("M_P");
    });

    it("should not allow user to multi-lock when parameters are malformed, invalid sif addresses", async function () {
      // Add the tokens into white list
      await batchAddTokensToEthWhitelist(state, [
        state.token1.address,
        state.token2.address,
        state.token3.address,
      ]);
      
      // Attempt to lock tokens
      await expect(
        state.bridgeBank.connect(userOne).multiLock(
          [state.sender + "ee", state.sender, state.sender],
          [state.token1.address, state.token2.address, state.token3.address],
          [state.amount, state.amount, state.amount]
        )
      ).to.be.revertedWith("INV_SIF_ADDR");
    });

    it("should not allow user to multi-lock when bridgebank is paused", async function () {
      await state.bridgeBank.connect(pauser).pause();
      // Attempt to lock tokens
      await expect(
        state.bridgeBank.connect(userOne).multiLock(
          [state.sender, state.sender, state.sender],
          [state.token1.address, state.token2.address, state.token3.address],
          [state.amount, state.amount, state.amount]
        )
      ).to.be.revertedWith("Pausable: paused");
    });

    it("should not allow user to multi-lock ERC20 tokens and Eth in the same call", async function () {
      // Add the tokens into white list
      await batchAddTokensToEthWhitelist(state, [
        state.token1.address,
        state.token2.address,
        state.token3.address,
      ]);

      // Attempt to lock tokens and Ether in the same call
      await expect(state.bridgeBank.connect(userOne).multiLock(
        [state.sender, state.sender, state.sender, state.sender],
        [state.token1.address, state.token2.address, state.token3.address, state.constants.zeroAddress],
        [state.amount, state.amount, state.amount, state.amount]
      ), { value: 100 })
        .to.be.revertedWith('Address: call to non-contract');
    });

    it("should not allow user to multi-burn tokens and Eth in the same call", async function () {
      // Add the tokens into whitelist
      // Also, add Ether into whitelist, which shouldn't be done but
      // we'll indulge in this scenario to bypass the whitelist requirements
      await state.bridgeBank.connect(owner).batchAddExistingBridgeTokens([
        state.token1.address,
        state.token2.address,
        state.token3.address,
        state.constants.zeroAddress
      ]);

      // Attempt to burn tokens and Ether in the same call
      await expect(state.bridgeBank.connect(userOne).multiLockBurn(
        [state.sender, state.sender, state.sender, state.sender],
        [state.token1.address, state.token2.address, state.token3.address, state.constants.zeroAddress],
        [state.amount, state.amount, state.amount, state.amount],
        [true, true, true, true]
      )).to.be.reverted;
    });
  });

  describe("Multi Lock Burn ERC20 Tokens", function () {
    it("should revert when parameters are malformed, not enough token amounts", async function () {
      // Attempt to lock tokens
      await expect(
        state.bridgeBank.connect(userOne).multiLockBurn(
          [state.sender, state.sender, state.sender],
          [state.token1.address, state.token2.address, state.token3.address],
          [state.amount, state.amount],
          [false, false, false],

        )
      ).to.be.revertedWith("M_P");
    });

    it("should revert when multi-lock parameters are malformed, not enough token addresses", async function () {
      // Attempt to lock tokens
      await expect(
        state.bridgeBank.connect(userOne).multiLockBurn(
          [state.sender, state.sender, state.sender],
          [state.token1.address, state.token2.address],
          [state.amount, state.amount, state.amount],
          [false, false, false]
        )
      ).to.be.revertedWith("M_P");
    });

    it("should revert when multi-lock parameters are malformed, not enough sif addresses", async function () {
      // Attempt to lock tokens
      await expect(
        state.bridgeBank.connect(userOne).multiLockBurn(
          [state.sender, state.sender],
          [state.token1.address, state.token2.address, state.token3.address],
          [state.amount, state.amount, state.amount],
          [false, false, false]
        )
      ).to.be.revertedWith("M_P");
    });

    it("should revert when multi-lock parameters are malformed, not enough booleans", async function () {
      // Attempt to lock tokens
      await expect(
        state.bridgeBank.connect(userOne).multiLockBurn(
          [state.sender, state.sender],
          [state.token1.address, state.token2.address],
          [state.amount, state.amount],
          [false, false, false]
        )
      ).to.be.revertedWith("M_P");
    });

    it("should revert when multi-lock parameters are malformed, invalid sif addresses", async function () {
      // Add the tokens into white list
      await batchAddTokensToEthWhitelist(state, [
        state.token1.address,
        state.token2.address,
        state.token3.address,
      ]);
      
      // Attempt to lock tokens
      await expect(
        state.bridgeBank.connect(userOne).multiLockBurn(
          [state.sender + "ee", state.sender, state.sender],
          [state.token1.address, state.token2.address, state.token3.address],
          [state.amount, state.amount, state.amount],
          [false, false, false]
        )
      ).to.be.revertedWith("INV_SIF_ADDR");
    });

    it("should not allow user to multi-lock/burn when bridgebank is paused", async function () {
      await state.bridgeBank.connect(pauser).pause();
      // Attempt to lock tokens
      await expect(
        state.bridgeBank.connect(userOne).multiLockBurn(
          [state.sender, state.sender, state.sender],
          [state.token1.address, state.token2.address, state.token3.address],
          [state.amount, state.amount, state.amount],
          [false, false, false]
        )
      ).to.be.revertedWith("Pausable: paused");
    });
  });

  describe("Whitelist", function () {
    it("should allow the operator to add a token to the whitelist", async function () {
      // Add the tokens into white list
      await expect(state.bridgeBank.connect(operator)
        .updateEthWhiteList(state.token1.address, true))
        .to.emit(state.bridgeBank, 'LogWhiteListUpdate')
        .withArgs(state.token1.address, true);

      expect(await state.bridgeBank.getTokenInEthWhiteList(state.token1.address)).to.be.equal(true);
    });

    it("should not allow user to add a token to the whitelist", async function () {
      // Add the tokens into white list
      await expect(state.bridgeBank.connect(userOne)
        .updateEthWhiteList(state.token1.address, true))
        .to.be.revertedWith("!operator");

      expect(await state.bridgeBank.getTokenInEthWhiteList(state.token1.address)).to.be.equal(false);
    });

    it("should allow the operator to remove a token from the whitelist", async function () {
      // Add the tokens into white list
      await state.bridgeBank.connect(operator)
        .updateEthWhiteList(state.token1.address, true)
        .should.be.fulfilled;

      // Remove the token from whitelist
      await expect(state.bridgeBank.connect(operator)
        .updateEthWhiteList(state.token1.address, false))
        .to.emit(state.bridgeBank, 'LogWhiteListUpdate')
        .withArgs(state.token1.address, false);

      expect(await state.bridgeBank.getTokenInEthWhiteList(state.token1.address)).to.be.equal(false);
    });

    it("should not allow user to remove a token from the whitelist", async function () {
      // Add the tokens into white list
      await state.bridgeBank.connect(operator)
        .updateEthWhiteList(state.token1.address, true)
        .should.be.fulfilled;

      await expect(state.bridgeBank.connect(userOne)
        .updateEthWhiteList(state.token1.address, false))
        .to.be.revertedWith("!operator");

      expect(await state.bridgeBank.getTokenInEthWhiteList(state.token1.address)).to.be.equal(true);
    });

    it("should not allow user to lock ERC20 tokens that are in Cosmos whitelist", async function () {
      // add token as BridgeToken
      await state.bridgeBank.connect(owner).addExistingBridgeToken(state.token1.address);

      // Attempt to lock tokens
      await expect(state.bridgeBank.connect(userOne).lock(
        state.sender,
        state.token1.address,
        state.amount
      )).to.be.revertedWith("Only token not in cosmos whitelist can be locked");
    });

    it("should not allow user to multi-lock ERC20 tokens if at least one of them is in cosmos whitelist", async function () {
      // add token1 as BridgeToken
      await state.bridgeBank.connect(owner).addExistingBridgeToken(state.token1.address);

      // Attempt to lock tokens
      await expect(state.bridgeBank.connect(userOne).multiLock(
        [state.sender, state.sender, state.sender],
        [state.token1.address, state.token2.address, state.token3.address],
        [state.amount, state.amount, state.amount]
      )).to.be.revertedWith("Only token not in cosmos whitelist can be locked");
    });

    it("should not allow user to multi-lock ERC20 tokens with multiLockBurn method if one of them is cosmos whitelist", async function () {
      // add token1 as BridgeToken
      await state.bridgeBank.connect(owner).addExistingBridgeToken(state.token1.address);

      // Attempt to lock tokens
      await expect(state.bridgeBank.connect(userOne).multiLockBurn(
        [state.sender, state.sender, state.sender],
        [state.token1.address, state.token2.address, state.token3.address],
        [state.amount, state.amount, state.amount],
        [false, false, false]
      )).to.be.revertedWith("Only token not in cosmos whitelist can be locked");
    });

    it("should not allow user to multi-lock and burn ERC20 tokens and rowan with multiLockBurn method if at least one of them is in cosmos whitelist ", async function () {
      // add token1 as BridgeToken
      await state.bridgeBank.connect(owner).addExistingBridgeToken(state.token1.address);

      // approve bridgebank to spend rowan
      await state.rowan.connect(userOne).approve(
        state.bridgeBank.address,
        state.amount
      );

      // Lock & burn tokens
      await expect(state.bridgeBank.connect(userOne).multiLockBurn(
        [state.sender, state.sender, state.sender],
        [state.token1.address, state.token2.address, state.rowan.address],
        [state.amount, state.amount, state.amount],
        [false, false, true]
      )).to.be.revertedWith("Only token not in cosmos whitelist can be locked");
    });
  });
});