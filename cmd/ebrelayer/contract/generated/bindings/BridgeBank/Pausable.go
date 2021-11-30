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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pausers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
