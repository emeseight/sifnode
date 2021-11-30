// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CosmosBridgeStorage

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

// CosmosBridgeStorageMetaData contains all meta data concerning the CosmosBridgeStorage contract.
var CosmosBridgeStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"bridgeBank\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasBridgeBank\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CosmosBridgeStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use CosmosBridgeStorageMetaData.ABI instead.
var CosmosBridgeStorageABI = CosmosBridgeStorageMetaData.ABI

// CosmosBridgeStorage is an auto generated Go binding around an Ethereum contract.
type CosmosBridgeStorage struct {
	CosmosBridgeStorageCaller     // Read-only binding to the contract
	CosmosBridgeStorageTransactor // Write-only binding to the contract
	CosmosBridgeStorageFilterer   // Log filterer for contract events
}

// CosmosBridgeStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type CosmosBridgeStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmosBridgeStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CosmosBridgeStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmosBridgeStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CosmosBridgeStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmosBridgeStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CosmosBridgeStorageSession struct {
	Contract     *CosmosBridgeStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CosmosBridgeStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CosmosBridgeStorageCallerSession struct {
	Contract *CosmosBridgeStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// CosmosBridgeStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CosmosBridgeStorageTransactorSession struct {
	Contract     *CosmosBridgeStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// CosmosBridgeStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type CosmosBridgeStorageRaw struct {
	Contract *CosmosBridgeStorage // Generic contract binding to access the raw methods on
}

// CosmosBridgeStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CosmosBridgeStorageCallerRaw struct {
	Contract *CosmosBridgeStorageCaller // Generic read-only contract binding to access the raw methods on
}

// CosmosBridgeStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CosmosBridgeStorageTransactorRaw struct {
	Contract *CosmosBridgeStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCosmosBridgeStorage creates a new instance of CosmosBridgeStorage, bound to a specific deployed contract.
func NewCosmosBridgeStorage(address common.Address, backend bind.ContractBackend) (*CosmosBridgeStorage, error) {
	contract, err := bindCosmosBridgeStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CosmosBridgeStorage{CosmosBridgeStorageCaller: CosmosBridgeStorageCaller{contract: contract}, CosmosBridgeStorageTransactor: CosmosBridgeStorageTransactor{contract: contract}, CosmosBridgeStorageFilterer: CosmosBridgeStorageFilterer{contract: contract}}, nil
}

// NewCosmosBridgeStorageCaller creates a new read-only instance of CosmosBridgeStorage, bound to a specific deployed contract.
func NewCosmosBridgeStorageCaller(address common.Address, caller bind.ContractCaller) (*CosmosBridgeStorageCaller, error) {
	contract, err := bindCosmosBridgeStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CosmosBridgeStorageCaller{contract: contract}, nil
}

// NewCosmosBridgeStorageTransactor creates a new write-only instance of CosmosBridgeStorage, bound to a specific deployed contract.
func NewCosmosBridgeStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*CosmosBridgeStorageTransactor, error) {
	contract, err := bindCosmosBridgeStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CosmosBridgeStorageTransactor{contract: contract}, nil
}

// NewCosmosBridgeStorageFilterer creates a new log filterer instance of CosmosBridgeStorage, bound to a specific deployed contract.
func NewCosmosBridgeStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*CosmosBridgeStorageFilterer, error) {
	contract, err := bindCosmosBridgeStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CosmosBridgeStorageFilterer{contract: contract}, nil
}

// bindCosmosBridgeStorage binds a generic wrapper to an already deployed contract.
func bindCosmosBridgeStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CosmosBridgeStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosmosBridgeStorage *CosmosBridgeStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CosmosBridgeStorage.Contract.CosmosBridgeStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosmosBridgeStorage *CosmosBridgeStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmosBridgeStorage.Contract.CosmosBridgeStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosmosBridgeStorage *CosmosBridgeStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosmosBridgeStorage.Contract.CosmosBridgeStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosmosBridgeStorage *CosmosBridgeStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CosmosBridgeStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosmosBridgeStorage *CosmosBridgeStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmosBridgeStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosmosBridgeStorage *CosmosBridgeStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosmosBridgeStorage.Contract.contract.Transact(opts, method, params...)
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_CosmosBridgeStorage *CosmosBridgeStorageCaller) BridgeBank(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmosBridgeStorage.contract.Call(opts, &out, "bridgeBank")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_CosmosBridgeStorage *CosmosBridgeStorageSession) BridgeBank() (common.Address, error) {
	return _CosmosBridgeStorage.Contract.BridgeBank(&_CosmosBridgeStorage.CallOpts)
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_CosmosBridgeStorage *CosmosBridgeStorageCallerSession) BridgeBank() (common.Address, error) {
	return _CosmosBridgeStorage.Contract.BridgeBank(&_CosmosBridgeStorage.CallOpts)
}

// HasBridgeBank is a free data retrieval call binding the contract method 0x69294a4e.
//
// Solidity: function hasBridgeBank() view returns(bool)
func (_CosmosBridgeStorage *CosmosBridgeStorageCaller) HasBridgeBank(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CosmosBridgeStorage.contract.Call(opts, &out, "hasBridgeBank")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasBridgeBank is a free data retrieval call binding the contract method 0x69294a4e.
//
// Solidity: function hasBridgeBank() view returns(bool)
func (_CosmosBridgeStorage *CosmosBridgeStorageSession) HasBridgeBank() (bool, error) {
	return _CosmosBridgeStorage.Contract.HasBridgeBank(&_CosmosBridgeStorage.CallOpts)
}

// HasBridgeBank is a free data retrieval call binding the contract method 0x69294a4e.
//
// Solidity: function hasBridgeBank() view returns(bool)
func (_CosmosBridgeStorage *CosmosBridgeStorageCallerSession) HasBridgeBank() (bool, error) {
	return _CosmosBridgeStorage.Contract.HasBridgeBank(&_CosmosBridgeStorage.CallOpts)
}
