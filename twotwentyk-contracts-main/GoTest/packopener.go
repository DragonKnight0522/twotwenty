// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// PackopenerMetaData contains all meta data concerning the Packopener contract.
var PackopenerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_cardPackAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_categoryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_yearAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dayMonthAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_craftingCardAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_triggerContract\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"_tokenCategories\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_tokenYears\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_tokenDayMonths\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_tokenCraftingCards\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_tokenTriggers\",\"type\":\"string[]\"}],\"name\":\"openPack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PackopenerABI is the input ABI used to generate the binding from.
// Deprecated: Use PackopenerMetaData.ABI instead.
var PackopenerABI = PackopenerMetaData.ABI

// Packopener is an auto generated Go binding around an Ethereum contract.
type Packopener struct {
	PackopenerCaller     // Read-only binding to the contract
	PackopenerTransactor // Write-only binding to the contract
	PackopenerFilterer   // Log filterer for contract events
}

// PackopenerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PackopenerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PackopenerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PackopenerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PackopenerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PackopenerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PackopenerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PackopenerSession struct {
	Contract     *Packopener       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PackopenerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PackopenerCallerSession struct {
	Contract *PackopenerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PackopenerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PackopenerTransactorSession struct {
	Contract     *PackopenerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PackopenerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PackopenerRaw struct {
	Contract *Packopener // Generic contract binding to access the raw methods on
}

// PackopenerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PackopenerCallerRaw struct {
	Contract *PackopenerCaller // Generic read-only contract binding to access the raw methods on
}

// PackopenerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PackopenerTransactorRaw struct {
	Contract *PackopenerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPackopener creates a new instance of Packopener, bound to a specific deployed contract.
func NewPackopener(address common.Address, backend bind.ContractBackend) (*Packopener, error) {
	contract, err := bindPackopener(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Packopener{PackopenerCaller: PackopenerCaller{contract: contract}, PackopenerTransactor: PackopenerTransactor{contract: contract}, PackopenerFilterer: PackopenerFilterer{contract: contract}}, nil
}

// NewPackopenerCaller creates a new read-only instance of Packopener, bound to a specific deployed contract.
func NewPackopenerCaller(address common.Address, caller bind.ContractCaller) (*PackopenerCaller, error) {
	contract, err := bindPackopener(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PackopenerCaller{contract: contract}, nil
}

// NewPackopenerTransactor creates a new write-only instance of Packopener, bound to a specific deployed contract.
func NewPackopenerTransactor(address common.Address, transactor bind.ContractTransactor) (*PackopenerTransactor, error) {
	contract, err := bindPackopener(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PackopenerTransactor{contract: contract}, nil
}

// NewPackopenerFilterer creates a new log filterer instance of Packopener, bound to a specific deployed contract.
func NewPackopenerFilterer(address common.Address, filterer bind.ContractFilterer) (*PackopenerFilterer, error) {
	contract, err := bindPackopener(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PackopenerFilterer{contract: contract}, nil
}

// bindPackopener binds a generic wrapper to an already deployed contract.
func bindPackopener(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PackopenerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Packopener *PackopenerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Packopener.Contract.PackopenerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Packopener *PackopenerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Packopener.Contract.PackopenerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Packopener *PackopenerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Packopener.Contract.PackopenerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Packopener *PackopenerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Packopener.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Packopener *PackopenerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Packopener.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Packopener *PackopenerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Packopener.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Packopener *PackopenerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Packopener.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Packopener *PackopenerSession) Owner() (common.Address, error) {
	return _Packopener.Contract.Owner(&_Packopener.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Packopener *PackopenerCallerSession) Owner() (common.Address, error) {
	return _Packopener.Contract.Owner(&_Packopener.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Packopener *PackopenerTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _Packopener.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Packopener *PackopenerSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Packopener.Contract.ChangeAdmin(&_Packopener.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Packopener *PackopenerTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Packopener.Contract.ChangeAdmin(&_Packopener.TransactOpts, _newAdmin)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Packopener *PackopenerTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Packopener.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Packopener *PackopenerSession) Initialize() (*types.Transaction, error) {
	return _Packopener.Contract.Initialize(&_Packopener.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Packopener *PackopenerTransactorSession) Initialize() (*types.Transaction, error) {
	return _Packopener.Contract.Initialize(&_Packopener.TransactOpts)
}

// OpenPack is a paid mutator transaction binding the contract method 0x3874820b.
//
// Solidity: function openPack(uint256 _tokenId, address _cardPackAddress, address _categoryAddress, address _yearAddress, address _dayMonthAddress, address _craftingCardAddress, address _triggerContract, string[] _tokenCategories, string[] _tokenYears, string[] _tokenDayMonths, string[] _tokenCraftingCards, string[] _tokenTriggers) returns()
func (_Packopener *PackopenerTransactor) OpenPack(opts *bind.TransactOpts, _tokenId *big.Int, _cardPackAddress common.Address, _categoryAddress common.Address, _yearAddress common.Address, _dayMonthAddress common.Address, _craftingCardAddress common.Address, _triggerContract common.Address, _tokenCategories []string, _tokenYears []string, _tokenDayMonths []string, _tokenCraftingCards []string, _tokenTriggers []string) (*types.Transaction, error) {
	return _Packopener.contract.Transact(opts, "openPack", _tokenId, _cardPackAddress, _categoryAddress, _yearAddress, _dayMonthAddress, _craftingCardAddress, _triggerContract, _tokenCategories, _tokenYears, _tokenDayMonths, _tokenCraftingCards, _tokenTriggers)
}

// OpenPack is a paid mutator transaction binding the contract method 0x3874820b.
//
// Solidity: function openPack(uint256 _tokenId, address _cardPackAddress, address _categoryAddress, address _yearAddress, address _dayMonthAddress, address _craftingCardAddress, address _triggerContract, string[] _tokenCategories, string[] _tokenYears, string[] _tokenDayMonths, string[] _tokenCraftingCards, string[] _tokenTriggers) returns()
func (_Packopener *PackopenerSession) OpenPack(_tokenId *big.Int, _cardPackAddress common.Address, _categoryAddress common.Address, _yearAddress common.Address, _dayMonthAddress common.Address, _craftingCardAddress common.Address, _triggerContract common.Address, _tokenCategories []string, _tokenYears []string, _tokenDayMonths []string, _tokenCraftingCards []string, _tokenTriggers []string) (*types.Transaction, error) {
	return _Packopener.Contract.OpenPack(&_Packopener.TransactOpts, _tokenId, _cardPackAddress, _categoryAddress, _yearAddress, _dayMonthAddress, _craftingCardAddress, _triggerContract, _tokenCategories, _tokenYears, _tokenDayMonths, _tokenCraftingCards, _tokenTriggers)
}

// OpenPack is a paid mutator transaction binding the contract method 0x3874820b.
//
// Solidity: function openPack(uint256 _tokenId, address _cardPackAddress, address _categoryAddress, address _yearAddress, address _dayMonthAddress, address _craftingCardAddress, address _triggerContract, string[] _tokenCategories, string[] _tokenYears, string[] _tokenDayMonths, string[] _tokenCraftingCards, string[] _tokenTriggers) returns()
func (_Packopener *PackopenerTransactorSession) OpenPack(_tokenId *big.Int, _cardPackAddress common.Address, _categoryAddress common.Address, _yearAddress common.Address, _dayMonthAddress common.Address, _craftingCardAddress common.Address, _triggerContract common.Address, _tokenCategories []string, _tokenYears []string, _tokenDayMonths []string, _tokenCraftingCards []string, _tokenTriggers []string) (*types.Transaction, error) {
	return _Packopener.Contract.OpenPack(&_Packopener.TransactOpts, _tokenId, _cardPackAddress, _categoryAddress, _yearAddress, _dayMonthAddress, _craftingCardAddress, _triggerContract, _tokenCategories, _tokenYears, _tokenDayMonths, _tokenCraftingCards, _tokenTriggers)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Packopener *PackopenerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Packopener.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Packopener *PackopenerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Packopener.Contract.RenounceOwnership(&_Packopener.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Packopener *PackopenerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Packopener.Contract.RenounceOwnership(&_Packopener.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Packopener *PackopenerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Packopener.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Packopener *PackopenerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Packopener.Contract.TransferOwnership(&_Packopener.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Packopener *PackopenerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Packopener.Contract.TransferOwnership(&_Packopener.TransactOpts, newOwner)
}

// PackopenerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Packopener contract.
type PackopenerOwnershipTransferredIterator struct {
	Event *PackopenerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PackopenerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackopenerOwnershipTransferred)
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
		it.Event = new(PackopenerOwnershipTransferred)
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
func (it *PackopenerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackopenerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackopenerOwnershipTransferred represents a OwnershipTransferred event raised by the Packopener contract.
type PackopenerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Packopener *PackopenerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PackopenerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Packopener.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PackopenerOwnershipTransferredIterator{contract: _Packopener.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Packopener *PackopenerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PackopenerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Packopener.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackopenerOwnershipTransferred)
				if err := _Packopener.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Packopener *PackopenerFilterer) ParseOwnershipTransferred(log types.Log) (*PackopenerOwnershipTransferred, error) {
	event := new(PackopenerOwnershipTransferred)
	if err := _Packopener.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
