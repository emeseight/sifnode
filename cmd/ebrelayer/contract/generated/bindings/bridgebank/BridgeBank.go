// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BridgeBank

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

// BridgeBankMetaData contains all meta data concerning the BridgeBank contract.
var BridgeBankMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"LogBridgeTokenMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_to\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"int32\",\"name\":\"_networkDescriptor\",\"type\":\"int32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_denom\",\"type\":\"string\"}],\"name\":\"LogBurn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_to\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int32\",\"name\":\"_networkDescriptor\",\"type\":\"int32\"}],\"name\":\"LogLock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_cosmosDenom\",\"type\":\"string\"}],\"name\":\"LogNewBridgeToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"LogUnlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_value\",\"type\":\"bool\"}],\"name\":\"LogWhiteListUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_sifAddress\",\"type\":\"bytes\"}],\"name\":\"VSA\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"addExistingBridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"contractsAddresses\",\"type\":\"address[]\"}],\"name\":\"batchAddExistingBridgeTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"batchForceSetBridgeTokenDenom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"_denoms\",\"type\":\"string[]\"}],\"name\":\"batchSetBridgeTokenDenom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blocklist\",\"outputs\":[{\"internalType\":\"contractIBlocklist\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOperator\",\"type\":\"address\"}],\"name\":\"changeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"contractDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"contractDenom\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"contractName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"contractSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cosmosBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"cosmosDenom\",\"type\":\"string\"}],\"name\":\"createNewBridgeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"forceSetBridgeTokenDenom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getCosmosTokenInWhiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"ethereumReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"handleUnpeg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasBlocklist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cosmosBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pauser\",\"type\":\"address\"},{\"internalType\":\"int32\",\"name\":\"_networkDescriptor\",\"type\":\"int32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockBurnNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"recipient\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"name\":\"multiLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"recipient\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"isBurn\",\"type\":\"bool[]\"}],\"name\":\"multiLockBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkDescriptor\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pausers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cosmosBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pauser\",\"type\":\"address\"},{\"internalType\":\"int32\",\"name\":\"_networkDescriptor\",\"type\":\"int32\"}],\"name\":\"reinitialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"blocklistAddress\",\"type\":\"address\"}],\"name\":\"setBlocklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_denom\",\"type\":\"string\"}],\"name\":\"setBridgeTokenDenom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BridgeBankABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeBankMetaData.ABI instead.
var BridgeBankABI = BridgeBankMetaData.ABI

// BridgeBank is an auto generated Go binding around an Ethereum contract.
type BridgeBank struct {
	BridgeBankCaller     // Read-only binding to the contract
	BridgeBankTransactor // Write-only binding to the contract
	BridgeBankFilterer   // Log filterer for contract events
}

// BridgeBankCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeBankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeBankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeBankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeBankSession struct {
	Contract     *BridgeBank       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeBankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeBankCallerSession struct {
	Contract *BridgeBankCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BridgeBankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeBankTransactorSession struct {
	Contract     *BridgeBankTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BridgeBankRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeBankRaw struct {
	Contract *BridgeBank // Generic contract binding to access the raw methods on
}

// BridgeBankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeBankCallerRaw struct {
	Contract *BridgeBankCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeBankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeBankTransactorRaw struct {
	Contract *BridgeBankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeBank creates a new instance of BridgeBank, bound to a specific deployed contract.
func NewBridgeBank(address common.Address, backend bind.ContractBackend) (*BridgeBank, error) {
	contract, err := bindBridgeBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeBank{BridgeBankCaller: BridgeBankCaller{contract: contract}, BridgeBankTransactor: BridgeBankTransactor{contract: contract}, BridgeBankFilterer: BridgeBankFilterer{contract: contract}}, nil
}

// NewBridgeBankCaller creates a new read-only instance of BridgeBank, bound to a specific deployed contract.
func NewBridgeBankCaller(address common.Address, caller bind.ContractCaller) (*BridgeBankCaller, error) {
	contract, err := bindBridgeBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeBankCaller{contract: contract}, nil
}

// NewBridgeBankTransactor creates a new write-only instance of BridgeBank, bound to a specific deployed contract.
func NewBridgeBankTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeBankTransactor, error) {
	contract, err := bindBridgeBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeBankTransactor{contract: contract}, nil
}

// NewBridgeBankFilterer creates a new log filterer instance of BridgeBank, bound to a specific deployed contract.
func NewBridgeBankFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeBankFilterer, error) {
	contract, err := bindBridgeBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeBankFilterer{contract: contract}, nil
}

// bindBridgeBank binds a generic wrapper to an already deployed contract.
func bindBridgeBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeBankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeBank *BridgeBankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeBank.Contract.BridgeBankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeBank *BridgeBankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBank.Contract.BridgeBankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeBank *BridgeBankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeBank.Contract.BridgeBankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeBank *BridgeBankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeBank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeBank *BridgeBankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeBank *BridgeBankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeBank.Contract.contract.Transact(opts, method, params...)
}

// VSA is a free data retrieval call binding the contract method 0xc228979d.
//
// Solidity: function VSA(bytes _sifAddress) pure returns(bool)
func (_BridgeBank *BridgeBankCaller) VSA(opts *bind.CallOpts, _sifAddress []byte) (bool, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "VSA", _sifAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VSA is a free data retrieval call binding the contract method 0xc228979d.
//
// Solidity: function VSA(bytes _sifAddress) pure returns(bool)
func (_BridgeBank *BridgeBankSession) VSA(_sifAddress []byte) (bool, error) {
	return _BridgeBank.Contract.VSA(&_BridgeBank.CallOpts, _sifAddress)
}

// VSA is a free data retrieval call binding the contract method 0xc228979d.
//
// Solidity: function VSA(bytes _sifAddress) pure returns(bool)
func (_BridgeBank *BridgeBankCallerSession) VSA(_sifAddress []byte) (bool, error) {
	return _BridgeBank.Contract.VSA(&_BridgeBank.CallOpts, _sifAddress)
}

// Blocklist is a free data retrieval call binding the contract method 0xd64e5396.
//
// Solidity: function blocklist() view returns(address)
func (_BridgeBank *BridgeBankCaller) Blocklist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "blocklist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Blocklist is a free data retrieval call binding the contract method 0xd64e5396.
//
// Solidity: function blocklist() view returns(address)
func (_BridgeBank *BridgeBankSession) Blocklist() (common.Address, error) {
	return _BridgeBank.Contract.Blocklist(&_BridgeBank.CallOpts)
}

// Blocklist is a free data retrieval call binding the contract method 0xd64e5396.
//
// Solidity: function blocklist() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) Blocklist() (common.Address, error) {
	return _BridgeBank.Contract.Blocklist(&_BridgeBank.CallOpts)
}

// ContractDecimals is a free data retrieval call binding the contract method 0xfc093d45.
//
// Solidity: function contractDecimals(address ) view returns(uint8)
func (_BridgeBank *BridgeBankCaller) ContractDecimals(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "contractDecimals", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ContractDecimals is a free data retrieval call binding the contract method 0xfc093d45.
//
// Solidity: function contractDecimals(address ) view returns(uint8)
func (_BridgeBank *BridgeBankSession) ContractDecimals(arg0 common.Address) (uint8, error) {
	return _BridgeBank.Contract.ContractDecimals(&_BridgeBank.CallOpts, arg0)
}

// ContractDecimals is a free data retrieval call binding the contract method 0xfc093d45.
//
// Solidity: function contractDecimals(address ) view returns(uint8)
func (_BridgeBank *BridgeBankCallerSession) ContractDecimals(arg0 common.Address) (uint8, error) {
	return _BridgeBank.Contract.ContractDecimals(&_BridgeBank.CallOpts, arg0)
}

// ContractDenom is a free data retrieval call binding the contract method 0xa6045927.
//
// Solidity: function contractDenom(address ) view returns(string)
func (_BridgeBank *BridgeBankCaller) ContractDenom(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "contractDenom", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractDenom is a free data retrieval call binding the contract method 0xa6045927.
//
// Solidity: function contractDenom(address ) view returns(string)
func (_BridgeBank *BridgeBankSession) ContractDenom(arg0 common.Address) (string, error) {
	return _BridgeBank.Contract.ContractDenom(&_BridgeBank.CallOpts, arg0)
}

// ContractDenom is a free data retrieval call binding the contract method 0xa6045927.
//
// Solidity: function contractDenom(address ) view returns(string)
func (_BridgeBank *BridgeBankCallerSession) ContractDenom(arg0 common.Address) (string, error) {
	return _BridgeBank.Contract.ContractDenom(&_BridgeBank.CallOpts, arg0)
}

// ContractName is a free data retrieval call binding the contract method 0xf33b2a1c.
//
// Solidity: function contractName(address ) view returns(string)
func (_BridgeBank *BridgeBankCaller) ContractName(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "contractName", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractName is a free data retrieval call binding the contract method 0xf33b2a1c.
//
// Solidity: function contractName(address ) view returns(string)
func (_BridgeBank *BridgeBankSession) ContractName(arg0 common.Address) (string, error) {
	return _BridgeBank.Contract.ContractName(&_BridgeBank.CallOpts, arg0)
}

// ContractName is a free data retrieval call binding the contract method 0xf33b2a1c.
//
// Solidity: function contractName(address ) view returns(string)
func (_BridgeBank *BridgeBankCallerSession) ContractName(arg0 common.Address) (string, error) {
	return _BridgeBank.Contract.ContractName(&_BridgeBank.CallOpts, arg0)
}

// ContractSymbol is a free data retrieval call binding the contract method 0x89398e0b.
//
// Solidity: function contractSymbol(address ) view returns(string)
func (_BridgeBank *BridgeBankCaller) ContractSymbol(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "contractSymbol", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractSymbol is a free data retrieval call binding the contract method 0x89398e0b.
//
// Solidity: function contractSymbol(address ) view returns(string)
func (_BridgeBank *BridgeBankSession) ContractSymbol(arg0 common.Address) (string, error) {
	return _BridgeBank.Contract.ContractSymbol(&_BridgeBank.CallOpts, arg0)
}

// ContractSymbol is a free data retrieval call binding the contract method 0x89398e0b.
//
// Solidity: function contractSymbol(address ) view returns(string)
func (_BridgeBank *BridgeBankCallerSession) ContractSymbol(arg0 common.Address) (string, error) {
	return _BridgeBank.Contract.ContractSymbol(&_BridgeBank.CallOpts, arg0)
}

// CosmosBridge is a free data retrieval call binding the contract method 0xb0e9ef71.
//
// Solidity: function cosmosBridge() view returns(address)
func (_BridgeBank *BridgeBankCaller) CosmosBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "cosmosBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CosmosBridge is a free data retrieval call binding the contract method 0xb0e9ef71.
//
// Solidity: function cosmosBridge() view returns(address)
func (_BridgeBank *BridgeBankSession) CosmosBridge() (common.Address, error) {
	return _BridgeBank.Contract.CosmosBridge(&_BridgeBank.CallOpts)
}

// CosmosBridge is a free data retrieval call binding the contract method 0xb0e9ef71.
//
// Solidity: function cosmosBridge() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) CosmosBridge() (common.Address, error) {
	return _BridgeBank.Contract.CosmosBridge(&_BridgeBank.CallOpts)
}

// GetCosmosTokenInWhiteList is a free data retrieval call binding the contract method 0x96e5c356.
//
// Solidity: function getCosmosTokenInWhiteList(address _token) view returns(bool)
func (_BridgeBank *BridgeBankCaller) GetCosmosTokenInWhiteList(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "getCosmosTokenInWhiteList", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetCosmosTokenInWhiteList is a free data retrieval call binding the contract method 0x96e5c356.
//
// Solidity: function getCosmosTokenInWhiteList(address _token) view returns(bool)
func (_BridgeBank *BridgeBankSession) GetCosmosTokenInWhiteList(_token common.Address) (bool, error) {
	return _BridgeBank.Contract.GetCosmosTokenInWhiteList(&_BridgeBank.CallOpts, _token)
}

// GetCosmosTokenInWhiteList is a free data retrieval call binding the contract method 0x96e5c356.
//
// Solidity: function getCosmosTokenInWhiteList(address _token) view returns(bool)
func (_BridgeBank *BridgeBankCallerSession) GetCosmosTokenInWhiteList(_token common.Address) (bool, error) {
	return _BridgeBank.Contract.GetCosmosTokenInWhiteList(&_BridgeBank.CallOpts, _token)
}

// HasBlocklist is a free data retrieval call binding the contract method 0x634210f4.
//
// Solidity: function hasBlocklist() view returns(bool)
func (_BridgeBank *BridgeBankCaller) HasBlocklist(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "hasBlocklist")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasBlocklist is a free data retrieval call binding the contract method 0x634210f4.
//
// Solidity: function hasBlocklist() view returns(bool)
func (_BridgeBank *BridgeBankSession) HasBlocklist() (bool, error) {
	return _BridgeBank.Contract.HasBlocklist(&_BridgeBank.CallOpts)
}

// HasBlocklist is a free data retrieval call binding the contract method 0x634210f4.
//
// Solidity: function hasBlocklist() view returns(bool)
func (_BridgeBank *BridgeBankCallerSession) HasBlocklist() (bool, error) {
	return _BridgeBank.Contract.HasBlocklist(&_BridgeBank.CallOpts)
}

// LockBurnNonce is a free data retrieval call binding the contract method 0x1deed3bb.
//
// Solidity: function lockBurnNonce() view returns(uint256)
func (_BridgeBank *BridgeBankCaller) LockBurnNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "lockBurnNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockBurnNonce is a free data retrieval call binding the contract method 0x1deed3bb.
//
// Solidity: function lockBurnNonce() view returns(uint256)
func (_BridgeBank *BridgeBankSession) LockBurnNonce() (*big.Int, error) {
	return _BridgeBank.Contract.LockBurnNonce(&_BridgeBank.CallOpts)
}

// LockBurnNonce is a free data retrieval call binding the contract method 0x1deed3bb.
//
// Solidity: function lockBurnNonce() view returns(uint256)
func (_BridgeBank *BridgeBankCallerSession) LockBurnNonce() (*big.Int, error) {
	return _BridgeBank.Contract.LockBurnNonce(&_BridgeBank.CallOpts)
}

// NetworkDescriptor is a free data retrieval call binding the contract method 0x909d4a1c.
//
// Solidity: function networkDescriptor() view returns(int32)
func (_BridgeBank *BridgeBankCaller) NetworkDescriptor(opts *bind.CallOpts) (int32, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "networkDescriptor")

	if err != nil {
		return *new(int32), err
	}

	out0 := *abi.ConvertType(out[0], new(int32)).(*int32)

	return out0, err

}

// NetworkDescriptor is a free data retrieval call binding the contract method 0x909d4a1c.
//
// Solidity: function networkDescriptor() view returns(int32)
func (_BridgeBank *BridgeBankSession) NetworkDescriptor() (int32, error) {
	return _BridgeBank.Contract.NetworkDescriptor(&_BridgeBank.CallOpts)
}

// NetworkDescriptor is a free data retrieval call binding the contract method 0x909d4a1c.
//
// Solidity: function networkDescriptor() view returns(int32)
func (_BridgeBank *BridgeBankCallerSession) NetworkDescriptor() (int32, error) {
	return _BridgeBank.Contract.NetworkDescriptor(&_BridgeBank.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_BridgeBank *BridgeBankCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_BridgeBank *BridgeBankSession) Operator() (common.Address, error) {
	return _BridgeBank.Contract.Operator(&_BridgeBank.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) Operator() (common.Address, error) {
	return _BridgeBank.Contract.Operator(&_BridgeBank.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeBank *BridgeBankCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeBank *BridgeBankSession) Owner() (common.Address, error) {
	return _BridgeBank.Contract.Owner(&_BridgeBank.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) Owner() (common.Address, error) {
	return _BridgeBank.Contract.Owner(&_BridgeBank.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeBank *BridgeBankCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeBank *BridgeBankSession) Paused() (bool, error) {
	return _BridgeBank.Contract.Paused(&_BridgeBank.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeBank *BridgeBankCallerSession) Paused() (bool, error) {
	return _BridgeBank.Contract.Paused(&_BridgeBank.CallOpts)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_BridgeBank *BridgeBankCaller) Pausers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "pausers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_BridgeBank *BridgeBankSession) Pausers(arg0 common.Address) (bool, error) {
	return _BridgeBank.Contract.Pausers(&_BridgeBank.CallOpts, arg0)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_BridgeBank *BridgeBankCallerSession) Pausers(arg0 common.Address) (bool, error) {
	return _BridgeBank.Contract.Pausers(&_BridgeBank.CallOpts, arg0)
}

// AddExistingBridgeToken is a paid mutator transaction binding the contract method 0xfcf9cf6e.
//
// Solidity: function addExistingBridgeToken(address contractAddress) returns(bool)
func (_BridgeBank *BridgeBankTransactor) AddExistingBridgeToken(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "addExistingBridgeToken", contractAddress)
}

// AddExistingBridgeToken is a paid mutator transaction binding the contract method 0xfcf9cf6e.
//
// Solidity: function addExistingBridgeToken(address contractAddress) returns(bool)
func (_BridgeBank *BridgeBankSession) AddExistingBridgeToken(contractAddress common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.AddExistingBridgeToken(&_BridgeBank.TransactOpts, contractAddress)
}

// AddExistingBridgeToken is a paid mutator transaction binding the contract method 0xfcf9cf6e.
//
// Solidity: function addExistingBridgeToken(address contractAddress) returns(bool)
func (_BridgeBank *BridgeBankTransactorSession) AddExistingBridgeToken(contractAddress common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.AddExistingBridgeToken(&_BridgeBank.TransactOpts, contractAddress)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_BridgeBank *BridgeBankTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_BridgeBank *BridgeBankSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.AddPauser(&_BridgeBank.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_BridgeBank *BridgeBankTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.AddPauser(&_BridgeBank.TransactOpts, account)
}

// BatchAddExistingBridgeTokens is a paid mutator transaction binding the contract method 0xda3b4415.
//
// Solidity: function batchAddExistingBridgeTokens(address[] contractsAddresses) returns(bool)
func (_BridgeBank *BridgeBankTransactor) BatchAddExistingBridgeTokens(opts *bind.TransactOpts, contractsAddresses []common.Address) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "batchAddExistingBridgeTokens", contractsAddresses)
}

// BatchAddExistingBridgeTokens is a paid mutator transaction binding the contract method 0xda3b4415.
//
// Solidity: function batchAddExistingBridgeTokens(address[] contractsAddresses) returns(bool)
func (_BridgeBank *BridgeBankSession) BatchAddExistingBridgeTokens(contractsAddresses []common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.BatchAddExistingBridgeTokens(&_BridgeBank.TransactOpts, contractsAddresses)
}

// BatchAddExistingBridgeTokens is a paid mutator transaction binding the contract method 0xda3b4415.
//
// Solidity: function batchAddExistingBridgeTokens(address[] contractsAddresses) returns(bool)
func (_BridgeBank *BridgeBankTransactorSession) BatchAddExistingBridgeTokens(contractsAddresses []common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.BatchAddExistingBridgeTokens(&_BridgeBank.TransactOpts, contractsAddresses)
}

// BatchForceSetBridgeTokenDenom is a paid mutator transaction binding the contract method 0x37ac4118.
//
// Solidity: function batchForceSetBridgeTokenDenom(address[] _tokens) returns(bool)
func (_BridgeBank *BridgeBankTransactor) BatchForceSetBridgeTokenDenom(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "batchForceSetBridgeTokenDenom", _tokens)
}

// BatchForceSetBridgeTokenDenom is a paid mutator transaction binding the contract method 0x37ac4118.
//
// Solidity: function batchForceSetBridgeTokenDenom(address[] _tokens) returns(bool)
func (_BridgeBank *BridgeBankSession) BatchForceSetBridgeTokenDenom(_tokens []common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.BatchForceSetBridgeTokenDenom(&_BridgeBank.TransactOpts, _tokens)
}

// BatchForceSetBridgeTokenDenom is a paid mutator transaction binding the contract method 0x37ac4118.
//
// Solidity: function batchForceSetBridgeTokenDenom(address[] _tokens) returns(bool)
func (_BridgeBank *BridgeBankTransactorSession) BatchForceSetBridgeTokenDenom(_tokens []common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.BatchForceSetBridgeTokenDenom(&_BridgeBank.TransactOpts, _tokens)
}

// BatchSetBridgeTokenDenom is a paid mutator transaction binding the contract method 0xf671ed3c.
//
// Solidity: function batchSetBridgeTokenDenom(address[] _tokens, string[] _denoms) returns(bool)
func (_BridgeBank *BridgeBankTransactor) BatchSetBridgeTokenDenom(opts *bind.TransactOpts, _tokens []common.Address, _denoms []string) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "batchSetBridgeTokenDenom", _tokens, _denoms)
}

// BatchSetBridgeTokenDenom is a paid mutator transaction binding the contract method 0xf671ed3c.
//
// Solidity: function batchSetBridgeTokenDenom(address[] _tokens, string[] _denoms) returns(bool)
func (_BridgeBank *BridgeBankSession) BatchSetBridgeTokenDenom(_tokens []common.Address, _denoms []string) (*types.Transaction, error) {
	return _BridgeBank.Contract.BatchSetBridgeTokenDenom(&_BridgeBank.TransactOpts, _tokens, _denoms)
}

// BatchSetBridgeTokenDenom is a paid mutator transaction binding the contract method 0xf671ed3c.
//
// Solidity: function batchSetBridgeTokenDenom(address[] _tokens, string[] _denoms) returns(bool)
func (_BridgeBank *BridgeBankTransactorSession) BatchSetBridgeTokenDenom(_tokens []common.Address, _denoms []string) (*types.Transaction, error) {
	return _BridgeBank.Contract.BatchSetBridgeTokenDenom(&_BridgeBank.TransactOpts, _tokens, _denoms)
}

// Burn is a paid mutator transaction binding the contract method 0xdc9ae17d.
//
// Solidity: function burn(bytes recipient, address token, uint256 amount) returns()
func (_BridgeBank *BridgeBankTransactor) Burn(opts *bind.TransactOpts, recipient []byte, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "burn", recipient, token, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xdc9ae17d.
//
// Solidity: function burn(bytes recipient, address token, uint256 amount) returns()
func (_BridgeBank *BridgeBankSession) Burn(recipient []byte, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.Burn(&_BridgeBank.TransactOpts, recipient, token, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xdc9ae17d.
//
// Solidity: function burn(bytes recipient, address token, uint256 amount) returns()
func (_BridgeBank *BridgeBankTransactorSession) Burn(recipient []byte, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.Burn(&_BridgeBank.TransactOpts, recipient, token, amount)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address _newOperator) returns()
func (_BridgeBank *BridgeBankTransactor) ChangeOperator(opts *bind.TransactOpts, _newOperator common.Address) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "changeOperator", _newOperator)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address _newOperator) returns()
func (_BridgeBank *BridgeBankSession) ChangeOperator(_newOperator common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.ChangeOperator(&_BridgeBank.TransactOpts, _newOperator)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address _newOperator) returns()
func (_BridgeBank *BridgeBankTransactorSession) ChangeOperator(_newOperator common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.ChangeOperator(&_BridgeBank.TransactOpts, _newOperator)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_BridgeBank *BridgeBankTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_BridgeBank *BridgeBankSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.ChangeOwner(&_BridgeBank.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_BridgeBank *BridgeBankTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.ChangeOwner(&_BridgeBank.TransactOpts, newOwner)
}

// CreateNewBridgeToken is a paid mutator transaction binding the contract method 0xd51c0e95.
//
// Solidity: function createNewBridgeToken(string name, string symbol, uint8 decimals, string cosmosDenom) returns(address)
func (_BridgeBank *BridgeBankTransactor) CreateNewBridgeToken(opts *bind.TransactOpts, name string, symbol string, decimals uint8, cosmosDenom string) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "createNewBridgeToken", name, symbol, decimals, cosmosDenom)
}

// CreateNewBridgeToken is a paid mutator transaction binding the contract method 0xd51c0e95.
//
// Solidity: function createNewBridgeToken(string name, string symbol, uint8 decimals, string cosmosDenom) returns(address)
func (_BridgeBank *BridgeBankSession) CreateNewBridgeToken(name string, symbol string, decimals uint8, cosmosDenom string) (*types.Transaction, error) {
	return _BridgeBank.Contract.CreateNewBridgeToken(&_BridgeBank.TransactOpts, name, symbol, decimals, cosmosDenom)
}

// CreateNewBridgeToken is a paid mutator transaction binding the contract method 0xd51c0e95.
//
// Solidity: function createNewBridgeToken(string name, string symbol, uint8 decimals, string cosmosDenom) returns(address)
func (_BridgeBank *BridgeBankTransactorSession) CreateNewBridgeToken(name string, symbol string, decimals uint8, cosmosDenom string) (*types.Transaction, error) {
	return _BridgeBank.Contract.CreateNewBridgeToken(&_BridgeBank.TransactOpts, name, symbol, decimals, cosmosDenom)
}

// ForceSetBridgeTokenDenom is a paid mutator transaction binding the contract method 0xf9fc3ccb.
//
// Solidity: function forceSetBridgeTokenDenom(address _token) returns(bool)
func (_BridgeBank *BridgeBankTransactor) ForceSetBridgeTokenDenom(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "forceSetBridgeTokenDenom", _token)
}

// ForceSetBridgeTokenDenom is a paid mutator transaction binding the contract method 0xf9fc3ccb.
//
// Solidity: function forceSetBridgeTokenDenom(address _token) returns(bool)
func (_BridgeBank *BridgeBankSession) ForceSetBridgeTokenDenom(_token common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.ForceSetBridgeTokenDenom(&_BridgeBank.TransactOpts, _token)
}

// ForceSetBridgeTokenDenom is a paid mutator transaction binding the contract method 0xf9fc3ccb.
//
// Solidity: function forceSetBridgeTokenDenom(address _token) returns(bool)
func (_BridgeBank *BridgeBankTransactorSession) ForceSetBridgeTokenDenom(_token common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.ForceSetBridgeTokenDenom(&_BridgeBank.TransactOpts, _token)
}

// HandleUnpeg is a paid mutator transaction binding the contract method 0xe4cf380a.
//
// Solidity: function handleUnpeg(address ethereumReceiver, address tokenAddress, uint256 amount) returns()
func (_BridgeBank *BridgeBankTransactor) HandleUnpeg(opts *bind.TransactOpts, ethereumReceiver common.Address, tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "handleUnpeg", ethereumReceiver, tokenAddress, amount)
}

// HandleUnpeg is a paid mutator transaction binding the contract method 0xe4cf380a.
//
// Solidity: function handleUnpeg(address ethereumReceiver, address tokenAddress, uint256 amount) returns()
func (_BridgeBank *BridgeBankSession) HandleUnpeg(ethereumReceiver common.Address, tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.HandleUnpeg(&_BridgeBank.TransactOpts, ethereumReceiver, tokenAddress, amount)
}

// HandleUnpeg is a paid mutator transaction binding the contract method 0xe4cf380a.
//
// Solidity: function handleUnpeg(address ethereumReceiver, address tokenAddress, uint256 amount) returns()
func (_BridgeBank *BridgeBankTransactorSession) HandleUnpeg(ethereumReceiver common.Address, tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.HandleUnpeg(&_BridgeBank.TransactOpts, ethereumReceiver, tokenAddress, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x6e187272.
//
// Solidity: function initialize(address _operator, address _cosmosBridgeAddress, address _owner, address _pauser, int32 _networkDescriptor) returns()
func (_BridgeBank *BridgeBankTransactor) Initialize(opts *bind.TransactOpts, _operator common.Address, _cosmosBridgeAddress common.Address, _owner common.Address, _pauser common.Address, _networkDescriptor int32) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "initialize", _operator, _cosmosBridgeAddress, _owner, _pauser, _networkDescriptor)
}

// Initialize is a paid mutator transaction binding the contract method 0x6e187272.
//
// Solidity: function initialize(address _operator, address _cosmosBridgeAddress, address _owner, address _pauser, int32 _networkDescriptor) returns()
func (_BridgeBank *BridgeBankSession) Initialize(_operator common.Address, _cosmosBridgeAddress common.Address, _owner common.Address, _pauser common.Address, _networkDescriptor int32) (*types.Transaction, error) {
	return _BridgeBank.Contract.Initialize(&_BridgeBank.TransactOpts, _operator, _cosmosBridgeAddress, _owner, _pauser, _networkDescriptor)
}

// Initialize is a paid mutator transaction binding the contract method 0x6e187272.
//
// Solidity: function initialize(address _operator, address _cosmosBridgeAddress, address _owner, address _pauser, int32 _networkDescriptor) returns()
func (_BridgeBank *BridgeBankTransactorSession) Initialize(_operator common.Address, _cosmosBridgeAddress common.Address, _owner common.Address, _pauser common.Address, _networkDescriptor int32) (*types.Transaction, error) {
	return _BridgeBank.Contract.Initialize(&_BridgeBank.TransactOpts, _operator, _cosmosBridgeAddress, _owner, _pauser, _networkDescriptor)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_BridgeBank *BridgeBankTransactor) Initialize0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "initialize0")
}

// Initialize0 is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_BridgeBank *BridgeBankSession) Initialize0() (*types.Transaction, error) {
	return _BridgeBank.Contract.Initialize0(&_BridgeBank.TransactOpts)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_BridgeBank *BridgeBankTransactorSession) Initialize0() (*types.Transaction, error) {
	return _BridgeBank.Contract.Initialize0(&_BridgeBank.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0x9df2a385.
//
// Solidity: function lock(bytes recipient, address token, uint256 amount) payable returns()
func (_BridgeBank *BridgeBankTransactor) Lock(opts *bind.TransactOpts, recipient []byte, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "lock", recipient, token, amount)
}

// Lock is a paid mutator transaction binding the contract method 0x9df2a385.
//
// Solidity: function lock(bytes recipient, address token, uint256 amount) payable returns()
func (_BridgeBank *BridgeBankSession) Lock(recipient []byte, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.Lock(&_BridgeBank.TransactOpts, recipient, token, amount)
}

// Lock is a paid mutator transaction binding the contract method 0x9df2a385.
//
// Solidity: function lock(bytes recipient, address token, uint256 amount) payable returns()
func (_BridgeBank *BridgeBankTransactorSession) Lock(recipient []byte, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.Lock(&_BridgeBank.TransactOpts, recipient, token, amount)
}

// MultiLock is a paid mutator transaction binding the contract method 0x715283c1.
//
// Solidity: function multiLock(bytes[] recipient, address[] token, uint256[] amount) returns()
func (_BridgeBank *BridgeBankTransactor) MultiLock(opts *bind.TransactOpts, recipient [][]byte, token []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "multiLock", recipient, token, amount)
}

// MultiLock is a paid mutator transaction binding the contract method 0x715283c1.
//
// Solidity: function multiLock(bytes[] recipient, address[] token, uint256[] amount) returns()
func (_BridgeBank *BridgeBankSession) MultiLock(recipient [][]byte, token []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.MultiLock(&_BridgeBank.TransactOpts, recipient, token, amount)
}

// MultiLock is a paid mutator transaction binding the contract method 0x715283c1.
//
// Solidity: function multiLock(bytes[] recipient, address[] token, uint256[] amount) returns()
func (_BridgeBank *BridgeBankTransactorSession) MultiLock(recipient [][]byte, token []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.MultiLock(&_BridgeBank.TransactOpts, recipient, token, amount)
}

// MultiLockBurn is a paid mutator transaction binding the contract method 0x846b6fb8.
//
// Solidity: function multiLockBurn(bytes[] recipient, address[] token, uint256[] amount, bool[] isBurn) returns()
func (_BridgeBank *BridgeBankTransactor) MultiLockBurn(opts *bind.TransactOpts, recipient [][]byte, token []common.Address, amount []*big.Int, isBurn []bool) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "multiLockBurn", recipient, token, amount, isBurn)
}

// MultiLockBurn is a paid mutator transaction binding the contract method 0x846b6fb8.
//
// Solidity: function multiLockBurn(bytes[] recipient, address[] token, uint256[] amount, bool[] isBurn) returns()
func (_BridgeBank *BridgeBankSession) MultiLockBurn(recipient [][]byte, token []common.Address, amount []*big.Int, isBurn []bool) (*types.Transaction, error) {
	return _BridgeBank.Contract.MultiLockBurn(&_BridgeBank.TransactOpts, recipient, token, amount, isBurn)
}

// MultiLockBurn is a paid mutator transaction binding the contract method 0x846b6fb8.
//
// Solidity: function multiLockBurn(bytes[] recipient, address[] token, uint256[] amount, bool[] isBurn) returns()
func (_BridgeBank *BridgeBankTransactorSession) MultiLockBurn(recipient [][]byte, token []common.Address, amount []*big.Int, isBurn []bool) (*types.Transaction, error) {
	return _BridgeBank.Contract.MultiLockBurn(&_BridgeBank.TransactOpts, recipient, token, amount, isBurn)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeBank *BridgeBankTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeBank *BridgeBankSession) Pause() (*types.Transaction, error) {
	return _BridgeBank.Contract.Pause(&_BridgeBank.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeBank *BridgeBankTransactorSession) Pause() (*types.Transaction, error) {
	return _BridgeBank.Contract.Pause(&_BridgeBank.TransactOpts)
}

// Reinitialize is a paid mutator transaction binding the contract method 0x813d0d33.
//
// Solidity: function reinitialize(address _operator, address _cosmosBridgeAddress, address _owner, address _pauser, int32 _networkDescriptor) returns()
func (_BridgeBank *BridgeBankTransactor) Reinitialize(opts *bind.TransactOpts, _operator common.Address, _cosmosBridgeAddress common.Address, _owner common.Address, _pauser common.Address, _networkDescriptor int32) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "reinitialize", _operator, _cosmosBridgeAddress, _owner, _pauser, _networkDescriptor)
}

// Reinitialize is a paid mutator transaction binding the contract method 0x813d0d33.
//
// Solidity: function reinitialize(address _operator, address _cosmosBridgeAddress, address _owner, address _pauser, int32 _networkDescriptor) returns()
func (_BridgeBank *BridgeBankSession) Reinitialize(_operator common.Address, _cosmosBridgeAddress common.Address, _owner common.Address, _pauser common.Address, _networkDescriptor int32) (*types.Transaction, error) {
	return _BridgeBank.Contract.Reinitialize(&_BridgeBank.TransactOpts, _operator, _cosmosBridgeAddress, _owner, _pauser, _networkDescriptor)
}

// Reinitialize is a paid mutator transaction binding the contract method 0x813d0d33.
//
// Solidity: function reinitialize(address _operator, address _cosmosBridgeAddress, address _owner, address _pauser, int32 _networkDescriptor) returns()
func (_BridgeBank *BridgeBankTransactorSession) Reinitialize(_operator common.Address, _cosmosBridgeAddress common.Address, _owner common.Address, _pauser common.Address, _networkDescriptor int32) (*types.Transaction, error) {
	return _BridgeBank.Contract.Reinitialize(&_BridgeBank.TransactOpts, _operator, _cosmosBridgeAddress, _owner, _pauser, _networkDescriptor)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_BridgeBank *BridgeBankTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_BridgeBank *BridgeBankSession) RenouncePauser() (*types.Transaction, error) {
	return _BridgeBank.Contract.RenouncePauser(&_BridgeBank.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_BridgeBank *BridgeBankTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _BridgeBank.Contract.RenouncePauser(&_BridgeBank.TransactOpts)
}

// SetBlocklist is a paid mutator transaction binding the contract method 0xaef18ae7.
//
// Solidity: function setBlocklist(address blocklistAddress) returns()
func (_BridgeBank *BridgeBankTransactor) SetBlocklist(opts *bind.TransactOpts, blocklistAddress common.Address) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "setBlocklist", blocklistAddress)
}

// SetBlocklist is a paid mutator transaction binding the contract method 0xaef18ae7.
//
// Solidity: function setBlocklist(address blocklistAddress) returns()
func (_BridgeBank *BridgeBankSession) SetBlocklist(blocklistAddress common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.SetBlocklist(&_BridgeBank.TransactOpts, blocklistAddress)
}

// SetBlocklist is a paid mutator transaction binding the contract method 0xaef18ae7.
//
// Solidity: function setBlocklist(address blocklistAddress) returns()
func (_BridgeBank *BridgeBankTransactorSession) SetBlocklist(blocklistAddress common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.SetBlocklist(&_BridgeBank.TransactOpts, blocklistAddress)
}

// SetBridgeTokenDenom is a paid mutator transaction binding the contract method 0x05d4f0f6.
//
// Solidity: function setBridgeTokenDenom(address _token, string _denom) returns(bool)
func (_BridgeBank *BridgeBankTransactor) SetBridgeTokenDenom(opts *bind.TransactOpts, _token common.Address, _denom string) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "setBridgeTokenDenom", _token, _denom)
}

// SetBridgeTokenDenom is a paid mutator transaction binding the contract method 0x05d4f0f6.
//
// Solidity: function setBridgeTokenDenom(address _token, string _denom) returns(bool)
func (_BridgeBank *BridgeBankSession) SetBridgeTokenDenom(_token common.Address, _denom string) (*types.Transaction, error) {
	return _BridgeBank.Contract.SetBridgeTokenDenom(&_BridgeBank.TransactOpts, _token, _denom)
}

// SetBridgeTokenDenom is a paid mutator transaction binding the contract method 0x05d4f0f6.
//
// Solidity: function setBridgeTokenDenom(address _token, string _denom) returns(bool)
func (_BridgeBank *BridgeBankTransactorSession) SetBridgeTokenDenom(_token common.Address, _denom string) (*types.Transaction, error) {
	return _BridgeBank.Contract.SetBridgeTokenDenom(&_BridgeBank.TransactOpts, _token, _denom)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeBank *BridgeBankTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeBank *BridgeBankSession) Unpause() (*types.Transaction, error) {
	return _BridgeBank.Contract.Unpause(&_BridgeBank.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeBank *BridgeBankTransactorSession) Unpause() (*types.Transaction, error) {
	return _BridgeBank.Contract.Unpause(&_BridgeBank.TransactOpts)
}

// BridgeBankLogBridgeTokenMintIterator is returned from FilterLogBridgeTokenMint and is used to iterate over the raw logs and unpacked data for LogBridgeTokenMint events raised by the BridgeBank contract.
type BridgeBankLogBridgeTokenMintIterator struct {
	Event *BridgeBankLogBridgeTokenMint // Event containing the contract specifics and raw log

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
func (it *BridgeBankLogBridgeTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankLogBridgeTokenMint)
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
		it.Event = new(BridgeBankLogBridgeTokenMint)
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
func (it *BridgeBankLogBridgeTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankLogBridgeTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankLogBridgeTokenMint represents a LogBridgeTokenMint event raised by the BridgeBank contract.
type BridgeBankLogBridgeTokenMint struct {
	Token       common.Address
	Amount      *big.Int
	Beneficiary common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogBridgeTokenMint is a free log retrieval operation binding the contract event 0x4ce045d816af1004ed2e4feb414f2bdaf7e3c92341dd98275e665e861d19551c.
//
// Solidity: event LogBridgeTokenMint(address _token, uint256 _amount, address _beneficiary)
func (_BridgeBank *BridgeBankFilterer) FilterLogBridgeTokenMint(opts *bind.FilterOpts) (*BridgeBankLogBridgeTokenMintIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "LogBridgeTokenMint")
	if err != nil {
		return nil, err
	}
	return &BridgeBankLogBridgeTokenMintIterator{contract: _BridgeBank.contract, event: "LogBridgeTokenMint", logs: logs, sub: sub}, nil
}

// WatchLogBridgeTokenMint is a free log subscription operation binding the contract event 0x4ce045d816af1004ed2e4feb414f2bdaf7e3c92341dd98275e665e861d19551c.
//
// Solidity: event LogBridgeTokenMint(address _token, uint256 _amount, address _beneficiary)
func (_BridgeBank *BridgeBankFilterer) WatchLogBridgeTokenMint(opts *bind.WatchOpts, sink chan<- *BridgeBankLogBridgeTokenMint) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "LogBridgeTokenMint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankLogBridgeTokenMint)
				if err := _BridgeBank.contract.UnpackLog(event, "LogBridgeTokenMint", log); err != nil {
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

// ParseLogBridgeTokenMint is a log parse operation binding the contract event 0x4ce045d816af1004ed2e4feb414f2bdaf7e3c92341dd98275e665e861d19551c.
//
// Solidity: event LogBridgeTokenMint(address _token, uint256 _amount, address _beneficiary)
func (_BridgeBank *BridgeBankFilterer) ParseLogBridgeTokenMint(log types.Log) (*BridgeBankLogBridgeTokenMint, error) {
	event := new(BridgeBankLogBridgeTokenMint)
	if err := _BridgeBank.contract.UnpackLog(event, "LogBridgeTokenMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBankLogBurnIterator is returned from FilterLogBurn and is used to iterate over the raw logs and unpacked data for LogBurn events raised by the BridgeBank contract.
type BridgeBankLogBurnIterator struct {
	Event *BridgeBankLogBurn // Event containing the contract specifics and raw log

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
func (it *BridgeBankLogBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankLogBurn)
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
		it.Event = new(BridgeBankLogBurn)
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
func (it *BridgeBankLogBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankLogBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankLogBurn represents a LogBurn event raised by the BridgeBank contract.
type BridgeBankLogBurn struct {
	From              common.Address
	To                []byte
	Token             common.Address
	Value             *big.Int
	Nonce             *big.Int
	Decimals          uint8
	NetworkDescriptor int32
	Denom             string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLogBurn is a free log retrieval operation binding the contract event 0x86b9aeb12786b553000528f6918cdd9cc7ef53005a650c68f6b4eba274c6a232.
//
// Solidity: event LogBurn(address _from, bytes _to, address _token, uint256 _value, uint256 indexed _nonce, uint8 _decimals, int32 _networkDescriptor, string _denom)
func (_BridgeBank *BridgeBankFilterer) FilterLogBurn(opts *bind.FilterOpts, _nonce []*big.Int) (*BridgeBankLogBurnIterator, error) {

	var _nonceRule []interface{}
	for _, _nonceItem := range _nonce {
		_nonceRule = append(_nonceRule, _nonceItem)
	}

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "LogBurn", _nonceRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBankLogBurnIterator{contract: _BridgeBank.contract, event: "LogBurn", logs: logs, sub: sub}, nil
}

// WatchLogBurn is a free log subscription operation binding the contract event 0x86b9aeb12786b553000528f6918cdd9cc7ef53005a650c68f6b4eba274c6a232.
//
// Solidity: event LogBurn(address _from, bytes _to, address _token, uint256 _value, uint256 indexed _nonce, uint8 _decimals, int32 _networkDescriptor, string _denom)
func (_BridgeBank *BridgeBankFilterer) WatchLogBurn(opts *bind.WatchOpts, sink chan<- *BridgeBankLogBurn, _nonce []*big.Int) (event.Subscription, error) {

	var _nonceRule []interface{}
	for _, _nonceItem := range _nonce {
		_nonceRule = append(_nonceRule, _nonceItem)
	}

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "LogBurn", _nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankLogBurn)
				if err := _BridgeBank.contract.UnpackLog(event, "LogBurn", log); err != nil {
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

// ParseLogBurn is a log parse operation binding the contract event 0x86b9aeb12786b553000528f6918cdd9cc7ef53005a650c68f6b4eba274c6a232.
//
// Solidity: event LogBurn(address _from, bytes _to, address _token, uint256 _value, uint256 indexed _nonce, uint8 _decimals, int32 _networkDescriptor, string _denom)
func (_BridgeBank *BridgeBankFilterer) ParseLogBurn(log types.Log) (*BridgeBankLogBurn, error) {
	event := new(BridgeBankLogBurn)
	if err := _BridgeBank.contract.UnpackLog(event, "LogBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBankLogLockIterator is returned from FilterLogLock and is used to iterate over the raw logs and unpacked data for LogLock events raised by the BridgeBank contract.
type BridgeBankLogLockIterator struct {
	Event *BridgeBankLogLock // Event containing the contract specifics and raw log

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
func (it *BridgeBankLogLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankLogLock)
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
		it.Event = new(BridgeBankLogLock)
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
func (it *BridgeBankLogLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankLogLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankLogLock represents a LogLock event raised by the BridgeBank contract.
type BridgeBankLogLock struct {
	From              common.Address
	To                []byte
	Token             common.Address
	Value             *big.Int
	Nonce             *big.Int
	Decimals          uint8
	Symbol            string
	Name              string
	NetworkDescriptor int32
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLogLock is a free log retrieval operation binding the contract event 0x38da50ece741cf1ee3e6f7da39cd068935085be46091e50cb29ec0293c82a1c2.
//
// Solidity: event LogLock(address _from, bytes _to, address _token, uint256 _value, uint256 indexed _nonce, uint8 _decimals, string _symbol, string _name, int32 _networkDescriptor)
func (_BridgeBank *BridgeBankFilterer) FilterLogLock(opts *bind.FilterOpts, _nonce []*big.Int) (*BridgeBankLogLockIterator, error) {

	var _nonceRule []interface{}
	for _, _nonceItem := range _nonce {
		_nonceRule = append(_nonceRule, _nonceItem)
	}

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "LogLock", _nonceRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBankLogLockIterator{contract: _BridgeBank.contract, event: "LogLock", logs: logs, sub: sub}, nil
}

// WatchLogLock is a free log subscription operation binding the contract event 0x38da50ece741cf1ee3e6f7da39cd068935085be46091e50cb29ec0293c82a1c2.
//
// Solidity: event LogLock(address _from, bytes _to, address _token, uint256 _value, uint256 indexed _nonce, uint8 _decimals, string _symbol, string _name, int32 _networkDescriptor)
func (_BridgeBank *BridgeBankFilterer) WatchLogLock(opts *bind.WatchOpts, sink chan<- *BridgeBankLogLock, _nonce []*big.Int) (event.Subscription, error) {

	var _nonceRule []interface{}
	for _, _nonceItem := range _nonce {
		_nonceRule = append(_nonceRule, _nonceItem)
	}

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "LogLock", _nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankLogLock)
				if err := _BridgeBank.contract.UnpackLog(event, "LogLock", log); err != nil {
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

// ParseLogLock is a log parse operation binding the contract event 0x38da50ece741cf1ee3e6f7da39cd068935085be46091e50cb29ec0293c82a1c2.
//
// Solidity: event LogLock(address _from, bytes _to, address _token, uint256 _value, uint256 indexed _nonce, uint8 _decimals, string _symbol, string _name, int32 _networkDescriptor)
func (_BridgeBank *BridgeBankFilterer) ParseLogLock(log types.Log) (*BridgeBankLogLock, error) {
	event := new(BridgeBankLogLock)
	if err := _BridgeBank.contract.UnpackLog(event, "LogLock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBankLogNewBridgeTokenIterator is returned from FilterLogNewBridgeToken and is used to iterate over the raw logs and unpacked data for LogNewBridgeToken events raised by the BridgeBank contract.
type BridgeBankLogNewBridgeTokenIterator struct {
	Event *BridgeBankLogNewBridgeToken // Event containing the contract specifics and raw log

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
func (it *BridgeBankLogNewBridgeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankLogNewBridgeToken)
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
		it.Event = new(BridgeBankLogNewBridgeToken)
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
func (it *BridgeBankLogNewBridgeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankLogNewBridgeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankLogNewBridgeToken represents a LogNewBridgeToken event raised by the BridgeBank contract.
type BridgeBankLogNewBridgeToken struct {
	Token       common.Address
	Symbol      common.Hash
	CosmosDenom common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogNewBridgeToken is a free log retrieval operation binding the contract event 0x93cfe31d5f3f010792a839d81994f04ccd1ba4711cc428b7029e6ab3d522848f.
//
// Solidity: event LogNewBridgeToken(address indexed _token, string indexed _symbol, string indexed _cosmosDenom)
func (_BridgeBank *BridgeBankFilterer) FilterLogNewBridgeToken(opts *bind.FilterOpts, _token []common.Address, _symbol []string, _cosmosDenom []string) (*BridgeBankLogNewBridgeTokenIterator, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}
	var _symbolRule []interface{}
	for _, _symbolItem := range _symbol {
		_symbolRule = append(_symbolRule, _symbolItem)
	}
	var _cosmosDenomRule []interface{}
	for _, _cosmosDenomItem := range _cosmosDenom {
		_cosmosDenomRule = append(_cosmosDenomRule, _cosmosDenomItem)
	}

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "LogNewBridgeToken", _tokenRule, _symbolRule, _cosmosDenomRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBankLogNewBridgeTokenIterator{contract: _BridgeBank.contract, event: "LogNewBridgeToken", logs: logs, sub: sub}, nil
}

// WatchLogNewBridgeToken is a free log subscription operation binding the contract event 0x93cfe31d5f3f010792a839d81994f04ccd1ba4711cc428b7029e6ab3d522848f.
//
// Solidity: event LogNewBridgeToken(address indexed _token, string indexed _symbol, string indexed _cosmosDenom)
func (_BridgeBank *BridgeBankFilterer) WatchLogNewBridgeToken(opts *bind.WatchOpts, sink chan<- *BridgeBankLogNewBridgeToken, _token []common.Address, _symbol []string, _cosmosDenom []string) (event.Subscription, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}
	var _symbolRule []interface{}
	for _, _symbolItem := range _symbol {
		_symbolRule = append(_symbolRule, _symbolItem)
	}
	var _cosmosDenomRule []interface{}
	for _, _cosmosDenomItem := range _cosmosDenom {
		_cosmosDenomRule = append(_cosmosDenomRule, _cosmosDenomItem)
	}

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "LogNewBridgeToken", _tokenRule, _symbolRule, _cosmosDenomRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankLogNewBridgeToken)
				if err := _BridgeBank.contract.UnpackLog(event, "LogNewBridgeToken", log); err != nil {
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

// ParseLogNewBridgeToken is a log parse operation binding the contract event 0x93cfe31d5f3f010792a839d81994f04ccd1ba4711cc428b7029e6ab3d522848f.
//
// Solidity: event LogNewBridgeToken(address indexed _token, string indexed _symbol, string indexed _cosmosDenom)
func (_BridgeBank *BridgeBankFilterer) ParseLogNewBridgeToken(log types.Log) (*BridgeBankLogNewBridgeToken, error) {
	event := new(BridgeBankLogNewBridgeToken)
	if err := _BridgeBank.contract.UnpackLog(event, "LogNewBridgeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBankLogUnlockIterator is returned from FilterLogUnlock and is used to iterate over the raw logs and unpacked data for LogUnlock events raised by the BridgeBank contract.
type BridgeBankLogUnlockIterator struct {
	Event *BridgeBankLogUnlock // Event containing the contract specifics and raw log

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
func (it *BridgeBankLogUnlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankLogUnlock)
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
		it.Event = new(BridgeBankLogUnlock)
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
func (it *BridgeBankLogUnlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankLogUnlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankLogUnlock represents a LogUnlock event raised by the BridgeBank contract.
type BridgeBankLogUnlock struct {
	To    common.Address
	Token common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogUnlock is a free log retrieval operation binding the contract event 0xc2c64ff0cfc4d626042b306aa9b2f79227fcf39aeb429429a4d98d573fd009a4.
//
// Solidity: event LogUnlock(address _to, address _token, uint256 _value)
func (_BridgeBank *BridgeBankFilterer) FilterLogUnlock(opts *bind.FilterOpts) (*BridgeBankLogUnlockIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "LogUnlock")
	if err != nil {
		return nil, err
	}
	return &BridgeBankLogUnlockIterator{contract: _BridgeBank.contract, event: "LogUnlock", logs: logs, sub: sub}, nil
}

// WatchLogUnlock is a free log subscription operation binding the contract event 0xc2c64ff0cfc4d626042b306aa9b2f79227fcf39aeb429429a4d98d573fd009a4.
//
// Solidity: event LogUnlock(address _to, address _token, uint256 _value)
func (_BridgeBank *BridgeBankFilterer) WatchLogUnlock(opts *bind.WatchOpts, sink chan<- *BridgeBankLogUnlock) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "LogUnlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankLogUnlock)
				if err := _BridgeBank.contract.UnpackLog(event, "LogUnlock", log); err != nil {
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

// ParseLogUnlock is a log parse operation binding the contract event 0xc2c64ff0cfc4d626042b306aa9b2f79227fcf39aeb429429a4d98d573fd009a4.
//
// Solidity: event LogUnlock(address _to, address _token, uint256 _value)
func (_BridgeBank *BridgeBankFilterer) ParseLogUnlock(log types.Log) (*BridgeBankLogUnlock, error) {
	event := new(BridgeBankLogUnlock)
	if err := _BridgeBank.contract.UnpackLog(event, "LogUnlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBankLogWhiteListUpdateIterator is returned from FilterLogWhiteListUpdate and is used to iterate over the raw logs and unpacked data for LogWhiteListUpdate events raised by the BridgeBank contract.
type BridgeBankLogWhiteListUpdateIterator struct {
	Event *BridgeBankLogWhiteListUpdate // Event containing the contract specifics and raw log

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
func (it *BridgeBankLogWhiteListUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankLogWhiteListUpdate)
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
		it.Event = new(BridgeBankLogWhiteListUpdate)
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
func (it *BridgeBankLogWhiteListUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankLogWhiteListUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankLogWhiteListUpdate represents a LogWhiteListUpdate event raised by the BridgeBank contract.
type BridgeBankLogWhiteListUpdate struct {
	Token common.Address
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogWhiteListUpdate is a free log retrieval operation binding the contract event 0xbabeec6613676f4972214f70ddf47c77eb8360e09b54278f21d6f999faf7c15c.
//
// Solidity: event LogWhiteListUpdate(address _token, bool _value)
func (_BridgeBank *BridgeBankFilterer) FilterLogWhiteListUpdate(opts *bind.FilterOpts) (*BridgeBankLogWhiteListUpdateIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "LogWhiteListUpdate")
	if err != nil {
		return nil, err
	}
	return &BridgeBankLogWhiteListUpdateIterator{contract: _BridgeBank.contract, event: "LogWhiteListUpdate", logs: logs, sub: sub}, nil
}

// WatchLogWhiteListUpdate is a free log subscription operation binding the contract event 0xbabeec6613676f4972214f70ddf47c77eb8360e09b54278f21d6f999faf7c15c.
//
// Solidity: event LogWhiteListUpdate(address _token, bool _value)
func (_BridgeBank *BridgeBankFilterer) WatchLogWhiteListUpdate(opts *bind.WatchOpts, sink chan<- *BridgeBankLogWhiteListUpdate) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "LogWhiteListUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankLogWhiteListUpdate)
				if err := _BridgeBank.contract.UnpackLog(event, "LogWhiteListUpdate", log); err != nil {
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

// ParseLogWhiteListUpdate is a log parse operation binding the contract event 0xbabeec6613676f4972214f70ddf47c77eb8360e09b54278f21d6f999faf7c15c.
//
// Solidity: event LogWhiteListUpdate(address _token, bool _value)
func (_BridgeBank *BridgeBankFilterer) ParseLogWhiteListUpdate(log types.Log) (*BridgeBankLogWhiteListUpdate, error) {
	event := new(BridgeBankLogWhiteListUpdate)
	if err := _BridgeBank.contract.UnpackLog(event, "LogWhiteListUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBankPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BridgeBank contract.
type BridgeBankPausedIterator struct {
	Event *BridgeBankPaused // Event containing the contract specifics and raw log

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
func (it *BridgeBankPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankPaused)
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
		it.Event = new(BridgeBankPaused)
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
func (it *BridgeBankPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankPaused represents a Paused event raised by the BridgeBank contract.
type BridgeBankPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BridgeBank *BridgeBankFilterer) FilterPaused(opts *bind.FilterOpts) (*BridgeBankPausedIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BridgeBankPausedIterator{contract: _BridgeBank.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BridgeBank *BridgeBankFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BridgeBankPaused) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankPaused)
				if err := _BridgeBank.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BridgeBank *BridgeBankFilterer) ParsePaused(log types.Log) (*BridgeBankPaused, error) {
	event := new(BridgeBankPaused)
	if err := _BridgeBank.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBankUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BridgeBank contract.
type BridgeBankUnpausedIterator struct {
	Event *BridgeBankUnpaused // Event containing the contract specifics and raw log

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
func (it *BridgeBankUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankUnpaused)
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
		it.Event = new(BridgeBankUnpaused)
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
func (it *BridgeBankUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankUnpaused represents a Unpaused event raised by the BridgeBank contract.
type BridgeBankUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BridgeBank *BridgeBankFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BridgeBankUnpausedIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BridgeBankUnpausedIterator{contract: _BridgeBank.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BridgeBank *BridgeBankFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BridgeBankUnpaused) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankUnpaused)
				if err := _BridgeBank.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BridgeBank *BridgeBankFilterer) ParseUnpaused(log types.Log) (*BridgeBankUnpaused, error) {
	event := new(BridgeBankUnpaused)
	if err := _BridgeBank.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
