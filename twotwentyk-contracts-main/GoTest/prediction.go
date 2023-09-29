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

// PredictionMetaData contains all meta data concerning the Prediction contract.
var PredictionMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"identityTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"triggerTokenId\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"CollectionCrafted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_triggersContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_craftingCardContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identityTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"triggerTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"craftingCardTokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"}],\"name\":\"craftCollection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"crafted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"craftingCardContract\",\"outputs\":[{\"internalType\":\"contractIPart\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"}],\"name\":\"getCollectionTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"identitiesContract\",\"outputs\":[{\"internalType\":\"contractIPart\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identitiesContract\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"triggersContract\",\"outputs\":[{\"internalType\":\"contractIPart\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"win\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"won\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PredictionABI is the input ABI used to generate the binding from.
// Deprecated: Use PredictionMetaData.ABI instead.
var PredictionABI = PredictionMetaData.ABI

// Prediction is an auto generated Go binding around an Ethereum contract.
type Prediction struct {
	PredictionCaller     // Read-only binding to the contract
	PredictionTransactor // Write-only binding to the contract
	PredictionFilterer   // Log filterer for contract events
}

// PredictionCaller is an auto generated read-only Go binding around an Ethereum contract.
type PredictionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PredictionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PredictionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PredictionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PredictionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PredictionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PredictionSession struct {
	Contract     *Prediction       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PredictionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PredictionCallerSession struct {
	Contract *PredictionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PredictionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PredictionTransactorSession struct {
	Contract     *PredictionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PredictionRaw is an auto generated low-level Go binding around an Ethereum contract.
type PredictionRaw struct {
	Contract *Prediction // Generic contract binding to access the raw methods on
}

// PredictionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PredictionCallerRaw struct {
	Contract *PredictionCaller // Generic read-only contract binding to access the raw methods on
}

// PredictionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PredictionTransactorRaw struct {
	Contract *PredictionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrediction creates a new instance of Prediction, bound to a specific deployed contract.
func NewPrediction(address common.Address, backend bind.ContractBackend) (*Prediction, error) {
	contract, err := bindPrediction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Prediction{PredictionCaller: PredictionCaller{contract: contract}, PredictionTransactor: PredictionTransactor{contract: contract}, PredictionFilterer: PredictionFilterer{contract: contract}}, nil
}

// NewPredictionCaller creates a new read-only instance of Prediction, bound to a specific deployed contract.
func NewPredictionCaller(address common.Address, caller bind.ContractCaller) (*PredictionCaller, error) {
	contract, err := bindPrediction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PredictionCaller{contract: contract}, nil
}

// NewPredictionTransactor creates a new write-only instance of Prediction, bound to a specific deployed contract.
func NewPredictionTransactor(address common.Address, transactor bind.ContractTransactor) (*PredictionTransactor, error) {
	contract, err := bindPrediction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PredictionTransactor{contract: contract}, nil
}

// NewPredictionFilterer creates a new log filterer instance of Prediction, bound to a specific deployed contract.
func NewPredictionFilterer(address common.Address, filterer bind.ContractFilterer) (*PredictionFilterer, error) {
	contract, err := bindPrediction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PredictionFilterer{contract: contract}, nil
}

// bindPrediction binds a generic wrapper to an already deployed contract.
func bindPrediction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PredictionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Prediction *PredictionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Prediction.Contract.PredictionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Prediction *PredictionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Prediction.Contract.PredictionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Prediction *PredictionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Prediction.Contract.PredictionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Prediction *PredictionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Prediction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Prediction *PredictionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Prediction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Prediction *PredictionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Prediction.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Prediction *PredictionCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Prediction *PredictionSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Prediction.Contract.BalanceOf(&_Prediction.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Prediction *PredictionCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Prediction.Contract.BalanceOf(&_Prediction.CallOpts, owner)
}

// Crafted is a free data retrieval call binding the contract method 0xcac4cc29.
//
// Solidity: function crafted(uint256 ) view returns(bool)
func (_Prediction *PredictionCaller) Crafted(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "crafted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Crafted is a free data retrieval call binding the contract method 0xcac4cc29.
//
// Solidity: function crafted(uint256 ) view returns(bool)
func (_Prediction *PredictionSession) Crafted(arg0 *big.Int) (bool, error) {
	return _Prediction.Contract.Crafted(&_Prediction.CallOpts, arg0)
}

// Crafted is a free data retrieval call binding the contract method 0xcac4cc29.
//
// Solidity: function crafted(uint256 ) view returns(bool)
func (_Prediction *PredictionCallerSession) Crafted(arg0 *big.Int) (bool, error) {
	return _Prediction.Contract.Crafted(&_Prediction.CallOpts, arg0)
}

// CraftingCardContract is a free data retrieval call binding the contract method 0x2c948296.
//
// Solidity: function craftingCardContract() view returns(address)
func (_Prediction *PredictionCaller) CraftingCardContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "craftingCardContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CraftingCardContract is a free data retrieval call binding the contract method 0x2c948296.
//
// Solidity: function craftingCardContract() view returns(address)
func (_Prediction *PredictionSession) CraftingCardContract() (common.Address, error) {
	return _Prediction.Contract.CraftingCardContract(&_Prediction.CallOpts)
}

// CraftingCardContract is a free data retrieval call binding the contract method 0x2c948296.
//
// Solidity: function craftingCardContract() view returns(address)
func (_Prediction *PredictionCallerSession) CraftingCardContract() (common.Address, error) {
	return _Prediction.Contract.CraftingCardContract(&_Prediction.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Prediction *PredictionCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Prediction *PredictionSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Prediction.Contract.GetApproved(&_Prediction.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Prediction *PredictionCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Prediction.Contract.GetApproved(&_Prediction.CallOpts, tokenId)
}

// GetCollectionTokens is a free data retrieval call binding the contract method 0x2bb88056.
//
// Solidity: function getCollectionTokens(uint256 collectionId) view returns(uint256, uint256[])
func (_Prediction *PredictionCaller) GetCollectionTokens(opts *bind.CallOpts, collectionId *big.Int) (*big.Int, []*big.Int, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "getCollectionTokens", collectionId)

	if err != nil {
		return *new(*big.Int), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetCollectionTokens is a free data retrieval call binding the contract method 0x2bb88056.
//
// Solidity: function getCollectionTokens(uint256 collectionId) view returns(uint256, uint256[])
func (_Prediction *PredictionSession) GetCollectionTokens(collectionId *big.Int) (*big.Int, []*big.Int, error) {
	return _Prediction.Contract.GetCollectionTokens(&_Prediction.CallOpts, collectionId)
}

// GetCollectionTokens is a free data retrieval call binding the contract method 0x2bb88056.
//
// Solidity: function getCollectionTokens(uint256 collectionId) view returns(uint256, uint256[])
func (_Prediction *PredictionCallerSession) GetCollectionTokens(collectionId *big.Int) (*big.Int, []*big.Int, error) {
	return _Prediction.Contract.GetCollectionTokens(&_Prediction.CallOpts, collectionId)
}

// IdentitiesContract is a free data retrieval call binding the contract method 0xbff246c2.
//
// Solidity: function identitiesContract() view returns(address)
func (_Prediction *PredictionCaller) IdentitiesContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "identitiesContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdentitiesContract is a free data retrieval call binding the contract method 0xbff246c2.
//
// Solidity: function identitiesContract() view returns(address)
func (_Prediction *PredictionSession) IdentitiesContract() (common.Address, error) {
	return _Prediction.Contract.IdentitiesContract(&_Prediction.CallOpts)
}

// IdentitiesContract is a free data retrieval call binding the contract method 0xbff246c2.
//
// Solidity: function identitiesContract() view returns(address)
func (_Prediction *PredictionCallerSession) IdentitiesContract() (common.Address, error) {
	return _Prediction.Contract.IdentitiesContract(&_Prediction.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Prediction *PredictionCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Prediction *PredictionSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Prediction.Contract.IsApprovedForAll(&_Prediction.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Prediction *PredictionCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Prediction.Contract.IsApprovedForAll(&_Prediction.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Prediction *PredictionCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Prediction *PredictionSession) Name() (string, error) {
	return _Prediction.Contract.Name(&_Prediction.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Prediction *PredictionCallerSession) Name() (string, error) {
	return _Prediction.Contract.Name(&_Prediction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Prediction *PredictionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Prediction *PredictionSession) Owner() (common.Address, error) {
	return _Prediction.Contract.Owner(&_Prediction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Prediction *PredictionCallerSession) Owner() (common.Address, error) {
	return _Prediction.Contract.Owner(&_Prediction.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Prediction *PredictionCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Prediction *PredictionSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Prediction.Contract.OwnerOf(&_Prediction.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Prediction *PredictionCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Prediction.Contract.OwnerOf(&_Prediction.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Prediction *PredictionCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Prediction *PredictionSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Prediction.Contract.SupportsInterface(&_Prediction.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Prediction *PredictionCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Prediction.Contract.SupportsInterface(&_Prediction.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Prediction *PredictionCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Prediction *PredictionSession) Symbol() (string, error) {
	return _Prediction.Contract.Symbol(&_Prediction.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Prediction *PredictionCallerSession) Symbol() (string, error) {
	return _Prediction.Contract.Symbol(&_Prediction.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Prediction *PredictionCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Prediction *PredictionSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Prediction.Contract.TokenURI(&_Prediction.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Prediction *PredictionCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Prediction.Contract.TokenURI(&_Prediction.CallOpts, tokenId)
}

// TriggersContract is a free data retrieval call binding the contract method 0x3f254c60.
//
// Solidity: function triggersContract() view returns(address)
func (_Prediction *PredictionCaller) TriggersContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "triggersContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TriggersContract is a free data retrieval call binding the contract method 0x3f254c60.
//
// Solidity: function triggersContract() view returns(address)
func (_Prediction *PredictionSession) TriggersContract() (common.Address, error) {
	return _Prediction.Contract.TriggersContract(&_Prediction.CallOpts)
}

// TriggersContract is a free data retrieval call binding the contract method 0x3f254c60.
//
// Solidity: function triggersContract() view returns(address)
func (_Prediction *PredictionCallerSession) TriggersContract() (common.Address, error) {
	return _Prediction.Contract.TriggersContract(&_Prediction.CallOpts)
}

// Won is a free data retrieval call binding the contract method 0x71ee1aba.
//
// Solidity: function won(address ) view returns(bool)
func (_Prediction *PredictionCaller) Won(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "won", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Won is a free data retrieval call binding the contract method 0x71ee1aba.
//
// Solidity: function won(address ) view returns(bool)
func (_Prediction *PredictionSession) Won(arg0 common.Address) (bool, error) {
	return _Prediction.Contract.Won(&_Prediction.CallOpts, arg0)
}

// Won is a free data retrieval call binding the contract method 0x71ee1aba.
//
// Solidity: function won(address ) view returns(bool)
func (_Prediction *PredictionCallerSession) Won(arg0 common.Address) (bool, error) {
	return _Prediction.Contract.Won(&_Prediction.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Prediction *PredictionTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Prediction *PredictionSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Prediction.Contract.Approve(&_Prediction.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Prediction *PredictionTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Prediction.Contract.Approve(&_Prediction.TransactOpts, to, tokenId)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Prediction *PredictionTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Prediction *PredictionSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Prediction.Contract.ChangeAdmin(&_Prediction.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Prediction *PredictionTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Prediction.Contract.ChangeAdmin(&_Prediction.TransactOpts, _newAdmin)
}

// CraftCollection is a paid mutator transaction binding the contract method 0x92781fd0.
//
// Solidity: function craftCollection(address _triggersContract, address _craftingCardContract, uint256 identityTokenId, uint256[] triggerTokenIds, uint256 craftingCardTokenId, string _tokenURI) returns()
func (_Prediction *PredictionTransactor) CraftCollection(opts *bind.TransactOpts, _triggersContract common.Address, _craftingCardContract common.Address, identityTokenId *big.Int, triggerTokenIds []*big.Int, craftingCardTokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "craftCollection", _triggersContract, _craftingCardContract, identityTokenId, triggerTokenIds, craftingCardTokenId, _tokenURI)
}

// CraftCollection is a paid mutator transaction binding the contract method 0x92781fd0.
//
// Solidity: function craftCollection(address _triggersContract, address _craftingCardContract, uint256 identityTokenId, uint256[] triggerTokenIds, uint256 craftingCardTokenId, string _tokenURI) returns()
func (_Prediction *PredictionSession) CraftCollection(_triggersContract common.Address, _craftingCardContract common.Address, identityTokenId *big.Int, triggerTokenIds []*big.Int, craftingCardTokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _Prediction.Contract.CraftCollection(&_Prediction.TransactOpts, _triggersContract, _craftingCardContract, identityTokenId, triggerTokenIds, craftingCardTokenId, _tokenURI)
}

// CraftCollection is a paid mutator transaction binding the contract method 0x92781fd0.
//
// Solidity: function craftCollection(address _triggersContract, address _craftingCardContract, uint256 identityTokenId, uint256[] triggerTokenIds, uint256 craftingCardTokenId, string _tokenURI) returns()
func (_Prediction *PredictionTransactorSession) CraftCollection(_triggersContract common.Address, _craftingCardContract common.Address, identityTokenId *big.Int, triggerTokenIds []*big.Int, craftingCardTokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _Prediction.Contract.CraftCollection(&_Prediction.TransactOpts, _triggersContract, _craftingCardContract, identityTokenId, triggerTokenIds, craftingCardTokenId, _tokenURI)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _identitiesContract) returns()
func (_Prediction *PredictionTransactor) Initialize(opts *bind.TransactOpts, _identitiesContract common.Address) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "initialize", _identitiesContract)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _identitiesContract) returns()
func (_Prediction *PredictionSession) Initialize(_identitiesContract common.Address) (*types.Transaction, error) {
	return _Prediction.Contract.Initialize(&_Prediction.TransactOpts, _identitiesContract)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _identitiesContract) returns()
func (_Prediction *PredictionTransactorSession) Initialize(_identitiesContract common.Address) (*types.Transaction, error) {
	return _Prediction.Contract.Initialize(&_Prediction.TransactOpts, _identitiesContract)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Prediction *PredictionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Prediction *PredictionSession) RenounceOwnership() (*types.Transaction, error) {
	return _Prediction.Contract.RenounceOwnership(&_Prediction.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Prediction *PredictionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Prediction.Contract.RenounceOwnership(&_Prediction.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Prediction *PredictionTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Prediction *PredictionSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Prediction.Contract.SafeTransferFrom(&_Prediction.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Prediction *PredictionTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Prediction.Contract.SafeTransferFrom(&_Prediction.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Prediction *PredictionTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Prediction *PredictionSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Prediction.Contract.SafeTransferFrom0(&_Prediction.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Prediction *PredictionTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Prediction.Contract.SafeTransferFrom0(&_Prediction.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Prediction *PredictionTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Prediction *PredictionSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Prediction.Contract.SetApprovalForAll(&_Prediction.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Prediction *PredictionTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Prediction.Contract.SetApprovalForAll(&_Prediction.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Prediction *PredictionTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Prediction *PredictionSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Prediction.Contract.TransferFrom(&_Prediction.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Prediction *PredictionTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Prediction.Contract.TransferFrom(&_Prediction.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Prediction *PredictionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Prediction *PredictionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Prediction.Contract.TransferOwnership(&_Prediction.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Prediction *PredictionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Prediction.Contract.TransferOwnership(&_Prediction.TransactOpts, newOwner)
}

// Win is a paid mutator transaction binding the contract method 0xa34cc845.
//
// Solidity: function win(address _address) returns()
func (_Prediction *PredictionTransactor) Win(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "win", _address)
}

// Win is a paid mutator transaction binding the contract method 0xa34cc845.
//
// Solidity: function win(address _address) returns()
func (_Prediction *PredictionSession) Win(_address common.Address) (*types.Transaction, error) {
	return _Prediction.Contract.Win(&_Prediction.TransactOpts, _address)
}

// Win is a paid mutator transaction binding the contract method 0xa34cc845.
//
// Solidity: function win(address _address) returns()
func (_Prediction *PredictionTransactorSession) Win(_address common.Address) (*types.Transaction, error) {
	return _Prediction.Contract.Win(&_Prediction.TransactOpts, _address)
}

// PredictionApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Prediction contract.
type PredictionApprovalIterator struct {
	Event *PredictionApproval // Event containing the contract specifics and raw log

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
func (it *PredictionApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionApproval)
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
		it.Event = new(PredictionApproval)
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
func (it *PredictionApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionApproval represents a Approval event raised by the Prediction contract.
type PredictionApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Prediction *PredictionFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*PredictionApprovalIterator, error) {

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

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PredictionApprovalIterator{contract: _Prediction.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Prediction *PredictionFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PredictionApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionApproval)
				if err := _Prediction.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Prediction *PredictionFilterer) ParseApproval(log types.Log) (*PredictionApproval, error) {
	event := new(PredictionApproval)
	if err := _Prediction.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Prediction contract.
type PredictionApprovalForAllIterator struct {
	Event *PredictionApprovalForAll // Event containing the contract specifics and raw log

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
func (it *PredictionApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionApprovalForAll)
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
		it.Event = new(PredictionApprovalForAll)
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
func (it *PredictionApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionApprovalForAll represents a ApprovalForAll event raised by the Prediction contract.
type PredictionApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Prediction *PredictionFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*PredictionApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &PredictionApprovalForAllIterator{contract: _Prediction.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Prediction *PredictionFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *PredictionApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionApprovalForAll)
				if err := _Prediction.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Prediction *PredictionFilterer) ParseApprovalForAll(log types.Log) (*PredictionApprovalForAll, error) {
	event := new(PredictionApprovalForAll)
	if err := _Prediction.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionCollectionCraftedIterator is returned from FilterCollectionCrafted and is used to iterate over the raw logs and unpacked data for CollectionCrafted events raised by the Prediction contract.
type PredictionCollectionCraftedIterator struct {
	Event *PredictionCollectionCrafted // Event containing the contract specifics and raw log

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
func (it *PredictionCollectionCraftedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionCollectionCrafted)
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
		it.Event = new(PredictionCollectionCrafted)
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
func (it *PredictionCollectionCraftedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionCollectionCraftedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionCollectionCrafted represents a CollectionCrafted event raised by the Prediction contract.
type PredictionCollectionCrafted struct {
	CollectionId    *big.Int
	IdentityTokenId *big.Int
	TriggerTokenId  []*big.Int
	TokenURI        string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCollectionCrafted is a free log retrieval operation binding the contract event 0x59392a125204831e61d97876a45b9667b68def2a5e303e6068932be0fd52372d.
//
// Solidity: event CollectionCrafted(uint256 indexed collectionId, uint256 identityTokenId, uint256[] triggerTokenId, string tokenURI)
func (_Prediction *PredictionFilterer) FilterCollectionCrafted(opts *bind.FilterOpts, collectionId []*big.Int) (*PredictionCollectionCraftedIterator, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "CollectionCrafted", collectionIdRule)
	if err != nil {
		return nil, err
	}
	return &PredictionCollectionCraftedIterator{contract: _Prediction.contract, event: "CollectionCrafted", logs: logs, sub: sub}, nil
}

// WatchCollectionCrafted is a free log subscription operation binding the contract event 0x59392a125204831e61d97876a45b9667b68def2a5e303e6068932be0fd52372d.
//
// Solidity: event CollectionCrafted(uint256 indexed collectionId, uint256 identityTokenId, uint256[] triggerTokenId, string tokenURI)
func (_Prediction *PredictionFilterer) WatchCollectionCrafted(opts *bind.WatchOpts, sink chan<- *PredictionCollectionCrafted, collectionId []*big.Int) (event.Subscription, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "CollectionCrafted", collectionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionCollectionCrafted)
				if err := _Prediction.contract.UnpackLog(event, "CollectionCrafted", log); err != nil {
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

// ParseCollectionCrafted is a log parse operation binding the contract event 0x59392a125204831e61d97876a45b9667b68def2a5e303e6068932be0fd52372d.
//
// Solidity: event CollectionCrafted(uint256 indexed collectionId, uint256 identityTokenId, uint256[] triggerTokenId, string tokenURI)
func (_Prediction *PredictionFilterer) ParseCollectionCrafted(log types.Log) (*PredictionCollectionCrafted, error) {
	event := new(PredictionCollectionCrafted)
	if err := _Prediction.contract.UnpackLog(event, "CollectionCrafted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Prediction contract.
type PredictionOwnershipTransferredIterator struct {
	Event *PredictionOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PredictionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionOwnershipTransferred)
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
		it.Event = new(PredictionOwnershipTransferred)
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
func (it *PredictionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionOwnershipTransferred represents a OwnershipTransferred event raised by the Prediction contract.
type PredictionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Prediction *PredictionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PredictionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PredictionOwnershipTransferredIterator{contract: _Prediction.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Prediction *PredictionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PredictionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionOwnershipTransferred)
				if err := _Prediction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Prediction *PredictionFilterer) ParseOwnershipTransferred(log types.Log) (*PredictionOwnershipTransferred, error) {
	event := new(PredictionOwnershipTransferred)
	if err := _Prediction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Prediction contract.
type PredictionTransferIterator struct {
	Event *PredictionTransfer // Event containing the contract specifics and raw log

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
func (it *PredictionTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionTransfer)
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
		it.Event = new(PredictionTransfer)
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
func (it *PredictionTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionTransfer represents a Transfer event raised by the Prediction contract.
type PredictionTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Prediction *PredictionFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*PredictionTransferIterator, error) {

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

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PredictionTransferIterator{contract: _Prediction.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Prediction *PredictionFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PredictionTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionTransfer)
				if err := _Prediction.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Prediction *PredictionFilterer) ParseTransfer(log types.Log) (*PredictionTransfer, error) {
	event := new(PredictionTransfer)
	if err := _Prediction.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
