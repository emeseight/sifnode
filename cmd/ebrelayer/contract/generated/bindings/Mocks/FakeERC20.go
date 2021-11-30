// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Mocks

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

// MocksMetaData contains all meta data concerning the Mocks contract.
var MocksMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MocksABI is the input ABI used to generate the binding from.
// Deprecated: Use MocksMetaData.ABI instead.
var MocksABI = MocksMetaData.ABI

// Mocks is an auto generated Go binding around an Ethereum contract.
type Mocks struct {
	MocksCaller     // Read-only binding to the contract
	MocksTransactor // Write-only binding to the contract
	MocksFilterer   // Log filterer for contract events
}

// MocksCaller is an auto generated read-only Go binding around an Ethereum contract.
type MocksCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MocksTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MocksTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MocksFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MocksFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MocksSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MocksSession struct {
	Contract     *Mocks            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MocksCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MocksCallerSession struct {
	Contract *MocksCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MocksTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MocksTransactorSession struct {
	Contract     *MocksTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MocksRaw is an auto generated low-level Go binding around an Ethereum contract.
type MocksRaw struct {
	Contract *Mocks // Generic contract binding to access the raw methods on
}

// MocksCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MocksCallerRaw struct {
	Contract *MocksCaller // Generic read-only contract binding to access the raw methods on
}

// MocksTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MocksTransactorRaw struct {
	Contract *MocksTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMocks creates a new instance of Mocks, bound to a specific deployed contract.
func NewMocks(address common.Address, backend bind.ContractBackend) (*Mocks, error) {
	contract, err := bindMocks(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mocks{MocksCaller: MocksCaller{contract: contract}, MocksTransactor: MocksTransactor{contract: contract}, MocksFilterer: MocksFilterer{contract: contract}}, nil
}

// NewMocksCaller creates a new read-only instance of Mocks, bound to a specific deployed contract.
func NewMocksCaller(address common.Address, caller bind.ContractCaller) (*MocksCaller, error) {
	contract, err := bindMocks(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MocksCaller{contract: contract}, nil
}

// NewMocksTransactor creates a new write-only instance of Mocks, bound to a specific deployed contract.
func NewMocksTransactor(address common.Address, transactor bind.ContractTransactor) (*MocksTransactor, error) {
	contract, err := bindMocks(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MocksTransactor{contract: contract}, nil
}

// NewMocksFilterer creates a new log filterer instance of Mocks, bound to a specific deployed contract.
func NewMocksFilterer(address common.Address, filterer bind.ContractFilterer) (*MocksFilterer, error) {
	contract, err := bindMocks(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MocksFilterer{contract: contract}, nil
}

// bindMocks binds a generic wrapper to an already deployed contract.
func bindMocks(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MocksABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mocks *MocksRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mocks.Contract.MocksCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mocks *MocksRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mocks.Contract.MocksTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mocks *MocksRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mocks.Contract.MocksTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mocks *MocksCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mocks.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mocks *MocksTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mocks.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mocks *MocksTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mocks.Contract.contract.Transact(opts, method, params...)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Mocks *MocksTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Mocks.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Mocks *MocksSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Mocks.Contract.TransferFrom(&_Mocks.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Mocks *MocksTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Mocks.Contract.TransferFrom(&_Mocks.TransactOpts, from, to, value)
}
