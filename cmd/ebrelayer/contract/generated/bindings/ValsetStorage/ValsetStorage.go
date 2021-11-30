// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ValsetStorage

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

// ValsetStorageMetaData contains all meta data concerning the ValsetStorage contract.
var ValsetStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"currentValsetVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"powers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ValsetStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use ValsetStorageMetaData.ABI instead.
var ValsetStorageABI = ValsetStorageMetaData.ABI

// ValsetStorage is an auto generated Go binding around an Ethereum contract.
type ValsetStorage struct {
	ValsetStorageCaller     // Read-only binding to the contract
	ValsetStorageTransactor // Write-only binding to the contract
	ValsetStorageFilterer   // Log filterer for contract events
}

// ValsetStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValsetStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValsetStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValsetStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValsetStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValsetStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValsetStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValsetStorageSession struct {
	Contract     *ValsetStorage    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValsetStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValsetStorageCallerSession struct {
	Contract *ValsetStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ValsetStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValsetStorageTransactorSession struct {
	Contract     *ValsetStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ValsetStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValsetStorageRaw struct {
	Contract *ValsetStorage // Generic contract binding to access the raw methods on
}

// ValsetStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValsetStorageCallerRaw struct {
	Contract *ValsetStorageCaller // Generic read-only contract binding to access the raw methods on
}

// ValsetStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValsetStorageTransactorRaw struct {
	Contract *ValsetStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValsetStorage creates a new instance of ValsetStorage, bound to a specific deployed contract.
func NewValsetStorage(address common.Address, backend bind.ContractBackend) (*ValsetStorage, error) {
	contract, err := bindValsetStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValsetStorage{ValsetStorageCaller: ValsetStorageCaller{contract: contract}, ValsetStorageTransactor: ValsetStorageTransactor{contract: contract}, ValsetStorageFilterer: ValsetStorageFilterer{contract: contract}}, nil
}

// NewValsetStorageCaller creates a new read-only instance of ValsetStorage, bound to a specific deployed contract.
func NewValsetStorageCaller(address common.Address, caller bind.ContractCaller) (*ValsetStorageCaller, error) {
	contract, err := bindValsetStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValsetStorageCaller{contract: contract}, nil
}

// NewValsetStorageTransactor creates a new write-only instance of ValsetStorage, bound to a specific deployed contract.
func NewValsetStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*ValsetStorageTransactor, error) {
	contract, err := bindValsetStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValsetStorageTransactor{contract: contract}, nil
}

// NewValsetStorageFilterer creates a new log filterer instance of ValsetStorage, bound to a specific deployed contract.
func NewValsetStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*ValsetStorageFilterer, error) {
	contract, err := bindValsetStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValsetStorageFilterer{contract: contract}, nil
}

// bindValsetStorage binds a generic wrapper to an already deployed contract.
func bindValsetStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValsetStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValsetStorage *ValsetStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValsetStorage.Contract.ValsetStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValsetStorage *ValsetStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValsetStorage.Contract.ValsetStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValsetStorage *ValsetStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValsetStorage.Contract.ValsetStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValsetStorage *ValsetStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValsetStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValsetStorage *ValsetStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValsetStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValsetStorage *ValsetStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValsetStorage.Contract.contract.Transact(opts, method, params...)
}

// CurrentValsetVersion is a free data retrieval call binding the contract method 0x8d56c37d.
//
// Solidity: function currentValsetVersion() view returns(uint256)
func (_ValsetStorage *ValsetStorageCaller) CurrentValsetVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValsetStorage.contract.Call(opts, &out, "currentValsetVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentValsetVersion is a free data retrieval call binding the contract method 0x8d56c37d.
//
// Solidity: function currentValsetVersion() view returns(uint256)
func (_ValsetStorage *ValsetStorageSession) CurrentValsetVersion() (*big.Int, error) {
	return _ValsetStorage.Contract.CurrentValsetVersion(&_ValsetStorage.CallOpts)
}

// CurrentValsetVersion is a free data retrieval call binding the contract method 0x8d56c37d.
//
// Solidity: function currentValsetVersion() view returns(uint256)
func (_ValsetStorage *ValsetStorageCallerSession) CurrentValsetVersion() (*big.Int, error) {
	return _ValsetStorage.Contract.CurrentValsetVersion(&_ValsetStorage.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_ValsetStorage *ValsetStorageCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValsetStorage.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_ValsetStorage *ValsetStorageSession) Operator() (common.Address, error) {
	return _ValsetStorage.Contract.Operator(&_ValsetStorage.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_ValsetStorage *ValsetStorageCallerSession) Operator() (common.Address, error) {
	return _ValsetStorage.Contract.Operator(&_ValsetStorage.CallOpts)
}

// Powers is a free data retrieval call binding the contract method 0x850f43dd.
//
// Solidity: function powers(address , uint256 ) view returns(uint256)
func (_ValsetStorage *ValsetStorageCaller) Powers(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ValsetStorage.contract.Call(opts, &out, "powers", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Powers is a free data retrieval call binding the contract method 0x850f43dd.
//
// Solidity: function powers(address , uint256 ) view returns(uint256)
func (_ValsetStorage *ValsetStorageSession) Powers(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _ValsetStorage.Contract.Powers(&_ValsetStorage.CallOpts, arg0, arg1)
}

// Powers is a free data retrieval call binding the contract method 0x850f43dd.
//
// Solidity: function powers(address , uint256 ) view returns(uint256)
func (_ValsetStorage *ValsetStorageCallerSession) Powers(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _ValsetStorage.Contract.Powers(&_ValsetStorage.CallOpts, arg0, arg1)
}

// TotalPower is a free data retrieval call binding the contract method 0xdb3ad22c.
//
// Solidity: function totalPower() view returns(uint256)
func (_ValsetStorage *ValsetStorageCaller) TotalPower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValsetStorage.contract.Call(opts, &out, "totalPower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPower is a free data retrieval call binding the contract method 0xdb3ad22c.
//
// Solidity: function totalPower() view returns(uint256)
func (_ValsetStorage *ValsetStorageSession) TotalPower() (*big.Int, error) {
	return _ValsetStorage.Contract.TotalPower(&_ValsetStorage.CallOpts)
}

// TotalPower is a free data retrieval call binding the contract method 0xdb3ad22c.
//
// Solidity: function totalPower() view returns(uint256)
func (_ValsetStorage *ValsetStorageCallerSession) TotalPower() (*big.Int, error) {
	return _ValsetStorage.Contract.TotalPower(&_ValsetStorage.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_ValsetStorage *ValsetStorageCaller) ValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValsetStorage.contract.Call(opts, &out, "validatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_ValsetStorage *ValsetStorageSession) ValidatorCount() (*big.Int, error) {
	return _ValsetStorage.Contract.ValidatorCount(&_ValsetStorage.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_ValsetStorage *ValsetStorageCallerSession) ValidatorCount() (*big.Int, error) {
	return _ValsetStorage.Contract.ValidatorCount(&_ValsetStorage.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0x45aaf18c.
//
// Solidity: function validators(address , uint256 ) view returns(bool)
func (_ValsetStorage *ValsetStorageCaller) Validators(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _ValsetStorage.contract.Call(opts, &out, "validators", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Validators is a free data retrieval call binding the contract method 0x45aaf18c.
//
// Solidity: function validators(address , uint256 ) view returns(bool)
func (_ValsetStorage *ValsetStorageSession) Validators(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _ValsetStorage.Contract.Validators(&_ValsetStorage.CallOpts, arg0, arg1)
}

// Validators is a free data retrieval call binding the contract method 0x45aaf18c.
//
// Solidity: function validators(address , uint256 ) view returns(bool)
func (_ValsetStorage *ValsetStorageCallerSession) Validators(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _ValsetStorage.Contract.Validators(&_ValsetStorage.CallOpts, arg0, arg1)
}
