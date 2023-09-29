// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package craftingFactory

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

// CraftingFactoryMetaData contains all meta data concerning the CraftingFactory contract.
var CraftingFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"crafting\",\"type\":\"address\"}],\"name\":\"CraftingCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"adminWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createCrafting\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"crafting\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CraftingFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use CraftingFactoryMetaData.ABI instead.
var CraftingFactoryABI = CraftingFactoryMetaData.ABI

// CraftingFactory is an auto generated Go binding around an Ethereum contract.
type CraftingFactory struct {
	CraftingFactoryCaller     // Read-only binding to the contract
	CraftingFactoryTransactor // Write-only binding to the contract
	CraftingFactoryFilterer   // Log filterer for contract events
}

// CraftingFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CraftingFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CraftingFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CraftingFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CraftingFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CraftingFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CraftingFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CraftingFactorySession struct {
	Contract     *CraftingFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CraftingFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CraftingFactoryCallerSession struct {
	Contract *CraftingFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CraftingFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CraftingFactoryTransactorSession struct {
	Contract     *CraftingFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CraftingFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CraftingFactoryRaw struct {
	Contract *CraftingFactory // Generic contract binding to access the raw methods on
}

// CraftingFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CraftingFactoryCallerRaw struct {
	Contract *CraftingFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// CraftingFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CraftingFactoryTransactorRaw struct {
	Contract *CraftingFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCraftingFactory creates a new instance of CraftingFactory, bound to a specific deployed contract.
func NewCraftingFactory(address common.Address, backend bind.ContractBackend) (*CraftingFactory, error) {
	contract, err := bindCraftingFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CraftingFactory{CraftingFactoryCaller: CraftingFactoryCaller{contract: contract}, CraftingFactoryTransactor: CraftingFactoryTransactor{contract: contract}, CraftingFactoryFilterer: CraftingFactoryFilterer{contract: contract}}, nil
}

// NewCraftingFactoryCaller creates a new read-only instance of CraftingFactory, bound to a specific deployed contract.
func NewCraftingFactoryCaller(address common.Address, caller bind.ContractCaller) (*CraftingFactoryCaller, error) {
	contract, err := bindCraftingFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CraftingFactoryCaller{contract: contract}, nil
}

// NewCraftingFactoryTransactor creates a new write-only instance of CraftingFactory, bound to a specific deployed contract.
func NewCraftingFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*CraftingFactoryTransactor, error) {
	contract, err := bindCraftingFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CraftingFactoryTransactor{contract: contract}, nil
}

// NewCraftingFactoryFilterer creates a new log filterer instance of CraftingFactory, bound to a specific deployed contract.
func NewCraftingFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*CraftingFactoryFilterer, error) {
	contract, err := bindCraftingFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CraftingFactoryFilterer{contract: contract}, nil
}

// bindCraftingFactory binds a generic wrapper to an already deployed contract.
func bindCraftingFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CraftingFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CraftingFactory *CraftingFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CraftingFactory.Contract.CraftingFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CraftingFactory *CraftingFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CraftingFactory.Contract.CraftingFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CraftingFactory *CraftingFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CraftingFactory.Contract.CraftingFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CraftingFactory *CraftingFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CraftingFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CraftingFactory *CraftingFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CraftingFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CraftingFactory *CraftingFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CraftingFactory.Contract.contract.Transact(opts, method, params...)
}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_CraftingFactory *CraftingFactoryCaller) AdminWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CraftingFactory.contract.Call(opts, &out, "adminWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_CraftingFactory *CraftingFactorySession) AdminWallet() (common.Address, error) {
	return _CraftingFactory.Contract.AdminWallet(&_CraftingFactory.CallOpts)
}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_CraftingFactory *CraftingFactoryCallerSession) AdminWallet() (common.Address, error) {
	return _CraftingFactory.Contract.AdminWallet(&_CraftingFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CraftingFactory *CraftingFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CraftingFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CraftingFactory *CraftingFactorySession) Owner() (common.Address, error) {
	return _CraftingFactory.Contract.Owner(&_CraftingFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CraftingFactory *CraftingFactoryCallerSession) Owner() (common.Address, error) {
	return _CraftingFactory.Contract.Owner(&_CraftingFactory.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_CraftingFactory *CraftingFactoryTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _CraftingFactory.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_CraftingFactory *CraftingFactorySession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _CraftingFactory.Contract.ChangeAdmin(&_CraftingFactory.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_CraftingFactory *CraftingFactoryTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _CraftingFactory.Contract.ChangeAdmin(&_CraftingFactory.TransactOpts, _newAdmin)
}

// CreateCrafting is a paid mutator transaction binding the contract method 0x222c116c.
//
// Solidity: function createCrafting(uint256 collectionId, address owner) returns(address crafting)
func (_CraftingFactory *CraftingFactoryTransactor) CreateCrafting(opts *bind.TransactOpts, collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _CraftingFactory.contract.Transact(opts, "createCrafting", collectionId, owner)
}

// CreateCrafting is a paid mutator transaction binding the contract method 0x222c116c.
//
// Solidity: function createCrafting(uint256 collectionId, address owner) returns(address crafting)
func (_CraftingFactory *CraftingFactorySession) CreateCrafting(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _CraftingFactory.Contract.CreateCrafting(&_CraftingFactory.TransactOpts, collectionId, owner)
}

// CreateCrafting is a paid mutator transaction binding the contract method 0x222c116c.
//
// Solidity: function createCrafting(uint256 collectionId, address owner) returns(address crafting)
func (_CraftingFactory *CraftingFactoryTransactorSession) CreateCrafting(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _CraftingFactory.Contract.CreateCrafting(&_CraftingFactory.TransactOpts, collectionId, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CraftingFactory *CraftingFactoryTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CraftingFactory.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CraftingFactory *CraftingFactorySession) Initialize() (*types.Transaction, error) {
	return _CraftingFactory.Contract.Initialize(&_CraftingFactory.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CraftingFactory *CraftingFactoryTransactorSession) Initialize() (*types.Transaction, error) {
	return _CraftingFactory.Contract.Initialize(&_CraftingFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CraftingFactory *CraftingFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CraftingFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CraftingFactory *CraftingFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _CraftingFactory.Contract.RenounceOwnership(&_CraftingFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CraftingFactory *CraftingFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CraftingFactory.Contract.RenounceOwnership(&_CraftingFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CraftingFactory *CraftingFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CraftingFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CraftingFactory *CraftingFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CraftingFactory.Contract.TransferOwnership(&_CraftingFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CraftingFactory *CraftingFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CraftingFactory.Contract.TransferOwnership(&_CraftingFactory.TransactOpts, newOwner)
}

// CraftingFactoryCraftingCreatedIterator is returned from FilterCraftingCreated and is used to iterate over the raw logs and unpacked data for CraftingCreated events raised by the CraftingFactory contract.
type CraftingFactoryCraftingCreatedIterator struct {
	Event *CraftingFactoryCraftingCreated // Event containing the contract specifics and raw log

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
func (it *CraftingFactoryCraftingCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CraftingFactoryCraftingCreated)
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
		it.Event = new(CraftingFactoryCraftingCreated)
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
func (it *CraftingFactoryCraftingCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CraftingFactoryCraftingCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CraftingFactoryCraftingCreated represents a CraftingCreated event raised by the CraftingFactory contract.
type CraftingFactoryCraftingCreated struct {
	Crafting common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCraftingCreated is a free log retrieval operation binding the contract event 0x5949f717c13342b6ef8d9818d4878e982e2568e41a7372c059ddca6fa3d5a385.
//
// Solidity: event CraftingCreated(address crafting)
func (_CraftingFactory *CraftingFactoryFilterer) FilterCraftingCreated(opts *bind.FilterOpts) (*CraftingFactoryCraftingCreatedIterator, error) {

	logs, sub, err := _CraftingFactory.contract.FilterLogs(opts, "CraftingCreated")
	if err != nil {
		return nil, err
	}
	return &CraftingFactoryCraftingCreatedIterator{contract: _CraftingFactory.contract, event: "CraftingCreated", logs: logs, sub: sub}, nil
}

// WatchCraftingCreated is a free log subscription operation binding the contract event 0x5949f717c13342b6ef8d9818d4878e982e2568e41a7372c059ddca6fa3d5a385.
//
// Solidity: event CraftingCreated(address crafting)
func (_CraftingFactory *CraftingFactoryFilterer) WatchCraftingCreated(opts *bind.WatchOpts, sink chan<- *CraftingFactoryCraftingCreated) (event.Subscription, error) {

	logs, sub, err := _CraftingFactory.contract.WatchLogs(opts, "CraftingCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CraftingFactoryCraftingCreated)
				if err := _CraftingFactory.contract.UnpackLog(event, "CraftingCreated", log); err != nil {
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

// ParseCraftingCreated is a log parse operation binding the contract event 0x5949f717c13342b6ef8d9818d4878e982e2568e41a7372c059ddca6fa3d5a385.
//
// Solidity: event CraftingCreated(address crafting)
func (_CraftingFactory *CraftingFactoryFilterer) ParseCraftingCreated(log types.Log) (*CraftingFactoryCraftingCreated, error) {
	event := new(CraftingFactoryCraftingCreated)
	if err := _CraftingFactory.contract.UnpackLog(event, "CraftingCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CraftingFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CraftingFactory contract.
type CraftingFactoryOwnershipTransferredIterator struct {
	Event *CraftingFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CraftingFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CraftingFactoryOwnershipTransferred)
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
		it.Event = new(CraftingFactoryOwnershipTransferred)
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
func (it *CraftingFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CraftingFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CraftingFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the CraftingFactory contract.
type CraftingFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CraftingFactory *CraftingFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CraftingFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CraftingFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CraftingFactoryOwnershipTransferredIterator{contract: _CraftingFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CraftingFactory *CraftingFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CraftingFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CraftingFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CraftingFactoryOwnershipTransferred)
				if err := _CraftingFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CraftingFactory *CraftingFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*CraftingFactoryOwnershipTransferred, error) {
	event := new(CraftingFactoryOwnershipTransferred)
	if err := _CraftingFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
