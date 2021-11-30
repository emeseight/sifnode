// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MockUpgrade

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CosmosBridgeClaimData is an auto generated low-level Go binding around an user-defined struct.
type CosmosBridgeClaimData struct {
	CosmosSender         []byte
	CosmosSenderSequence *big.Int
	EthereumReceiver     common.Address
	TokenAddress         common.Address
	Amount               *big.Int
	TokenName            string
	TokenSymbol          string
	TokenDecimals        uint8
	NetworkDescriptor    int32
	DoublePeg            bool
	Nonce                *big.Int
	CosmosDenom          string
}

// CosmosBridgeSignatureData is an auto generated low-level Go binding around an user-defined struct.
type CosmosBridgeSignatureData struct {
	Signer common.Address
	V      uint8
	R      [32]byte
	S      [32]byte
}

// MockUpgradeMetaData contains all meta data concerning the MockUpgrade contract.
var MockUpgradeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridgeBank\",\"type\":\"address\"}],\"name\":\"LogBridgeBankSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"int32\",\"name\":\"_networkDescriptor\",\"type\":\"int32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceContractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cosmosDenom\",\"type\":\"string\"}],\"name\":\"LogNewBridgeTokenCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_prophecyID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"}],\"name\":\"LogNewOracleClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prophecyID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ethereumReceiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogNewProphecyClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prophecyID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"LogProphecyCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_prophecyID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_prophecyPowerCurrent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_prophecyPowerThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_submitter\",\"type\":\"address\"}],\"name\":\"LogProphecyProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_power\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_currentValsetVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalPower\",\"type\":\"uint256\"}],\"name\":\"LogValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_power\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_currentValsetVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalPower\",\"type\":\"uint256\"}],\"name\":\"LogValidatorPowerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_power\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_currentValsetVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalPower\",\"type\":\"uint256\"}],\"name\":\"LogValidatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newValsetVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalPower\",\"type\":\"uint256\"}],\"name\":\"LogValsetReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newValsetVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalPower\",\"type\":\"uint256\"}],\"name\":\"LogValsetUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_validatorPower\",\"type\":\"uint256\"}],\"name\":\"addValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"sigs\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"cosmosSender\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"cosmosSenderSequence\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"ethereumReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"int32\",\"name\":\"networkDescriptor\",\"type\":\"int32\"},{\"internalType\":\"bool\",\"name\":\"doublePeg\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"nonce\",\"type\":\"uint128\"},{\"internalType\":\"string\",\"name\":\"cosmosDenom\",\"type\":\"string\"}],\"internalType\":\"structCosmosBridge.ClaimData[]\",\"name\":\"claims\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"internalType\":\"structCosmosBridge.SignatureData[][]\",\"name\":\"signatureData\",\"type\":\"tuple[][]\"}],\"name\":\"batchSubmitProphecyClaimAggregatedSigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeBank\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOperator\",\"type\":\"address\"}],\"name\":\"changeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"consensusThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cosmosBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentValsetVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"cosmosSender\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"cosmosSenderSequence\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"ethereumReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"int32\",\"name\":\"_networkDescriptor\",\"type\":\"int32\"},{\"internalType\":\"bool\",\"name\":\"doublePeg\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"nonce\",\"type\":\"uint128\"},{\"internalType\":\"string\",\"name\":\"cosmosDenom\",\"type\":\"string\"}],\"name\":\"getProphecyID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"signedPower\",\"type\":\"uint256\"}],\"name\":\"getProphecyStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"}],\"name\":\"getValidatorPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasBridgeBank\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_consensusThreshold\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_initValidators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_initPowers\",\"type\":\"uint256[]\"},{\"internalType\":\"int32\",\"name\":\"_networkDescriptor\",\"type\":\"int32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"}],\"name\":\"isActiveValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastNonceSubmitted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkDescriptor\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"powers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_valsetVersion\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"}],\"name\":\"recoverGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_bridgeBank\",\"type\":\"address\"}],\"name\":\"setBridgeBank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sourceAddressToDestinationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashDigest\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"cosmosSender\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"cosmosSenderSequence\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"ethereumReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"int32\",\"name\":\"networkDescriptor\",\"type\":\"int32\"},{\"internalType\":\"bool\",\"name\":\"doublePeg\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"nonce\",\"type\":\"uint128\"},{\"internalType\":\"string\",\"name\":\"cosmosDenom\",\"type\":\"string\"}],\"internalType\":\"structCosmosBridge.ClaimData\",\"name\":\"claimData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"internalType\":\"structCosmosBridge.SignatureData[]\",\"name\":\"signatureData\",\"type\":\"tuple[]\"}],\"name\":\"submitProphecyClaimAggregatedSigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenFaucet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newValidatorPower\",\"type\":\"uint256\"}],\"name\":\"updateValidatorPower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"updateValset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MockUpgradeABI is the input ABI used to generate the binding from.
// Deprecated: Use MockUpgradeMetaData.ABI instead.
var MockUpgradeABI = MockUpgradeMetaData.ABI

// MockUpgrade is an auto generated Go binding around an Ethereum contract.
type MockUpgrade struct {
	MockUpgradeCaller     // Read-only binding to the contract
	MockUpgradeTransactor // Write-only binding to the contract
	MockUpgradeFilterer   // Log filterer for contract events
}

// MockUpgradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockUpgradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockUpgradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockUpgradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockUpgradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockUpgradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockUpgradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockUpgradeSession struct {
	Contract     *MockUpgrade      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockUpgradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockUpgradeCallerSession struct {
	Contract *MockUpgradeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MockUpgradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockUpgradeTransactorSession struct {
	Contract     *MockUpgradeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MockUpgradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockUpgradeRaw struct {
	Contract *MockUpgrade // Generic contract binding to access the raw methods on
}

// MockUpgradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockUpgradeCallerRaw struct {
	Contract *MockUpgradeCaller // Generic read-only contract binding to access the raw methods on
}

// MockUpgradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockUpgradeTransactorRaw struct {
	Contract *MockUpgradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockUpgrade creates a new instance of MockUpgrade, bound to a specific deployed contract.
func NewMockUpgrade(address common.Address, backend bind.ContractBackend) (*MockUpgrade, error) {
	contract, err := bindMockUpgrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockUpgrade{MockUpgradeCaller: MockUpgradeCaller{contract: contract}, MockUpgradeTransactor: MockUpgradeTransactor{contract: contract}, MockUpgradeFilterer: MockUpgradeFilterer{contract: contract}}, nil
}

// NewMockUpgradeCaller creates a new read-only instance of MockUpgrade, bound to a specific deployed contract.
func NewMockUpgradeCaller(address common.Address, caller bind.ContractCaller) (*MockUpgradeCaller, error) {
	contract, err := bindMockUpgrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockUpgradeCaller{contract: contract}, nil
}

// NewMockUpgradeTransactor creates a new write-only instance of MockUpgrade, bound to a specific deployed contract.
func NewMockUpgradeTransactor(address common.Address, transactor bind.ContractTransactor) (*MockUpgradeTransactor, error) {
	contract, err := bindMockUpgrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockUpgradeTransactor{contract: contract}, nil
}

// NewMockUpgradeFilterer creates a new log filterer instance of MockUpgrade, bound to a specific deployed contract.
func NewMockUpgradeFilterer(address common.Address, filterer bind.ContractFilterer) (*MockUpgradeFilterer, error) {
	contract, err := bindMockUpgrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockUpgradeFilterer{contract: contract}, nil
}

// bindMockUpgrade binds a generic wrapper to an already deployed contract.
func bindMockUpgrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockUpgradeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockUpgrade *MockUpgradeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockUpgrade.Contract.MockUpgradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockUpgrade *MockUpgradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockUpgrade.Contract.MockUpgradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockUpgrade *MockUpgradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockUpgrade.Contract.MockUpgradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockUpgrade *MockUpgradeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockUpgrade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockUpgrade *MockUpgradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockUpgrade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockUpgrade *MockUpgradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockUpgrade.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MockUpgrade *MockUpgradeCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MockUpgrade.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MockUpgrade *MockUpgradeSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MockUpgrade.Contract.BalanceOf(&_MockUpgrade.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MockUpgrade *MockUpgradeCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MockUpgrade.Contract.BalanceOf(&_MockUpgrade.CallOpts, account)
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_MockUpgrade *MockUpgradeCaller) BridgeBank(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockUpgrade.contract.Call(opts, &out, "bridgeBank")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_MockUpgrade *MockUpgradeSession) BridgeBank() (common.Address, error) {
	return _MockUpgrade.Contract.BridgeBank(&_MockUpgrade.CallOpts)
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_MockUpgrade *MockUpgradeCallerSession) BridgeBank() (common.Address, error) {
	return _MockUpgrade.Contract.BridgeBank(&_MockUpgrade.CallOpts)
}

// ConsensusThreshold is a free data retrieval call binding the contract method 0xf9b0b5b9.
//
// Solidity: function consensusThreshold() view returns(uint256)
func (_MockUpgrade *MockUpgradeCaller) ConsensusThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockUpgrade.contract.Call(opts, &out, "consensusThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConsensusThreshold is a free data retrieval call binding the contract method 0xf9b0b5b9.
//
// Solidity: function consensusThreshold() view returns(uint256)
func (_MockUpgrade *MockUpgradeSession) ConsensusThreshold() (*big.Int, error) {
	return _MockUpgrade.Contract.ConsensusThreshold(&_MockUpgrade.CallOpts)
}

// ConsensusThreshold is a free data retrieval call binding the contract method 0xf9b0b5b9.
//
// Solidity: function consensusThreshold() view returns(uint256)
func (_MockUpgrade *MockUpgradeCallerSession) ConsensusThreshold() (*big.Int, error) {
	return _MockUpgrade.Contract.ConsensusThreshold(&_MockUpgrade.CallOpts)
}

// CosmosBridge is a free data retrieval call binding the contract method 0xb0e9ef71.
//
// Solidity: function cosmosBridge() view returns(address)
func (_MockUpgrade *MockUpgradeCaller) CosmosBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockUpgrade.contract.Call(opts, &out, "cosmosBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CosmosBridge is a free data retrieval call binding the contract method 0xb0e9ef71.
//
// Solidity: function cosmosBridge() view returns(address)
func (_MockUpgrade *MockUpgradeSession) CosmosBridge() (common.Address, error) {
	return _MockUpgrade.Contract.CosmosBridge(&_MockUpgrade.CallOpts)
}

// CosmosBridge is a free data retrieval call binding the contract method 0xb0e9ef71.
//
// Solidity: function cosmosBridge() view returns(address)
func (_MockUpgrade *MockUpgradeCallerSession) CosmosBridge() (common.Address, error) {
	return _MockUpgrade.Contract.CosmosBridge(&_MockUpgrade.CallOpts)
}

// CurrentValsetVersion is a free data retrieval call binding the contract method 0x8d56c37d.
//
// Solidity: function currentValsetVersion() view returns(uint256)
func (_MockUpgrade *MockUpgradeCaller) CurrentValsetVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockUpgrade.contract.Call(opts, &out, "currentValsetVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentValsetVersion is a free data retrieval call binding the contract method 0x8d56c37d.
//
// Solidity: function currentValsetVersion() view returns(uint256)
func (_MockUpgrade *MockUpgradeSession) CurrentValsetVersion() (*big.Int, error) {
	return _MockUpgrade.Contract.CurrentValsetVersion(&_MockUpgrade.CallOpts)
}

// CurrentValsetVersion is a free data retrieval call binding the contract method 0x8d56c37d.
//
// Solidity: function currentValsetVersion() view returns(uint256)
func (_MockUpgrade *MockUpgradeCallerSession) CurrentValsetVersion() (*big.Int, error) {
	return _MockUpgrade.Contract.CurrentValsetVersion(&_MockUpgrade.CallOpts)
}

// GetProphecyID is a free data retrieval call binding the contract method 0x257266a8.
//
// Solidity: function getProphecyID(bytes cosmosSender, uint256 cosmosSenderSequence, address ethereumReceiver, address tokenAddress, uint256 amount, string tokenName, string tokenSymbol, uint8 tokenDecimals, int32 _networkDescriptor, bool doublePeg, uint128 nonce, string cosmosDenom) pure returns(uint256)
func (_MockUpgrade *MockUpgradeCaller) GetProphecyID(opts *bind.CallOpts, cosmosSender []byte, cosmosSenderSequence *big.Int, ethereumReceiver common.Address, tokenAddress common.Address, amount *big.Int, tokenName string, tokenSymbol string, tokenDecimals uint8, _networkDescriptor int32, doublePeg bool, nonce *big.Int, cosmosDenom string) (*big.Int, error) {
	var out []interface{}
	err := _MockUpgrade.contract.Call(opts, &out, "getProphecyID", cosmosSender, cosmosSenderSequence, ethereumReceiver, tokenAddress, amount, tokenName, tokenSymbol, tokenDecimals, _networkDescriptor, doublePeg, nonce, cosmosDenom)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProphecyID is a free data retrieval call binding the contract method 0x257266a8.
//
// Solidity: function getProphecyID(bytes cosmosSender, uint256 cosmosSenderSequence, address ethereumReceiver, address tokenAddress, uint256 amount, string tokenName, string tokenSymbol, uint8 tokenDecimals, int32 _networkDescriptor, bool doublePeg, uint128 nonce, string cosmosDenom) pure returns(uint256)
func (_MockUpgrade *MockUpgradeSession) GetProphecyID(cosmosSender []byte, cosmosSenderSequence *big.Int, ethereumReceiver common.Address, tokenAddress common.Address, amount *big.Int, tokenName string, tokenSymbol string, tokenDecimals uint8, _networkDescriptor int32, doublePeg bool, nonce *big.Int, cosmosDenom string) (*big.Int, error) {
	return _MockUpgrade.Contract.GetProphecyID(&_MockUpgrade.CallOpts, cosmosSender, cosmosSenderSequence, ethereumReceiver, tokenAddress, amount, tokenName, tokenSymbol, tokenDecimals, _networkDescriptor, doublePeg, nonce, cosmosDenom)
}

// GetProphecyID is a free data retrieval call binding the contract method 0x257266a8.
//
// Solidity: function getProphecyID(bytes cosmosSender, uint256 cosmosSenderSequence, address ethereumReceiver, address tokenAddress, uint256 amount, string tokenName, string tokenSymbol, uint8 tokenDecimals, int32 _networkDescriptor, bool doublePeg, uint128 nonce, string cosmosDenom) pure returns(uint256)
func (_MockUpgrade *MockUpgradeCallerSession) GetProphecyID(cosmosSender []byte, cosmosSenderSequence *big.Int, ethereumReceiver common.Address, tokenAddress common.Address, amount *big.Int, tokenName string, tokenSymbol string, tokenDecimals uint8, _networkDescriptor int32, doublePeg bool, nonce *big.Int, cosmosDenom string) (*big.Int, error) {
	return _MockUpgrade.Contract.GetProphecyID(&_MockUpgrade.CallOpts, cosmosSender, cosmosSenderSequence, ethereumReceiver, tokenAddress, amount, tokenName, tokenSymbol, tokenDecimals, _networkDescriptor, doublePeg, nonce, cosmosDenom)
}

// GetProphecyStatus is a free data retrieval call binding the contract method 0x77491c75.
//
// Solidity: function getProphecyStatus(uint256 signedPower) view returns(bool)
func (_MockUpgrade *MockUpgradeCaller) GetProphecyStatus(opts *bind.CallOpts, signedPower *big.Int) (bool, error) {
	var out []interface{}
	err := _MockUpgrade.contract.Call(opts, &out, "getProphecyStatus", signedPower)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetProphecyStatus is a free data retrieval call binding the contract method 0x77491c75.
//
// Solidity: function getProphecyStatus(uint256 signedPower) view returns(bool)
func (_MockUpgrade *MockUpgradeSession) GetProphecyStatus(signedPower *big.Int) (bool, error) {
	return _MockUpgrade.Contract.GetProphecyStatus(&_MockUpgrade.CallOpts, signedPower)
}

// GetProphecyStatus is a free data retrieval call binding the contract method 0x77491c75.
//
// Solidity: function getProphecyStatus(uint256 signedPower) view returns(bool)
func (_MockUpgrade *MockUpgradeCallerSession) GetProphecyStatus(signedPower *big.Int) (bool, error) {
	return _MockUpgrade.Contract.GetProphecyStatus(&_MockUpgrade.CallOpts, signedPower)
}

// GetValidatorPower is a free data retrieval call binding the contract method 0x473691a4.
//
// Solidity: function getValidatorPower(address _validatorAddress) view returns(uint256)
func (_MockUpgrade *MockUpgradeCaller) GetValidatorPower(opts *bind.CallOpts, _validatorAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MockUpgrade.contract.Call(opts, &out, "getValidatorPower", _validatorAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorPower is a free data retrieval call binding the contract method 0x473691a4.
//
// Solidity: function getValidatorPower(address _validatorAddress) view returns(uint256)
func (_MockUpgrade *MockUpgradeSession) GetValidatorPower(_validatorAddress common.Address) (*big.Int, error) {
	return _MockUpgrade.Contract.GetValidatorPower(&_MockUpgrade.CallOpts, _validatorAddress)
}

// GetValidatorPower is a free data retrieval call binding the contract method 0x473691a4.
//
// Solidity: function getValidatorPower(address _validatorAddress) view returns(uint256)
func (_MockUpgrade *MockUpgradeCallerSession) GetValidatorPower(_validatorAddress common.Address) (*big.Int, error) {
	return _MockUpgrade.Contract.GetValidatorPower(&_MockUpgrade.CallOpts, _validatorAddress)
}

// HasBridgeBank is a free data retrieval call binding the contract method 0x69294a4e.
//
// Solidity: function hasBridgeBank() view returns(bool)
func (_MockUpgrade *MockUpgradeCaller) HasBridgeBank(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MockUpgrade.contract.Call(opts, &out, "hasBridgeBank")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasBridgeBank is a free data retrieval call binding the contract method 0x69294a4e.
//
// Solidity: function hasBridgeBank() view returns(bool)
func (_MockUpgrade *MockUpgradeSession) HasBridgeBank() (bool, error) {
	return _MockUpgrade.Contract.HasBridgeBank(&_MockUpgrade.CallOpts)
}

// HasBridgeBank is a free data retrieval call binding the contract method 0x69294a4e.
//
// Solidity: function hasBridgeBank() view returns(bool)
func (_MockUpgrade *MockUpgradeCallerSession) HasBridgeBank() (bool, error) {
	return _MockUpgrade.Contract.HasBridgeBank(&_MockUpgrade.CallOpts)
}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address _validatorAddress) view returns(bool)
func (_MockUpgrade *MockUpgradeCaller) IsActiveValidator(opts *bind.CallOpts, _validatorAddress common.Address) (bool, error) {
	var out []interface{}
	err := _MockUpgrade.contract.Call(opts, &out, "isActiveValidator", _validatorAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address _validatorAddress) view returns(bool)
func (_MockUpgrade *MockUpgradeSession) IsActiveValidator(_validatorAddress common.Address) (bool, error) {
	return _MockUpgrade.Contract.IsActiveValidator(&_MockUpgrade.CallOpts, _validatorAddress)
}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address _validatorAddress) view returns(bool)
func (_MockUpgrade *MockUpgradeCallerSession) IsActiveValidator(_validatorAddress common.Address) (bool, error) {
	return _MockUpgrade.Contract.IsActiveValidator(&_MockUpgrade.CallOpts, _validatorAddress)
}

// LastNonceSubmitted is a free data retrieval call binding the contract method 0x457c1288.
//
// Solidity: function lastNonceSubmitted() view returns(uint256)
func (_MockUpgrade *MockUpgradeCaller) LastNonceSubmitted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockUpgrade.contract.Call(opts, &out, "lastNonceSubmitted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastNonceSubmitted is a free data retrieval call binding the contract method 0x457c1288.
//
// Solidity: function lastNonceSubmitted() view returns(uint256)
func (_MockUpgrade *MockUpgradeSession) LastNonceSubmitted() (*big.Int, error) {
	return _MockUpgrade.Contract.LastNonceSubmitted(&_MockUpgrade.CallOpts)
}

// LastNonceSubmitted is a free data retrieval call binding the contract method 0x457c1288.
//
// Solidity: function lastNonceSubmitted() view returns(uint256)
func (_MockUpgrade *MockUpgradeCallerSession) LastNonceSubmitted() (*big.Int, error) {
	return _MockUpgrade.Contract.LastNonceSubmitted(&_MockUpgrade.CallOpts)
}

// NetworkDescriptor is a free data retrieval call binding the contract method 0x909d4a1c.
//
// Solidity: function networkDescriptor() view returns(int32)
func (_MockUpgrade *MockUpgradeCaller) NetworkDescriptor(opts *bind.CallOpts) (int32, error) {
	var out []interface{}
	err := _MockUpgrade.contract.Call(opts, &out, "networkDescriptor")

	if err != nil {
		return *new(int32), err
	}

	out0 := *abi.ConvertType(out[0], new(int32)).(*int32)

	return out0, err

}

// NetworkDescriptor is a free data retrieval call binding the contract method 0x909d4a1c.
//
// Solidity: function networkDescriptor() view returns(int32)
func (_MockUpgrade *MockUpgradeSession) NetworkDescriptor() (int32, error) {
	return _MockUpgrade.Contract.NetworkDescriptor(&_MockUpgrade.CallOpts)
}

// NetworkDescriptor is a free data retrieval call binding the contract method 0x909d4a1c.
//
// Solidity: function networkDescriptor() view returns(int32)
func (_MockUpgrade *MockUpgradeCallerSession) NetworkDescriptor() (int32, error) {
	return _MockUpgrade.Contract.NetworkDescriptor(&_MockUpgrade.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_MockUpgrade *MockUpgradeCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockUpgrade.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_MockUpgrade *MockUpgradeSession) Operator() (common.Address, error) {
	return _MockUpgrade.Contract.Operator(&_MockUpgrade.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_MockUpgrade *MockUpgradeCallerSession) Operator() (common.Address, error) {
	return _MockUpgrade.Contract.Operator(&_MockUpgrade.CallOpts)
}

// Powers is a free data retrieval call binding the contract method 0x850f43dd.
//
// Solidity: function powers(address , uint256 ) view returns(uint256)
func (_MockUpgrade *MockUpgradeCaller) Powers(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MockUpgrade.contract.Call(opts, &out, "powers", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Powers is a free data retrieval call binding the contract method 0x850f43dd.
//
// Solidity: function powers(address , uint256 ) view returns(uint256)
func (_MockUpgrade *MockUpgradeSession) Powers(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _MockUpgrade.Contract.Powers(&_MockUpgrade.CallOpts, arg0, arg1)
}

// Powers is a free data retrieval call binding the contract method 0x850f43dd.
//
// Solidity: function powers(address , uint256 ) view returns(uint256)
func (_MockUpgrade *MockUpgradeCallerSession) Powers(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _MockUpgrade.Contract.Powers(&_MockUpgrade.CallOpts, arg0, arg1)
}

// SourceAddressToDestinationAddress is a free data retrieval call binding the contract method 0x7b010263.
//
// Solidity: function sourceAddressToDestinationAddress(address ) view returns(address)
func (_MockUpgrade *MockUpgradeCaller) SourceAddressToDestinationAddress(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _MockUpgrade.contract.Call(opts, &out, "sourceAddressToDestinationAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SourceAddressToDestinationAddress is a free data retrieval call binding the contract method 0x7b010263.
//
// Solidity: function sourceAddressToDestinationAddress(address ) view returns(address)
func (_MockUpgrade *MockUpgradeSession) SourceAddressToDestinationAddress(arg0 common.Address) (common.Address, error) {
	return _MockUpgrade.Contract.SourceAddressToDestinationAddress(&_MockUpgrade.CallOpts, arg0)
}

// SourceAddressToDestinationAddress is a free data retrieval call binding the contract method 0x7b010263.
//
// Solidity: function sourceAddressToDestinationAddress(address ) view returns(address)
func (_MockUpgrade *MockUpgradeCallerSession) SourceAddressToDestinationAddress(arg0 common.Address) (common.Address, error) {
	return _MockUpgrade.Contract.SourceAddressToDestinationAddress(&_MockUpgrade.CallOpts, arg0)
}

// TotalPower is a free data retrieval call binding the contract method 0xdb3ad22c.
//
// Solidity: function totalPower() view returns(uint256)
func (_MockUpgrade *MockUpgradeCaller) TotalPower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockUpgrade.contract.Call(opts, &out, "totalPower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPower is a free data retrieval call binding the contract method 0xdb3ad22c.
//
// Solidity: function totalPower() view returns(uint256)
func (_MockUpgrade *MockUpgradeSession) TotalPower() (*big.Int, error) {
	return _MockUpgrade.Contract.TotalPower(&_MockUpgrade.CallOpts)
}

// TotalPower is a free data retrieval call binding the contract method 0xdb3ad22c.
//
// Solidity: function totalPower() view returns(uint256)
func (_MockUpgrade *MockUpgradeCallerSession) TotalPower() (*big.Int, error) {
	return _MockUpgrade.Contract.TotalPower(&_MockUpgrade.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_MockUpgrade *MockUpgradeCaller) ValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockUpgrade.contract.Call(opts, &out, "validatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_MockUpgrade *MockUpgradeSession) ValidatorCount() (*big.Int, error) {
	return _MockUpgrade.Contract.ValidatorCount(&_MockUpgrade.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_MockUpgrade *MockUpgradeCallerSession) ValidatorCount() (*big.Int, error) {
	return _MockUpgrade.Contract.ValidatorCount(&_MockUpgrade.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0x45aaf18c.
//
// Solidity: function validators(address , uint256 ) view returns(bool)
func (_MockUpgrade *MockUpgradeCaller) Validators(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _MockUpgrade.contract.Call(opts, &out, "validators", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Validators is a free data retrieval call binding the contract method 0x45aaf18c.
//
// Solidity: function validators(address , uint256 ) view returns(bool)
func (_MockUpgrade *MockUpgradeSession) Validators(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _MockUpgrade.Contract.Validators(&_MockUpgrade.CallOpts, arg0, arg1)
}

// Validators is a free data retrieval call binding the contract method 0x45aaf18c.
//
// Solidity: function validators(address , uint256 ) view returns(bool)
func (_MockUpgrade *MockUpgradeCallerSession) Validators(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _MockUpgrade.Contract.Validators(&_MockUpgrade.CallOpts, arg0, arg1)
}

// AddValidator is a paid mutator transaction binding the contract method 0xfc6c1f02.
//
// Solidity: function addValidator(address _validatorAddress, uint256 _validatorPower) returns()
func (_MockUpgrade *MockUpgradeTransactor) AddValidator(opts *bind.TransactOpts, _validatorAddress common.Address, _validatorPower *big.Int) (*types.Transaction, error) {
	return _MockUpgrade.contract.Transact(opts, "addValidator", _validatorAddress, _validatorPower)
}

// AddValidator is a paid mutator transaction binding the contract method 0xfc6c1f02.
//
// Solidity: function addValidator(address _validatorAddress, uint256 _validatorPower) returns()
func (_MockUpgrade *MockUpgradeSession) AddValidator(_validatorAddress common.Address, _validatorPower *big.Int) (*types.Transaction, error) {
	return _MockUpgrade.Contract.AddValidator(&_MockUpgrade.TransactOpts, _validatorAddress, _validatorPower)
}

// AddValidator is a paid mutator transaction binding the contract method 0xfc6c1f02.
//
// Solidity: function addValidator(address _validatorAddress, uint256 _validatorPower) returns()
func (_MockUpgrade *MockUpgradeTransactorSession) AddValidator(_validatorAddress common.Address, _validatorPower *big.Int) (*types.Transaction, error) {
	return _MockUpgrade.Contract.AddValidator(&_MockUpgrade.TransactOpts, _validatorAddress, _validatorPower)
}

// BatchSubmitProphecyClaimAggregatedSigs is a paid mutator transaction binding the contract method 0x5078f631.
//
// Solidity: function batchSubmitProphecyClaimAggregatedSigs(bytes32[] sigs, (bytes,uint256,address,address,uint256,string,string,uint8,int32,bool,uint128,string)[] claims, (address,uint8,bytes32,bytes32)[][] signatureData) returns()
func (_MockUpgrade *MockUpgradeTransactor) BatchSubmitProphecyClaimAggregatedSigs(opts *bind.TransactOpts, sigs [][32]byte, claims []CosmosBridgeClaimData, signatureData [][]CosmosBridgeSignatureData) (*types.Transaction, error) {
	return _MockUpgrade.contract.Transact(opts, "batchSubmitProphecyClaimAggregatedSigs", sigs, claims, signatureData)
}

// BatchSubmitProphecyClaimAggregatedSigs is a paid mutator transaction binding the contract method 0x5078f631.
//
// Solidity: function batchSubmitProphecyClaimAggregatedSigs(bytes32[] sigs, (bytes,uint256,address,address,uint256,string,string,uint8,int32,bool,uint128,string)[] claims, (address,uint8,bytes32,bytes32)[][] signatureData) returns()
func (_MockUpgrade *MockUpgradeSession) BatchSubmitProphecyClaimAggregatedSigs(sigs [][32]byte, claims []CosmosBridgeClaimData, signatureData [][]CosmosBridgeSignatureData) (*types.Transaction, error) {
	return _MockUpgrade.Contract.BatchSubmitProphecyClaimAggregatedSigs(&_MockUpgrade.TransactOpts, sigs, claims, signatureData)
}

// BatchSubmitProphecyClaimAggregatedSigs is a paid mutator transaction binding the contract method 0x5078f631.
//
// Solidity: function batchSubmitProphecyClaimAggregatedSigs(bytes32[] sigs, (bytes,uint256,address,address,uint256,string,string,uint8,int32,bool,uint128,string)[] claims, (address,uint8,bytes32,bytes32)[][] signatureData) returns()
func (_MockUpgrade *MockUpgradeTransactorSession) BatchSubmitProphecyClaimAggregatedSigs(sigs [][32]byte, claims []CosmosBridgeClaimData, signatureData [][]CosmosBridgeSignatureData) (*types.Transaction, error) {
	return _MockUpgrade.Contract.BatchSubmitProphecyClaimAggregatedSigs(&_MockUpgrade.TransactOpts, sigs, claims, signatureData)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address _newOperator) returns()
func (_MockUpgrade *MockUpgradeTransactor) ChangeOperator(opts *bind.TransactOpts, _newOperator common.Address) (*types.Transaction, error) {
	return _MockUpgrade.contract.Transact(opts, "changeOperator", _newOperator)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address _newOperator) returns()
func (_MockUpgrade *MockUpgradeSession) ChangeOperator(_newOperator common.Address) (*types.Transaction, error) {
	return _MockUpgrade.Contract.ChangeOperator(&_MockUpgrade.TransactOpts, _newOperator)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address _newOperator) returns()
func (_MockUpgrade *MockUpgradeTransactorSession) ChangeOperator(_newOperator common.Address) (*types.Transaction, error) {
	return _MockUpgrade.Contract.ChangeOperator(&_MockUpgrade.TransactOpts, _newOperator)
}

// Initialize is a paid mutator transaction binding the contract method 0xdf3ca886.
//
// Solidity: function initialize(address _operator, uint256 _consensusThreshold, address[] _initValidators, uint256[] _initPowers, int32 _networkDescriptor) returns()
func (_MockUpgrade *MockUpgradeTransactor) Initialize(opts *bind.TransactOpts, _operator common.Address, _consensusThreshold *big.Int, _initValidators []common.Address, _initPowers []*big.Int, _networkDescriptor int32) (*types.Transaction, error) {
	return _MockUpgrade.contract.Transact(opts, "initialize", _operator, _consensusThreshold, _initValidators, _initPowers, _networkDescriptor)
}

// Initialize is a paid mutator transaction binding the contract method 0xdf3ca886.
//
// Solidity: function initialize(address _operator, uint256 _consensusThreshold, address[] _initValidators, uint256[] _initPowers, int32 _networkDescriptor) returns()
func (_MockUpgrade *MockUpgradeSession) Initialize(_operator common.Address, _consensusThreshold *big.Int, _initValidators []common.Address, _initPowers []*big.Int, _networkDescriptor int32) (*types.Transaction, error) {
	return _MockUpgrade.Contract.Initialize(&_MockUpgrade.TransactOpts, _operator, _consensusThreshold, _initValidators, _initPowers, _networkDescriptor)
}

// Initialize is a paid mutator transaction binding the contract method 0xdf3ca886.
//
// Solidity: function initialize(address _operator, uint256 _consensusThreshold, address[] _initValidators, uint256[] _initPowers, int32 _networkDescriptor) returns()
func (_MockUpgrade *MockUpgradeTransactorSession) Initialize(_operator common.Address, _consensusThreshold *big.Int, _initValidators []common.Address, _initPowers []*big.Int, _networkDescriptor int32) (*types.Transaction, error) {
	return _MockUpgrade.Contract.Initialize(&_MockUpgrade.TransactOpts, _operator, _consensusThreshold, _initValidators, _initPowers, _networkDescriptor)
}

// RecoverGas is a paid mutator transaction binding the contract method 0xb5672be3.
//
// Solidity: function recoverGas(uint256 _valsetVersion, address _validatorAddress) returns()
func (_MockUpgrade *MockUpgradeTransactor) RecoverGas(opts *bind.TransactOpts, _valsetVersion *big.Int, _validatorAddress common.Address) (*types.Transaction, error) {
	return _MockUpgrade.contract.Transact(opts, "recoverGas", _valsetVersion, _validatorAddress)
}

// RecoverGas is a paid mutator transaction binding the contract method 0xb5672be3.
//
// Solidity: function recoverGas(uint256 _valsetVersion, address _validatorAddress) returns()
func (_MockUpgrade *MockUpgradeSession) RecoverGas(_valsetVersion *big.Int, _validatorAddress common.Address) (*types.Transaction, error) {
	return _MockUpgrade.Contract.RecoverGas(&_MockUpgrade.TransactOpts, _valsetVersion, _validatorAddress)
}

// RecoverGas is a paid mutator transaction binding the contract method 0xb5672be3.
//
// Solidity: function recoverGas(uint256 _valsetVersion, address _validatorAddress) returns()
func (_MockUpgrade *MockUpgradeTransactorSession) RecoverGas(_valsetVersion *big.Int, _validatorAddress common.Address) (*types.Transaction, error) {
	return _MockUpgrade.Contract.RecoverGas(&_MockUpgrade.TransactOpts, _valsetVersion, _validatorAddress)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validatorAddress) returns()
func (_MockUpgrade *MockUpgradeTransactor) RemoveValidator(opts *bind.TransactOpts, _validatorAddress common.Address) (*types.Transaction, error) {
	return _MockUpgrade.contract.Transact(opts, "removeValidator", _validatorAddress)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validatorAddress) returns()
func (_MockUpgrade *MockUpgradeSession) RemoveValidator(_validatorAddress common.Address) (*types.Transaction, error) {
	return _MockUpgrade.Contract.RemoveValidator(&_MockUpgrade.TransactOpts, _validatorAddress)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validatorAddress) returns()
func (_MockUpgrade *MockUpgradeTransactorSession) RemoveValidator(_validatorAddress common.Address) (*types.Transaction, error) {
	return _MockUpgrade.Contract.RemoveValidator(&_MockUpgrade.TransactOpts, _validatorAddress)
}

// SetBridgeBank is a paid mutator transaction binding the contract method 0x814c92c3.
//
// Solidity: function setBridgeBank(address _bridgeBank) returns()
func (_MockUpgrade *MockUpgradeTransactor) SetBridgeBank(opts *bind.TransactOpts, _bridgeBank common.Address) (*types.Transaction, error) {
	return _MockUpgrade.contract.Transact(opts, "setBridgeBank", _bridgeBank)
}

// SetBridgeBank is a paid mutator transaction binding the contract method 0x814c92c3.
//
// Solidity: function setBridgeBank(address _bridgeBank) returns()
func (_MockUpgrade *MockUpgradeSession) SetBridgeBank(_bridgeBank common.Address) (*types.Transaction, error) {
	return _MockUpgrade.Contract.SetBridgeBank(&_MockUpgrade.TransactOpts, _bridgeBank)
}

// SetBridgeBank is a paid mutator transaction binding the contract method 0x814c92c3.
//
// Solidity: function setBridgeBank(address _bridgeBank) returns()
func (_MockUpgrade *MockUpgradeTransactorSession) SetBridgeBank(_bridgeBank common.Address) (*types.Transaction, error) {
	return _MockUpgrade.Contract.SetBridgeBank(&_MockUpgrade.TransactOpts, _bridgeBank)
}

// SubmitProphecyClaimAggregatedSigs is a paid mutator transaction binding the contract method 0x72cbd7b9.
//
// Solidity: function submitProphecyClaimAggregatedSigs(bytes32 hashDigest, (bytes,uint256,address,address,uint256,string,string,uint8,int32,bool,uint128,string) claimData, (address,uint8,bytes32,bytes32)[] signatureData) returns()
func (_MockUpgrade *MockUpgradeTransactor) SubmitProphecyClaimAggregatedSigs(opts *bind.TransactOpts, hashDigest [32]byte, claimData CosmosBridgeClaimData, signatureData []CosmosBridgeSignatureData) (*types.Transaction, error) {
	return _MockUpgrade.contract.Transact(opts, "submitProphecyClaimAggregatedSigs", hashDigest, claimData, signatureData)
}

// SubmitProphecyClaimAggregatedSigs is a paid mutator transaction binding the contract method 0x72cbd7b9.
//
// Solidity: function submitProphecyClaimAggregatedSigs(bytes32 hashDigest, (bytes,uint256,address,address,uint256,string,string,uint8,int32,bool,uint128,string) claimData, (address,uint8,bytes32,bytes32)[] signatureData) returns()
func (_MockUpgrade *MockUpgradeSession) SubmitProphecyClaimAggregatedSigs(hashDigest [32]byte, claimData CosmosBridgeClaimData, signatureData []CosmosBridgeSignatureData) (*types.Transaction, error) {
	return _MockUpgrade.Contract.SubmitProphecyClaimAggregatedSigs(&_MockUpgrade.TransactOpts, hashDigest, claimData, signatureData)
}

// SubmitProphecyClaimAggregatedSigs is a paid mutator transaction binding the contract method 0x72cbd7b9.
//
// Solidity: function submitProphecyClaimAggregatedSigs(bytes32 hashDigest, (bytes,uint256,address,address,uint256,string,string,uint8,int32,bool,uint128,string) claimData, (address,uint8,bytes32,bytes32)[] signatureData) returns()
func (_MockUpgrade *MockUpgradeTransactorSession) SubmitProphecyClaimAggregatedSigs(hashDigest [32]byte, claimData CosmosBridgeClaimData, signatureData []CosmosBridgeSignatureData) (*types.Transaction, error) {
	return _MockUpgrade.Contract.SubmitProphecyClaimAggregatedSigs(&_MockUpgrade.TransactOpts, hashDigest, claimData, signatureData)
}

// TokenFaucet is a paid mutator transaction binding the contract method 0x89eb068a.
//
// Solidity: function tokenFaucet() returns()
func (_MockUpgrade *MockUpgradeTransactor) TokenFaucet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockUpgrade.contract.Transact(opts, "tokenFaucet")
}

// TokenFaucet is a paid mutator transaction binding the contract method 0x89eb068a.
//
// Solidity: function tokenFaucet() returns()
func (_MockUpgrade *MockUpgradeSession) TokenFaucet() (*types.Transaction, error) {
	return _MockUpgrade.Contract.TokenFaucet(&_MockUpgrade.TransactOpts)
}

// TokenFaucet is a paid mutator transaction binding the contract method 0x89eb068a.
//
// Solidity: function tokenFaucet() returns()
func (_MockUpgrade *MockUpgradeTransactorSession) TokenFaucet() (*types.Transaction, error) {
	return _MockUpgrade.Contract.TokenFaucet(&_MockUpgrade.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MockUpgrade *MockUpgradeTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockUpgrade.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MockUpgrade *MockUpgradeSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockUpgrade.Contract.Transfer(&_MockUpgrade.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MockUpgrade *MockUpgradeTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockUpgrade.Contract.Transfer(&_MockUpgrade.TransactOpts, recipient, amount)
}

// UpdateValidatorPower is a paid mutator transaction binding the contract method 0x2e75293b.
//
// Solidity: function updateValidatorPower(address _validatorAddress, uint256 _newValidatorPower) returns()
func (_MockUpgrade *MockUpgradeTransactor) UpdateValidatorPower(opts *bind.TransactOpts, _validatorAddress common.Address, _newValidatorPower *big.Int) (*types.Transaction, error) {
	return _MockUpgrade.contract.Transact(opts, "updateValidatorPower", _validatorAddress, _newValidatorPower)
}

// UpdateValidatorPower is a paid mutator transaction binding the contract method 0x2e75293b.
//
// Solidity: function updateValidatorPower(address _validatorAddress, uint256 _newValidatorPower) returns()
func (_MockUpgrade *MockUpgradeSession) UpdateValidatorPower(_validatorAddress common.Address, _newValidatorPower *big.Int) (*types.Transaction, error) {
	return _MockUpgrade.Contract.UpdateValidatorPower(&_MockUpgrade.TransactOpts, _validatorAddress, _newValidatorPower)
}

// UpdateValidatorPower is a paid mutator transaction binding the contract method 0x2e75293b.
//
// Solidity: function updateValidatorPower(address _validatorAddress, uint256 _newValidatorPower) returns()
func (_MockUpgrade *MockUpgradeTransactorSession) UpdateValidatorPower(_validatorAddress common.Address, _newValidatorPower *big.Int) (*types.Transaction, error) {
	return _MockUpgrade.Contract.UpdateValidatorPower(&_MockUpgrade.TransactOpts, _validatorAddress, _newValidatorPower)
}

// UpdateValset is a paid mutator transaction binding the contract method 0x788cf92f.
//
// Solidity: function updateValset(address[] _validators, uint256[] _powers) returns()
func (_MockUpgrade *MockUpgradeTransactor) UpdateValset(opts *bind.TransactOpts, _validators []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MockUpgrade.contract.Transact(opts, "updateValset", _validators, _powers)
}

// UpdateValset is a paid mutator transaction binding the contract method 0x788cf92f.
//
// Solidity: function updateValset(address[] _validators, uint256[] _powers) returns()
func (_MockUpgrade *MockUpgradeSession) UpdateValset(_validators []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MockUpgrade.Contract.UpdateValset(&_MockUpgrade.TransactOpts, _validators, _powers)
}

// UpdateValset is a paid mutator transaction binding the contract method 0x788cf92f.
//
// Solidity: function updateValset(address[] _validators, uint256[] _powers) returns()
func (_MockUpgrade *MockUpgradeTransactorSession) UpdateValset(_validators []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MockUpgrade.Contract.UpdateValset(&_MockUpgrade.TransactOpts, _validators, _powers)
}

// MockUpgradeLogBridgeBankSetIterator is returned from FilterLogBridgeBankSet and is used to iterate over the raw logs and unpacked data for LogBridgeBankSet events raised by the MockUpgrade contract.
type MockUpgradeLogBridgeBankSetIterator struct {
	Event *MockUpgradeLogBridgeBankSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MockUpgradeLogBridgeBankSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockUpgradeLogBridgeBankSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MockUpgradeLogBridgeBankSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MockUpgradeLogBridgeBankSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockUpgradeLogBridgeBankSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockUpgradeLogBridgeBankSet represents a LogBridgeBankSet event raised by the MockUpgrade contract.
type MockUpgradeLogBridgeBankSet struct {
	BridgeBank common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogBridgeBankSet is a free log retrieval operation binding the contract event 0xc8b65043fb196ac032b79a435397d1d14a96b4e9d12e366c3b1f550cb01d2dfa.
//
// Solidity: event LogBridgeBankSet(address bridgeBank)
func (_MockUpgrade *MockUpgradeFilterer) FilterLogBridgeBankSet(opts *bind.FilterOpts) (*MockUpgradeLogBridgeBankSetIterator, error) {

	logs, sub, err := _MockUpgrade.contract.FilterLogs(opts, "LogBridgeBankSet")
	if err != nil {
		return nil, err
	}
	return &MockUpgradeLogBridgeBankSetIterator{contract: _MockUpgrade.contract, event: "LogBridgeBankSet", logs: logs, sub: sub}, nil
}

// WatchLogBridgeBankSet is a free log subscription operation binding the contract event 0xc8b65043fb196ac032b79a435397d1d14a96b4e9d12e366c3b1f550cb01d2dfa.
//
// Solidity: event LogBridgeBankSet(address bridgeBank)
func (_MockUpgrade *MockUpgradeFilterer) WatchLogBridgeBankSet(opts *bind.WatchOpts, sink chan<- *MockUpgradeLogBridgeBankSet) (event.Subscription, error) {

	logs, sub, err := _MockUpgrade.contract.WatchLogs(opts, "LogBridgeBankSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockUpgradeLogBridgeBankSet)
				if err := _MockUpgrade.contract.UnpackLog(event, "LogBridgeBankSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogBridgeBankSet is a log parse operation binding the contract event 0xc8b65043fb196ac032b79a435397d1d14a96b4e9d12e366c3b1f550cb01d2dfa.
//
// Solidity: event LogBridgeBankSet(address bridgeBank)
func (_MockUpgrade *MockUpgradeFilterer) ParseLogBridgeBankSet(log types.Log) (*MockUpgradeLogBridgeBankSet, error) {
	event := new(MockUpgradeLogBridgeBankSet)
	if err := _MockUpgrade.contract.UnpackLog(event, "LogBridgeBankSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockUpgradeLogNewBridgeTokenCreatedIterator is returned from FilterLogNewBridgeTokenCreated and is used to iterate over the raw logs and unpacked data for LogNewBridgeTokenCreated events raised by the MockUpgrade contract.
type MockUpgradeLogNewBridgeTokenCreatedIterator struct {
	Event *MockUpgradeLogNewBridgeTokenCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MockUpgradeLogNewBridgeTokenCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockUpgradeLogNewBridgeTokenCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MockUpgradeLogNewBridgeTokenCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MockUpgradeLogNewBridgeTokenCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockUpgradeLogNewBridgeTokenCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockUpgradeLogNewBridgeTokenCreated represents a LogNewBridgeTokenCreated event raised by the MockUpgrade contract.
type MockUpgradeLogNewBridgeTokenCreated struct {
	Decimals              uint8
	NetworkDescriptor     int32
	Name                  string
	Symbol                string
	SourceContractAddress common.Address
	BridgeTokenAddress    common.Address
	CosmosDenom           string
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterLogNewBridgeTokenCreated is a free log retrieval operation binding the contract event 0xc6d838144c81411c2f7015ec7a5e7552549d723d48a5cdc1b349c0a7749a1ca3.
//
// Solidity: event LogNewBridgeTokenCreated(uint8 decimals, int32 indexed _networkDescriptor, string name, string symbol, address indexed sourceContractAddress, address indexed bridgeTokenAddress, string cosmosDenom)
func (_MockUpgrade *MockUpgradeFilterer) FilterLogNewBridgeTokenCreated(opts *bind.FilterOpts, _networkDescriptor []int32, sourceContractAddress []common.Address, bridgeTokenAddress []common.Address) (*MockUpgradeLogNewBridgeTokenCreatedIterator, error) {

	var _networkDescriptorRule []interface{}
	for _, _networkDescriptorItem := range _networkDescriptor {
		_networkDescriptorRule = append(_networkDescriptorRule, _networkDescriptorItem)
	}

	var sourceContractAddressRule []interface{}
	for _, sourceContractAddressItem := range sourceContractAddress {
		sourceContractAddressRule = append(sourceContractAddressRule, sourceContractAddressItem)
	}
	var bridgeTokenAddressRule []interface{}
	for _, bridgeTokenAddressItem := range bridgeTokenAddress {
		bridgeTokenAddressRule = append(bridgeTokenAddressRule, bridgeTokenAddressItem)
	}

	logs, sub, err := _MockUpgrade.contract.FilterLogs(opts, "LogNewBridgeTokenCreated", _networkDescriptorRule, sourceContractAddressRule, bridgeTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &MockUpgradeLogNewBridgeTokenCreatedIterator{contract: _MockUpgrade.contract, event: "LogNewBridgeTokenCreated", logs: logs, sub: sub}, nil
}

// WatchLogNewBridgeTokenCreated is a free log subscription operation binding the contract event 0xc6d838144c81411c2f7015ec7a5e7552549d723d48a5cdc1b349c0a7749a1ca3.
//
// Solidity: event LogNewBridgeTokenCreated(uint8 decimals, int32 indexed _networkDescriptor, string name, string symbol, address indexed sourceContractAddress, address indexed bridgeTokenAddress, string cosmosDenom)
func (_MockUpgrade *MockUpgradeFilterer) WatchLogNewBridgeTokenCreated(opts *bind.WatchOpts, sink chan<- *MockUpgradeLogNewBridgeTokenCreated, _networkDescriptor []int32, sourceContractAddress []common.Address, bridgeTokenAddress []common.Address) (event.Subscription, error) {

	var _networkDescriptorRule []interface{}
	for _, _networkDescriptorItem := range _networkDescriptor {
		_networkDescriptorRule = append(_networkDescriptorRule, _networkDescriptorItem)
	}

	var sourceContractAddressRule []interface{}
	for _, sourceContractAddressItem := range sourceContractAddress {
		sourceContractAddressRule = append(sourceContractAddressRule, sourceContractAddressItem)
	}
	var bridgeTokenAddressRule []interface{}
	for _, bridgeTokenAddressItem := range bridgeTokenAddress {
		bridgeTokenAddressRule = append(bridgeTokenAddressRule, bridgeTokenAddressItem)
	}

	logs, sub, err := _MockUpgrade.contract.WatchLogs(opts, "LogNewBridgeTokenCreated", _networkDescriptorRule, sourceContractAddressRule, bridgeTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockUpgradeLogNewBridgeTokenCreated)
				if err := _MockUpgrade.contract.UnpackLog(event, "LogNewBridgeTokenCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNewBridgeTokenCreated is a log parse operation binding the contract event 0xc6d838144c81411c2f7015ec7a5e7552549d723d48a5cdc1b349c0a7749a1ca3.
//
// Solidity: event LogNewBridgeTokenCreated(uint8 decimals, int32 indexed _networkDescriptor, string name, string symbol, address indexed sourceContractAddress, address indexed bridgeTokenAddress, string cosmosDenom)
func (_MockUpgrade *MockUpgradeFilterer) ParseLogNewBridgeTokenCreated(log types.Log) (*MockUpgradeLogNewBridgeTokenCreated, error) {
	event := new(MockUpgradeLogNewBridgeTokenCreated)
	if err := _MockUpgrade.contract.UnpackLog(event, "LogNewBridgeTokenCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockUpgradeLogNewOracleClaimIterator is returned from FilterLogNewOracleClaim and is used to iterate over the raw logs and unpacked data for LogNewOracleClaim events raised by the MockUpgrade contract.
type MockUpgradeLogNewOracleClaimIterator struct {
	Event *MockUpgradeLogNewOracleClaim // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MockUpgradeLogNewOracleClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockUpgradeLogNewOracleClaim)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MockUpgradeLogNewOracleClaim)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MockUpgradeLogNewOracleClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockUpgradeLogNewOracleClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockUpgradeLogNewOracleClaim represents a LogNewOracleClaim event raised by the MockUpgrade contract.
type MockUpgradeLogNewOracleClaim struct {
	ProphecyID       *big.Int
	ValidatorAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogNewOracleClaim is a free log retrieval operation binding the contract event 0x668fce9833323940537a9000d512a6c580a1c0797d2b526db0078ee9c5a087a9.
//
// Solidity: event LogNewOracleClaim(uint256 _prophecyID, address _validatorAddress)
func (_MockUpgrade *MockUpgradeFilterer) FilterLogNewOracleClaim(opts *bind.FilterOpts) (*MockUpgradeLogNewOracleClaimIterator, error) {

	logs, sub, err := _MockUpgrade.contract.FilterLogs(opts, "LogNewOracleClaim")
	if err != nil {
		return nil, err
	}
	return &MockUpgradeLogNewOracleClaimIterator{contract: _MockUpgrade.contract, event: "LogNewOracleClaim", logs: logs, sub: sub}, nil
}

// WatchLogNewOracleClaim is a free log subscription operation binding the contract event 0x668fce9833323940537a9000d512a6c580a1c0797d2b526db0078ee9c5a087a9.
//
// Solidity: event LogNewOracleClaim(uint256 _prophecyID, address _validatorAddress)
func (_MockUpgrade *MockUpgradeFilterer) WatchLogNewOracleClaim(opts *bind.WatchOpts, sink chan<- *MockUpgradeLogNewOracleClaim) (event.Subscription, error) {

	logs, sub, err := _MockUpgrade.contract.WatchLogs(opts, "LogNewOracleClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockUpgradeLogNewOracleClaim)
				if err := _MockUpgrade.contract.UnpackLog(event, "LogNewOracleClaim", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNewOracleClaim is a log parse operation binding the contract event 0x668fce9833323940537a9000d512a6c580a1c0797d2b526db0078ee9c5a087a9.
//
// Solidity: event LogNewOracleClaim(uint256 _prophecyID, address _validatorAddress)
func (_MockUpgrade *MockUpgradeFilterer) ParseLogNewOracleClaim(log types.Log) (*MockUpgradeLogNewOracleClaim, error) {
	event := new(MockUpgradeLogNewOracleClaim)
	if err := _MockUpgrade.contract.UnpackLog(event, "LogNewOracleClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockUpgradeLogNewProphecyClaimIterator is returned from FilterLogNewProphecyClaim and is used to iterate over the raw logs and unpacked data for LogNewProphecyClaim events raised by the MockUpgrade contract.
type MockUpgradeLogNewProphecyClaimIterator struct {
	Event *MockUpgradeLogNewProphecyClaim // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MockUpgradeLogNewProphecyClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockUpgradeLogNewProphecyClaim)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MockUpgradeLogNewProphecyClaim)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MockUpgradeLogNewProphecyClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockUpgradeLogNewProphecyClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockUpgradeLogNewProphecyClaim represents a LogNewProphecyClaim event raised by the MockUpgrade contract.
type MockUpgradeLogNewProphecyClaim struct {
	ProphecyID       *big.Int
	EthereumReceiver common.Address
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogNewProphecyClaim is a free log retrieval operation binding the contract event 0x392e2c08a6c5b92e1d4775a1e43829652f156f1310bd3fe2935f586e2d4ce36e.
//
// Solidity: event LogNewProphecyClaim(uint256 indexed prophecyID, address indexed ethereumReceiver, uint256 indexed amount)
func (_MockUpgrade *MockUpgradeFilterer) FilterLogNewProphecyClaim(opts *bind.FilterOpts, prophecyID []*big.Int, ethereumReceiver []common.Address, amount []*big.Int) (*MockUpgradeLogNewProphecyClaimIterator, error) {

	var prophecyIDRule []interface{}
	for _, prophecyIDItem := range prophecyID {
		prophecyIDRule = append(prophecyIDRule, prophecyIDItem)
	}
	var ethereumReceiverRule []interface{}
	for _, ethereumReceiverItem := range ethereumReceiver {
		ethereumReceiverRule = append(ethereumReceiverRule, ethereumReceiverItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _MockUpgrade.contract.FilterLogs(opts, "LogNewProphecyClaim", prophecyIDRule, ethereumReceiverRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &MockUpgradeLogNewProphecyClaimIterator{contract: _MockUpgrade.contract, event: "LogNewProphecyClaim", logs: logs, sub: sub}, nil
}

// WatchLogNewProphecyClaim is a free log subscription operation binding the contract event 0x392e2c08a6c5b92e1d4775a1e43829652f156f1310bd3fe2935f586e2d4ce36e.
//
// Solidity: event LogNewProphecyClaim(uint256 indexed prophecyID, address indexed ethereumReceiver, uint256 indexed amount)
func (_MockUpgrade *MockUpgradeFilterer) WatchLogNewProphecyClaim(opts *bind.WatchOpts, sink chan<- *MockUpgradeLogNewProphecyClaim, prophecyID []*big.Int, ethereumReceiver []common.Address, amount []*big.Int) (event.Subscription, error) {

	var prophecyIDRule []interface{}
	for _, prophecyIDItem := range prophecyID {
		prophecyIDRule = append(prophecyIDRule, prophecyIDItem)
	}
	var ethereumReceiverRule []interface{}
	for _, ethereumReceiverItem := range ethereumReceiver {
		ethereumReceiverRule = append(ethereumReceiverRule, ethereumReceiverItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _MockUpgrade.contract.WatchLogs(opts, "LogNewProphecyClaim", prophecyIDRule, ethereumReceiverRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockUpgradeLogNewProphecyClaim)
				if err := _MockUpgrade.contract.UnpackLog(event, "LogNewProphecyClaim", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNewProphecyClaim is a log parse operation binding the contract event 0x392e2c08a6c5b92e1d4775a1e43829652f156f1310bd3fe2935f586e2d4ce36e.
//
// Solidity: event LogNewProphecyClaim(uint256 indexed prophecyID, address indexed ethereumReceiver, uint256 indexed amount)
func (_MockUpgrade *MockUpgradeFilterer) ParseLogNewProphecyClaim(log types.Log) (*MockUpgradeLogNewProphecyClaim, error) {
	event := new(MockUpgradeLogNewProphecyClaim)
	if err := _MockUpgrade.contract.UnpackLog(event, "LogNewProphecyClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockUpgradeLogProphecyCompletedIterator is returned from FilterLogProphecyCompleted and is used to iterate over the raw logs and unpacked data for LogProphecyCompleted events raised by the MockUpgrade contract.
type MockUpgradeLogProphecyCompletedIterator struct {
	Event *MockUpgradeLogProphecyCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MockUpgradeLogProphecyCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockUpgradeLogProphecyCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MockUpgradeLogProphecyCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MockUpgradeLogProphecyCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockUpgradeLogProphecyCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockUpgradeLogProphecyCompleted represents a LogProphecyCompleted event raised by the MockUpgrade contract.
type MockUpgradeLogProphecyCompleted struct {
	ProphecyID *big.Int
	Success    bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogProphecyCompleted is a free log retrieval operation binding the contract event 0x6cf2aa1395dae0a852879973e322e6d1f5a48485a3a41c28a2fd8f82d3aae487.
//
// Solidity: event LogProphecyCompleted(uint256 prophecyID, bool success)
func (_MockUpgrade *MockUpgradeFilterer) FilterLogProphecyCompleted(opts *bind.FilterOpts) (*MockUpgradeLogProphecyCompletedIterator, error) {

	logs, sub, err := _MockUpgrade.contract.FilterLogs(opts, "LogProphecyCompleted")
	if err != nil {
		return nil, err
	}
	return &MockUpgradeLogProphecyCompletedIterator{contract: _MockUpgrade.contract, event: "LogProphecyCompleted", logs: logs, sub: sub}, nil
}

// WatchLogProphecyCompleted is a free log subscription operation binding the contract event 0x6cf2aa1395dae0a852879973e322e6d1f5a48485a3a41c28a2fd8f82d3aae487.
//
// Solidity: event LogProphecyCompleted(uint256 prophecyID, bool success)
func (_MockUpgrade *MockUpgradeFilterer) WatchLogProphecyCompleted(opts *bind.WatchOpts, sink chan<- *MockUpgradeLogProphecyCompleted) (event.Subscription, error) {

	logs, sub, err := _MockUpgrade.contract.WatchLogs(opts, "LogProphecyCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockUpgradeLogProphecyCompleted)
				if err := _MockUpgrade.contract.UnpackLog(event, "LogProphecyCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogProphecyCompleted is a log parse operation binding the contract event 0x6cf2aa1395dae0a852879973e322e6d1f5a48485a3a41c28a2fd8f82d3aae487.
//
// Solidity: event LogProphecyCompleted(uint256 prophecyID, bool success)
func (_MockUpgrade *MockUpgradeFilterer) ParseLogProphecyCompleted(log types.Log) (*MockUpgradeLogProphecyCompleted, error) {
	event := new(MockUpgradeLogProphecyCompleted)
	if err := _MockUpgrade.contract.UnpackLog(event, "LogProphecyCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockUpgradeLogProphecyProcessedIterator is returned from FilterLogProphecyProcessed and is used to iterate over the raw logs and unpacked data for LogProphecyProcessed events raised by the MockUpgrade contract.
type MockUpgradeLogProphecyProcessedIterator struct {
	Event *MockUpgradeLogProphecyProcessed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MockUpgradeLogProphecyProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockUpgradeLogProphecyProcessed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MockUpgradeLogProphecyProcessed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MockUpgradeLogProphecyProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockUpgradeLogProphecyProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockUpgradeLogProphecyProcessed represents a LogProphecyProcessed event raised by the MockUpgrade contract.
type MockUpgradeLogProphecyProcessed struct {
	ProphecyID             *big.Int
	ProphecyPowerCurrent   *big.Int
	ProphecyPowerThreshold *big.Int
	Submitter              common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterLogProphecyProcessed is a free log retrieval operation binding the contract event 0x1d8e3fbd601d9d92db7022fb97f75e132841b94db732dcecb0c93cb31852fcbc.
//
// Solidity: event LogProphecyProcessed(uint256 _prophecyID, uint256 _prophecyPowerCurrent, uint256 _prophecyPowerThreshold, address _submitter)
func (_MockUpgrade *MockUpgradeFilterer) FilterLogProphecyProcessed(opts *bind.FilterOpts) (*MockUpgradeLogProphecyProcessedIterator, error) {

	logs, sub, err := _MockUpgrade.contract.FilterLogs(opts, "LogProphecyProcessed")
	if err != nil {
		return nil, err
	}
	return &MockUpgradeLogProphecyProcessedIterator{contract: _MockUpgrade.contract, event: "LogProphecyProcessed", logs: logs, sub: sub}, nil
}

// WatchLogProphecyProcessed is a free log subscription operation binding the contract event 0x1d8e3fbd601d9d92db7022fb97f75e132841b94db732dcecb0c93cb31852fcbc.
//
// Solidity: event LogProphecyProcessed(uint256 _prophecyID, uint256 _prophecyPowerCurrent, uint256 _prophecyPowerThreshold, address _submitter)
func (_MockUpgrade *MockUpgradeFilterer) WatchLogProphecyProcessed(opts *bind.WatchOpts, sink chan<- *MockUpgradeLogProphecyProcessed) (event.Subscription, error) {

	logs, sub, err := _MockUpgrade.contract.WatchLogs(opts, "LogProphecyProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockUpgradeLogProphecyProcessed)
				if err := _MockUpgrade.contract.UnpackLog(event, "LogProphecyProcessed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogProphecyProcessed is a log parse operation binding the contract event 0x1d8e3fbd601d9d92db7022fb97f75e132841b94db732dcecb0c93cb31852fcbc.
//
// Solidity: event LogProphecyProcessed(uint256 _prophecyID, uint256 _prophecyPowerCurrent, uint256 _prophecyPowerThreshold, address _submitter)
func (_MockUpgrade *MockUpgradeFilterer) ParseLogProphecyProcessed(log types.Log) (*MockUpgradeLogProphecyProcessed, error) {
	event := new(MockUpgradeLogProphecyProcessed)
	if err := _MockUpgrade.contract.UnpackLog(event, "LogProphecyProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockUpgradeLogValidatorAddedIterator is returned from FilterLogValidatorAdded and is used to iterate over the raw logs and unpacked data for LogValidatorAdded events raised by the MockUpgrade contract.
type MockUpgradeLogValidatorAddedIterator struct {
	Event *MockUpgradeLogValidatorAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MockUpgradeLogValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockUpgradeLogValidatorAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MockUpgradeLogValidatorAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MockUpgradeLogValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockUpgradeLogValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockUpgradeLogValidatorAdded represents a LogValidatorAdded event raised by the MockUpgrade contract.
type MockUpgradeLogValidatorAdded struct {
	Validator            common.Address
	Power                *big.Int
	CurrentValsetVersion *big.Int
	ValidatorCount       *big.Int
	TotalPower           *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterLogValidatorAdded is a free log retrieval operation binding the contract event 0x1a396bcf647502e902dce665d58a0c1b25f982f193ab9a1d0f1500d8d927bf2a.
//
// Solidity: event LogValidatorAdded(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_MockUpgrade *MockUpgradeFilterer) FilterLogValidatorAdded(opts *bind.FilterOpts) (*MockUpgradeLogValidatorAddedIterator, error) {

	logs, sub, err := _MockUpgrade.contract.FilterLogs(opts, "LogValidatorAdded")
	if err != nil {
		return nil, err
	}
	return &MockUpgradeLogValidatorAddedIterator{contract: _MockUpgrade.contract, event: "LogValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchLogValidatorAdded is a free log subscription operation binding the contract event 0x1a396bcf647502e902dce665d58a0c1b25f982f193ab9a1d0f1500d8d927bf2a.
//
// Solidity: event LogValidatorAdded(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_MockUpgrade *MockUpgradeFilterer) WatchLogValidatorAdded(opts *bind.WatchOpts, sink chan<- *MockUpgradeLogValidatorAdded) (event.Subscription, error) {

	logs, sub, err := _MockUpgrade.contract.WatchLogs(opts, "LogValidatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockUpgradeLogValidatorAdded)
				if err := _MockUpgrade.contract.UnpackLog(event, "LogValidatorAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogValidatorAdded is a log parse operation binding the contract event 0x1a396bcf647502e902dce665d58a0c1b25f982f193ab9a1d0f1500d8d927bf2a.
//
// Solidity: event LogValidatorAdded(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_MockUpgrade *MockUpgradeFilterer) ParseLogValidatorAdded(log types.Log) (*MockUpgradeLogValidatorAdded, error) {
	event := new(MockUpgradeLogValidatorAdded)
	if err := _MockUpgrade.contract.UnpackLog(event, "LogValidatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockUpgradeLogValidatorPowerUpdatedIterator is returned from FilterLogValidatorPowerUpdated and is used to iterate over the raw logs and unpacked data for LogValidatorPowerUpdated events raised by the MockUpgrade contract.
type MockUpgradeLogValidatorPowerUpdatedIterator struct {
	Event *MockUpgradeLogValidatorPowerUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MockUpgradeLogValidatorPowerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockUpgradeLogValidatorPowerUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MockUpgradeLogValidatorPowerUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MockUpgradeLogValidatorPowerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockUpgradeLogValidatorPowerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockUpgradeLogValidatorPowerUpdated represents a LogValidatorPowerUpdated event raised by the MockUpgrade contract.
type MockUpgradeLogValidatorPowerUpdated struct {
	Validator            common.Address
	Power                *big.Int
	CurrentValsetVersion *big.Int
	ValidatorCount       *big.Int
	TotalPower           *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterLogValidatorPowerUpdated is a free log retrieval operation binding the contract event 0x335940ce4119f8aae891d73dba74510a3d51f6210134d058237f26e6a31d5340.
//
// Solidity: event LogValidatorPowerUpdated(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_MockUpgrade *MockUpgradeFilterer) FilterLogValidatorPowerUpdated(opts *bind.FilterOpts) (*MockUpgradeLogValidatorPowerUpdatedIterator, error) {

	logs, sub, err := _MockUpgrade.contract.FilterLogs(opts, "LogValidatorPowerUpdated")
	if err != nil {
		return nil, err
	}
	return &MockUpgradeLogValidatorPowerUpdatedIterator{contract: _MockUpgrade.contract, event: "LogValidatorPowerUpdated", logs: logs, sub: sub}, nil
}

// WatchLogValidatorPowerUpdated is a free log subscription operation binding the contract event 0x335940ce4119f8aae891d73dba74510a3d51f6210134d058237f26e6a31d5340.
//
// Solidity: event LogValidatorPowerUpdated(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_MockUpgrade *MockUpgradeFilterer) WatchLogValidatorPowerUpdated(opts *bind.WatchOpts, sink chan<- *MockUpgradeLogValidatorPowerUpdated) (event.Subscription, error) {

	logs, sub, err := _MockUpgrade.contract.WatchLogs(opts, "LogValidatorPowerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockUpgradeLogValidatorPowerUpdated)
				if err := _MockUpgrade.contract.UnpackLog(event, "LogValidatorPowerUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogValidatorPowerUpdated is a log parse operation binding the contract event 0x335940ce4119f8aae891d73dba74510a3d51f6210134d058237f26e6a31d5340.
//
// Solidity: event LogValidatorPowerUpdated(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_MockUpgrade *MockUpgradeFilterer) ParseLogValidatorPowerUpdated(log types.Log) (*MockUpgradeLogValidatorPowerUpdated, error) {
	event := new(MockUpgradeLogValidatorPowerUpdated)
	if err := _MockUpgrade.contract.UnpackLog(event, "LogValidatorPowerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockUpgradeLogValidatorRemovedIterator is returned from FilterLogValidatorRemoved and is used to iterate over the raw logs and unpacked data for LogValidatorRemoved events raised by the MockUpgrade contract.
type MockUpgradeLogValidatorRemovedIterator struct {
	Event *MockUpgradeLogValidatorRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MockUpgradeLogValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockUpgradeLogValidatorRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MockUpgradeLogValidatorRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MockUpgradeLogValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockUpgradeLogValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockUpgradeLogValidatorRemoved represents a LogValidatorRemoved event raised by the MockUpgrade contract.
type MockUpgradeLogValidatorRemoved struct {
	Validator            common.Address
	Power                *big.Int
	CurrentValsetVersion *big.Int
	ValidatorCount       *big.Int
	TotalPower           *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterLogValidatorRemoved is a free log retrieval operation binding the contract event 0x1241fb43a101ff98ab819a1882097d4ccada51ba60f326c1981cc48840f55b8c.
//
// Solidity: event LogValidatorRemoved(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_MockUpgrade *MockUpgradeFilterer) FilterLogValidatorRemoved(opts *bind.FilterOpts) (*MockUpgradeLogValidatorRemovedIterator, error) {

	logs, sub, err := _MockUpgrade.contract.FilterLogs(opts, "LogValidatorRemoved")
	if err != nil {
		return nil, err
	}
	return &MockUpgradeLogValidatorRemovedIterator{contract: _MockUpgrade.contract, event: "LogValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchLogValidatorRemoved is a free log subscription operation binding the contract event 0x1241fb43a101ff98ab819a1882097d4ccada51ba60f326c1981cc48840f55b8c.
//
// Solidity: event LogValidatorRemoved(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_MockUpgrade *MockUpgradeFilterer) WatchLogValidatorRemoved(opts *bind.WatchOpts, sink chan<- *MockUpgradeLogValidatorRemoved) (event.Subscription, error) {

	logs, sub, err := _MockUpgrade.contract.WatchLogs(opts, "LogValidatorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockUpgradeLogValidatorRemoved)
				if err := _MockUpgrade.contract.UnpackLog(event, "LogValidatorRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogValidatorRemoved is a log parse operation binding the contract event 0x1241fb43a101ff98ab819a1882097d4ccada51ba60f326c1981cc48840f55b8c.
//
// Solidity: event LogValidatorRemoved(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_MockUpgrade *MockUpgradeFilterer) ParseLogValidatorRemoved(log types.Log) (*MockUpgradeLogValidatorRemoved, error) {
	event := new(MockUpgradeLogValidatorRemoved)
	if err := _MockUpgrade.contract.UnpackLog(event, "LogValidatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockUpgradeLogValsetResetIterator is returned from FilterLogValsetReset and is used to iterate over the raw logs and unpacked data for LogValsetReset events raised by the MockUpgrade contract.
type MockUpgradeLogValsetResetIterator struct {
	Event *MockUpgradeLogValsetReset // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MockUpgradeLogValsetResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockUpgradeLogValsetReset)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MockUpgradeLogValsetReset)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MockUpgradeLogValsetResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockUpgradeLogValsetResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockUpgradeLogValsetReset represents a LogValsetReset event raised by the MockUpgrade contract.
type MockUpgradeLogValsetReset struct {
	NewValsetVersion *big.Int
	ValidatorCount   *big.Int
	TotalPower       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogValsetReset is a free log retrieval operation binding the contract event 0xd870653e19f161500290fd0c4ca41bf5cf2bcb1ba66448f41c66c512dabd65f2.
//
// Solidity: event LogValsetReset(uint256 _newValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_MockUpgrade *MockUpgradeFilterer) FilterLogValsetReset(opts *bind.FilterOpts) (*MockUpgradeLogValsetResetIterator, error) {

	logs, sub, err := _MockUpgrade.contract.FilterLogs(opts, "LogValsetReset")
	if err != nil {
		return nil, err
	}
	return &MockUpgradeLogValsetResetIterator{contract: _MockUpgrade.contract, event: "LogValsetReset", logs: logs, sub: sub}, nil
}

// WatchLogValsetReset is a free log subscription operation binding the contract event 0xd870653e19f161500290fd0c4ca41bf5cf2bcb1ba66448f41c66c512dabd65f2.
//
// Solidity: event LogValsetReset(uint256 _newValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_MockUpgrade *MockUpgradeFilterer) WatchLogValsetReset(opts *bind.WatchOpts, sink chan<- *MockUpgradeLogValsetReset) (event.Subscription, error) {

	logs, sub, err := _MockUpgrade.contract.WatchLogs(opts, "LogValsetReset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockUpgradeLogValsetReset)
				if err := _MockUpgrade.contract.UnpackLog(event, "LogValsetReset", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogValsetReset is a log parse operation binding the contract event 0xd870653e19f161500290fd0c4ca41bf5cf2bcb1ba66448f41c66c512dabd65f2.
//
// Solidity: event LogValsetReset(uint256 _newValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_MockUpgrade *MockUpgradeFilterer) ParseLogValsetReset(log types.Log) (*MockUpgradeLogValsetReset, error) {
	event := new(MockUpgradeLogValsetReset)
	if err := _MockUpgrade.contract.UnpackLog(event, "LogValsetReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockUpgradeLogValsetUpdatedIterator is returned from FilterLogValsetUpdated and is used to iterate over the raw logs and unpacked data for LogValsetUpdated events raised by the MockUpgrade contract.
type MockUpgradeLogValsetUpdatedIterator struct {
	Event *MockUpgradeLogValsetUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MockUpgradeLogValsetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockUpgradeLogValsetUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MockUpgradeLogValsetUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MockUpgradeLogValsetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockUpgradeLogValsetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockUpgradeLogValsetUpdated represents a LogValsetUpdated event raised by the MockUpgrade contract.
type MockUpgradeLogValsetUpdated struct {
	NewValsetVersion *big.Int
	ValidatorCount   *big.Int
	TotalPower       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogValsetUpdated is a free log retrieval operation binding the contract event 0x3a7ef0da3179668af8114719645585b5a37092ef2d66f187dcf63d83a221eaa6.
//
// Solidity: event LogValsetUpdated(uint256 _newValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_MockUpgrade *MockUpgradeFilterer) FilterLogValsetUpdated(opts *bind.FilterOpts) (*MockUpgradeLogValsetUpdatedIterator, error) {

	logs, sub, err := _MockUpgrade.contract.FilterLogs(opts, "LogValsetUpdated")
	if err != nil {
		return nil, err
	}
	return &MockUpgradeLogValsetUpdatedIterator{contract: _MockUpgrade.contract, event: "LogValsetUpdated", logs: logs, sub: sub}, nil
}

// WatchLogValsetUpdated is a free log subscription operation binding the contract event 0x3a7ef0da3179668af8114719645585b5a37092ef2d66f187dcf63d83a221eaa6.
//
// Solidity: event LogValsetUpdated(uint256 _newValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_MockUpgrade *MockUpgradeFilterer) WatchLogValsetUpdated(opts *bind.WatchOpts, sink chan<- *MockUpgradeLogValsetUpdated) (event.Subscription, error) {

	logs, sub, err := _MockUpgrade.contract.WatchLogs(opts, "LogValsetUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockUpgradeLogValsetUpdated)
				if err := _MockUpgrade.contract.UnpackLog(event, "LogValsetUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogValsetUpdated is a log parse operation binding the contract event 0x3a7ef0da3179668af8114719645585b5a37092ef2d66f187dcf63d83a221eaa6.
//
// Solidity: event LogValsetUpdated(uint256 _newValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_MockUpgrade *MockUpgradeFilterer) ParseLogValsetUpdated(log types.Log) (*MockUpgradeLogValsetUpdated, error) {
	event := new(MockUpgradeLogValsetUpdated)
	if err := _MockUpgrade.contract.UnpackLog(event, "LogValsetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
