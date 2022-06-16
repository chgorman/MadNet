// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PublicStakingABI is the input ABI used to generate the binding from.
const PublicStakingABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"payoutEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payoutAToken\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"}],\"name\":\"burnTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"payoutEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payoutAToken\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"circuitBreakerState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"}],\"name\":\"collectEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"payout\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"}],\"name\":\"collectEthTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"payout\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"}],\"name\":\"collectToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"payout\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"}],\"name\":\"collectTokenTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"payout\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"magic_\",\"type\":\"uint8\"}],\"name\":\"depositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"magic_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"depositToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"}],\"name\":\"estimateEthCollection\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"payout\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"estimateExcessEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"excess\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"estimateExcessToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"excess\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"}],\"name\":\"estimateTokenCollection\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"payout\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccumulatorScaleFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEthAccumulator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"accumulator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slush\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxMintLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"getMetamorphicContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"}],\"name\":\"getPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weightedShares\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"lockedStakingPosition\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freeAfter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawFreeAfter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatorEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatorToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenAccumulator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"accumulator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slush\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalReserveAToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalReserveEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSharesEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSharesToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockDuration_\",\"type\":\"uint256\"}],\"name\":\"lockOwnPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockDuration_\",\"type\":\"uint256\"}],\"name\":\"lockPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockDuration_\",\"type\":\"uint256\"}],\"name\":\"lockStakingPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockDuration_\",\"type\":\"uint256\"}],\"name\":\"lockWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockDuration_\",\"type\":\"uint256\"}],\"name\":\"mintTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"}],\"name\":\"skimExcessEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"excess\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"}],\"name\":\"skimExcessToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"excess\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tripCB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PublicStaking is an auto generated Go binding around an Ethereum contract.
type PublicStaking struct {
	PublicStakingCaller     // Read-only binding to the contract
	PublicStakingTransactor // Write-only binding to the contract
	PublicStakingFilterer   // Log filterer for contract events
}

// PublicStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicStakingSession struct {
	Contract     *PublicStaking    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PublicStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicStakingCallerSession struct {
	Contract *PublicStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PublicStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicStakingTransactorSession struct {
	Contract     *PublicStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PublicStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicStakingRaw struct {
	Contract *PublicStaking // Generic contract binding to access the raw methods on
}

// PublicStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicStakingCallerRaw struct {
	Contract *PublicStakingCaller // Generic read-only contract binding to access the raw methods on
}

// PublicStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicStakingTransactorRaw struct {
	Contract *PublicStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicStaking creates a new instance of PublicStaking, bound to a specific deployed contract.
func NewPublicStaking(address common.Address, backend bind.ContractBackend) (*PublicStaking, error) {
	contract, err := bindPublicStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PublicStaking{PublicStakingCaller: PublicStakingCaller{contract: contract}, PublicStakingTransactor: PublicStakingTransactor{contract: contract}, PublicStakingFilterer: PublicStakingFilterer{contract: contract}}, nil
}

// NewPublicStakingCaller creates a new read-only instance of PublicStaking, bound to a specific deployed contract.
func NewPublicStakingCaller(address common.Address, caller bind.ContractCaller) (*PublicStakingCaller, error) {
	contract, err := bindPublicStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicStakingCaller{contract: contract}, nil
}

// NewPublicStakingTransactor creates a new write-only instance of PublicStaking, bound to a specific deployed contract.
func NewPublicStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicStakingTransactor, error) {
	contract, err := bindPublicStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicStakingTransactor{contract: contract}, nil
}

// NewPublicStakingFilterer creates a new log filterer instance of PublicStaking, bound to a specific deployed contract.
func NewPublicStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicStakingFilterer, error) {
	contract, err := bindPublicStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicStakingFilterer{contract: contract}, nil
}

// bindPublicStaking binds a generic wrapper to an already deployed contract.
func bindPublicStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PublicStakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicStaking *PublicStakingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PublicStaking.Contract.PublicStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicStaking *PublicStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicStaking.Contract.PublicStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicStaking *PublicStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicStaking.Contract.PublicStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicStaking *PublicStakingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PublicStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicStaking *PublicStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicStaking *PublicStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicStaking.Contract.contract.Transact(opts, method, params...)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_PublicStaking *PublicStakingTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_PublicStaking *PublicStakingSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.Approve(&_PublicStaking.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_PublicStaking *PublicStakingTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.Approve(&_PublicStaking.TransactOpts, to, tokenId)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) returns(uint256)
func (_PublicStaking *PublicStakingTransactor) BalanceOf(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "balanceOf", owner)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) returns(uint256)
func (_PublicStaking *PublicStakingSession) BalanceOf(owner common.Address) (*types.Transaction, error) {
	return _PublicStaking.Contract.BalanceOf(&_PublicStaking.TransactOpts, owner)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) returns(uint256)
func (_PublicStaking *PublicStakingTransactorSession) BalanceOf(owner common.Address) (*types.Transaction, error) {
	return _PublicStaking.Contract.BalanceOf(&_PublicStaking.TransactOpts, owner)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenID_) returns(uint256 payoutEth, uint256 payoutAToken)
func (_PublicStaking *PublicStakingTransactor) Burn(opts *bind.TransactOpts, tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "burn", tokenID_)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenID_) returns(uint256 payoutEth, uint256 payoutAToken)
func (_PublicStaking *PublicStakingSession) Burn(tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.Burn(&_PublicStaking.TransactOpts, tokenID_)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenID_) returns(uint256 payoutEth, uint256 payoutAToken)
func (_PublicStaking *PublicStakingTransactorSession) Burn(tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.Burn(&_PublicStaking.TransactOpts, tokenID_)
}

// BurnTo is a paid mutator transaction binding the contract method 0xea785a5e.
//
// Solidity: function burnTo(address to_, uint256 tokenID_) returns(uint256 payoutEth, uint256 payoutAToken)
func (_PublicStaking *PublicStakingTransactor) BurnTo(opts *bind.TransactOpts, to_ common.Address, tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "burnTo", to_, tokenID_)
}

// BurnTo is a paid mutator transaction binding the contract method 0xea785a5e.
//
// Solidity: function burnTo(address to_, uint256 tokenID_) returns(uint256 payoutEth, uint256 payoutAToken)
func (_PublicStaking *PublicStakingSession) BurnTo(to_ common.Address, tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.BurnTo(&_PublicStaking.TransactOpts, to_, tokenID_)
}

// BurnTo is a paid mutator transaction binding the contract method 0xea785a5e.
//
// Solidity: function burnTo(address to_, uint256 tokenID_) returns(uint256 payoutEth, uint256 payoutAToken)
func (_PublicStaking *PublicStakingTransactorSession) BurnTo(to_ common.Address, tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.BurnTo(&_PublicStaking.TransactOpts, to_, tokenID_)
}

// CircuitBreakerState is a paid mutator transaction binding the contract method 0x89465c62.
//
// Solidity: function circuitBreakerState() returns(bool)
func (_PublicStaking *PublicStakingTransactor) CircuitBreakerState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "circuitBreakerState")
}

// CircuitBreakerState is a paid mutator transaction binding the contract method 0x89465c62.
//
// Solidity: function circuitBreakerState() returns(bool)
func (_PublicStaking *PublicStakingSession) CircuitBreakerState() (*types.Transaction, error) {
	return _PublicStaking.Contract.CircuitBreakerState(&_PublicStaking.TransactOpts)
}

// CircuitBreakerState is a paid mutator transaction binding the contract method 0x89465c62.
//
// Solidity: function circuitBreakerState() returns(bool)
func (_PublicStaking *PublicStakingTransactorSession) CircuitBreakerState() (*types.Transaction, error) {
	return _PublicStaking.Contract.CircuitBreakerState(&_PublicStaking.TransactOpts)
}

// CollectEth is a paid mutator transaction binding the contract method 0x2a0d8bd1.
//
// Solidity: function collectEth(uint256 tokenID_) returns(uint256 payout)
func (_PublicStaking *PublicStakingTransactor) CollectEth(opts *bind.TransactOpts, tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "collectEth", tokenID_)
}

// CollectEth is a paid mutator transaction binding the contract method 0x2a0d8bd1.
//
// Solidity: function collectEth(uint256 tokenID_) returns(uint256 payout)
func (_PublicStaking *PublicStakingSession) CollectEth(tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.CollectEth(&_PublicStaking.TransactOpts, tokenID_)
}

// CollectEth is a paid mutator transaction binding the contract method 0x2a0d8bd1.
//
// Solidity: function collectEth(uint256 tokenID_) returns(uint256 payout)
func (_PublicStaking *PublicStakingTransactorSession) CollectEth(tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.CollectEth(&_PublicStaking.TransactOpts, tokenID_)
}

// CollectEthTo is a paid mutator transaction binding the contract method 0xbe444379.
//
// Solidity: function collectEthTo(address to_, uint256 tokenID_) returns(uint256 payout)
func (_PublicStaking *PublicStakingTransactor) CollectEthTo(opts *bind.TransactOpts, to_ common.Address, tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "collectEthTo", to_, tokenID_)
}

// CollectEthTo is a paid mutator transaction binding the contract method 0xbe444379.
//
// Solidity: function collectEthTo(address to_, uint256 tokenID_) returns(uint256 payout)
func (_PublicStaking *PublicStakingSession) CollectEthTo(to_ common.Address, tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.CollectEthTo(&_PublicStaking.TransactOpts, to_, tokenID_)
}

// CollectEthTo is a paid mutator transaction binding the contract method 0xbe444379.
//
// Solidity: function collectEthTo(address to_, uint256 tokenID_) returns(uint256 payout)
func (_PublicStaking *PublicStakingTransactorSession) CollectEthTo(to_ common.Address, tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.CollectEthTo(&_PublicStaking.TransactOpts, to_, tokenID_)
}

// CollectToken is a paid mutator transaction binding the contract method 0xe35c3e28.
//
// Solidity: function collectToken(uint256 tokenID_) returns(uint256 payout)
func (_PublicStaking *PublicStakingTransactor) CollectToken(opts *bind.TransactOpts, tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "collectToken", tokenID_)
}

// CollectToken is a paid mutator transaction binding the contract method 0xe35c3e28.
//
// Solidity: function collectToken(uint256 tokenID_) returns(uint256 payout)
func (_PublicStaking *PublicStakingSession) CollectToken(tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.CollectToken(&_PublicStaking.TransactOpts, tokenID_)
}

// CollectToken is a paid mutator transaction binding the contract method 0xe35c3e28.
//
// Solidity: function collectToken(uint256 tokenID_) returns(uint256 payout)
func (_PublicStaking *PublicStakingTransactorSession) CollectToken(tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.CollectToken(&_PublicStaking.TransactOpts, tokenID_)
}

// CollectTokenTo is a paid mutator transaction binding the contract method 0x8853b950.
//
// Solidity: function collectTokenTo(address to_, uint256 tokenID_) returns(uint256 payout)
func (_PublicStaking *PublicStakingTransactor) CollectTokenTo(opts *bind.TransactOpts, to_ common.Address, tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "collectTokenTo", to_, tokenID_)
}

// CollectTokenTo is a paid mutator transaction binding the contract method 0x8853b950.
//
// Solidity: function collectTokenTo(address to_, uint256 tokenID_) returns(uint256 payout)
func (_PublicStaking *PublicStakingSession) CollectTokenTo(to_ common.Address, tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.CollectTokenTo(&_PublicStaking.TransactOpts, to_, tokenID_)
}

// CollectTokenTo is a paid mutator transaction binding the contract method 0x8853b950.
//
// Solidity: function collectTokenTo(address to_, uint256 tokenID_) returns(uint256 payout)
func (_PublicStaking *PublicStakingTransactorSession) CollectTokenTo(to_ common.Address, tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.CollectTokenTo(&_PublicStaking.TransactOpts, to_, tokenID_)
}

// DepositEth is a paid mutator transaction binding the contract method 0x99a89ecc.
//
// Solidity: function depositEth(uint8 magic_) returns()
func (_PublicStaking *PublicStakingTransactor) DepositEth(opts *bind.TransactOpts, magic_ uint8) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "depositEth", magic_)
}

// DepositEth is a paid mutator transaction binding the contract method 0x99a89ecc.
//
// Solidity: function depositEth(uint8 magic_) returns()
func (_PublicStaking *PublicStakingSession) DepositEth(magic_ uint8) (*types.Transaction, error) {
	return _PublicStaking.Contract.DepositEth(&_PublicStaking.TransactOpts, magic_)
}

// DepositEth is a paid mutator transaction binding the contract method 0x99a89ecc.
//
// Solidity: function depositEth(uint8 magic_) returns()
func (_PublicStaking *PublicStakingTransactorSession) DepositEth(magic_ uint8) (*types.Transaction, error) {
	return _PublicStaking.Contract.DepositEth(&_PublicStaking.TransactOpts, magic_)
}

// DepositToken is a paid mutator transaction binding the contract method 0x8191f5e5.
//
// Solidity: function depositToken(uint8 magic_, uint256 amount_) returns()
func (_PublicStaking *PublicStakingTransactor) DepositToken(opts *bind.TransactOpts, magic_ uint8, amount_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "depositToken", magic_, amount_)
}

// DepositToken is a paid mutator transaction binding the contract method 0x8191f5e5.
//
// Solidity: function depositToken(uint8 magic_, uint256 amount_) returns()
func (_PublicStaking *PublicStakingSession) DepositToken(magic_ uint8, amount_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.DepositToken(&_PublicStaking.TransactOpts, magic_, amount_)
}

// DepositToken is a paid mutator transaction binding the contract method 0x8191f5e5.
//
// Solidity: function depositToken(uint8 magic_, uint256 amount_) returns()
func (_PublicStaking *PublicStakingTransactorSession) DepositToken(magic_ uint8, amount_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.DepositToken(&_PublicStaking.TransactOpts, magic_, amount_)
}

// EstimateEthCollection is a paid mutator transaction binding the contract method 0x20ea0d48.
//
// Solidity: function estimateEthCollection(uint256 tokenID_) returns(uint256 payout)
func (_PublicStaking *PublicStakingTransactor) EstimateEthCollection(opts *bind.TransactOpts, tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "estimateEthCollection", tokenID_)
}

// EstimateEthCollection is a paid mutator transaction binding the contract method 0x20ea0d48.
//
// Solidity: function estimateEthCollection(uint256 tokenID_) returns(uint256 payout)
func (_PublicStaking *PublicStakingSession) EstimateEthCollection(tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.EstimateEthCollection(&_PublicStaking.TransactOpts, tokenID_)
}

// EstimateEthCollection is a paid mutator transaction binding the contract method 0x20ea0d48.
//
// Solidity: function estimateEthCollection(uint256 tokenID_) returns(uint256 payout)
func (_PublicStaking *PublicStakingTransactorSession) EstimateEthCollection(tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.EstimateEthCollection(&_PublicStaking.TransactOpts, tokenID_)
}

// EstimateExcessEth is a paid mutator transaction binding the contract method 0x905953ed.
//
// Solidity: function estimateExcessEth() returns(uint256 excess)
func (_PublicStaking *PublicStakingTransactor) EstimateExcessEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "estimateExcessEth")
}

// EstimateExcessEth is a paid mutator transaction binding the contract method 0x905953ed.
//
// Solidity: function estimateExcessEth() returns(uint256 excess)
func (_PublicStaking *PublicStakingSession) EstimateExcessEth() (*types.Transaction, error) {
	return _PublicStaking.Contract.EstimateExcessEth(&_PublicStaking.TransactOpts)
}

// EstimateExcessEth is a paid mutator transaction binding the contract method 0x905953ed.
//
// Solidity: function estimateExcessEth() returns(uint256 excess)
func (_PublicStaking *PublicStakingTransactorSession) EstimateExcessEth() (*types.Transaction, error) {
	return _PublicStaking.Contract.EstimateExcessEth(&_PublicStaking.TransactOpts)
}

// EstimateExcessToken is a paid mutator transaction binding the contract method 0x3eed3eff.
//
// Solidity: function estimateExcessToken() returns(uint256 excess)
func (_PublicStaking *PublicStakingTransactor) EstimateExcessToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "estimateExcessToken")
}

// EstimateExcessToken is a paid mutator transaction binding the contract method 0x3eed3eff.
//
// Solidity: function estimateExcessToken() returns(uint256 excess)
func (_PublicStaking *PublicStakingSession) EstimateExcessToken() (*types.Transaction, error) {
	return _PublicStaking.Contract.EstimateExcessToken(&_PublicStaking.TransactOpts)
}

// EstimateExcessToken is a paid mutator transaction binding the contract method 0x3eed3eff.
//
// Solidity: function estimateExcessToken() returns(uint256 excess)
func (_PublicStaking *PublicStakingTransactorSession) EstimateExcessToken() (*types.Transaction, error) {
	return _PublicStaking.Contract.EstimateExcessToken(&_PublicStaking.TransactOpts)
}

// EstimateTokenCollection is a paid mutator transaction binding the contract method 0x93c5748d.
//
// Solidity: function estimateTokenCollection(uint256 tokenID_) returns(uint256 payout)
func (_PublicStaking *PublicStakingTransactor) EstimateTokenCollection(opts *bind.TransactOpts, tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "estimateTokenCollection", tokenID_)
}

// EstimateTokenCollection is a paid mutator transaction binding the contract method 0x93c5748d.
//
// Solidity: function estimateTokenCollection(uint256 tokenID_) returns(uint256 payout)
func (_PublicStaking *PublicStakingSession) EstimateTokenCollection(tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.EstimateTokenCollection(&_PublicStaking.TransactOpts, tokenID_)
}

// EstimateTokenCollection is a paid mutator transaction binding the contract method 0x93c5748d.
//
// Solidity: function estimateTokenCollection(uint256 tokenID_) returns(uint256 payout)
func (_PublicStaking *PublicStakingTransactorSession) EstimateTokenCollection(tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.EstimateTokenCollection(&_PublicStaking.TransactOpts, tokenID_)
}

// GetAccumulatorScaleFactor is a paid mutator transaction binding the contract method 0x99785132.
//
// Solidity: function getAccumulatorScaleFactor() returns(uint256)
func (_PublicStaking *PublicStakingTransactor) GetAccumulatorScaleFactor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "getAccumulatorScaleFactor")
}

// GetAccumulatorScaleFactor is a paid mutator transaction binding the contract method 0x99785132.
//
// Solidity: function getAccumulatorScaleFactor() returns(uint256)
func (_PublicStaking *PublicStakingSession) GetAccumulatorScaleFactor() (*types.Transaction, error) {
	return _PublicStaking.Contract.GetAccumulatorScaleFactor(&_PublicStaking.TransactOpts)
}

// GetAccumulatorScaleFactor is a paid mutator transaction binding the contract method 0x99785132.
//
// Solidity: function getAccumulatorScaleFactor() returns(uint256)
func (_PublicStaking *PublicStakingTransactorSession) GetAccumulatorScaleFactor() (*types.Transaction, error) {
	return _PublicStaking.Contract.GetAccumulatorScaleFactor(&_PublicStaking.TransactOpts)
}

// GetApproved is a paid mutator transaction binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) returns(address)
func (_PublicStaking *PublicStakingTransactor) GetApproved(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "getApproved", tokenId)
}

// GetApproved is a paid mutator transaction binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) returns(address)
func (_PublicStaking *PublicStakingSession) GetApproved(tokenId *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.GetApproved(&_PublicStaking.TransactOpts, tokenId)
}

// GetApproved is a paid mutator transaction binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) returns(address)
func (_PublicStaking *PublicStakingTransactorSession) GetApproved(tokenId *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.GetApproved(&_PublicStaking.TransactOpts, tokenId)
}

// GetEthAccumulator is a paid mutator transaction binding the contract method 0x548652d2.
//
// Solidity: function getEthAccumulator() returns(uint256 accumulator, uint256 slush)
func (_PublicStaking *PublicStakingTransactor) GetEthAccumulator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "getEthAccumulator")
}

// GetEthAccumulator is a paid mutator transaction binding the contract method 0x548652d2.
//
// Solidity: function getEthAccumulator() returns(uint256 accumulator, uint256 slush)
func (_PublicStaking *PublicStakingSession) GetEthAccumulator() (*types.Transaction, error) {
	return _PublicStaking.Contract.GetEthAccumulator(&_PublicStaking.TransactOpts)
}

// GetEthAccumulator is a paid mutator transaction binding the contract method 0x548652d2.
//
// Solidity: function getEthAccumulator() returns(uint256 accumulator, uint256 slush)
func (_PublicStaking *PublicStakingTransactorSession) GetEthAccumulator() (*types.Transaction, error) {
	return _PublicStaking.Contract.GetEthAccumulator(&_PublicStaking.TransactOpts)
}

// GetMaxMintLock is a paid mutator transaction binding the contract method 0x090f70f0.
//
// Solidity: function getMaxMintLock() returns(uint256)
func (_PublicStaking *PublicStakingTransactor) GetMaxMintLock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "getMaxMintLock")
}

// GetMaxMintLock is a paid mutator transaction binding the contract method 0x090f70f0.
//
// Solidity: function getMaxMintLock() returns(uint256)
func (_PublicStaking *PublicStakingSession) GetMaxMintLock() (*types.Transaction, error) {
	return _PublicStaking.Contract.GetMaxMintLock(&_PublicStaking.TransactOpts)
}

// GetMaxMintLock is a paid mutator transaction binding the contract method 0x090f70f0.
//
// Solidity: function getMaxMintLock() returns(uint256)
func (_PublicStaking *PublicStakingTransactorSession) GetMaxMintLock() (*types.Transaction, error) {
	return _PublicStaking.Contract.GetMaxMintLock(&_PublicStaking.TransactOpts)
}

// GetMetamorphicContractAddress is a paid mutator transaction binding the contract method 0x8653a465.
//
// Solidity: function getMetamorphicContractAddress(bytes32 _salt, address _factory) returns(address)
func (_PublicStaking *PublicStakingTransactor) GetMetamorphicContractAddress(opts *bind.TransactOpts, _salt [32]byte, _factory common.Address) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "getMetamorphicContractAddress", _salt, _factory)
}

// GetMetamorphicContractAddress is a paid mutator transaction binding the contract method 0x8653a465.
//
// Solidity: function getMetamorphicContractAddress(bytes32 _salt, address _factory) returns(address)
func (_PublicStaking *PublicStakingSession) GetMetamorphicContractAddress(_salt [32]byte, _factory common.Address) (*types.Transaction, error) {
	return _PublicStaking.Contract.GetMetamorphicContractAddress(&_PublicStaking.TransactOpts, _salt, _factory)
}

// GetMetamorphicContractAddress is a paid mutator transaction binding the contract method 0x8653a465.
//
// Solidity: function getMetamorphicContractAddress(bytes32 _salt, address _factory) returns(address)
func (_PublicStaking *PublicStakingTransactorSession) GetMetamorphicContractAddress(_salt [32]byte, _factory common.Address) (*types.Transaction, error) {
	return _PublicStaking.Contract.GetMetamorphicContractAddress(&_PublicStaking.TransactOpts, _salt, _factory)
}

// GetPosition is a paid mutator transaction binding the contract method 0xeb02c301.
//
// Solidity: function getPosition(uint256 tokenID_) returns(uint256 weightedShares, bool lockedStakingPosition, uint256 shares, uint256 freeAfter, uint256 withdrawFreeAfter, uint256 accumulatorEth, uint256 accumulatorToken)
func (_PublicStaking *PublicStakingTransactor) GetPosition(opts *bind.TransactOpts, tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "getPosition", tokenID_)
}

// GetPosition is a paid mutator transaction binding the contract method 0xeb02c301.
//
// Solidity: function getPosition(uint256 tokenID_) returns(uint256 weightedShares, bool lockedStakingPosition, uint256 shares, uint256 freeAfter, uint256 withdrawFreeAfter, uint256 accumulatorEth, uint256 accumulatorToken)
func (_PublicStaking *PublicStakingSession) GetPosition(tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.GetPosition(&_PublicStaking.TransactOpts, tokenID_)
}

// GetPosition is a paid mutator transaction binding the contract method 0xeb02c301.
//
// Solidity: function getPosition(uint256 tokenID_) returns(uint256 weightedShares, bool lockedStakingPosition, uint256 shares, uint256 freeAfter, uint256 withdrawFreeAfter, uint256 accumulatorEth, uint256 accumulatorToken)
func (_PublicStaking *PublicStakingTransactorSession) GetPosition(tokenID_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.GetPosition(&_PublicStaking.TransactOpts, tokenID_)
}

// GetTokenAccumulator is a paid mutator transaction binding the contract method 0xc47c6e14.
//
// Solidity: function getTokenAccumulator() returns(uint256 accumulator, uint256 slush)
func (_PublicStaking *PublicStakingTransactor) GetTokenAccumulator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "getTokenAccumulator")
}

// GetTokenAccumulator is a paid mutator transaction binding the contract method 0xc47c6e14.
//
// Solidity: function getTokenAccumulator() returns(uint256 accumulator, uint256 slush)
func (_PublicStaking *PublicStakingSession) GetTokenAccumulator() (*types.Transaction, error) {
	return _PublicStaking.Contract.GetTokenAccumulator(&_PublicStaking.TransactOpts)
}

// GetTokenAccumulator is a paid mutator transaction binding the contract method 0xc47c6e14.
//
// Solidity: function getTokenAccumulator() returns(uint256 accumulator, uint256 slush)
func (_PublicStaking *PublicStakingTransactorSession) GetTokenAccumulator() (*types.Transaction, error) {
	return _PublicStaking.Contract.GetTokenAccumulator(&_PublicStaking.TransactOpts)
}

// GetTotalReserveAToken is a paid mutator transaction binding the contract method 0x856de8d2.
//
// Solidity: function getTotalReserveAToken() returns(uint256)
func (_PublicStaking *PublicStakingTransactor) GetTotalReserveAToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "getTotalReserveAToken")
}

// GetTotalReserveAToken is a paid mutator transaction binding the contract method 0x856de8d2.
//
// Solidity: function getTotalReserveAToken() returns(uint256)
func (_PublicStaking *PublicStakingSession) GetTotalReserveAToken() (*types.Transaction, error) {
	return _PublicStaking.Contract.GetTotalReserveAToken(&_PublicStaking.TransactOpts)
}

// GetTotalReserveAToken is a paid mutator transaction binding the contract method 0x856de8d2.
//
// Solidity: function getTotalReserveAToken() returns(uint256)
func (_PublicStaking *PublicStakingTransactorSession) GetTotalReserveAToken() (*types.Transaction, error) {
	return _PublicStaking.Contract.GetTotalReserveAToken(&_PublicStaking.TransactOpts)
}

// GetTotalReserveEth is a paid mutator transaction binding the contract method 0x19b8be2f.
//
// Solidity: function getTotalReserveEth() returns(uint256)
func (_PublicStaking *PublicStakingTransactor) GetTotalReserveEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "getTotalReserveEth")
}

// GetTotalReserveEth is a paid mutator transaction binding the contract method 0x19b8be2f.
//
// Solidity: function getTotalReserveEth() returns(uint256)
func (_PublicStaking *PublicStakingSession) GetTotalReserveEth() (*types.Transaction, error) {
	return _PublicStaking.Contract.GetTotalReserveEth(&_PublicStaking.TransactOpts)
}

// GetTotalReserveEth is a paid mutator transaction binding the contract method 0x19b8be2f.
//
// Solidity: function getTotalReserveEth() returns(uint256)
func (_PublicStaking *PublicStakingTransactorSession) GetTotalReserveEth() (*types.Transaction, error) {
	return _PublicStaking.Contract.GetTotalReserveEth(&_PublicStaking.TransactOpts)
}

// GetTotalSharesEth is a paid mutator transaction binding the contract method 0x6b0b5eae.
//
// Solidity: function getTotalSharesEth() returns(uint256)
func (_PublicStaking *PublicStakingTransactor) GetTotalSharesEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "getTotalSharesEth")
}

// GetTotalSharesEth is a paid mutator transaction binding the contract method 0x6b0b5eae.
//
// Solidity: function getTotalSharesEth() returns(uint256)
func (_PublicStaking *PublicStakingSession) GetTotalSharesEth() (*types.Transaction, error) {
	return _PublicStaking.Contract.GetTotalSharesEth(&_PublicStaking.TransactOpts)
}

// GetTotalSharesEth is a paid mutator transaction binding the contract method 0x6b0b5eae.
//
// Solidity: function getTotalSharesEth() returns(uint256)
func (_PublicStaking *PublicStakingTransactorSession) GetTotalSharesEth() (*types.Transaction, error) {
	return _PublicStaking.Contract.GetTotalSharesEth(&_PublicStaking.TransactOpts)
}

// GetTotalSharesToken is a paid mutator transaction binding the contract method 0x114be736.
//
// Solidity: function getTotalSharesToken() returns(uint256)
func (_PublicStaking *PublicStakingTransactor) GetTotalSharesToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "getTotalSharesToken")
}

// GetTotalSharesToken is a paid mutator transaction binding the contract method 0x114be736.
//
// Solidity: function getTotalSharesToken() returns(uint256)
func (_PublicStaking *PublicStakingSession) GetTotalSharesToken() (*types.Transaction, error) {
	return _PublicStaking.Contract.GetTotalSharesToken(&_PublicStaking.TransactOpts)
}

// GetTotalSharesToken is a paid mutator transaction binding the contract method 0x114be736.
//
// Solidity: function getTotalSharesToken() returns(uint256)
func (_PublicStaking *PublicStakingTransactorSession) GetTotalSharesToken() (*types.Transaction, error) {
	return _PublicStaking.Contract.GetTotalSharesToken(&_PublicStaking.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PublicStaking *PublicStakingTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PublicStaking *PublicStakingSession) Initialize() (*types.Transaction, error) {
	return _PublicStaking.Contract.Initialize(&_PublicStaking.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PublicStaking *PublicStakingTransactorSession) Initialize() (*types.Transaction, error) {
	return _PublicStaking.Contract.Initialize(&_PublicStaking.TransactOpts)
}

// IsApprovedForAll is a paid mutator transaction binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) returns(bool)
func (_PublicStaking *PublicStakingTransactor) IsApprovedForAll(opts *bind.TransactOpts, owner common.Address, operator common.Address) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "isApprovedForAll", owner, operator)
}

// IsApprovedForAll is a paid mutator transaction binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) returns(bool)
func (_PublicStaking *PublicStakingSession) IsApprovedForAll(owner common.Address, operator common.Address) (*types.Transaction, error) {
	return _PublicStaking.Contract.IsApprovedForAll(&_PublicStaking.TransactOpts, owner, operator)
}

// IsApprovedForAll is a paid mutator transaction binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) returns(bool)
func (_PublicStaking *PublicStakingTransactorSession) IsApprovedForAll(owner common.Address, operator common.Address) (*types.Transaction, error) {
	return _PublicStaking.Contract.IsApprovedForAll(&_PublicStaking.TransactOpts, owner, operator)
}

// LockOwnPosition is a paid mutator transaction binding the contract method 0xe42a673c.
//
// Solidity: function lockOwnPosition(uint256 tokenID_, uint256 lockDuration_) returns(uint256)
func (_PublicStaking *PublicStakingTransactor) LockOwnPosition(opts *bind.TransactOpts, tokenID_ *big.Int, lockDuration_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "lockOwnPosition", tokenID_, lockDuration_)
}

// LockOwnPosition is a paid mutator transaction binding the contract method 0xe42a673c.
//
// Solidity: function lockOwnPosition(uint256 tokenID_, uint256 lockDuration_) returns(uint256)
func (_PublicStaking *PublicStakingSession) LockOwnPosition(tokenID_ *big.Int, lockDuration_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.LockOwnPosition(&_PublicStaking.TransactOpts, tokenID_, lockDuration_)
}

// LockOwnPosition is a paid mutator transaction binding the contract method 0xe42a673c.
//
// Solidity: function lockOwnPosition(uint256 tokenID_, uint256 lockDuration_) returns(uint256)
func (_PublicStaking *PublicStakingTransactorSession) LockOwnPosition(tokenID_ *big.Int, lockDuration_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.LockOwnPosition(&_PublicStaking.TransactOpts, tokenID_, lockDuration_)
}

// LockPosition is a paid mutator transaction binding the contract method 0x0cc65dfb.
//
// Solidity: function lockPosition(address caller_, uint256 tokenID_, uint256 lockDuration_) returns(uint256)
func (_PublicStaking *PublicStakingTransactor) LockPosition(opts *bind.TransactOpts, caller_ common.Address, tokenID_ *big.Int, lockDuration_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "lockPosition", caller_, tokenID_, lockDuration_)
}

// LockPosition is a paid mutator transaction binding the contract method 0x0cc65dfb.
//
// Solidity: function lockPosition(address caller_, uint256 tokenID_, uint256 lockDuration_) returns(uint256)
func (_PublicStaking *PublicStakingSession) LockPosition(caller_ common.Address, tokenID_ *big.Int, lockDuration_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.LockPosition(&_PublicStaking.TransactOpts, caller_, tokenID_, lockDuration_)
}

// LockPosition is a paid mutator transaction binding the contract method 0x0cc65dfb.
//
// Solidity: function lockPosition(address caller_, uint256 tokenID_, uint256 lockDuration_) returns(uint256)
func (_PublicStaking *PublicStakingTransactorSession) LockPosition(caller_ common.Address, tokenID_ *big.Int, lockDuration_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.LockPosition(&_PublicStaking.TransactOpts, caller_, tokenID_, lockDuration_)
}

// LockStakingPosition is a paid mutator transaction binding the contract method 0xa8e876b2.
//
// Solidity: function lockStakingPosition(uint256 tokenID_, uint256 lockDuration_) returns()
func (_PublicStaking *PublicStakingTransactor) LockStakingPosition(opts *bind.TransactOpts, tokenID_ *big.Int, lockDuration_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "lockStakingPosition", tokenID_, lockDuration_)
}

// LockStakingPosition is a paid mutator transaction binding the contract method 0xa8e876b2.
//
// Solidity: function lockStakingPosition(uint256 tokenID_, uint256 lockDuration_) returns()
func (_PublicStaking *PublicStakingSession) LockStakingPosition(tokenID_ *big.Int, lockDuration_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.LockStakingPosition(&_PublicStaking.TransactOpts, tokenID_, lockDuration_)
}

// LockStakingPosition is a paid mutator transaction binding the contract method 0xa8e876b2.
//
// Solidity: function lockStakingPosition(uint256 tokenID_, uint256 lockDuration_) returns()
func (_PublicStaking *PublicStakingTransactorSession) LockStakingPosition(tokenID_ *big.Int, lockDuration_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.LockStakingPosition(&_PublicStaking.TransactOpts, tokenID_, lockDuration_)
}

// LockWithdraw is a paid mutator transaction binding the contract method 0x0e4eb15b.
//
// Solidity: function lockWithdraw(uint256 tokenID_, uint256 lockDuration_) returns(uint256)
func (_PublicStaking *PublicStakingTransactor) LockWithdraw(opts *bind.TransactOpts, tokenID_ *big.Int, lockDuration_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "lockWithdraw", tokenID_, lockDuration_)
}

// LockWithdraw is a paid mutator transaction binding the contract method 0x0e4eb15b.
//
// Solidity: function lockWithdraw(uint256 tokenID_, uint256 lockDuration_) returns(uint256)
func (_PublicStaking *PublicStakingSession) LockWithdraw(tokenID_ *big.Int, lockDuration_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.LockWithdraw(&_PublicStaking.TransactOpts, tokenID_, lockDuration_)
}

// LockWithdraw is a paid mutator transaction binding the contract method 0x0e4eb15b.
//
// Solidity: function lockWithdraw(uint256 tokenID_, uint256 lockDuration_) returns(uint256)
func (_PublicStaking *PublicStakingTransactorSession) LockWithdraw(tokenID_ *big.Int, lockDuration_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.LockWithdraw(&_PublicStaking.TransactOpts, tokenID_, lockDuration_)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount_) returns(uint256 tokenID)
func (_PublicStaking *PublicStakingTransactor) Mint(opts *bind.TransactOpts, amount_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "mint", amount_)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount_) returns(uint256 tokenID)
func (_PublicStaking *PublicStakingSession) Mint(amount_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.Mint(&_PublicStaking.TransactOpts, amount_)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount_) returns(uint256 tokenID)
func (_PublicStaking *PublicStakingTransactorSession) Mint(amount_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.Mint(&_PublicStaking.TransactOpts, amount_)
}

// MintTo is a paid mutator transaction binding the contract method 0x2baf2acb.
//
// Solidity: function mintTo(address to_, uint256 amount_, uint256 lockDuration_) returns(uint256 tokenID)
func (_PublicStaking *PublicStakingTransactor) MintTo(opts *bind.TransactOpts, to_ common.Address, amount_ *big.Int, lockDuration_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "mintTo", to_, amount_, lockDuration_)
}

// MintTo is a paid mutator transaction binding the contract method 0x2baf2acb.
//
// Solidity: function mintTo(address to_, uint256 amount_, uint256 lockDuration_) returns(uint256 tokenID)
func (_PublicStaking *PublicStakingSession) MintTo(to_ common.Address, amount_ *big.Int, lockDuration_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.MintTo(&_PublicStaking.TransactOpts, to_, amount_, lockDuration_)
}

// MintTo is a paid mutator transaction binding the contract method 0x2baf2acb.
//
// Solidity: function mintTo(address to_, uint256 amount_, uint256 lockDuration_) returns(uint256 tokenID)
func (_PublicStaking *PublicStakingTransactorSession) MintTo(to_ common.Address, amount_ *big.Int, lockDuration_ *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.MintTo(&_PublicStaking.TransactOpts, to_, amount_, lockDuration_)
}

// Name is a paid mutator transaction binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_PublicStaking *PublicStakingTransactor) Name(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "name")
}

// Name is a paid mutator transaction binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_PublicStaking *PublicStakingSession) Name() (*types.Transaction, error) {
	return _PublicStaking.Contract.Name(&_PublicStaking.TransactOpts)
}

// Name is a paid mutator transaction binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_PublicStaking *PublicStakingTransactorSession) Name() (*types.Transaction, error) {
	return _PublicStaking.Contract.Name(&_PublicStaking.TransactOpts)
}

// OwnerOf is a paid mutator transaction binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) returns(address)
func (_PublicStaking *PublicStakingTransactor) OwnerOf(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "ownerOf", tokenId)
}

// OwnerOf is a paid mutator transaction binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) returns(address)
func (_PublicStaking *PublicStakingSession) OwnerOf(tokenId *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.OwnerOf(&_PublicStaking.TransactOpts, tokenId)
}

// OwnerOf is a paid mutator transaction binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) returns(address)
func (_PublicStaking *PublicStakingTransactorSession) OwnerOf(tokenId *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.OwnerOf(&_PublicStaking.TransactOpts, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_PublicStaking *PublicStakingTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_PublicStaking *PublicStakingSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.SafeTransferFrom(&_PublicStaking.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_PublicStaking *PublicStakingTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.SafeTransferFrom(&_PublicStaking.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_PublicStaking *PublicStakingTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_PublicStaking *PublicStakingSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _PublicStaking.Contract.SafeTransferFrom0(&_PublicStaking.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_PublicStaking *PublicStakingTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _PublicStaking.Contract.SafeTransferFrom0(&_PublicStaking.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PublicStaking *PublicStakingTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PublicStaking *PublicStakingSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _PublicStaking.Contract.SetApprovalForAll(&_PublicStaking.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PublicStaking *PublicStakingTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _PublicStaking.Contract.SetApprovalForAll(&_PublicStaking.TransactOpts, operator, approved)
}

// SkimExcessEth is a paid mutator transaction binding the contract method 0x971b505b.
//
// Solidity: function skimExcessEth(address to_) returns(uint256 excess)
func (_PublicStaking *PublicStakingTransactor) SkimExcessEth(opts *bind.TransactOpts, to_ common.Address) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "skimExcessEth", to_)
}

// SkimExcessEth is a paid mutator transaction binding the contract method 0x971b505b.
//
// Solidity: function skimExcessEth(address to_) returns(uint256 excess)
func (_PublicStaking *PublicStakingSession) SkimExcessEth(to_ common.Address) (*types.Transaction, error) {
	return _PublicStaking.Contract.SkimExcessEth(&_PublicStaking.TransactOpts, to_)
}

// SkimExcessEth is a paid mutator transaction binding the contract method 0x971b505b.
//
// Solidity: function skimExcessEth(address to_) returns(uint256 excess)
func (_PublicStaking *PublicStakingTransactorSession) SkimExcessEth(to_ common.Address) (*types.Transaction, error) {
	return _PublicStaking.Contract.SkimExcessEth(&_PublicStaking.TransactOpts, to_)
}

// SkimExcessToken is a paid mutator transaction binding the contract method 0x7aa507fb.
//
// Solidity: function skimExcessToken(address to_) returns(uint256 excess)
func (_PublicStaking *PublicStakingTransactor) SkimExcessToken(opts *bind.TransactOpts, to_ common.Address) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "skimExcessToken", to_)
}

// SkimExcessToken is a paid mutator transaction binding the contract method 0x7aa507fb.
//
// Solidity: function skimExcessToken(address to_) returns(uint256 excess)
func (_PublicStaking *PublicStakingSession) SkimExcessToken(to_ common.Address) (*types.Transaction, error) {
	return _PublicStaking.Contract.SkimExcessToken(&_PublicStaking.TransactOpts, to_)
}

// SkimExcessToken is a paid mutator transaction binding the contract method 0x7aa507fb.
//
// Solidity: function skimExcessToken(address to_) returns(uint256 excess)
func (_PublicStaking *PublicStakingTransactorSession) SkimExcessToken(to_ common.Address) (*types.Transaction, error) {
	return _PublicStaking.Contract.SkimExcessToken(&_PublicStaking.TransactOpts, to_)
}

// SupportsInterface is a paid mutator transaction binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) returns(bool)
func (_PublicStaking *PublicStakingTransactor) SupportsInterface(opts *bind.TransactOpts, interfaceId [4]byte) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "supportsInterface", interfaceId)
}

// SupportsInterface is a paid mutator transaction binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) returns(bool)
func (_PublicStaking *PublicStakingSession) SupportsInterface(interfaceId [4]byte) (*types.Transaction, error) {
	return _PublicStaking.Contract.SupportsInterface(&_PublicStaking.TransactOpts, interfaceId)
}

// SupportsInterface is a paid mutator transaction binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) returns(bool)
func (_PublicStaking *PublicStakingTransactorSession) SupportsInterface(interfaceId [4]byte) (*types.Transaction, error) {
	return _PublicStaking.Contract.SupportsInterface(&_PublicStaking.TransactOpts, interfaceId)
}

// Symbol is a paid mutator transaction binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_PublicStaking *PublicStakingTransactor) Symbol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "symbol")
}

// Symbol is a paid mutator transaction binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_PublicStaking *PublicStakingSession) Symbol() (*types.Transaction, error) {
	return _PublicStaking.Contract.Symbol(&_PublicStaking.TransactOpts)
}

// Symbol is a paid mutator transaction binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_PublicStaking *PublicStakingTransactorSession) Symbol() (*types.Transaction, error) {
	return _PublicStaking.Contract.Symbol(&_PublicStaking.TransactOpts)
}

// TokenURI is a paid mutator transaction binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) returns(string)
func (_PublicStaking *PublicStakingTransactor) TokenURI(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "tokenURI", tokenId)
}

// TokenURI is a paid mutator transaction binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) returns(string)
func (_PublicStaking *PublicStakingSession) TokenURI(tokenId *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.TokenURI(&_PublicStaking.TransactOpts, tokenId)
}

// TokenURI is a paid mutator transaction binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) returns(string)
func (_PublicStaking *PublicStakingTransactorSession) TokenURI(tokenId *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.TokenURI(&_PublicStaking.TransactOpts, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_PublicStaking *PublicStakingTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_PublicStaking *PublicStakingSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.TransferFrom(&_PublicStaking.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_PublicStaking *PublicStakingTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PublicStaking.Contract.TransferFrom(&_PublicStaking.TransactOpts, from, to, tokenId)
}

// TripCB is a paid mutator transaction binding the contract method 0xadfdc03f.
//
// Solidity: function tripCB() returns()
func (_PublicStaking *PublicStakingTransactor) TripCB(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicStaking.contract.Transact(opts, "tripCB")
}

// TripCB is a paid mutator transaction binding the contract method 0xadfdc03f.
//
// Solidity: function tripCB() returns()
func (_PublicStaking *PublicStakingSession) TripCB() (*types.Transaction, error) {
	return _PublicStaking.Contract.TripCB(&_PublicStaking.TransactOpts)
}

// TripCB is a paid mutator transaction binding the contract method 0xadfdc03f.
//
// Solidity: function tripCB() returns()
func (_PublicStaking *PublicStakingTransactorSession) TripCB() (*types.Transaction, error) {
	return _PublicStaking.Contract.TripCB(&_PublicStaking.TransactOpts)
}

// PublicStakingApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PublicStaking contract.
type PublicStakingApprovalIterator struct {
	Event *PublicStakingApproval // Event containing the contract specifics and raw log

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
func (it *PublicStakingApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicStakingApproval)
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
		it.Event = new(PublicStakingApproval)
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
func (it *PublicStakingApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicStakingApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicStakingApproval represents a Approval event raised by the PublicStaking contract.
type PublicStakingApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_PublicStaking *PublicStakingFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*PublicStakingApprovalIterator, error) {

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

	logs, sub, err := _PublicStaking.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PublicStakingApprovalIterator{contract: _PublicStaking.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_PublicStaking *PublicStakingFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PublicStakingApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _PublicStaking.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicStakingApproval)
				if err := _PublicStaking.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_PublicStaking *PublicStakingFilterer) ParseApproval(log types.Log) (*PublicStakingApproval, error) {
	event := new(PublicStakingApproval)
	if err := _PublicStaking.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PublicStakingApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the PublicStaking contract.
type PublicStakingApprovalForAllIterator struct {
	Event *PublicStakingApprovalForAll // Event containing the contract specifics and raw log

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
func (it *PublicStakingApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicStakingApprovalForAll)
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
		it.Event = new(PublicStakingApprovalForAll)
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
func (it *PublicStakingApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicStakingApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicStakingApprovalForAll represents a ApprovalForAll event raised by the PublicStaking contract.
type PublicStakingApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_PublicStaking *PublicStakingFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*PublicStakingApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PublicStaking.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &PublicStakingApprovalForAllIterator{contract: _PublicStaking.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_PublicStaking *PublicStakingFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *PublicStakingApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PublicStaking.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicStakingApprovalForAll)
				if err := _PublicStaking.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_PublicStaking *PublicStakingFilterer) ParseApprovalForAll(log types.Log) (*PublicStakingApprovalForAll, error) {
	event := new(PublicStakingApprovalForAll)
	if err := _PublicStaking.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PublicStakingTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PublicStaking contract.
type PublicStakingTransferIterator struct {
	Event *PublicStakingTransfer // Event containing the contract specifics and raw log

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
func (it *PublicStakingTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicStakingTransfer)
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
		it.Event = new(PublicStakingTransfer)
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
func (it *PublicStakingTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicStakingTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicStakingTransfer represents a Transfer event raised by the PublicStaking contract.
type PublicStakingTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_PublicStaking *PublicStakingFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*PublicStakingTransferIterator, error) {

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

	logs, sub, err := _PublicStaking.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PublicStakingTransferIterator{contract: _PublicStaking.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_PublicStaking *PublicStakingFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PublicStakingTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _PublicStaking.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicStakingTransfer)
				if err := _PublicStaking.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_PublicStaking *PublicStakingFilterer) ParseTransfer(log types.Log) (*PublicStakingTransfer, error) {
	event := new(PublicStakingTransfer)
	if err := _PublicStaking.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
