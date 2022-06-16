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

// SnapshotsABI is the input ABI used to generate the binding from.
const SnapshotsABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochLength_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSafeToProceedConsensus\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signatureRaw\",\"type\":\"bytes\"}],\"name\":\"SnapshotTaken\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getAliceNetHeightFromLatestSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch_\",\"type\":\"uint256\"}],\"name\":\"getAliceNetHeightFromSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockClaimsFromLatestSnapshot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"chainId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"height\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"txCount\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"prevBlock\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"headerRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structBClaimsParserLibrary.BClaims\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch_\",\"type\":\"uint256\"}],\"name\":\"getBlockClaimsFromSnapshot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"chainId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"height\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"txCount\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"prevBlock\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"headerRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structBClaimsParserLibrary.BClaims\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainIdFromLatestSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch_\",\"type\":\"uint256\"}],\"name\":\"getChainIdFromSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommittedHeightFromLatestSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch_\",\"type\":\"uint256\"}],\"name\":\"getCommittedHeightFromSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"name\":\"getEpochFromHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEpochLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestSnapshot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"committedAt\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"chainId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"height\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"txCount\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"prevBlock\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"headerRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structBClaimsParserLibrary.BClaims\",\"name\":\"blockClaims\",\"type\":\"tuple\"}],\"internalType\":\"structSnapshot\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"getMetamorphicContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinimumIntervalBetweenSnapshots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch_\",\"type\":\"uint256\"}],\"name\":\"getSnapshot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"committedAt\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"chainId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"height\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"txCount\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"prevBlock\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"headerRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structBClaimsParserLibrary.BClaims\",\"name\":\"blockClaims\",\"type\":\"tuple\"}],\"internalType\":\"structSnapshot\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSnapshotDesperationDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSnapshotDesperationFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"desperationDelay_\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"desperationFactor_\",\"type\":\"uint32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"myIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocksSinceDesperation\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blsig\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"desperationFactor\",\"type\":\"uint256\"}],\"name\":\"mayValidatorSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"groupSignature_\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"bClaims_\",\"type\":\"bytes[]\"}],\"name\":\"migrateSnapshots\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"minimumIntervalBetweenSnapshots_\",\"type\":\"uint32\"}],\"name\":\"setMinimumIntervalBetweenSnapshots\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"desperationDelay_\",\"type\":\"uint32\"}],\"name\":\"setSnapshotDesperationDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"desperationFactor_\",\"type\":\"uint32\"}],\"name\":\"setSnapshotDesperationFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"groupSignature_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bClaims_\",\"type\":\"bytes\"}],\"name\":\"snapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Snapshots is an auto generated Go binding around an Ethereum contract.
type Snapshots struct {
	SnapshotsCaller     // Read-only binding to the contract
	SnapshotsTransactor // Write-only binding to the contract
	SnapshotsFilterer   // Log filterer for contract events
}

// SnapshotsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SnapshotsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SnapshotsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SnapshotsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SnapshotsSession struct {
	Contract     *Snapshots        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SnapshotsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SnapshotsCallerSession struct {
	Contract *SnapshotsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SnapshotsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SnapshotsTransactorSession struct {
	Contract     *SnapshotsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SnapshotsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SnapshotsRaw struct {
	Contract *Snapshots // Generic contract binding to access the raw methods on
}

// SnapshotsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SnapshotsCallerRaw struct {
	Contract *SnapshotsCaller // Generic read-only contract binding to access the raw methods on
}

// SnapshotsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SnapshotsTransactorRaw struct {
	Contract *SnapshotsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSnapshots creates a new instance of Snapshots, bound to a specific deployed contract.
func NewSnapshots(address common.Address, backend bind.ContractBackend) (*Snapshots, error) {
	contract, err := bindSnapshots(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Snapshots{SnapshotsCaller: SnapshotsCaller{contract: contract}, SnapshotsTransactor: SnapshotsTransactor{contract: contract}, SnapshotsFilterer: SnapshotsFilterer{contract: contract}}, nil
}

// NewSnapshotsCaller creates a new read-only instance of Snapshots, bound to a specific deployed contract.
func NewSnapshotsCaller(address common.Address, caller bind.ContractCaller) (*SnapshotsCaller, error) {
	contract, err := bindSnapshots(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SnapshotsCaller{contract: contract}, nil
}

// NewSnapshotsTransactor creates a new write-only instance of Snapshots, bound to a specific deployed contract.
func NewSnapshotsTransactor(address common.Address, transactor bind.ContractTransactor) (*SnapshotsTransactor, error) {
	contract, err := bindSnapshots(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SnapshotsTransactor{contract: contract}, nil
}

// NewSnapshotsFilterer creates a new log filterer instance of Snapshots, bound to a specific deployed contract.
func NewSnapshotsFilterer(address common.Address, filterer bind.ContractFilterer) (*SnapshotsFilterer, error) {
	contract, err := bindSnapshots(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SnapshotsFilterer{contract: contract}, nil
}

// bindSnapshots binds a generic wrapper to an already deployed contract.
func bindSnapshots(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SnapshotsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Snapshots *SnapshotsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Snapshots.Contract.SnapshotsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Snapshots *SnapshotsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Snapshots.Contract.SnapshotsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Snapshots *SnapshotsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Snapshots.Contract.SnapshotsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Snapshots *SnapshotsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Snapshots.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Snapshots *SnapshotsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Snapshots.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Snapshots *SnapshotsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Snapshots.Contract.contract.Transact(opts, method, params...)
}

// BClaimsParserLibraryBClaims is an auto generated low-level Go binding around an user-defined struct.
type BClaimsParserLibraryBClaims struct {
	ChainId    uint32
	Height     uint32
	TxCount    uint32
	PrevBlock  [32]byte
	TxRoot     [32]byte
	StateRoot  [32]byte
	HeaderRoot [32]byte
}

// Snapshot is an auto generated low-level Go binding around an user-defined struct.
type Snapshot struct {
	CommittedAt *big.Int
	BlockClaims BClaimsParserLibraryBClaims
}

// GetAliceNetHeightFromLatestSnapshot is a paid mutator transaction binding the contract method 0xff07fc0e.
//
// Solidity: function getAliceNetHeightFromLatestSnapshot() returns(uint256)
func (_Snapshots *SnapshotsTransactor) GetAliceNetHeightFromLatestSnapshot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Snapshots.contract.Transact(opts, "getAliceNetHeightFromLatestSnapshot")
}

// GetAliceNetHeightFromLatestSnapshot is a paid mutator transaction binding the contract method 0xff07fc0e.
//
// Solidity: function getAliceNetHeightFromLatestSnapshot() returns(uint256)
func (_Snapshots *SnapshotsSession) GetAliceNetHeightFromLatestSnapshot() (*types.Transaction, error) {
	return _Snapshots.Contract.GetAliceNetHeightFromLatestSnapshot(&_Snapshots.TransactOpts)
}

// GetAliceNetHeightFromLatestSnapshot is a paid mutator transaction binding the contract method 0xff07fc0e.
//
// Solidity: function getAliceNetHeightFromLatestSnapshot() returns(uint256)
func (_Snapshots *SnapshotsTransactorSession) GetAliceNetHeightFromLatestSnapshot() (*types.Transaction, error) {
	return _Snapshots.Contract.GetAliceNetHeightFromLatestSnapshot(&_Snapshots.TransactOpts)
}

// GetAliceNetHeightFromSnapshot is a paid mutator transaction binding the contract method 0xc5e8fde1.
//
// Solidity: function getAliceNetHeightFromSnapshot(uint256 epoch_) returns(uint256)
func (_Snapshots *SnapshotsTransactor) GetAliceNetHeightFromSnapshot(opts *bind.TransactOpts, epoch_ *big.Int) (*types.Transaction, error) {
	return _Snapshots.contract.Transact(opts, "getAliceNetHeightFromSnapshot", epoch_)
}

// GetAliceNetHeightFromSnapshot is a paid mutator transaction binding the contract method 0xc5e8fde1.
//
// Solidity: function getAliceNetHeightFromSnapshot(uint256 epoch_) returns(uint256)
func (_Snapshots *SnapshotsSession) GetAliceNetHeightFromSnapshot(epoch_ *big.Int) (*types.Transaction, error) {
	return _Snapshots.Contract.GetAliceNetHeightFromSnapshot(&_Snapshots.TransactOpts, epoch_)
}

// GetAliceNetHeightFromSnapshot is a paid mutator transaction binding the contract method 0xc5e8fde1.
//
// Solidity: function getAliceNetHeightFromSnapshot(uint256 epoch_) returns(uint256)
func (_Snapshots *SnapshotsTransactorSession) GetAliceNetHeightFromSnapshot(epoch_ *big.Int) (*types.Transaction, error) {
	return _Snapshots.Contract.GetAliceNetHeightFromSnapshot(&_Snapshots.TransactOpts, epoch_)
}

// GetBlockClaimsFromLatestSnapshot is a paid mutator transaction binding the contract method 0xc2ea6603.
//
// Solidity: function getBlockClaimsFromLatestSnapshot() returns(BClaimsParserLibraryBClaims)
func (_Snapshots *SnapshotsTransactor) GetBlockClaimsFromLatestSnapshot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Snapshots.contract.Transact(opts, "getBlockClaimsFromLatestSnapshot")
}

// GetBlockClaimsFromLatestSnapshot is a paid mutator transaction binding the contract method 0xc2ea6603.
//
// Solidity: function getBlockClaimsFromLatestSnapshot() returns(BClaimsParserLibraryBClaims)
func (_Snapshots *SnapshotsSession) GetBlockClaimsFromLatestSnapshot() (*types.Transaction, error) {
	return _Snapshots.Contract.GetBlockClaimsFromLatestSnapshot(&_Snapshots.TransactOpts)
}

// GetBlockClaimsFromLatestSnapshot is a paid mutator transaction binding the contract method 0xc2ea6603.
//
// Solidity: function getBlockClaimsFromLatestSnapshot() returns(BClaimsParserLibraryBClaims)
func (_Snapshots *SnapshotsTransactorSession) GetBlockClaimsFromLatestSnapshot() (*types.Transaction, error) {
	return _Snapshots.Contract.GetBlockClaimsFromLatestSnapshot(&_Snapshots.TransactOpts)
}

// GetBlockClaimsFromSnapshot is a paid mutator transaction binding the contract method 0x45dfc599.
//
// Solidity: function getBlockClaimsFromSnapshot(uint256 epoch_) returns(BClaimsParserLibraryBClaims)
func (_Snapshots *SnapshotsTransactor) GetBlockClaimsFromSnapshot(opts *bind.TransactOpts, epoch_ *big.Int) (*types.Transaction, error) {
	return _Snapshots.contract.Transact(opts, "getBlockClaimsFromSnapshot", epoch_)
}

// GetBlockClaimsFromSnapshot is a paid mutator transaction binding the contract method 0x45dfc599.
//
// Solidity: function getBlockClaimsFromSnapshot(uint256 epoch_) returns(BClaimsParserLibraryBClaims)
func (_Snapshots *SnapshotsSession) GetBlockClaimsFromSnapshot(epoch_ *big.Int) (*types.Transaction, error) {
	return _Snapshots.Contract.GetBlockClaimsFromSnapshot(&_Snapshots.TransactOpts, epoch_)
}

// GetBlockClaimsFromSnapshot is a paid mutator transaction binding the contract method 0x45dfc599.
//
// Solidity: function getBlockClaimsFromSnapshot(uint256 epoch_) returns(BClaimsParserLibraryBClaims)
func (_Snapshots *SnapshotsTransactorSession) GetBlockClaimsFromSnapshot(epoch_ *big.Int) (*types.Transaction, error) {
	return _Snapshots.Contract.GetBlockClaimsFromSnapshot(&_Snapshots.TransactOpts, epoch_)
}

// GetChainId is a paid mutator transaction binding the contract method 0x3408e470.
//
// Solidity: function getChainId() returns(uint256)
func (_Snapshots *SnapshotsTransactor) GetChainId(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Snapshots.contract.Transact(opts, "getChainId")
}

// GetChainId is a paid mutator transaction binding the contract method 0x3408e470.
//
// Solidity: function getChainId() returns(uint256)
func (_Snapshots *SnapshotsSession) GetChainId() (*types.Transaction, error) {
	return _Snapshots.Contract.GetChainId(&_Snapshots.TransactOpts)
}

// GetChainId is a paid mutator transaction binding the contract method 0x3408e470.
//
// Solidity: function getChainId() returns(uint256)
func (_Snapshots *SnapshotsTransactorSession) GetChainId() (*types.Transaction, error) {
	return _Snapshots.Contract.GetChainId(&_Snapshots.TransactOpts)
}

// GetChainIdFromLatestSnapshot is a paid mutator transaction binding the contract method 0xd9c11657.
//
// Solidity: function getChainIdFromLatestSnapshot() returns(uint256)
func (_Snapshots *SnapshotsTransactor) GetChainIdFromLatestSnapshot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Snapshots.contract.Transact(opts, "getChainIdFromLatestSnapshot")
}

// GetChainIdFromLatestSnapshot is a paid mutator transaction binding the contract method 0xd9c11657.
//
// Solidity: function getChainIdFromLatestSnapshot() returns(uint256)
func (_Snapshots *SnapshotsSession) GetChainIdFromLatestSnapshot() (*types.Transaction, error) {
	return _Snapshots.Contract.GetChainIdFromLatestSnapshot(&_Snapshots.TransactOpts)
}

// GetChainIdFromLatestSnapshot is a paid mutator transaction binding the contract method 0xd9c11657.
//
// Solidity: function getChainIdFromLatestSnapshot() returns(uint256)
func (_Snapshots *SnapshotsTransactorSession) GetChainIdFromLatestSnapshot() (*types.Transaction, error) {
	return _Snapshots.Contract.GetChainIdFromLatestSnapshot(&_Snapshots.TransactOpts)
}

// GetChainIdFromSnapshot is a paid mutator transaction binding the contract method 0x19f74669.
//
// Solidity: function getChainIdFromSnapshot(uint256 epoch_) returns(uint256)
func (_Snapshots *SnapshotsTransactor) GetChainIdFromSnapshot(opts *bind.TransactOpts, epoch_ *big.Int) (*types.Transaction, error) {
	return _Snapshots.contract.Transact(opts, "getChainIdFromSnapshot", epoch_)
}

// GetChainIdFromSnapshot is a paid mutator transaction binding the contract method 0x19f74669.
//
// Solidity: function getChainIdFromSnapshot(uint256 epoch_) returns(uint256)
func (_Snapshots *SnapshotsSession) GetChainIdFromSnapshot(epoch_ *big.Int) (*types.Transaction, error) {
	return _Snapshots.Contract.GetChainIdFromSnapshot(&_Snapshots.TransactOpts, epoch_)
}

// GetChainIdFromSnapshot is a paid mutator transaction binding the contract method 0x19f74669.
//
// Solidity: function getChainIdFromSnapshot(uint256 epoch_) returns(uint256)
func (_Snapshots *SnapshotsTransactorSession) GetChainIdFromSnapshot(epoch_ *big.Int) (*types.Transaction, error) {
	return _Snapshots.Contract.GetChainIdFromSnapshot(&_Snapshots.TransactOpts, epoch_)
}

// GetCommittedHeightFromLatestSnapshot is a paid mutator transaction binding the contract method 0x026c2b7e.
//
// Solidity: function getCommittedHeightFromLatestSnapshot() returns(uint256)
func (_Snapshots *SnapshotsTransactor) GetCommittedHeightFromLatestSnapshot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Snapshots.contract.Transact(opts, "getCommittedHeightFromLatestSnapshot")
}

// GetCommittedHeightFromLatestSnapshot is a paid mutator transaction binding the contract method 0x026c2b7e.
//
// Solidity: function getCommittedHeightFromLatestSnapshot() returns(uint256)
func (_Snapshots *SnapshotsSession) GetCommittedHeightFromLatestSnapshot() (*types.Transaction, error) {
	return _Snapshots.Contract.GetCommittedHeightFromLatestSnapshot(&_Snapshots.TransactOpts)
}

// GetCommittedHeightFromLatestSnapshot is a paid mutator transaction binding the contract method 0x026c2b7e.
//
// Solidity: function getCommittedHeightFromLatestSnapshot() returns(uint256)
func (_Snapshots *SnapshotsTransactorSession) GetCommittedHeightFromLatestSnapshot() (*types.Transaction, error) {
	return _Snapshots.Contract.GetCommittedHeightFromLatestSnapshot(&_Snapshots.TransactOpts)
}

// GetCommittedHeightFromSnapshot is a paid mutator transaction binding the contract method 0xe18c697a.
//
// Solidity: function getCommittedHeightFromSnapshot(uint256 epoch_) returns(uint256)
func (_Snapshots *SnapshotsTransactor) GetCommittedHeightFromSnapshot(opts *bind.TransactOpts, epoch_ *big.Int) (*types.Transaction, error) {
	return _Snapshots.contract.Transact(opts, "getCommittedHeightFromSnapshot", epoch_)
}

// GetCommittedHeightFromSnapshot is a paid mutator transaction binding the contract method 0xe18c697a.
//
// Solidity: function getCommittedHeightFromSnapshot(uint256 epoch_) returns(uint256)
func (_Snapshots *SnapshotsSession) GetCommittedHeightFromSnapshot(epoch_ *big.Int) (*types.Transaction, error) {
	return _Snapshots.Contract.GetCommittedHeightFromSnapshot(&_Snapshots.TransactOpts, epoch_)
}

// GetCommittedHeightFromSnapshot is a paid mutator transaction binding the contract method 0xe18c697a.
//
// Solidity: function getCommittedHeightFromSnapshot(uint256 epoch_) returns(uint256)
func (_Snapshots *SnapshotsTransactorSession) GetCommittedHeightFromSnapshot(epoch_ *big.Int) (*types.Transaction, error) {
	return _Snapshots.Contract.GetCommittedHeightFromSnapshot(&_Snapshots.TransactOpts, epoch_)
}

// GetEpoch is a paid mutator transaction binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() returns(uint256)
func (_Snapshots *SnapshotsTransactor) GetEpoch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Snapshots.contract.Transact(opts, "getEpoch")
}

// GetEpoch is a paid mutator transaction binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() returns(uint256)
func (_Snapshots *SnapshotsSession) GetEpoch() (*types.Transaction, error) {
	return _Snapshots.Contract.GetEpoch(&_Snapshots.TransactOpts)
}

// GetEpoch is a paid mutator transaction binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() returns(uint256)
func (_Snapshots *SnapshotsTransactorSession) GetEpoch() (*types.Transaction, error) {
	return _Snapshots.Contract.GetEpoch(&_Snapshots.TransactOpts)
}

// GetEpochFromHeight is a paid mutator transaction binding the contract method 0x2eee30ce.
//
// Solidity: function getEpochFromHeight(uint256 height) returns(uint256)
func (_Snapshots *SnapshotsTransactor) GetEpochFromHeight(opts *bind.TransactOpts, height *big.Int) (*types.Transaction, error) {
	return _Snapshots.contract.Transact(opts, "getEpochFromHeight", height)
}

// GetEpochFromHeight is a paid mutator transaction binding the contract method 0x2eee30ce.
//
// Solidity: function getEpochFromHeight(uint256 height) returns(uint256)
func (_Snapshots *SnapshotsSession) GetEpochFromHeight(height *big.Int) (*types.Transaction, error) {
	return _Snapshots.Contract.GetEpochFromHeight(&_Snapshots.TransactOpts, height)
}

// GetEpochFromHeight is a paid mutator transaction binding the contract method 0x2eee30ce.
//
// Solidity: function getEpochFromHeight(uint256 height) returns(uint256)
func (_Snapshots *SnapshotsTransactorSession) GetEpochFromHeight(height *big.Int) (*types.Transaction, error) {
	return _Snapshots.Contract.GetEpochFromHeight(&_Snapshots.TransactOpts, height)
}

// GetEpochLength is a paid mutator transaction binding the contract method 0xcfe8a73b.
//
// Solidity: function getEpochLength() returns(uint256)
func (_Snapshots *SnapshotsTransactor) GetEpochLength(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Snapshots.contract.Transact(opts, "getEpochLength")
}

// GetEpochLength is a paid mutator transaction binding the contract method 0xcfe8a73b.
//
// Solidity: function getEpochLength() returns(uint256)
func (_Snapshots *SnapshotsSession) GetEpochLength() (*types.Transaction, error) {
	return _Snapshots.Contract.GetEpochLength(&_Snapshots.TransactOpts)
}

// GetEpochLength is a paid mutator transaction binding the contract method 0xcfe8a73b.
//
// Solidity: function getEpochLength() returns(uint256)
func (_Snapshots *SnapshotsTransactorSession) GetEpochLength() (*types.Transaction, error) {
	return _Snapshots.Contract.GetEpochLength(&_Snapshots.TransactOpts)
}

// GetLatestSnapshot is a paid mutator transaction binding the contract method 0xd518f243.
//
// Solidity: function getLatestSnapshot() returns(Snapshot)
func (_Snapshots *SnapshotsTransactor) GetLatestSnapshot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Snapshots.contract.Transact(opts, "getLatestSnapshot")
}

// GetLatestSnapshot is a paid mutator transaction binding the contract method 0xd518f243.
//
// Solidity: function getLatestSnapshot() returns(Snapshot)
func (_Snapshots *SnapshotsSession) GetLatestSnapshot() (*types.Transaction, error) {
	return _Snapshots.Contract.GetLatestSnapshot(&_Snapshots.TransactOpts)
}

// GetLatestSnapshot is a paid mutator transaction binding the contract method 0xd518f243.
//
// Solidity: function getLatestSnapshot() returns(Snapshot)
func (_Snapshots *SnapshotsTransactorSession) GetLatestSnapshot() (*types.Transaction, error) {
	return _Snapshots.Contract.GetLatestSnapshot(&_Snapshots.TransactOpts)
}

// GetMetamorphicContractAddress is a paid mutator transaction binding the contract method 0x8653a465.
//
// Solidity: function getMetamorphicContractAddress(bytes32 _salt, address _factory) returns(address)
func (_Snapshots *SnapshotsTransactor) GetMetamorphicContractAddress(opts *bind.TransactOpts, _salt [32]byte, _factory common.Address) (*types.Transaction, error) {
	return _Snapshots.contract.Transact(opts, "getMetamorphicContractAddress", _salt, _factory)
}

// GetMetamorphicContractAddress is a paid mutator transaction binding the contract method 0x8653a465.
//
// Solidity: function getMetamorphicContractAddress(bytes32 _salt, address _factory) returns(address)
func (_Snapshots *SnapshotsSession) GetMetamorphicContractAddress(_salt [32]byte, _factory common.Address) (*types.Transaction, error) {
	return _Snapshots.Contract.GetMetamorphicContractAddress(&_Snapshots.TransactOpts, _salt, _factory)
}

// GetMetamorphicContractAddress is a paid mutator transaction binding the contract method 0x8653a465.
//
// Solidity: function getMetamorphicContractAddress(bytes32 _salt, address _factory) returns(address)
func (_Snapshots *SnapshotsTransactorSession) GetMetamorphicContractAddress(_salt [32]byte, _factory common.Address) (*types.Transaction, error) {
	return _Snapshots.Contract.GetMetamorphicContractAddress(&_Snapshots.TransactOpts, _salt, _factory)
}

// GetMinimumIntervalBetweenSnapshots is a paid mutator transaction binding the contract method 0x42438d7b.
//
// Solidity: function getMinimumIntervalBetweenSnapshots() returns(uint256)
func (_Snapshots *SnapshotsTransactor) GetMinimumIntervalBetweenSnapshots(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Snapshots.contract.Transact(opts, "getMinimumIntervalBetweenSnapshots")
}

// GetMinimumIntervalBetweenSnapshots is a paid mutator transaction binding the contract method 0x42438d7b.
//
// Solidity: function getMinimumIntervalBetweenSnapshots() returns(uint256)
func (_Snapshots *SnapshotsSession) GetMinimumIntervalBetweenSnapshots() (*types.Transaction, error) {
	return _Snapshots.Contract.GetMinimumIntervalBetweenSnapshots(&_Snapshots.TransactOpts)
}

// GetMinimumIntervalBetweenSnapshots is a paid mutator transaction binding the contract method 0x42438d7b.
//
// Solidity: function getMinimumIntervalBetweenSnapshots() returns(uint256)
func (_Snapshots *SnapshotsTransactorSession) GetMinimumIntervalBetweenSnapshots() (*types.Transaction, error) {
	return _Snapshots.Contract.GetMinimumIntervalBetweenSnapshots(&_Snapshots.TransactOpts)
}

// GetSnapshot is a paid mutator transaction binding the contract method 0x76f10ad0.
//
// Solidity: function getSnapshot(uint256 epoch_) returns(Snapshot)
func (_Snapshots *SnapshotsTransactor) GetSnapshot(opts *bind.TransactOpts, epoch_ *big.Int) (*types.Transaction, error) {
	return _Snapshots.contract.Transact(opts, "getSnapshot", epoch_)
}

// GetSnapshot is a paid mutator transaction binding the contract method 0x76f10ad0.
//
// Solidity: function getSnapshot(uint256 epoch_) returns(Snapshot)
func (_Snapshots *SnapshotsSession) GetSnapshot(epoch_ *big.Int) (*types.Transaction, error) {
	return _Snapshots.Contract.GetSnapshot(&_Snapshots.TransactOpts, epoch_)
}

// GetSnapshot is a paid mutator transaction binding the contract method 0x76f10ad0.
//
// Solidity: function getSnapshot(uint256 epoch_) returns(Snapshot)
func (_Snapshots *SnapshotsTransactorSession) GetSnapshot(epoch_ *big.Int) (*types.Transaction, error) {
	return _Snapshots.Contract.GetSnapshot(&_Snapshots.TransactOpts, epoch_)
}

// GetSnapshotDesperationDelay is a paid mutator transaction binding the contract method 0xd17fcc56.
//
// Solidity: function getSnapshotDesperationDelay() returns(uint256)
func (_Snapshots *SnapshotsTransactor) GetSnapshotDesperationDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Snapshots.contract.Transact(opts, "getSnapshotDesperationDelay")
}

// GetSnapshotDesperationDelay is a paid mutator transaction binding the contract method 0xd17fcc56.
//
// Solidity: function getSnapshotDesperationDelay() returns(uint256)
func (_Snapshots *SnapshotsSession) GetSnapshotDesperationDelay() (*types.Transaction, error) {
	return _Snapshots.Contract.GetSnapshotDesperationDelay(&_Snapshots.TransactOpts)
}

// GetSnapshotDesperationDelay is a paid mutator transaction binding the contract method 0xd17fcc56.
//
// Solidity: function getSnapshotDesperationDelay() returns(uint256)
func (_Snapshots *SnapshotsTransactorSession) GetSnapshotDesperationDelay() (*types.Transaction, error) {
	return _Snapshots.Contract.GetSnapshotDesperationDelay(&_Snapshots.TransactOpts)
}

// GetSnapshotDesperationFactor is a paid mutator transaction binding the contract method 0x7cc4cce6.
//
// Solidity: function getSnapshotDesperationFactor() returns(uint256)
func (_Snapshots *SnapshotsTransactor) GetSnapshotDesperationFactor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Snapshots.contract.Transact(opts, "getSnapshotDesperationFactor")
}

// GetSnapshotDesperationFactor is a paid mutator transaction binding the contract method 0x7cc4cce6.
//
// Solidity: function getSnapshotDesperationFactor() returns(uint256)
func (_Snapshots *SnapshotsSession) GetSnapshotDesperationFactor() (*types.Transaction, error) {
	return _Snapshots.Contract.GetSnapshotDesperationFactor(&_Snapshots.TransactOpts)
}

// GetSnapshotDesperationFactor is a paid mutator transaction binding the contract method 0x7cc4cce6.
//
// Solidity: function getSnapshotDesperationFactor() returns(uint256)
func (_Snapshots *SnapshotsTransactorSession) GetSnapshotDesperationFactor() (*types.Transaction, error) {
	return _Snapshots.Contract.GetSnapshotDesperationFactor(&_Snapshots.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x3ecc0f5e.
//
// Solidity: function initialize(uint32 desperationDelay_, uint32 desperationFactor_) returns()
func (_Snapshots *SnapshotsTransactor) Initialize(opts *bind.TransactOpts, desperationDelay_ uint32, desperationFactor_ uint32) (*types.Transaction, error) {
	return _Snapshots.contract.Transact(opts, "initialize", desperationDelay_, desperationFactor_)
}

// Initialize is a paid mutator transaction binding the contract method 0x3ecc0f5e.
//
// Solidity: function initialize(uint32 desperationDelay_, uint32 desperationFactor_) returns()
func (_Snapshots *SnapshotsSession) Initialize(desperationDelay_ uint32, desperationFactor_ uint32) (*types.Transaction, error) {
	return _Snapshots.Contract.Initialize(&_Snapshots.TransactOpts, desperationDelay_, desperationFactor_)
}

// Initialize is a paid mutator transaction binding the contract method 0x3ecc0f5e.
//
// Solidity: function initialize(uint32 desperationDelay_, uint32 desperationFactor_) returns()
func (_Snapshots *SnapshotsTransactorSession) Initialize(desperationDelay_ uint32, desperationFactor_ uint32) (*types.Transaction, error) {
	return _Snapshots.Contract.Initialize(&_Snapshots.TransactOpts, desperationDelay_, desperationFactor_)
}

// MayValidatorSnapshot is a paid mutator transaction binding the contract method 0xf45fa246.
//
// Solidity: function mayValidatorSnapshot(uint256 numValidators, uint256 myIdx, uint256 blocksSinceDesperation, bytes32 blsig, uint256 desperationFactor) returns(bool)
func (_Snapshots *SnapshotsTransactor) MayValidatorSnapshot(opts *bind.TransactOpts, numValidators *big.Int, myIdx *big.Int, blocksSinceDesperation *big.Int, blsig [32]byte, desperationFactor *big.Int) (*types.Transaction, error) {
	return _Snapshots.contract.Transact(opts, "mayValidatorSnapshot", numValidators, myIdx, blocksSinceDesperation, blsig, desperationFactor)
}

// MayValidatorSnapshot is a paid mutator transaction binding the contract method 0xf45fa246.
//
// Solidity: function mayValidatorSnapshot(uint256 numValidators, uint256 myIdx, uint256 blocksSinceDesperation, bytes32 blsig, uint256 desperationFactor) returns(bool)
func (_Snapshots *SnapshotsSession) MayValidatorSnapshot(numValidators *big.Int, myIdx *big.Int, blocksSinceDesperation *big.Int, blsig [32]byte, desperationFactor *big.Int) (*types.Transaction, error) {
	return _Snapshots.Contract.MayValidatorSnapshot(&_Snapshots.TransactOpts, numValidators, myIdx, blocksSinceDesperation, blsig, desperationFactor)
}

// MayValidatorSnapshot is a paid mutator transaction binding the contract method 0xf45fa246.
//
// Solidity: function mayValidatorSnapshot(uint256 numValidators, uint256 myIdx, uint256 blocksSinceDesperation, bytes32 blsig, uint256 desperationFactor) returns(bool)
func (_Snapshots *SnapshotsTransactorSession) MayValidatorSnapshot(numValidators *big.Int, myIdx *big.Int, blocksSinceDesperation *big.Int, blsig [32]byte, desperationFactor *big.Int) (*types.Transaction, error) {
	return _Snapshots.Contract.MayValidatorSnapshot(&_Snapshots.TransactOpts, numValidators, myIdx, blocksSinceDesperation, blsig, desperationFactor)
}

// MigrateSnapshots is a paid mutator transaction binding the contract method 0xae2728ea.
//
// Solidity: function migrateSnapshots(bytes[] groupSignature_, bytes[] bClaims_) returns(bool)
func (_Snapshots *SnapshotsTransactor) MigrateSnapshots(opts *bind.TransactOpts, groupSignature_ [][]byte, bClaims_ [][]byte) (*types.Transaction, error) {
	return _Snapshots.contract.Transact(opts, "migrateSnapshots", groupSignature_, bClaims_)
}

// MigrateSnapshots is a paid mutator transaction binding the contract method 0xae2728ea.
//
// Solidity: function migrateSnapshots(bytes[] groupSignature_, bytes[] bClaims_) returns(bool)
func (_Snapshots *SnapshotsSession) MigrateSnapshots(groupSignature_ [][]byte, bClaims_ [][]byte) (*types.Transaction, error) {
	return _Snapshots.Contract.MigrateSnapshots(&_Snapshots.TransactOpts, groupSignature_, bClaims_)
}

// MigrateSnapshots is a paid mutator transaction binding the contract method 0xae2728ea.
//
// Solidity: function migrateSnapshots(bytes[] groupSignature_, bytes[] bClaims_) returns(bool)
func (_Snapshots *SnapshotsTransactorSession) MigrateSnapshots(groupSignature_ [][]byte, bClaims_ [][]byte) (*types.Transaction, error) {
	return _Snapshots.Contract.MigrateSnapshots(&_Snapshots.TransactOpts, groupSignature_, bClaims_)
}

// SetMinimumIntervalBetweenSnapshots is a paid mutator transaction binding the contract method 0xeb7c7afe.
//
// Solidity: function setMinimumIntervalBetweenSnapshots(uint32 minimumIntervalBetweenSnapshots_) returns()
func (_Snapshots *SnapshotsTransactor) SetMinimumIntervalBetweenSnapshots(opts *bind.TransactOpts, minimumIntervalBetweenSnapshots_ uint32) (*types.Transaction, error) {
	return _Snapshots.contract.Transact(opts, "setMinimumIntervalBetweenSnapshots", minimumIntervalBetweenSnapshots_)
}

// SetMinimumIntervalBetweenSnapshots is a paid mutator transaction binding the contract method 0xeb7c7afe.
//
// Solidity: function setMinimumIntervalBetweenSnapshots(uint32 minimumIntervalBetweenSnapshots_) returns()
func (_Snapshots *SnapshotsSession) SetMinimumIntervalBetweenSnapshots(minimumIntervalBetweenSnapshots_ uint32) (*types.Transaction, error) {
	return _Snapshots.Contract.SetMinimumIntervalBetweenSnapshots(&_Snapshots.TransactOpts, minimumIntervalBetweenSnapshots_)
}

// SetMinimumIntervalBetweenSnapshots is a paid mutator transaction binding the contract method 0xeb7c7afe.
//
// Solidity: function setMinimumIntervalBetweenSnapshots(uint32 minimumIntervalBetweenSnapshots_) returns()
func (_Snapshots *SnapshotsTransactorSession) SetMinimumIntervalBetweenSnapshots(minimumIntervalBetweenSnapshots_ uint32) (*types.Transaction, error) {
	return _Snapshots.Contract.SetMinimumIntervalBetweenSnapshots(&_Snapshots.TransactOpts, minimumIntervalBetweenSnapshots_)
}

// SetSnapshotDesperationDelay is a paid mutator transaction binding the contract method 0xc2e8fef2.
//
// Solidity: function setSnapshotDesperationDelay(uint32 desperationDelay_) returns()
func (_Snapshots *SnapshotsTransactor) SetSnapshotDesperationDelay(opts *bind.TransactOpts, desperationDelay_ uint32) (*types.Transaction, error) {
	return _Snapshots.contract.Transact(opts, "setSnapshotDesperationDelay", desperationDelay_)
}

// SetSnapshotDesperationDelay is a paid mutator transaction binding the contract method 0xc2e8fef2.
//
// Solidity: function setSnapshotDesperationDelay(uint32 desperationDelay_) returns()
func (_Snapshots *SnapshotsSession) SetSnapshotDesperationDelay(desperationDelay_ uint32) (*types.Transaction, error) {
	return _Snapshots.Contract.SetSnapshotDesperationDelay(&_Snapshots.TransactOpts, desperationDelay_)
}

// SetSnapshotDesperationDelay is a paid mutator transaction binding the contract method 0xc2e8fef2.
//
// Solidity: function setSnapshotDesperationDelay(uint32 desperationDelay_) returns()
func (_Snapshots *SnapshotsTransactorSession) SetSnapshotDesperationDelay(desperationDelay_ uint32) (*types.Transaction, error) {
	return _Snapshots.Contract.SetSnapshotDesperationDelay(&_Snapshots.TransactOpts, desperationDelay_)
}

// SetSnapshotDesperationFactor is a paid mutator transaction binding the contract method 0x3fa7a1ad.
//
// Solidity: function setSnapshotDesperationFactor(uint32 desperationFactor_) returns()
func (_Snapshots *SnapshotsTransactor) SetSnapshotDesperationFactor(opts *bind.TransactOpts, desperationFactor_ uint32) (*types.Transaction, error) {
	return _Snapshots.contract.Transact(opts, "setSnapshotDesperationFactor", desperationFactor_)
}

// SetSnapshotDesperationFactor is a paid mutator transaction binding the contract method 0x3fa7a1ad.
//
// Solidity: function setSnapshotDesperationFactor(uint32 desperationFactor_) returns()
func (_Snapshots *SnapshotsSession) SetSnapshotDesperationFactor(desperationFactor_ uint32) (*types.Transaction, error) {
	return _Snapshots.Contract.SetSnapshotDesperationFactor(&_Snapshots.TransactOpts, desperationFactor_)
}

// SetSnapshotDesperationFactor is a paid mutator transaction binding the contract method 0x3fa7a1ad.
//
// Solidity: function setSnapshotDesperationFactor(uint32 desperationFactor_) returns()
func (_Snapshots *SnapshotsTransactorSession) SetSnapshotDesperationFactor(desperationFactor_ uint32) (*types.Transaction, error) {
	return _Snapshots.Contract.SetSnapshotDesperationFactor(&_Snapshots.TransactOpts, desperationFactor_)
}

// Snapshot is a paid mutator transaction binding the contract method 0x08ca1f25.
//
// Solidity: function snapshot(bytes groupSignature_, bytes bClaims_) returns(bool)
func (_Snapshots *SnapshotsTransactor) Snapshot(opts *bind.TransactOpts, groupSignature_ []byte, bClaims_ []byte) (*types.Transaction, error) {
	return _Snapshots.contract.Transact(opts, "snapshot", groupSignature_, bClaims_)
}

// Snapshot is a paid mutator transaction binding the contract method 0x08ca1f25.
//
// Solidity: function snapshot(bytes groupSignature_, bytes bClaims_) returns(bool)
func (_Snapshots *SnapshotsSession) Snapshot(groupSignature_ []byte, bClaims_ []byte) (*types.Transaction, error) {
	return _Snapshots.Contract.Snapshot(&_Snapshots.TransactOpts, groupSignature_, bClaims_)
}

// Snapshot is a paid mutator transaction binding the contract method 0x08ca1f25.
//
// Solidity: function snapshot(bytes groupSignature_, bytes bClaims_) returns(bool)
func (_Snapshots *SnapshotsTransactorSession) Snapshot(groupSignature_ []byte, bClaims_ []byte) (*types.Transaction, error) {
	return _Snapshots.Contract.Snapshot(&_Snapshots.TransactOpts, groupSignature_, bClaims_)
}

// SnapshotsSnapshotTakenIterator is returned from FilterSnapshotTaken and is used to iterate over the raw logs and unpacked data for SnapshotTaken events raised by the Snapshots contract.
type SnapshotsSnapshotTakenIterator struct {
	Event *SnapshotsSnapshotTaken // Event containing the contract specifics and raw log

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
func (it *SnapshotsSnapshotTakenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SnapshotsSnapshotTaken)
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
		it.Event = new(SnapshotsSnapshotTaken)
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
func (it *SnapshotsSnapshotTakenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SnapshotsSnapshotTakenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SnapshotsSnapshotTaken represents a SnapshotTaken event raised by the Snapshots contract.
type SnapshotsSnapshotTaken struct {
	ChainId                  *big.Int
	Epoch                    *big.Int
	Height                   *big.Int
	Validator                common.Address
	IsSafeToProceedConsensus bool
	SignatureRaw             []byte
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSnapshotTaken is a free log retrieval operation binding the contract event 0x24b0dff7469a7007db81d741ef90d7966936fb78bc19d667f4575ecbf56ab350.
//
// Solidity: event SnapshotTaken(uint256 chainId, uint256 indexed epoch, uint256 height, address indexed validator, bool isSafeToProceedConsensus, bytes signatureRaw)
func (_Snapshots *SnapshotsFilterer) FilterSnapshotTaken(opts *bind.FilterOpts, epoch []*big.Int, validator []common.Address) (*SnapshotsSnapshotTakenIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Snapshots.contract.FilterLogs(opts, "SnapshotTaken", epochRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &SnapshotsSnapshotTakenIterator{contract: _Snapshots.contract, event: "SnapshotTaken", logs: logs, sub: sub}, nil
}

// WatchSnapshotTaken is a free log subscription operation binding the contract event 0x24b0dff7469a7007db81d741ef90d7966936fb78bc19d667f4575ecbf56ab350.
//
// Solidity: event SnapshotTaken(uint256 chainId, uint256 indexed epoch, uint256 height, address indexed validator, bool isSafeToProceedConsensus, bytes signatureRaw)
func (_Snapshots *SnapshotsFilterer) WatchSnapshotTaken(opts *bind.WatchOpts, sink chan<- *SnapshotsSnapshotTaken, epoch []*big.Int, validator []common.Address) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Snapshots.contract.WatchLogs(opts, "SnapshotTaken", epochRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SnapshotsSnapshotTaken)
				if err := _Snapshots.contract.UnpackLog(event, "SnapshotTaken", log); err != nil {
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

// ParseSnapshotTaken is a log parse operation binding the contract event 0x24b0dff7469a7007db81d741ef90d7966936fb78bc19d667f4575ecbf56ab350.
//
// Solidity: event SnapshotTaken(uint256 chainId, uint256 indexed epoch, uint256 height, address indexed validator, bool isSafeToProceedConsensus, bytes signatureRaw)
func (_Snapshots *SnapshotsFilterer) ParseSnapshotTaken(log types.Log) (*SnapshotsSnapshotTaken, error) {
	event := new(SnapshotsSnapshotTaken)
	if err := _Snapshots.contract.UnpackLog(event, "SnapshotTaken", log); err != nil {
		return nil, err
	}
	return event, nil
}
