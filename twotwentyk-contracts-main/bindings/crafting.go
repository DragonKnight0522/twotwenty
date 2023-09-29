// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crafting

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

// CraftingMetaData contains all meta data concerning the Crafting contract.
var CraftingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dayMonthTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yearTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"categoryTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"craftingCardTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"CollectionCrafted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousCrafter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newCrafter\",\"type\":\"address\"}],\"name\":\"CraftingContractChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"PartCrafted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"categoryContract\",\"outputs\":[{\"internalType\":\"contractIPart\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newCrafter\",\"type\":\"address\"}],\"name\":\"changeCrafter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"changeToCrafted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dayMonthContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_yearContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_categoryContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_craftingCardContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dayMonthTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"yearTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"categoryTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"craftingCardTokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"}],\"name\":\"craftCollection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"crafted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"craftingCardContract\",\"outputs\":[{\"internalType\":\"contractIPart\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dayMonthContract\",\"outputs\":[{\"internalType\":\"contractIPart\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"}],\"name\":\"getCollectionTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isCrafted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yearContract\",\"outputs\":[{\"internalType\":\"contractIPart\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CraftingABI is the input ABI used to generate the binding from.
// Deprecated: Use CraftingMetaData.ABI instead.
var CraftingABI = CraftingMetaData.ABI

// Crafting is an auto generated Go binding around an Ethereum contract.
type Crafting struct {
	CraftingCaller     // Read-only binding to the contract
	CraftingTransactor // Write-only binding to the contract
	CraftingFilterer   // Log filterer for contract events
}

// CraftingCaller is an auto generated read-only Go binding around an Ethereum contract.
type CraftingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CraftingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CraftingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CraftingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CraftingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CraftingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CraftingSession struct {
	Contract     *Crafting         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CraftingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CraftingCallerSession struct {
	Contract *CraftingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CraftingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CraftingTransactorSession struct {
	Contract     *CraftingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CraftingRaw is an auto generated low-level Go binding around an Ethereum contract.
type CraftingRaw struct {
	Contract *Crafting // Generic contract binding to access the raw methods on
}

// CraftingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CraftingCallerRaw struct {
	Contract *CraftingCaller // Generic read-only contract binding to access the raw methods on
}

// CraftingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CraftingTransactorRaw struct {
	Contract *CraftingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrafting creates a new instance of Crafting, bound to a specific deployed contract.
func NewCrafting(address common.Address, backend bind.ContractBackend) (*Crafting, error) {
	contract, err := bindCrafting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Crafting{CraftingCaller: CraftingCaller{contract: contract}, CraftingTransactor: CraftingTransactor{contract: contract}, CraftingFilterer: CraftingFilterer{contract: contract}}, nil
}

// NewCraftingCaller creates a new read-only instance of Crafting, bound to a specific deployed contract.
func NewCraftingCaller(address common.Address, caller bind.ContractCaller) (*CraftingCaller, error) {
	contract, err := bindCrafting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CraftingCaller{contract: contract}, nil
}

// NewCraftingTransactor creates a new write-only instance of Crafting, bound to a specific deployed contract.
func NewCraftingTransactor(address common.Address, transactor bind.ContractTransactor) (*CraftingTransactor, error) {
	contract, err := bindCrafting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CraftingTransactor{contract: contract}, nil
}

// NewCraftingFilterer creates a new log filterer instance of Crafting, bound to a specific deployed contract.
func NewCraftingFilterer(address common.Address, filterer bind.ContractFilterer) (*CraftingFilterer, error) {
	contract, err := bindCrafting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CraftingFilterer{contract: contract}, nil
}

// bindCrafting binds a generic wrapper to an already deployed contract.
func bindCrafting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CraftingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crafting *CraftingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Crafting.Contract.CraftingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crafting *CraftingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crafting.Contract.CraftingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crafting *CraftingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crafting.Contract.CraftingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crafting *CraftingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Crafting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crafting *CraftingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crafting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crafting *CraftingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crafting.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Crafting *CraftingCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Crafting.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Crafting *CraftingSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Crafting.Contract.BalanceOf(&_Crafting.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Crafting *CraftingCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Crafting.Contract.BalanceOf(&_Crafting.CallOpts, owner)
}

// CategoryContract is a free data retrieval call binding the contract method 0xe4f0bdcc.
//
// Solidity: function categoryContract() view returns(address)
func (_Crafting *CraftingCaller) CategoryContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crafting.contract.Call(opts, &out, "categoryContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CategoryContract is a free data retrieval call binding the contract method 0xe4f0bdcc.
//
// Solidity: function categoryContract() view returns(address)
func (_Crafting *CraftingSession) CategoryContract() (common.Address, error) {
	return _Crafting.Contract.CategoryContract(&_Crafting.CallOpts)
}

// CategoryContract is a free data retrieval call binding the contract method 0xe4f0bdcc.
//
// Solidity: function categoryContract() view returns(address)
func (_Crafting *CraftingCallerSession) CategoryContract() (common.Address, error) {
	return _Crafting.Contract.CategoryContract(&_Crafting.CallOpts)
}

// Crafted is a free data retrieval call binding the contract method 0xcac4cc29.
//
// Solidity: function crafted(uint256 ) view returns(bool)
func (_Crafting *CraftingCaller) Crafted(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Crafting.contract.Call(opts, &out, "crafted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Crafted is a free data retrieval call binding the contract method 0xcac4cc29.
//
// Solidity: function crafted(uint256 ) view returns(bool)
func (_Crafting *CraftingSession) Crafted(arg0 *big.Int) (bool, error) {
	return _Crafting.Contract.Crafted(&_Crafting.CallOpts, arg0)
}

// Crafted is a free data retrieval call binding the contract method 0xcac4cc29.
//
// Solidity: function crafted(uint256 ) view returns(bool)
func (_Crafting *CraftingCallerSession) Crafted(arg0 *big.Int) (bool, error) {
	return _Crafting.Contract.Crafted(&_Crafting.CallOpts, arg0)
}

// CraftingCardContract is a free data retrieval call binding the contract method 0x2c948296.
//
// Solidity: function craftingCardContract() view returns(address)
func (_Crafting *CraftingCaller) CraftingCardContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crafting.contract.Call(opts, &out, "craftingCardContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CraftingCardContract is a free data retrieval call binding the contract method 0x2c948296.
//
// Solidity: function craftingCardContract() view returns(address)
func (_Crafting *CraftingSession) CraftingCardContract() (common.Address, error) {
	return _Crafting.Contract.CraftingCardContract(&_Crafting.CallOpts)
}

// CraftingCardContract is a free data retrieval call binding the contract method 0x2c948296.
//
// Solidity: function craftingCardContract() view returns(address)
func (_Crafting *CraftingCallerSession) CraftingCardContract() (common.Address, error) {
	return _Crafting.Contract.CraftingCardContract(&_Crafting.CallOpts)
}

// DayMonthContract is a free data retrieval call binding the contract method 0x9829ba1c.
//
// Solidity: function dayMonthContract() view returns(address)
func (_Crafting *CraftingCaller) DayMonthContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crafting.contract.Call(opts, &out, "dayMonthContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DayMonthContract is a free data retrieval call binding the contract method 0x9829ba1c.
//
// Solidity: function dayMonthContract() view returns(address)
func (_Crafting *CraftingSession) DayMonthContract() (common.Address, error) {
	return _Crafting.Contract.DayMonthContract(&_Crafting.CallOpts)
}

// DayMonthContract is a free data retrieval call binding the contract method 0x9829ba1c.
//
// Solidity: function dayMonthContract() view returns(address)
func (_Crafting *CraftingCallerSession) DayMonthContract() (common.Address, error) {
	return _Crafting.Contract.DayMonthContract(&_Crafting.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Crafting *CraftingCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Crafting.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Crafting *CraftingSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Crafting.Contract.GetApproved(&_Crafting.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Crafting *CraftingCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Crafting.Contract.GetApproved(&_Crafting.CallOpts, tokenId)
}

// GetCollectionTokens is a free data retrieval call binding the contract method 0x2bb88056.
//
// Solidity: function getCollectionTokens(uint256 collectionId) view returns(uint256, uint256, uint256, uint256)
func (_Crafting *CraftingCaller) GetCollectionTokens(opts *bind.CallOpts, collectionId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Crafting.contract.Call(opts, &out, "getCollectionTokens", collectionId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetCollectionTokens is a free data retrieval call binding the contract method 0x2bb88056.
//
// Solidity: function getCollectionTokens(uint256 collectionId) view returns(uint256, uint256, uint256, uint256)
func (_Crafting *CraftingSession) GetCollectionTokens(collectionId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Crafting.Contract.GetCollectionTokens(&_Crafting.CallOpts, collectionId)
}

// GetCollectionTokens is a free data retrieval call binding the contract method 0x2bb88056.
//
// Solidity: function getCollectionTokens(uint256 collectionId) view returns(uint256, uint256, uint256, uint256)
func (_Crafting *CraftingCallerSession) GetCollectionTokens(collectionId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Crafting.Contract.GetCollectionTokens(&_Crafting.CallOpts, collectionId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Crafting *CraftingCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Crafting.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Crafting *CraftingSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Crafting.Contract.IsApprovedForAll(&_Crafting.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Crafting *CraftingCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Crafting.Contract.IsApprovedForAll(&_Crafting.CallOpts, owner, operator)
}

// IsCrafted is a free data retrieval call binding the contract method 0xfd96527d.
//
// Solidity: function isCrafted(uint256 tokenId) view returns(bool)
func (_Crafting *CraftingCaller) IsCrafted(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Crafting.contract.Call(opts, &out, "isCrafted", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCrafted is a free data retrieval call binding the contract method 0xfd96527d.
//
// Solidity: function isCrafted(uint256 tokenId) view returns(bool)
func (_Crafting *CraftingSession) IsCrafted(tokenId *big.Int) (bool, error) {
	return _Crafting.Contract.IsCrafted(&_Crafting.CallOpts, tokenId)
}

// IsCrafted is a free data retrieval call binding the contract method 0xfd96527d.
//
// Solidity: function isCrafted(uint256 tokenId) view returns(bool)
func (_Crafting *CraftingCallerSession) IsCrafted(tokenId *big.Int) (bool, error) {
	return _Crafting.Contract.IsCrafted(&_Crafting.CallOpts, tokenId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Crafting *CraftingCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Crafting.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Crafting *CraftingSession) Name() (string, error) {
	return _Crafting.Contract.Name(&_Crafting.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Crafting *CraftingCallerSession) Name() (string, error) {
	return _Crafting.Contract.Name(&_Crafting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Crafting *CraftingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crafting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Crafting *CraftingSession) Owner() (common.Address, error) {
	return _Crafting.Contract.Owner(&_Crafting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Crafting *CraftingCallerSession) Owner() (common.Address, error) {
	return _Crafting.Contract.Owner(&_Crafting.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Crafting *CraftingCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Crafting.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Crafting *CraftingSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Crafting.Contract.OwnerOf(&_Crafting.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Crafting *CraftingCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Crafting.Contract.OwnerOf(&_Crafting.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Crafting *CraftingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Crafting.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Crafting *CraftingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Crafting.Contract.SupportsInterface(&_Crafting.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Crafting *CraftingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Crafting.Contract.SupportsInterface(&_Crafting.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Crafting *CraftingCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Crafting.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Crafting *CraftingSession) Symbol() (string, error) {
	return _Crafting.Contract.Symbol(&_Crafting.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Crafting *CraftingCallerSession) Symbol() (string, error) {
	return _Crafting.Contract.Symbol(&_Crafting.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Crafting *CraftingCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Crafting.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Crafting *CraftingSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Crafting.Contract.TokenURI(&_Crafting.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Crafting *CraftingCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Crafting.Contract.TokenURI(&_Crafting.CallOpts, tokenId)
}

// YearContract is a free data retrieval call binding the contract method 0x7230f8d4.
//
// Solidity: function yearContract() view returns(address)
func (_Crafting *CraftingCaller) YearContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crafting.contract.Call(opts, &out, "yearContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YearContract is a free data retrieval call binding the contract method 0x7230f8d4.
//
// Solidity: function yearContract() view returns(address)
func (_Crafting *CraftingSession) YearContract() (common.Address, error) {
	return _Crafting.Contract.YearContract(&_Crafting.CallOpts)
}

// YearContract is a free data retrieval call binding the contract method 0x7230f8d4.
//
// Solidity: function yearContract() view returns(address)
func (_Crafting *CraftingCallerSession) YearContract() (common.Address, error) {
	return _Crafting.Contract.YearContract(&_Crafting.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Crafting *CraftingTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Crafting.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Crafting *CraftingSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Crafting.Contract.Approve(&_Crafting.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Crafting *CraftingTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Crafting.Contract.Approve(&_Crafting.TransactOpts, to, tokenId)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Crafting *CraftingTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _Crafting.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Crafting *CraftingSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Crafting.Contract.ChangeAdmin(&_Crafting.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Crafting *CraftingTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Crafting.Contract.ChangeAdmin(&_Crafting.TransactOpts, _newAdmin)
}

// ChangeCrafter is a paid mutator transaction binding the contract method 0xe7d3b069.
//
// Solidity: function changeCrafter(address _newCrafter) returns()
func (_Crafting *CraftingTransactor) ChangeCrafter(opts *bind.TransactOpts, _newCrafter common.Address) (*types.Transaction, error) {
	return _Crafting.contract.Transact(opts, "changeCrafter", _newCrafter)
}

// ChangeCrafter is a paid mutator transaction binding the contract method 0xe7d3b069.
//
// Solidity: function changeCrafter(address _newCrafter) returns()
func (_Crafting *CraftingSession) ChangeCrafter(_newCrafter common.Address) (*types.Transaction, error) {
	return _Crafting.Contract.ChangeCrafter(&_Crafting.TransactOpts, _newCrafter)
}

// ChangeCrafter is a paid mutator transaction binding the contract method 0xe7d3b069.
//
// Solidity: function changeCrafter(address _newCrafter) returns()
func (_Crafting *CraftingTransactorSession) ChangeCrafter(_newCrafter common.Address) (*types.Transaction, error) {
	return _Crafting.Contract.ChangeCrafter(&_Crafting.TransactOpts, _newCrafter)
}

// ChangeToCrafted is a paid mutator transaction binding the contract method 0xe6b60d9b.
//
// Solidity: function changeToCrafted(uint256 tokenId) returns()
func (_Crafting *CraftingTransactor) ChangeToCrafted(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Crafting.contract.Transact(opts, "changeToCrafted", tokenId)
}

// ChangeToCrafted is a paid mutator transaction binding the contract method 0xe6b60d9b.
//
// Solidity: function changeToCrafted(uint256 tokenId) returns()
func (_Crafting *CraftingSession) ChangeToCrafted(tokenId *big.Int) (*types.Transaction, error) {
	return _Crafting.Contract.ChangeToCrafted(&_Crafting.TransactOpts, tokenId)
}

// ChangeToCrafted is a paid mutator transaction binding the contract method 0xe6b60d9b.
//
// Solidity: function changeToCrafted(uint256 tokenId) returns()
func (_Crafting *CraftingTransactorSession) ChangeToCrafted(tokenId *big.Int) (*types.Transaction, error) {
	return _Crafting.Contract.ChangeToCrafted(&_Crafting.TransactOpts, tokenId)
}

// CraftCollection is a paid mutator transaction binding the contract method 0xbf0837fe.
//
// Solidity: function craftCollection(address _dayMonthContract, address _yearContract, address _categoryContract, address _craftingCardContract, uint256 dayMonthTokenId, uint256 yearTokenId, uint256 categoryTokenId, uint256 craftingCardTokenId, string _tokenURI) returns()
func (_Crafting *CraftingTransactor) CraftCollection(opts *bind.TransactOpts, _dayMonthContract common.Address, _yearContract common.Address, _categoryContract common.Address, _craftingCardContract common.Address, dayMonthTokenId *big.Int, yearTokenId *big.Int, categoryTokenId *big.Int, craftingCardTokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _Crafting.contract.Transact(opts, "craftCollection", _dayMonthContract, _yearContract, _categoryContract, _craftingCardContract, dayMonthTokenId, yearTokenId, categoryTokenId, craftingCardTokenId, _tokenURI)
}

// CraftCollection is a paid mutator transaction binding the contract method 0xbf0837fe.
//
// Solidity: function craftCollection(address _dayMonthContract, address _yearContract, address _categoryContract, address _craftingCardContract, uint256 dayMonthTokenId, uint256 yearTokenId, uint256 categoryTokenId, uint256 craftingCardTokenId, string _tokenURI) returns()
func (_Crafting *CraftingSession) CraftCollection(_dayMonthContract common.Address, _yearContract common.Address, _categoryContract common.Address, _craftingCardContract common.Address, dayMonthTokenId *big.Int, yearTokenId *big.Int, categoryTokenId *big.Int, craftingCardTokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _Crafting.Contract.CraftCollection(&_Crafting.TransactOpts, _dayMonthContract, _yearContract, _categoryContract, _craftingCardContract, dayMonthTokenId, yearTokenId, categoryTokenId, craftingCardTokenId, _tokenURI)
}

// CraftCollection is a paid mutator transaction binding the contract method 0xbf0837fe.
//
// Solidity: function craftCollection(address _dayMonthContract, address _yearContract, address _categoryContract, address _craftingCardContract, uint256 dayMonthTokenId, uint256 yearTokenId, uint256 categoryTokenId, uint256 craftingCardTokenId, string _tokenURI) returns()
func (_Crafting *CraftingTransactorSession) CraftCollection(_dayMonthContract common.Address, _yearContract common.Address, _categoryContract common.Address, _craftingCardContract common.Address, dayMonthTokenId *big.Int, yearTokenId *big.Int, categoryTokenId *big.Int, craftingCardTokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _Crafting.Contract.CraftCollection(&_Crafting.TransactOpts, _dayMonthContract, _yearContract, _categoryContract, _craftingCardContract, dayMonthTokenId, yearTokenId, categoryTokenId, craftingCardTokenId, _tokenURI)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Crafting *CraftingTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crafting.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Crafting *CraftingSession) Initialize() (*types.Transaction, error) {
	return _Crafting.Contract.Initialize(&_Crafting.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Crafting *CraftingTransactorSession) Initialize() (*types.Transaction, error) {
	return _Crafting.Contract.Initialize(&_Crafting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Crafting *CraftingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crafting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Crafting *CraftingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Crafting.Contract.RenounceOwnership(&_Crafting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Crafting *CraftingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Crafting.Contract.RenounceOwnership(&_Crafting.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Crafting *CraftingTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Crafting.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Crafting *CraftingSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Crafting.Contract.SafeTransferFrom(&_Crafting.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Crafting *CraftingTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Crafting.Contract.SafeTransferFrom(&_Crafting.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Crafting *CraftingTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Crafting.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Crafting *CraftingSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Crafting.Contract.SafeTransferFrom0(&_Crafting.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Crafting *CraftingTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Crafting.Contract.SafeTransferFrom0(&_Crafting.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Crafting *CraftingTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Crafting.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Crafting *CraftingSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Crafting.Contract.SetApprovalForAll(&_Crafting.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Crafting *CraftingTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Crafting.Contract.SetApprovalForAll(&_Crafting.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Crafting *CraftingTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Crafting.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Crafting *CraftingSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Crafting.Contract.TransferFrom(&_Crafting.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Crafting *CraftingTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Crafting.Contract.TransferFrom(&_Crafting.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Crafting *CraftingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Crafting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Crafting *CraftingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Crafting.Contract.TransferOwnership(&_Crafting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Crafting *CraftingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Crafting.Contract.TransferOwnership(&_Crafting.TransactOpts, newOwner)
}

// CraftingApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Crafting contract.
type CraftingApprovalIterator struct {
	Event *CraftingApproval // Event containing the contract specifics and raw log

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
func (it *CraftingApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CraftingApproval)
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
		it.Event = new(CraftingApproval)
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
func (it *CraftingApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CraftingApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CraftingApproval represents a Approval event raised by the Crafting contract.
type CraftingApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Crafting *CraftingFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*CraftingApprovalIterator, error) {

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

	logs, sub, err := _Crafting.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CraftingApprovalIterator{contract: _Crafting.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Crafting *CraftingFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CraftingApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Crafting.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CraftingApproval)
				if err := _Crafting.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Crafting *CraftingFilterer) ParseApproval(log types.Log) (*CraftingApproval, error) {
	event := new(CraftingApproval)
	if err := _Crafting.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CraftingApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Crafting contract.
type CraftingApprovalForAllIterator struct {
	Event *CraftingApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CraftingApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CraftingApprovalForAll)
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
		it.Event = new(CraftingApprovalForAll)
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
func (it *CraftingApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CraftingApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CraftingApprovalForAll represents a ApprovalForAll event raised by the Crafting contract.
type CraftingApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Crafting *CraftingFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CraftingApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Crafting.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CraftingApprovalForAllIterator{contract: _Crafting.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Crafting *CraftingFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CraftingApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Crafting.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CraftingApprovalForAll)
				if err := _Crafting.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Crafting *CraftingFilterer) ParseApprovalForAll(log types.Log) (*CraftingApprovalForAll, error) {
	event := new(CraftingApprovalForAll)
	if err := _Crafting.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CraftingCollectionCraftedIterator is returned from FilterCollectionCrafted and is used to iterate over the raw logs and unpacked data for CollectionCrafted events raised by the Crafting contract.
type CraftingCollectionCraftedIterator struct {
	Event *CraftingCollectionCrafted // Event containing the contract specifics and raw log

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
func (it *CraftingCollectionCraftedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CraftingCollectionCrafted)
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
		it.Event = new(CraftingCollectionCrafted)
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
func (it *CraftingCollectionCraftedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CraftingCollectionCraftedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CraftingCollectionCrafted represents a CollectionCrafted event raised by the Crafting contract.
type CraftingCollectionCrafted struct {
	CollectionId        *big.Int
	DayMonthTokenId     *big.Int
	YearTokenId         *big.Int
	CategoryTokenId     *big.Int
	CraftingCardTokenId *big.Int
	TokenURI            string
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterCollectionCrafted is a free log retrieval operation binding the contract event 0xe5979afd3a6266ac0cf01741228fe2b459458306a8d7a2c3152da07ee0c2f2be.
//
// Solidity: event CollectionCrafted(uint256 indexed collectionId, uint256 dayMonthTokenId, uint256 yearTokenId, uint256 categoryTokenId, uint256 craftingCardTokenId, string tokenURI)
func (_Crafting *CraftingFilterer) FilterCollectionCrafted(opts *bind.FilterOpts, collectionId []*big.Int) (*CraftingCollectionCraftedIterator, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}

	logs, sub, err := _Crafting.contract.FilterLogs(opts, "CollectionCrafted", collectionIdRule)
	if err != nil {
		return nil, err
	}
	return &CraftingCollectionCraftedIterator{contract: _Crafting.contract, event: "CollectionCrafted", logs: logs, sub: sub}, nil
}

// WatchCollectionCrafted is a free log subscription operation binding the contract event 0xe5979afd3a6266ac0cf01741228fe2b459458306a8d7a2c3152da07ee0c2f2be.
//
// Solidity: event CollectionCrafted(uint256 indexed collectionId, uint256 dayMonthTokenId, uint256 yearTokenId, uint256 categoryTokenId, uint256 craftingCardTokenId, string tokenURI)
func (_Crafting *CraftingFilterer) WatchCollectionCrafted(opts *bind.WatchOpts, sink chan<- *CraftingCollectionCrafted, collectionId []*big.Int) (event.Subscription, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}

	logs, sub, err := _Crafting.contract.WatchLogs(opts, "CollectionCrafted", collectionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CraftingCollectionCrafted)
				if err := _Crafting.contract.UnpackLog(event, "CollectionCrafted", log); err != nil {
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

// ParseCollectionCrafted is a log parse operation binding the contract event 0xe5979afd3a6266ac0cf01741228fe2b459458306a8d7a2c3152da07ee0c2f2be.
//
// Solidity: event CollectionCrafted(uint256 indexed collectionId, uint256 dayMonthTokenId, uint256 yearTokenId, uint256 categoryTokenId, uint256 craftingCardTokenId, string tokenURI)
func (_Crafting *CraftingFilterer) ParseCollectionCrafted(log types.Log) (*CraftingCollectionCrafted, error) {
	event := new(CraftingCollectionCrafted)
	if err := _Crafting.contract.UnpackLog(event, "CollectionCrafted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CraftingCraftingContractChangedIterator is returned from FilterCraftingContractChanged and is used to iterate over the raw logs and unpacked data for CraftingContractChanged events raised by the Crafting contract.
type CraftingCraftingContractChangedIterator struct {
	Event *CraftingCraftingContractChanged // Event containing the contract specifics and raw log

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
func (it *CraftingCraftingContractChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CraftingCraftingContractChanged)
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
		it.Event = new(CraftingCraftingContractChanged)
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
func (it *CraftingCraftingContractChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CraftingCraftingContractChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CraftingCraftingContractChanged represents a CraftingContractChanged event raised by the Crafting contract.
type CraftingCraftingContractChanged struct {
	PreviousCrafter common.Address
	NewCrafter      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCraftingContractChanged is a free log retrieval operation binding the contract event 0x83149e90b909cde0e8099cb6998203f8d8dce5fa29151688382650f5637d7d88.
//
// Solidity: event CraftingContractChanged(address indexed previousCrafter, address indexed newCrafter)
func (_Crafting *CraftingFilterer) FilterCraftingContractChanged(opts *bind.FilterOpts, previousCrafter []common.Address, newCrafter []common.Address) (*CraftingCraftingContractChangedIterator, error) {

	var previousCrafterRule []interface{}
	for _, previousCrafterItem := range previousCrafter {
		previousCrafterRule = append(previousCrafterRule, previousCrafterItem)
	}
	var newCrafterRule []interface{}
	for _, newCrafterItem := range newCrafter {
		newCrafterRule = append(newCrafterRule, newCrafterItem)
	}

	logs, sub, err := _Crafting.contract.FilterLogs(opts, "CraftingContractChanged", previousCrafterRule, newCrafterRule)
	if err != nil {
		return nil, err
	}
	return &CraftingCraftingContractChangedIterator{contract: _Crafting.contract, event: "CraftingContractChanged", logs: logs, sub: sub}, nil
}

// WatchCraftingContractChanged is a free log subscription operation binding the contract event 0x83149e90b909cde0e8099cb6998203f8d8dce5fa29151688382650f5637d7d88.
//
// Solidity: event CraftingContractChanged(address indexed previousCrafter, address indexed newCrafter)
func (_Crafting *CraftingFilterer) WatchCraftingContractChanged(opts *bind.WatchOpts, sink chan<- *CraftingCraftingContractChanged, previousCrafter []common.Address, newCrafter []common.Address) (event.Subscription, error) {

	var previousCrafterRule []interface{}
	for _, previousCrafterItem := range previousCrafter {
		previousCrafterRule = append(previousCrafterRule, previousCrafterItem)
	}
	var newCrafterRule []interface{}
	for _, newCrafterItem := range newCrafter {
		newCrafterRule = append(newCrafterRule, newCrafterItem)
	}

	logs, sub, err := _Crafting.contract.WatchLogs(opts, "CraftingContractChanged", previousCrafterRule, newCrafterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CraftingCraftingContractChanged)
				if err := _Crafting.contract.UnpackLog(event, "CraftingContractChanged", log); err != nil {
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
func (_Crafting *CraftingFilterer) ParseCraftingContractChanged(log types.Log) (*CraftingCraftingContractChanged, error) {
	event := new(CraftingCraftingContractChanged)
	if err := _Crafting.contract.UnpackLog(event, "CraftingContractChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CraftingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Crafting contract.
type CraftingOwnershipTransferredIterator struct {
	Event *CraftingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CraftingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CraftingOwnershipTransferred)
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
		it.Event = new(CraftingOwnershipTransferred)
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
func (it *CraftingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CraftingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CraftingOwnershipTransferred represents a OwnershipTransferred event raised by the Crafting contract.
type CraftingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Crafting *CraftingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CraftingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Crafting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CraftingOwnershipTransferredIterator{contract: _Crafting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Crafting *CraftingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CraftingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Crafting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CraftingOwnershipTransferred)
				if err := _Crafting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Crafting *CraftingFilterer) ParseOwnershipTransferred(log types.Log) (*CraftingOwnershipTransferred, error) {
	event := new(CraftingOwnershipTransferred)
	if err := _Crafting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CraftingPartCraftedIterator is returned from FilterPartCrafted and is used to iterate over the raw logs and unpacked data for PartCrafted events raised by the Crafting contract.
type CraftingPartCraftedIterator struct {
	Event *CraftingPartCrafted // Event containing the contract specifics and raw log

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
func (it *CraftingPartCraftedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CraftingPartCrafted)
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
		it.Event = new(CraftingPartCrafted)
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
func (it *CraftingPartCraftedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CraftingPartCraftedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CraftingPartCrafted represents a PartCrafted event raised by the Crafting contract.
type CraftingPartCrafted struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPartCrafted is a free log retrieval operation binding the contract event 0x71094059908fc1a5d676895c3ab6ec0e8957a09dc7a09914b791a2be4dc7fb15.
//
// Solidity: event PartCrafted(uint256 indexed tokenId)
func (_Crafting *CraftingFilterer) FilterPartCrafted(opts *bind.FilterOpts, tokenId []*big.Int) (*CraftingPartCraftedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Crafting.contract.FilterLogs(opts, "PartCrafted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CraftingPartCraftedIterator{contract: _Crafting.contract, event: "PartCrafted", logs: logs, sub: sub}, nil
}

// WatchPartCrafted is a free log subscription operation binding the contract event 0x71094059908fc1a5d676895c3ab6ec0e8957a09dc7a09914b791a2be4dc7fb15.
//
// Solidity: event PartCrafted(uint256 indexed tokenId)
func (_Crafting *CraftingFilterer) WatchPartCrafted(opts *bind.WatchOpts, sink chan<- *CraftingPartCrafted, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Crafting.contract.WatchLogs(opts, "PartCrafted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CraftingPartCrafted)
				if err := _Crafting.contract.UnpackLog(event, "PartCrafted", log); err != nil {
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
func (_Crafting *CraftingFilterer) ParsePartCrafted(log types.Log) (*CraftingPartCrafted, error) {
	event := new(CraftingPartCrafted)
	if err := _Crafting.contract.UnpackLog(event, "PartCrafted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CraftingTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Crafting contract.
type CraftingTransferIterator struct {
	Event *CraftingTransfer // Event containing the contract specifics and raw log

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
func (it *CraftingTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CraftingTransfer)
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
		it.Event = new(CraftingTransfer)
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
func (it *CraftingTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CraftingTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CraftingTransfer represents a Transfer event raised by the Crafting contract.
type CraftingTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Crafting *CraftingFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*CraftingTransferIterator, error) {

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

	logs, sub, err := _Crafting.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CraftingTransferIterator{contract: _Crafting.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Crafting *CraftingFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CraftingTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Crafting.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CraftingTransfer)
				if err := _Crafting.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Crafting *CraftingFilterer) ParseTransfer(log types.Log) (*CraftingTransfer, error) {
	event := new(CraftingTransfer)
	if err := _Crafting.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
