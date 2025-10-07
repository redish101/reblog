// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package copyright

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
	_ = abi.ConvertType
)

// ReblogCopyrightPost is an auto generated low-level Go binding around an user-defined struct.
type ReblogCopyrightPost struct {
	Title   string
	IpfsURL string
}

// CopyrightMetaData contains all meta data concerning the Copyright contract.
var CopyrightMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"siteName\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"slug\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ipfsURL\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"}],\"name\":\"PostAddedOrUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"}],\"name\":\"SiteNameUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"slug\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipfsURL\",\"type\":\"string\"}],\"name\":\"addOrUpdatePost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"slug\",\"type\":\"string\"}],\"name\":\"getPost\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipfsURL\",\"type\":\"string\"}],\"internalType\":\"structReblogCopyright.Post\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSiteName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"siteName\",\"type\":\"string\"}],\"name\":\"setSiteName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CopyrightABI is the input ABI used to generate the binding from.
// Deprecated: Use CopyrightMetaData.ABI instead.
var CopyrightABI = CopyrightMetaData.ABI

// Copyright is an auto generated Go binding around an Ethereum contract.
type Copyright struct {
	CopyrightCaller     // Read-only binding to the contract
	CopyrightTransactor // Write-only binding to the contract
	CopyrightFilterer   // Log filterer for contract events
}

// CopyrightCaller is an auto generated read-only Go binding around an Ethereum contract.
type CopyrightCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CopyrightTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CopyrightTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CopyrightFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CopyrightFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CopyrightSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CopyrightSession struct {
	Contract     *Copyright        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CopyrightCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CopyrightCallerSession struct {
	Contract *CopyrightCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CopyrightTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CopyrightTransactorSession struct {
	Contract     *CopyrightTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CopyrightRaw is an auto generated low-level Go binding around an Ethereum contract.
type CopyrightRaw struct {
	Contract *Copyright // Generic contract binding to access the raw methods on
}

// CopyrightCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CopyrightCallerRaw struct {
	Contract *CopyrightCaller // Generic read-only contract binding to access the raw methods on
}

// CopyrightTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CopyrightTransactorRaw struct {
	Contract *CopyrightTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCopyright creates a new instance of Copyright, bound to a specific deployed contract.
func NewCopyright(address common.Address, backend bind.ContractBackend) (*Copyright, error) {
	contract, err := bindCopyright(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Copyright{CopyrightCaller: CopyrightCaller{contract: contract}, CopyrightTransactor: CopyrightTransactor{contract: contract}, CopyrightFilterer: CopyrightFilterer{contract: contract}}, nil
}

// NewCopyrightCaller creates a new read-only instance of Copyright, bound to a specific deployed contract.
func NewCopyrightCaller(address common.Address, caller bind.ContractCaller) (*CopyrightCaller, error) {
	contract, err := bindCopyright(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CopyrightCaller{contract: contract}, nil
}

// NewCopyrightTransactor creates a new write-only instance of Copyright, bound to a specific deployed contract.
func NewCopyrightTransactor(address common.Address, transactor bind.ContractTransactor) (*CopyrightTransactor, error) {
	contract, err := bindCopyright(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CopyrightTransactor{contract: contract}, nil
}

// NewCopyrightFilterer creates a new log filterer instance of Copyright, bound to a specific deployed contract.
func NewCopyrightFilterer(address common.Address, filterer bind.ContractFilterer) (*CopyrightFilterer, error) {
	contract, err := bindCopyright(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CopyrightFilterer{contract: contract}, nil
}

// bindCopyright binds a generic wrapper to an already deployed contract.
func bindCopyright(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CopyrightMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Copyright *CopyrightRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Copyright.Contract.CopyrightCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Copyright *CopyrightRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Copyright.Contract.CopyrightTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Copyright *CopyrightRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Copyright.Contract.CopyrightTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Copyright *CopyrightCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Copyright.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Copyright *CopyrightTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Copyright.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Copyright *CopyrightTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Copyright.Contract.contract.Transact(opts, method, params...)
}

// GetPost is a free data retrieval call binding the contract method 0xdf4a259b.
//
// Solidity: function getPost(string slug) view returns((string,string))
func (_Copyright *CopyrightCaller) GetPost(opts *bind.CallOpts, slug string) (ReblogCopyrightPost, error) {
	var out []interface{}
	err := _Copyright.contract.Call(opts, &out, "getPost", slug)

	if err != nil {
		return *new(ReblogCopyrightPost), err
	}

	out0 := *abi.ConvertType(out[0], new(ReblogCopyrightPost)).(*ReblogCopyrightPost)

	return out0, err

}

// GetPost is a free data retrieval call binding the contract method 0xdf4a259b.
//
// Solidity: function getPost(string slug) view returns((string,string))
func (_Copyright *CopyrightSession) GetPost(slug string) (ReblogCopyrightPost, error) {
	return _Copyright.Contract.GetPost(&_Copyright.CallOpts, slug)
}

// GetPost is a free data retrieval call binding the contract method 0xdf4a259b.
//
// Solidity: function getPost(string slug) view returns((string,string))
func (_Copyright *CopyrightCallerSession) GetPost(slug string) (ReblogCopyrightPost, error) {
	return _Copyright.Contract.GetPost(&_Copyright.CallOpts, slug)
}

// GetSiteName is a free data retrieval call binding the contract method 0xf10e27fa.
//
// Solidity: function getSiteName() view returns(string)
func (_Copyright *CopyrightCaller) GetSiteName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Copyright.contract.Call(opts, &out, "getSiteName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetSiteName is a free data retrieval call binding the contract method 0xf10e27fa.
//
// Solidity: function getSiteName() view returns(string)
func (_Copyright *CopyrightSession) GetSiteName() (string, error) {
	return _Copyright.Contract.GetSiteName(&_Copyright.CallOpts)
}

// GetSiteName is a free data retrieval call binding the contract method 0xf10e27fa.
//
// Solidity: function getSiteName() view returns(string)
func (_Copyright *CopyrightCallerSession) GetSiteName() (string, error) {
	return _Copyright.Contract.GetSiteName(&_Copyright.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Copyright *CopyrightCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Copyright.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Copyright *CopyrightSession) Owner() (common.Address, error) {
	return _Copyright.Contract.Owner(&_Copyright.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Copyright *CopyrightCallerSession) Owner() (common.Address, error) {
	return _Copyright.Contract.Owner(&_Copyright.CallOpts)
}

// AddOrUpdatePost is a paid mutator transaction binding the contract method 0xc3bf4320.
//
// Solidity: function addOrUpdatePost(string slug, string title, string ipfsURL) returns()
func (_Copyright *CopyrightTransactor) AddOrUpdatePost(opts *bind.TransactOpts, slug string, title string, ipfsURL string) (*types.Transaction, error) {
	return _Copyright.contract.Transact(opts, "addOrUpdatePost", slug, title, ipfsURL)
}

// AddOrUpdatePost is a paid mutator transaction binding the contract method 0xc3bf4320.
//
// Solidity: function addOrUpdatePost(string slug, string title, string ipfsURL) returns()
func (_Copyright *CopyrightSession) AddOrUpdatePost(slug string, title string, ipfsURL string) (*types.Transaction, error) {
	return _Copyright.Contract.AddOrUpdatePost(&_Copyright.TransactOpts, slug, title, ipfsURL)
}

// AddOrUpdatePost is a paid mutator transaction binding the contract method 0xc3bf4320.
//
// Solidity: function addOrUpdatePost(string slug, string title, string ipfsURL) returns()
func (_Copyright *CopyrightTransactorSession) AddOrUpdatePost(slug string, title string, ipfsURL string) (*types.Transaction, error) {
	return _Copyright.Contract.AddOrUpdatePost(&_Copyright.TransactOpts, slug, title, ipfsURL)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Copyright *CopyrightTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Copyright.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Copyright *CopyrightSession) RenounceOwnership() (*types.Transaction, error) {
	return _Copyright.Contract.RenounceOwnership(&_Copyright.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Copyright *CopyrightTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Copyright.Contract.RenounceOwnership(&_Copyright.TransactOpts)
}

// SetSiteName is a paid mutator transaction binding the contract method 0x09ace394.
//
// Solidity: function setSiteName(string siteName) returns()
func (_Copyright *CopyrightTransactor) SetSiteName(opts *bind.TransactOpts, siteName string) (*types.Transaction, error) {
	return _Copyright.contract.Transact(opts, "setSiteName", siteName)
}

// SetSiteName is a paid mutator transaction binding the contract method 0x09ace394.
//
// Solidity: function setSiteName(string siteName) returns()
func (_Copyright *CopyrightSession) SetSiteName(siteName string) (*types.Transaction, error) {
	return _Copyright.Contract.SetSiteName(&_Copyright.TransactOpts, siteName)
}

// SetSiteName is a paid mutator transaction binding the contract method 0x09ace394.
//
// Solidity: function setSiteName(string siteName) returns()
func (_Copyright *CopyrightTransactorSession) SetSiteName(siteName string) (*types.Transaction, error) {
	return _Copyright.Contract.SetSiteName(&_Copyright.TransactOpts, siteName)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Copyright *CopyrightTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Copyright.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Copyright *CopyrightSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Copyright.Contract.TransferOwnership(&_Copyright.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Copyright *CopyrightTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Copyright.Contract.TransferOwnership(&_Copyright.TransactOpts, newOwner)
}

// CopyrightOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Copyright contract.
type CopyrightOwnershipTransferredIterator struct {
	Event *CopyrightOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CopyrightOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CopyrightOwnershipTransferred)
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
		it.Event = new(CopyrightOwnershipTransferred)
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
func (it *CopyrightOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CopyrightOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CopyrightOwnershipTransferred represents a OwnershipTransferred event raised by the Copyright contract.
type CopyrightOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Copyright *CopyrightFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CopyrightOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Copyright.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CopyrightOwnershipTransferredIterator{contract: _Copyright.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Copyright *CopyrightFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CopyrightOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Copyright.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CopyrightOwnershipTransferred)
				if err := _Copyright.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Copyright *CopyrightFilterer) ParseOwnershipTransferred(log types.Log) (*CopyrightOwnershipTransferred, error) {
	event := new(CopyrightOwnershipTransferred)
	if err := _Copyright.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CopyrightPostAddedOrUpdatedIterator is returned from FilterPostAddedOrUpdated and is used to iterate over the raw logs and unpacked data for PostAddedOrUpdated events raised by the Copyright contract.
type CopyrightPostAddedOrUpdatedIterator struct {
	Event *CopyrightPostAddedOrUpdated // Event containing the contract specifics and raw log

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
func (it *CopyrightPostAddedOrUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CopyrightPostAddedOrUpdated)
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
		it.Event = new(CopyrightPostAddedOrUpdated)
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
func (it *CopyrightPostAddedOrUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CopyrightPostAddedOrUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CopyrightPostAddedOrUpdated represents a PostAddedOrUpdated event raised by the Copyright contract.
type CopyrightPostAddedOrUpdated struct {
	Slug    string
	Title   string
	IpfsURL string
	Author  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPostAddedOrUpdated is a free log retrieval operation binding the contract event 0x6a7e9f408a9934436149c1bdf7f3f36f85076e8fbce4b4f681f5645a907da0a6.
//
// Solidity: event PostAddedOrUpdated(string slug, string title, string ipfsURL, address author)
func (_Copyright *CopyrightFilterer) FilterPostAddedOrUpdated(opts *bind.FilterOpts) (*CopyrightPostAddedOrUpdatedIterator, error) {

	logs, sub, err := _Copyright.contract.FilterLogs(opts, "PostAddedOrUpdated")
	if err != nil {
		return nil, err
	}
	return &CopyrightPostAddedOrUpdatedIterator{contract: _Copyright.contract, event: "PostAddedOrUpdated", logs: logs, sub: sub}, nil
}

// WatchPostAddedOrUpdated is a free log subscription operation binding the contract event 0x6a7e9f408a9934436149c1bdf7f3f36f85076e8fbce4b4f681f5645a907da0a6.
//
// Solidity: event PostAddedOrUpdated(string slug, string title, string ipfsURL, address author)
func (_Copyright *CopyrightFilterer) WatchPostAddedOrUpdated(opts *bind.WatchOpts, sink chan<- *CopyrightPostAddedOrUpdated) (event.Subscription, error) {

	logs, sub, err := _Copyright.contract.WatchLogs(opts, "PostAddedOrUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CopyrightPostAddedOrUpdated)
				if err := _Copyright.contract.UnpackLog(event, "PostAddedOrUpdated", log); err != nil {
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

// ParsePostAddedOrUpdated is a log parse operation binding the contract event 0x6a7e9f408a9934436149c1bdf7f3f36f85076e8fbce4b4f681f5645a907da0a6.
//
// Solidity: event PostAddedOrUpdated(string slug, string title, string ipfsURL, address author)
func (_Copyright *CopyrightFilterer) ParsePostAddedOrUpdated(log types.Log) (*CopyrightPostAddedOrUpdated, error) {
	event := new(CopyrightPostAddedOrUpdated)
	if err := _Copyright.contract.UnpackLog(event, "PostAddedOrUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CopyrightSiteNameUpdatedIterator is returned from FilterSiteNameUpdated and is used to iterate over the raw logs and unpacked data for SiteNameUpdated events raised by the Copyright contract.
type CopyrightSiteNameUpdatedIterator struct {
	Event *CopyrightSiteNameUpdated // Event containing the contract specifics and raw log

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
func (it *CopyrightSiteNameUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CopyrightSiteNameUpdated)
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
		it.Event = new(CopyrightSiteNameUpdated)
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
func (it *CopyrightSiteNameUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CopyrightSiteNameUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CopyrightSiteNameUpdated represents a SiteNameUpdated event raised by the Copyright contract.
type CopyrightSiteNameUpdated struct {
	NewName string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSiteNameUpdated is a free log retrieval operation binding the contract event 0xd9b10a8bdbf5acfdf5618dc8aa7b3658ef03094a9d5a8cfb83b8fa73d88de099.
//
// Solidity: event SiteNameUpdated(string newName)
func (_Copyright *CopyrightFilterer) FilterSiteNameUpdated(opts *bind.FilterOpts) (*CopyrightSiteNameUpdatedIterator, error) {

	logs, sub, err := _Copyright.contract.FilterLogs(opts, "SiteNameUpdated")
	if err != nil {
		return nil, err
	}
	return &CopyrightSiteNameUpdatedIterator{contract: _Copyright.contract, event: "SiteNameUpdated", logs: logs, sub: sub}, nil
}

// WatchSiteNameUpdated is a free log subscription operation binding the contract event 0xd9b10a8bdbf5acfdf5618dc8aa7b3658ef03094a9d5a8cfb83b8fa73d88de099.
//
// Solidity: event SiteNameUpdated(string newName)
func (_Copyright *CopyrightFilterer) WatchSiteNameUpdated(opts *bind.WatchOpts, sink chan<- *CopyrightSiteNameUpdated) (event.Subscription, error) {

	logs, sub, err := _Copyright.contract.WatchLogs(opts, "SiteNameUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CopyrightSiteNameUpdated)
				if err := _Copyright.contract.UnpackLog(event, "SiteNameUpdated", log); err != nil {
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

// ParseSiteNameUpdated is a log parse operation binding the contract event 0xd9b10a8bdbf5acfdf5618dc8aa7b3658ef03094a9d5a8cfb83b8fa73d88de099.
//
// Solidity: event SiteNameUpdated(string newName)
func (_Copyright *CopyrightFilterer) ParseSiteNameUpdated(log types.Log) (*CopyrightSiteNameUpdated, error) {
	event := new(CopyrightSiteNameUpdated)
	if err := _Copyright.contract.UnpackLog(event, "SiteNameUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
