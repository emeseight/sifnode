// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OracleStorage

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

// OracleStorageMetaData contains all meta data concerning the OracleStorage contract.
var OracleStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"consensusThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cosmosBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// OracleStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleStorageMetaData.ABI instead.
var OracleStorageABI = OracleStorageMetaData.ABI

// OracleStorage is an auto generated Go binding around an Ethereum contract.
type OracleStorage struct {
	OracleStorageCaller     // Read-only binding to the contract
	OracleStorageTransactor // Write-only binding to the contract
	OracleStorageFilterer   // Log filterer for contract events
}

// OracleStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleStorageSession struct {
	Contract     *OracleStorage    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleStorageCallerSession struct {
	Contract *OracleStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// OracleStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleStorageTransactorSession struct {
	Contract     *OracleStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// OracleStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleStorageRaw struct {
	Contract *OracleStorage // Generic contract binding to access the raw methods on
}

// OracleStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleStorageCallerRaw struct {
	Contract *OracleStorageCaller // Generic read-only contract binding to access the raw methods on
}

// OracleStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleStorageTransactorRaw struct {
	Contract *OracleStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleStorage creates a new instance of OracleStorage, bound to a specific deployed contract.
func NewOracleStorage(address common.Address, backend bind.ContractBackend) (*OracleStorage, error) {
	contract, err := bindOracleStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleStorage{OracleStorageCaller: OracleStorageCaller{contract: contract}, OracleStorageTransactor: OracleStorageTransactor{contract: contract}, OracleStorageFilterer: OracleStorageFilterer{contract: contract}}, nil
}

// NewOracleStorageCaller creates a new read-only instance of OracleStorage, bound to a specific deployed contract.
func NewOracleStorageCaller(address common.Address, caller bind.ContractCaller) (*OracleStorageCaller, error) {
	contract, err := bindOracleStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleStorageCaller{contract: contract}, nil
}

// NewOracleStorageTransactor creates a new write-only instance of OracleStorage, bound to a specific deployed contract.
func NewOracleStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleStorageTransactor, error) {
	contract, err := bindOracleStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleStorageTransactor{contract: contract}, nil
}

// NewOracleStorageFilterer creates a new log filterer instance of OracleStorage, bound to a specific deployed contract.
func NewOracleStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleStorageFilterer, error) {
	contract, err := bindOracleStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleStorageFilterer{contract: contract}, nil
}

// bindOracleStorage binds a generic wrapper to an already deployed contract.
func bindOracleStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleStorage *OracleStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleStorage.Contract.OracleStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleStorage *OracleStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleStorage.Contract.OracleStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleStorage *OracleStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleStorage.Contract.OracleStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleStorage *OracleStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleStorage *OracleStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleStorage *OracleStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleStorage.Contract.contract.Transact(opts, method, params...)
}

// ConsensusThreshold is a free data retrieval call binding the contract method 0xf9b0b5b9.
//
// Solidity: function consensusThreshold() view returns(uint256)
func (_OracleStorage *OracleStorageCaller) ConsensusThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleStorage.contract.Call(opts, &out, "consensusThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConsensusThreshold is a free data retrieval call binding the contract method 0xf9b0b5b9.
//
// Solidity: function consensusThreshold() view returns(uint256)
func (_OracleStorage *OracleStorageSession) ConsensusThreshold() (*big.Int, error) {
	return _OracleStorage.Contract.ConsensusThreshold(&_OracleStorage.CallOpts)
}

// ConsensusThreshold is a free data retrieval call binding the contract method 0xf9b0b5b9.
//
// Solidity: function consensusThreshold() view returns(uint256)
func (_OracleStorage *OracleStorageCallerSession) ConsensusThreshold() (*big.Int, error) {
	return _OracleStorage.Contract.ConsensusThreshold(&_OracleStorage.CallOpts)
}

// CosmosBridge is a free data retrieval call binding the contract method 0xb0e9ef71.
//
// Solidity: function cosmosBridge() view returns(address)
func (_OracleStorage *OracleStorageCaller) CosmosBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OracleStorage.contract.Call(opts, &out, "cosmosBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CosmosBridge is a free data retrieval call binding the contract method 0xb0e9ef71.
//
// Solidity: function cosmosBridge() view returns(address)
func (_OracleStorage *OracleStorageSession) CosmosBridge() (common.Address, error) {
	return _OracleStorage.Contract.CosmosBridge(&_OracleStorage.CallOpts)
}

// CosmosBridge is a free data retrieval call binding the contract method 0xb0e9ef71.
//
// Solidity: function cosmosBridge() view returns(address)
func (_OracleStorage *OracleStorageCallerSession) CosmosBridge() (common.Address, error) {
	return _OracleStorage.Contract.CosmosBridge(&_OracleStorage.CallOpts)
}
