// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package triggerFactory

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

// TriggerFactoryMetaData contains all meta data concerning the TriggerFactory contract.
var TriggerFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_collectionId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousCrafter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newCrafter\",\"type\":\"address\"}],\"name\":\"CraftingContractChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousMinter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMinter\",\"type\":\"address\"}],\"name\":\"MinterContractChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"PartCrafted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"TriggerMinted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newCrafter\",\"type\":\"address\"}],\"name\":\"changeCrafter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newMinter\",\"type\":\"address\"}],\"name\":\"changeMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"changeToCrafted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isCrafted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[][]\",\"name\":\"_tokenURIs\",\"type\":\"string[][]\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"mintTrigger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TriggerFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use TriggerFactoryMetaData.ABI instead.
var TriggerFactoryABI = TriggerFactoryMetaData.ABI

// TriggerFactory is an auto generated Go binding around an Ethereum contract.
type TriggerFactory struct {
	TriggerFactoryCaller     // Read-only binding to the contract
	TriggerFactoryTransactor // Write-only binding to the contract
	TriggerFactoryFilterer   // Log filterer for contract events
}

// TriggerFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TriggerFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TriggerFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TriggerFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TriggerFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TriggerFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TriggerFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TriggerFactorySession struct {
	Contract     *TriggerFactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TriggerFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TriggerFactoryCallerSession struct {
	Contract *TriggerFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TriggerFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TriggerFactoryTransactorSession struct {
	Contract     *TriggerFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TriggerFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TriggerFactoryRaw struct {
	Contract *TriggerFactory // Generic contract binding to access the raw methods on
}

// TriggerFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TriggerFactoryCallerRaw struct {
	Contract *TriggerFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// TriggerFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TriggerFactoryTransactorRaw struct {
	Contract *TriggerFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTriggerFactory creates a new instance of TriggerFactory, bound to a specific deployed contract.
func NewTriggerFactory(address common.Address, backend bind.ContractBackend) (*TriggerFactory, error) {
	contract, err := bindTriggerFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TriggerFactory{TriggerFactoryCaller: TriggerFactoryCaller{contract: contract}, TriggerFactoryTransactor: TriggerFactoryTransactor{contract: contract}, TriggerFactoryFilterer: TriggerFactoryFilterer{contract: contract}}, nil
}

// NewTriggerFactoryCaller creates a new read-only instance of TriggerFactory, bound to a specific deployed contract.
func NewTriggerFactoryCaller(address common.Address, caller bind.ContractCaller) (*TriggerFactoryCaller, error) {
	contract, err := bindTriggerFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TriggerFactoryCaller{contract: contract}, nil
}

// NewTriggerFactoryTransactor creates a new write-only instance of TriggerFactory, bound to a specific deployed contract.
func NewTriggerFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*TriggerFactoryTransactor, error) {
	contract, err := bindTriggerFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TriggerFactoryTransactor{contract: contract}, nil
}

// NewTriggerFactoryFilterer creates a new log filterer instance of TriggerFactory, bound to a specific deployed contract.
func NewTriggerFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*TriggerFactoryFilterer, error) {
	contract, err := bindTriggerFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TriggerFactoryFilterer{contract: contract}, nil
}

// bindTriggerFactory binds a generic wrapper to an already deployed contract.
func bindTriggerFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TriggerFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TriggerFactory *TriggerFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TriggerFactory.Contract.TriggerFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TriggerFactory *TriggerFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TriggerFactory.Contract.TriggerFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TriggerFactory *TriggerFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TriggerFactory.Contract.TriggerFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TriggerFactory *TriggerFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TriggerFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TriggerFactory *TriggerFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TriggerFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TriggerFactory *TriggerFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TriggerFactory.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TriggerFactory *TriggerFactoryCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TriggerFactory.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TriggerFactory *TriggerFactorySession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TriggerFactory.Contract.BalanceOf(&_TriggerFactory.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TriggerFactory *TriggerFactoryCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TriggerFactory.Contract.BalanceOf(&_TriggerFactory.CallOpts, owner)
}

// CollectionId is a free data retrieval call binding the contract method 0x3d26bb67.
//
// Solidity: function collectionId() view returns(uint256)
func (_TriggerFactory *TriggerFactoryCaller) CollectionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TriggerFactory.contract.Call(opts, &out, "collectionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollectionId is a free data retrieval call binding the contract method 0x3d26bb67.
//
// Solidity: function collectionId() view returns(uint256)
func (_TriggerFactory *TriggerFactorySession) CollectionId() (*big.Int, error) {
	return _TriggerFactory.Contract.CollectionId(&_TriggerFactory.CallOpts)
}

// CollectionId is a free data retrieval call binding the contract method 0x3d26bb67.
//
// Solidity: function collectionId() view returns(uint256)
func (_TriggerFactory *TriggerFactoryCallerSession) CollectionId() (*big.Int, error) {
	return _TriggerFactory.Contract.CollectionId(&_TriggerFactory.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TriggerFactory *TriggerFactoryCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TriggerFactory.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TriggerFactory *TriggerFactorySession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _TriggerFactory.Contract.GetApproved(&_TriggerFactory.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TriggerFactory *TriggerFactoryCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _TriggerFactory.Contract.GetApproved(&_TriggerFactory.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TriggerFactory *TriggerFactoryCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _TriggerFactory.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TriggerFactory *TriggerFactorySession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _TriggerFactory.Contract.IsApprovedForAll(&_TriggerFactory.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TriggerFactory *TriggerFactoryCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _TriggerFactory.Contract.IsApprovedForAll(&_TriggerFactory.CallOpts, owner, operator)
}

// IsCrafted is a free data retrieval call binding the contract method 0xfd96527d.
//
// Solidity: function isCrafted(uint256 tokenId) view returns(bool)
func (_TriggerFactory *TriggerFactoryCaller) IsCrafted(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _TriggerFactory.contract.Call(opts, &out, "isCrafted", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCrafted is a free data retrieval call binding the contract method 0xfd96527d.
//
// Solidity: function isCrafted(uint256 tokenId) view returns(bool)
func (_TriggerFactory *TriggerFactorySession) IsCrafted(tokenId *big.Int) (bool, error) {
	return _TriggerFactory.Contract.IsCrafted(&_TriggerFactory.CallOpts, tokenId)
}

// IsCrafted is a free data retrieval call binding the contract method 0xfd96527d.
//
// Solidity: function isCrafted(uint256 tokenId) view returns(bool)
func (_TriggerFactory *TriggerFactoryCallerSession) IsCrafted(tokenId *big.Int) (bool, error) {
	return _TriggerFactory.Contract.IsCrafted(&_TriggerFactory.CallOpts, tokenId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TriggerFactory *TriggerFactoryCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TriggerFactory.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TriggerFactory *TriggerFactorySession) Name() (string, error) {
	return _TriggerFactory.Contract.Name(&_TriggerFactory.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TriggerFactory *TriggerFactoryCallerSession) Name() (string, error) {
	return _TriggerFactory.Contract.Name(&_TriggerFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TriggerFactory *TriggerFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TriggerFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TriggerFactory *TriggerFactorySession) Owner() (common.Address, error) {
	return _TriggerFactory.Contract.Owner(&_TriggerFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TriggerFactory *TriggerFactoryCallerSession) Owner() (common.Address, error) {
	return _TriggerFactory.Contract.Owner(&_TriggerFactory.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TriggerFactory *TriggerFactoryCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TriggerFactory.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TriggerFactory *TriggerFactorySession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _TriggerFactory.Contract.OwnerOf(&_TriggerFactory.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TriggerFactory *TriggerFactoryCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _TriggerFactory.Contract.OwnerOf(&_TriggerFactory.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TriggerFactory *TriggerFactoryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TriggerFactory.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TriggerFactory *TriggerFactorySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TriggerFactory.Contract.SupportsInterface(&_TriggerFactory.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TriggerFactory *TriggerFactoryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TriggerFactory.Contract.SupportsInterface(&_TriggerFactory.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TriggerFactory *TriggerFactoryCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TriggerFactory.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TriggerFactory *TriggerFactorySession) Symbol() (string, error) {
	return _TriggerFactory.Contract.Symbol(&_TriggerFactory.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TriggerFactory *TriggerFactoryCallerSession) Symbol() (string, error) {
	return _TriggerFactory.Contract.Symbol(&_TriggerFactory.CallOpts)
}

// TokenCounter is a free data retrieval call binding the contract method 0xd082e381.
//
// Solidity: function tokenCounter() view returns(uint256)
func (_TriggerFactory *TriggerFactoryCaller) TokenCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TriggerFactory.contract.Call(opts, &out, "tokenCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCounter is a free data retrieval call binding the contract method 0xd082e381.
//
// Solidity: function tokenCounter() view returns(uint256)
func (_TriggerFactory *TriggerFactorySession) TokenCounter() (*big.Int, error) {
	return _TriggerFactory.Contract.TokenCounter(&_TriggerFactory.CallOpts)
}

// TokenCounter is a free data retrieval call binding the contract method 0xd082e381.
//
// Solidity: function tokenCounter() view returns(uint256)
func (_TriggerFactory *TriggerFactoryCallerSession) TokenCounter() (*big.Int, error) {
	return _TriggerFactory.Contract.TokenCounter(&_TriggerFactory.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_TriggerFactory *TriggerFactoryCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _TriggerFactory.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_TriggerFactory *TriggerFactorySession) TokenURI(tokenId *big.Int) (string, error) {
	return _TriggerFactory.Contract.TokenURI(&_TriggerFactory.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_TriggerFactory *TriggerFactoryCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _TriggerFactory.Contract.TokenURI(&_TriggerFactory.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TriggerFactory *TriggerFactoryTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TriggerFactory.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TriggerFactory *TriggerFactorySession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TriggerFactory.Contract.Approve(&_TriggerFactory.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TriggerFactory *TriggerFactoryTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TriggerFactory.Contract.Approve(&_TriggerFactory.TransactOpts, to, tokenId)
}

// ChangeCrafter is a paid mutator transaction binding the contract method 0xe7d3b069.
//
// Solidity: function changeCrafter(address _newCrafter) returns()
func (_TriggerFactory *TriggerFactoryTransactor) ChangeCrafter(opts *bind.TransactOpts, _newCrafter common.Address) (*types.Transaction, error) {
	return _TriggerFactory.contract.Transact(opts, "changeCrafter", _newCrafter)
}

// ChangeCrafter is a paid mutator transaction binding the contract method 0xe7d3b069.
//
// Solidity: function changeCrafter(address _newCrafter) returns()
func (_TriggerFactory *TriggerFactorySession) ChangeCrafter(_newCrafter common.Address) (*types.Transaction, error) {
	return _TriggerFactory.Contract.ChangeCrafter(&_TriggerFactory.TransactOpts, _newCrafter)
}

// ChangeCrafter is a paid mutator transaction binding the contract method 0xe7d3b069.
//
// Solidity: function changeCrafter(address _newCrafter) returns()
func (_TriggerFactory *TriggerFactoryTransactorSession) ChangeCrafter(_newCrafter common.Address) (*types.Transaction, error) {
	return _TriggerFactory.Contract.ChangeCrafter(&_TriggerFactory.TransactOpts, _newCrafter)
}

// ChangeMinter is a paid mutator transaction binding the contract method 0x2c4d4d18.
//
// Solidity: function changeMinter(address _newMinter) returns()
func (_TriggerFactory *TriggerFactoryTransactor) ChangeMinter(opts *bind.TransactOpts, _newMinter common.Address) (*types.Transaction, error) {
	return _TriggerFactory.contract.Transact(opts, "changeMinter", _newMinter)
}

// ChangeMinter is a paid mutator transaction binding the contract method 0x2c4d4d18.
//
// Solidity: function changeMinter(address _newMinter) returns()
func (_TriggerFactory *TriggerFactorySession) ChangeMinter(_newMinter common.Address) (*types.Transaction, error) {
	return _TriggerFactory.Contract.ChangeMinter(&_TriggerFactory.TransactOpts, _newMinter)
}

// ChangeMinter is a paid mutator transaction binding the contract method 0x2c4d4d18.
//
// Solidity: function changeMinter(address _newMinter) returns()
func (_TriggerFactory *TriggerFactoryTransactorSession) ChangeMinter(_newMinter common.Address) (*types.Transaction, error) {
	return _TriggerFactory.Contract.ChangeMinter(&_TriggerFactory.TransactOpts, _newMinter)
}

// ChangeToCrafted is a paid mutator transaction binding the contract method 0xe6b60d9b.
//
// Solidity: function changeToCrafted(uint256 tokenId) returns()
func (_TriggerFactory *TriggerFactoryTransactor) ChangeToCrafted(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _TriggerFactory.contract.Transact(opts, "changeToCrafted", tokenId)
}

// ChangeToCrafted is a paid mutator transaction binding the contract method 0xe6b60d9b.
//
// Solidity: function changeToCrafted(uint256 tokenId) returns()
func (_TriggerFactory *TriggerFactorySession) ChangeToCrafted(tokenId *big.Int) (*types.Transaction, error) {
	return _TriggerFactory.Contract.ChangeToCrafted(&_TriggerFactory.TransactOpts, tokenId)
}

// ChangeToCrafted is a paid mutator transaction binding the contract method 0xe6b60d9b.
//
// Solidity: function changeToCrafted(uint256 tokenId) returns()
func (_TriggerFactory *TriggerFactoryTransactorSession) ChangeToCrafted(tokenId *big.Int) (*types.Transaction, error) {
	return _TriggerFactory.Contract.ChangeToCrafted(&_TriggerFactory.TransactOpts, tokenId)
}

// MintTrigger is a paid mutator transaction binding the contract method 0x53710fcd.
//
// Solidity: function mintTrigger(string[][] _tokenURIs, address _owner) returns()
func (_TriggerFactory *TriggerFactoryTransactor) MintTrigger(opts *bind.TransactOpts, _tokenURIs [][]string, _owner common.Address) (*types.Transaction, error) {
	return _TriggerFactory.contract.Transact(opts, "mintTrigger", _tokenURIs, _owner)
}

// MintTrigger is a paid mutator transaction binding the contract method 0x53710fcd.
//
// Solidity: function mintTrigger(string[][] _tokenURIs, address _owner) returns()
func (_TriggerFactory *TriggerFactorySession) MintTrigger(_tokenURIs [][]string, _owner common.Address) (*types.Transaction, error) {
	return _TriggerFactory.Contract.MintTrigger(&_TriggerFactory.TransactOpts, _tokenURIs, _owner)
}

// MintTrigger is a paid mutator transaction binding the contract method 0x53710fcd.
//
// Solidity: function mintTrigger(string[][] _tokenURIs, address _owner) returns()
func (_TriggerFactory *TriggerFactoryTransactorSession) MintTrigger(_tokenURIs [][]string, _owner common.Address) (*types.Transaction, error) {
	return _TriggerFactory.Contract.MintTrigger(&_TriggerFactory.TransactOpts, _tokenURIs, _owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TriggerFactory *TriggerFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TriggerFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TriggerFactory *TriggerFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _TriggerFactory.Contract.RenounceOwnership(&_TriggerFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TriggerFactory *TriggerFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TriggerFactory.Contract.RenounceOwnership(&_TriggerFactory.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TriggerFactory *TriggerFactoryTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TriggerFactory.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TriggerFactory *TriggerFactorySession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TriggerFactory.Contract.SafeTransferFrom(&_TriggerFactory.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TriggerFactory *TriggerFactoryTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TriggerFactory.Contract.SafeTransferFrom(&_TriggerFactory.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_TriggerFactory *TriggerFactoryTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _TriggerFactory.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_TriggerFactory *TriggerFactorySession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _TriggerFactory.Contract.SafeTransferFrom0(&_TriggerFactory.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_TriggerFactory *TriggerFactoryTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _TriggerFactory.Contract.SafeTransferFrom0(&_TriggerFactory.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TriggerFactory *TriggerFactoryTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _TriggerFactory.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TriggerFactory *TriggerFactorySession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TriggerFactory.Contract.SetApprovalForAll(&_TriggerFactory.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TriggerFactory *TriggerFactoryTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TriggerFactory.Contract.SetApprovalForAll(&_TriggerFactory.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TriggerFactory *TriggerFactoryTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TriggerFactory.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TriggerFactory *TriggerFactorySession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TriggerFactory.Contract.TransferFrom(&_TriggerFactory.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TriggerFactory *TriggerFactoryTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TriggerFactory.Contract.TransferFrom(&_TriggerFactory.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TriggerFactory *TriggerFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TriggerFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TriggerFactory *TriggerFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TriggerFactory.Contract.TransferOwnership(&_TriggerFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TriggerFactory *TriggerFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TriggerFactory.Contract.TransferOwnership(&_TriggerFactory.TransactOpts, newOwner)
}

// TriggerFactoryApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TriggerFactory contract.
type TriggerFactoryApprovalIterator struct {
	Event *TriggerFactoryApproval // Event containing the contract specifics and raw log

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
func (it *TriggerFactoryApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerFactoryApproval)
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
		it.Event = new(TriggerFactoryApproval)
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
func (it *TriggerFactoryApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerFactoryApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerFactoryApproval represents a Approval event raised by the TriggerFactory contract.
type TriggerFactoryApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TriggerFactory *TriggerFactoryFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*TriggerFactoryApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TriggerFactory.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TriggerFactoryApprovalIterator{contract: _TriggerFactory.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TriggerFactory *TriggerFactoryFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TriggerFactoryApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TriggerFactory.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerFactoryApproval)
				if err := _TriggerFactory.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TriggerFactory *TriggerFactoryFilterer) ParseApproval(log types.Log) (*TriggerFactoryApproval, error) {
	event := new(TriggerFactoryApproval)
	if err := _TriggerFactory.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerFactoryApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the TriggerFactory contract.
type TriggerFactoryApprovalForAllIterator struct {
	Event *TriggerFactoryApprovalForAll // Event containing the contract specifics and raw log

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
func (it *TriggerFactoryApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerFactoryApprovalForAll)
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
		it.Event = new(TriggerFactoryApprovalForAll)
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
func (it *TriggerFactoryApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerFactoryApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerFactoryApprovalForAll represents a ApprovalForAll event raised by the TriggerFactory contract.
type TriggerFactoryApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TriggerFactory *TriggerFactoryFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*TriggerFactoryApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TriggerFactory.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TriggerFactoryApprovalForAllIterator{contract: _TriggerFactory.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TriggerFactory *TriggerFactoryFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *TriggerFactoryApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TriggerFactory.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerFactoryApprovalForAll)
				if err := _TriggerFactory.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TriggerFactory *TriggerFactoryFilterer) ParseApprovalForAll(log types.Log) (*TriggerFactoryApprovalForAll, error) {
	event := new(TriggerFactoryApprovalForAll)
	if err := _TriggerFactory.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerFactoryCraftingContractChangedIterator is returned from FilterCraftingContractChanged and is used to iterate over the raw logs and unpacked data for CraftingContractChanged events raised by the TriggerFactory contract.
type TriggerFactoryCraftingContractChangedIterator struct {
	Event *TriggerFactoryCraftingContractChanged // Event containing the contract specifics and raw log

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
func (it *TriggerFactoryCraftingContractChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerFactoryCraftingContractChanged)
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
		it.Event = new(TriggerFactoryCraftingContractChanged)
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
func (it *TriggerFactoryCraftingContractChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerFactoryCraftingContractChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerFactoryCraftingContractChanged represents a CraftingContractChanged event raised by the TriggerFactory contract.
type TriggerFactoryCraftingContractChanged struct {
	PreviousCrafter common.Address
	NewCrafter      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCraftingContractChanged is a free log retrieval operation binding the contract event 0x83149e90b909cde0e8099cb6998203f8d8dce5fa29151688382650f5637d7d88.
//
// Solidity: event CraftingContractChanged(address indexed previousCrafter, address indexed newCrafter)
func (_TriggerFactory *TriggerFactoryFilterer) FilterCraftingContractChanged(opts *bind.FilterOpts, previousCrafter []common.Address, newCrafter []common.Address) (*TriggerFactoryCraftingContractChangedIterator, error) {

	var previousCrafterRule []interface{}
	for _, previousCrafterItem := range previousCrafter {
		previousCrafterRule = append(previousCrafterRule, previousCrafterItem)
	}
	var newCrafterRule []interface{}
	for _, newCrafterItem := range newCrafter {
		newCrafterRule = append(newCrafterRule, newCrafterItem)
	}

	logs, sub, err := _TriggerFactory.contract.FilterLogs(opts, "CraftingContractChanged", previousCrafterRule, newCrafterRule)
	if err != nil {
		return nil, err
	}
	return &TriggerFactoryCraftingContractChangedIterator{contract: _TriggerFactory.contract, event: "CraftingContractChanged", logs: logs, sub: sub}, nil
}

// WatchCraftingContractChanged is a free log subscription operation binding the contract event 0x83149e90b909cde0e8099cb6998203f8d8dce5fa29151688382650f5637d7d88.
//
// Solidity: event CraftingContractChanged(address indexed previousCrafter, address indexed newCrafter)
func (_TriggerFactory *TriggerFactoryFilterer) WatchCraftingContractChanged(opts *bind.WatchOpts, sink chan<- *TriggerFactoryCraftingContractChanged, previousCrafter []common.Address, newCrafter []common.Address) (event.Subscription, error) {

	var previousCrafterRule []interface{}
	for _, previousCrafterItem := range previousCrafter {
		previousCrafterRule = append(previousCrafterRule, previousCrafterItem)
	}
	var newCrafterRule []interface{}
	for _, newCrafterItem := range newCrafter {
		newCrafterRule = append(newCrafterRule, newCrafterItem)
	}

	logs, sub, err := _TriggerFactory.contract.WatchLogs(opts, "CraftingContractChanged", previousCrafterRule, newCrafterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerFactoryCraftingContractChanged)
				if err := _TriggerFactory.contract.UnpackLog(event, "CraftingContractChanged", log); err != nil {
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

// ParseCraftingContractChanged is a log parse operation binding the contract event 0x83149e90b909cde0e8099cb6998203f8d8dce5fa29151688382650f5637d7d88.
//
// Solidity: event CraftingContractChanged(address indexed previousCrafter, address indexed newCrafter)
func (_TriggerFactory *TriggerFactoryFilterer) ParseCraftingContractChanged(log types.Log) (*TriggerFactoryCraftingContractChanged, error) {
	event := new(TriggerFactoryCraftingContractChanged)
	if err := _TriggerFactory.contract.UnpackLog(event, "CraftingContractChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerFactoryMinterContractChangedIterator is returned from FilterMinterContractChanged and is used to iterate over the raw logs and unpacked data for MinterContractChanged events raised by the TriggerFactory contract.
type TriggerFactoryMinterContractChangedIterator struct {
	Event *TriggerFactoryMinterContractChanged // Event containing the contract specifics and raw log

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
func (it *TriggerFactoryMinterContractChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerFactoryMinterContractChanged)
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
		it.Event = new(TriggerFactoryMinterContractChanged)
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
func (it *TriggerFactoryMinterContractChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerFactoryMinterContractChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerFactoryMinterContractChanged represents a MinterContractChanged event raised by the TriggerFactory contract.
type TriggerFactoryMinterContractChanged struct {
	PreviousMinter common.Address
	NewMinter      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMinterContractChanged is a free log retrieval operation binding the contract event 0x114d06280373bf8a65ca81a7d5ec748e7747747d23dbaa3e41f3c4ea46223079.
//
// Solidity: event MinterContractChanged(address indexed previousMinter, address indexed newMinter)
func (_TriggerFactory *TriggerFactoryFilterer) FilterMinterContractChanged(opts *bind.FilterOpts, previousMinter []common.Address, newMinter []common.Address) (*TriggerFactoryMinterContractChangedIterator, error) {

	var previousMinterRule []interface{}
	for _, previousMinterItem := range previousMinter {
		previousMinterRule = append(previousMinterRule, previousMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _TriggerFactory.contract.FilterLogs(opts, "MinterContractChanged", previousMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return &TriggerFactoryMinterContractChangedIterator{contract: _TriggerFactory.contract, event: "MinterContractChanged", logs: logs, sub: sub}, nil
}

// WatchMinterContractChanged is a free log subscription operation binding the contract event 0x114d06280373bf8a65ca81a7d5ec748e7747747d23dbaa3e41f3c4ea46223079.
//
// Solidity: event MinterContractChanged(address indexed previousMinter, address indexed newMinter)
func (_TriggerFactory *TriggerFactoryFilterer) WatchMinterContractChanged(opts *bind.WatchOpts, sink chan<- *TriggerFactoryMinterContractChanged, previousMinter []common.Address, newMinter []common.Address) (event.Subscription, error) {

	var previousMinterRule []interface{}
	for _, previousMinterItem := range previousMinter {
		previousMinterRule = append(previousMinterRule, previousMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _TriggerFactory.contract.WatchLogs(opts, "MinterContractChanged", previousMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerFactoryMinterContractChanged)
				if err := _TriggerFactory.contract.UnpackLog(event, "MinterContractChanged", log); err != nil {
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

// ParseMinterContractChanged is a log parse operation binding the contract event 0x114d06280373bf8a65ca81a7d5ec748e7747747d23dbaa3e41f3c4ea46223079.
//
// Solidity: event MinterContractChanged(address indexed previousMinter, address indexed newMinter)
func (_TriggerFactory *TriggerFactoryFilterer) ParseMinterContractChanged(log types.Log) (*TriggerFactoryMinterContractChanged, error) {
	event := new(TriggerFactoryMinterContractChanged)
	if err := _TriggerFactory.contract.UnpackLog(event, "MinterContractChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TriggerFactory contract.
type TriggerFactoryOwnershipTransferredIterator struct {
	Event *TriggerFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TriggerFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerFactoryOwnershipTransferred)
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
		it.Event = new(TriggerFactoryOwnershipTransferred)
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
func (it *TriggerFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the TriggerFactory contract.
type TriggerFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TriggerFactory *TriggerFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TriggerFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TriggerFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TriggerFactoryOwnershipTransferredIterator{contract: _TriggerFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TriggerFactory *TriggerFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TriggerFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TriggerFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerFactoryOwnershipTransferred)
				if err := _TriggerFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TriggerFactory *TriggerFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*TriggerFactoryOwnershipTransferred, error) {
	event := new(TriggerFactoryOwnershipTransferred)
	if err := _TriggerFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerFactoryPartCraftedIterator is returned from FilterPartCrafted and is used to iterate over the raw logs and unpacked data for PartCrafted events raised by the TriggerFactory contract.
type TriggerFactoryPartCraftedIterator struct {
	Event *TriggerFactoryPartCrafted // Event containing the contract specifics and raw log

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
func (it *TriggerFactoryPartCraftedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerFactoryPartCrafted)
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
		it.Event = new(TriggerFactoryPartCrafted)
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
func (it *TriggerFactoryPartCraftedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerFactoryPartCraftedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerFactoryPartCrafted represents a PartCrafted event raised by the TriggerFactory contract.
type TriggerFactoryPartCrafted struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPartCrafted is a free log retrieval operation binding the contract event 0x71094059908fc1a5d676895c3ab6ec0e8957a09dc7a09914b791a2be4dc7fb15.
//
// Solidity: event PartCrafted(uint256 indexed tokenId)
func (_TriggerFactory *TriggerFactoryFilterer) FilterPartCrafted(opts *bind.FilterOpts, tokenId []*big.Int) (*TriggerFactoryPartCraftedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TriggerFactory.contract.FilterLogs(opts, "PartCrafted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TriggerFactoryPartCraftedIterator{contract: _TriggerFactory.contract, event: "PartCrafted", logs: logs, sub: sub}, nil
}

// WatchPartCrafted is a free log subscription operation binding the contract event 0x71094059908fc1a5d676895c3ab6ec0e8957a09dc7a09914b791a2be4dc7fb15.
//
// Solidity: event PartCrafted(uint256 indexed tokenId)
func (_TriggerFactory *TriggerFactoryFilterer) WatchPartCrafted(opts *bind.WatchOpts, sink chan<- *TriggerFactoryPartCrafted, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TriggerFactory.contract.WatchLogs(opts, "PartCrafted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerFactoryPartCrafted)
				if err := _TriggerFactory.contract.UnpackLog(event, "PartCrafted", log); err != nil {
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

// ParsePartCrafted is a log parse operation binding the contract event 0x71094059908fc1a5d676895c3ab6ec0e8957a09dc7a09914b791a2be4dc7fb15.
//
// Solidity: event PartCrafted(uint256 indexed tokenId)
func (_TriggerFactory *TriggerFactoryFilterer) ParsePartCrafted(log types.Log) (*TriggerFactoryPartCrafted, error) {
	event := new(TriggerFactoryPartCrafted)
	if err := _TriggerFactory.contract.UnpackLog(event, "PartCrafted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerFactoryTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TriggerFactory contract.
type TriggerFactoryTransferIterator struct {
	Event *TriggerFactoryTransfer // Event containing the contract specifics and raw log

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
func (it *TriggerFactoryTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerFactoryTransfer)
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
		it.Event = new(TriggerFactoryTransfer)
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
func (it *TriggerFactoryTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerFactoryTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerFactoryTransfer represents a Transfer event raised by the TriggerFactory contract.
type TriggerFactoryTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TriggerFactory *TriggerFactoryFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*TriggerFactoryTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TriggerFactory.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TriggerFactoryTransferIterator{contract: _TriggerFactory.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TriggerFactory *TriggerFactoryFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TriggerFactoryTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TriggerFactory.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerFactoryTransfer)
				if err := _TriggerFactory.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TriggerFactory *TriggerFactoryFilterer) ParseTransfer(log types.Log) (*TriggerFactoryTransfer, error) {
	event := new(TriggerFactoryTransfer)
	if err := _TriggerFactory.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerFactoryTriggerMintedIterator is returned from FilterTriggerMinted and is used to iterate over the raw logs and unpacked data for TriggerMinted events raised by the TriggerFactory contract.
type TriggerFactoryTriggerMintedIterator struct {
	Event *TriggerFactoryTriggerMinted // Event containing the contract specifics and raw log

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
func (it *TriggerFactoryTriggerMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerFactoryTriggerMinted)
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
		it.Event = new(TriggerFactoryTriggerMinted)
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
func (it *TriggerFactoryTriggerMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerFactoryTriggerMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerFactoryTriggerMinted represents a TriggerMinted event raised by the TriggerFactory contract.
type TriggerFactoryTriggerMinted struct {
	CollectionId *big.Int
	TokenId      *big.Int
	TokenURI     string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTriggerMinted is a free log retrieval operation binding the contract event 0x61372f552903c29239005dd638aa00ac6da17b54079387e652d9209fc0b59052.
//
// Solidity: event TriggerMinted(uint256 collectionId, uint256 indexed tokenId, string tokenURI)
func (_TriggerFactory *TriggerFactoryFilterer) FilterTriggerMinted(opts *bind.FilterOpts, tokenId []*big.Int) (*TriggerFactoryTriggerMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TriggerFactory.contract.FilterLogs(opts, "TriggerMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TriggerFactoryTriggerMintedIterator{contract: _TriggerFactory.contract, event: "TriggerMinted", logs: logs, sub: sub}, nil
}

// WatchTriggerMinted is a free log subscription operation binding the contract event 0x61372f552903c29239005dd638aa00ac6da17b54079387e652d9209fc0b59052.
//
// Solidity: event TriggerMinted(uint256 collectionId, uint256 indexed tokenId, string tokenURI)
func (_TriggerFactory *TriggerFactoryFilterer) WatchTriggerMinted(opts *bind.WatchOpts, sink chan<- *TriggerFactoryTriggerMinted, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TriggerFactory.contract.WatchLogs(opts, "TriggerMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerFactoryTriggerMinted)
				if err := _TriggerFactory.contract.UnpackLog(event, "TriggerMinted", log); err != nil {
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

// ParseTriggerMinted is a log parse operation binding the contract event 0x61372f552903c29239005dd638aa00ac6da17b54079387e652d9209fc0b59052.
//
// Solidity: event TriggerMinted(uint256 collectionId, uint256 indexed tokenId, string tokenURI)
func (_TriggerFactory *TriggerFactoryFilterer) ParseTriggerMinted(log types.Log) (*TriggerFactoryTriggerMinted, error) {
	event := new(TriggerFactoryTriggerMinted)
	if err := _TriggerFactory.contract.UnpackLog(event, "TriggerMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
