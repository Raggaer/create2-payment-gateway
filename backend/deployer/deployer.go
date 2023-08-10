// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package deployer

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

// ContractFactoryMetaData contains all meta data concerning the ContractFactory contract.
var ContractFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Deployed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"code\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"deploy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"4af63f02": "deploy(bytes,bytes32)",
	},
	Bin: "0x608060405234801561001057600080fd5b5060405161030438038061030483398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b610271806100936000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80634af63f0214610030575b600080fd5b61004361003e366004610186565b61005f565b6040516001600160a01b03909116815260200160405180910390f35b600080546001600160a01b031633146100c95760405162461bcd60e51b815260206004820152602160248201527f4f6e6c79206f776e65722063616e2063616c6c20746869732066756e6374696f6044820152603760f91b60648201526084015b60405180910390fd5b6000828451602086016000f590506001600160a01b03811661012d5760405162461bcd60e51b815260206004820152601960248201527f4661696c656420746f206465706c6f7920636f6e74726163740000000000000060448201526064016100c0565b6040516001600160a01b03821681527ff40fcec21964ffb566044d083b4073f29f7f7929110ea19e1b3ebe375d89055e9060200160405180910390a19392505050565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561019957600080fd5b823567ffffffffffffffff808211156101b157600080fd5b818501915085601f8301126101c557600080fd5b8135818111156101d7576101d7610170565b604051601f8201601f19908116603f011681019083821181831017156101ff576101ff610170565b8160405282815288602084870101111561021857600080fd5b82602086016020830137600060209382018401529896909101359650505050505056fea2646970667358221220e6ae5c3865c8d02f06b0f40a6a943e87cd865bc557697921f73940f11f0a505c64736f6c63430008130033",
}

// ContractFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractFactoryMetaData.ABI instead.
var ContractFactoryABI = ContractFactoryMetaData.ABI

// Deprecated: Use ContractFactoryMetaData.Sigs instead.
// ContractFactoryFuncSigs maps the 4-byte function signature to its string representation.
var ContractFactoryFuncSigs = ContractFactoryMetaData.Sigs

// ContractFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractFactoryMetaData.Bin instead.
var ContractFactoryBin = ContractFactoryMetaData.Bin

// DeployContractFactory deploys a new Ethereum contract, binding an instance of ContractFactory to it.
func DeployContractFactory(auth *bind.TransactOpts, backend bind.ContractBackend, owner common.Address) (common.Address, *types.Transaction, *ContractFactory, error) {
	parsed, err := ContractFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractFactoryBin), backend, owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractFactory{ContractFactoryCaller: ContractFactoryCaller{contract: contract}, ContractFactoryTransactor: ContractFactoryTransactor{contract: contract}, ContractFactoryFilterer: ContractFactoryFilterer{contract: contract}}, nil
}

// ContractFactory is an auto generated Go binding around an Ethereum contract.
type ContractFactory struct {
	ContractFactoryCaller     // Read-only binding to the contract
	ContractFactoryTransactor // Write-only binding to the contract
	ContractFactoryFilterer   // Log filterer for contract events
}

// ContractFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractFactorySession struct {
	Contract     *ContractFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractFactoryCallerSession struct {
	Contract *ContractFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ContractFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractFactoryTransactorSession struct {
	Contract     *ContractFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ContractFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractFactoryRaw struct {
	Contract *ContractFactory // Generic contract binding to access the raw methods on
}

// ContractFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractFactoryCallerRaw struct {
	Contract *ContractFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ContractFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractFactoryTransactorRaw struct {
	Contract *ContractFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractFactory creates a new instance of ContractFactory, bound to a specific deployed contract.
func NewContractFactory(address common.Address, backend bind.ContractBackend) (*ContractFactory, error) {
	contract, err := bindContractFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractFactory{ContractFactoryCaller: ContractFactoryCaller{contract: contract}, ContractFactoryTransactor: ContractFactoryTransactor{contract: contract}, ContractFactoryFilterer: ContractFactoryFilterer{contract: contract}}, nil
}

// NewContractFactoryCaller creates a new read-only instance of ContractFactory, bound to a specific deployed contract.
func NewContractFactoryCaller(address common.Address, caller bind.ContractCaller) (*ContractFactoryCaller, error) {
	contract, err := bindContractFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractFactoryCaller{contract: contract}, nil
}

// NewContractFactoryTransactor creates a new write-only instance of ContractFactory, bound to a specific deployed contract.
func NewContractFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractFactoryTransactor, error) {
	contract, err := bindContractFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractFactoryTransactor{contract: contract}, nil
}

// NewContractFactoryFilterer creates a new log filterer instance of ContractFactory, bound to a specific deployed contract.
func NewContractFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFactoryFilterer, error) {
	contract, err := bindContractFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFactoryFilterer{contract: contract}, nil
}

// bindContractFactory binds a generic wrapper to an already deployed contract.
func bindContractFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractFactory *ContractFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractFactory.Contract.ContractFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractFactory *ContractFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractFactory.Contract.ContractFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractFactory *ContractFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractFactory.Contract.ContractFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractFactory *ContractFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractFactory *ContractFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractFactory *ContractFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractFactory.Contract.contract.Transact(opts, method, params...)
}

// Deploy is a paid mutator transaction binding the contract method 0x4af63f02.
//
// Solidity: function deploy(bytes code, bytes32 salt) returns(address)
func (_ContractFactory *ContractFactoryTransactor) Deploy(opts *bind.TransactOpts, code []byte, salt [32]byte) (*types.Transaction, error) {
	return _ContractFactory.contract.Transact(opts, "deploy", code, salt)
}

// Deploy is a paid mutator transaction binding the contract method 0x4af63f02.
//
// Solidity: function deploy(bytes code, bytes32 salt) returns(address)
func (_ContractFactory *ContractFactorySession) Deploy(code []byte, salt [32]byte) (*types.Transaction, error) {
	return _ContractFactory.Contract.Deploy(&_ContractFactory.TransactOpts, code, salt)
}

// Deploy is a paid mutator transaction binding the contract method 0x4af63f02.
//
// Solidity: function deploy(bytes code, bytes32 salt) returns(address)
func (_ContractFactory *ContractFactoryTransactorSession) Deploy(code []byte, salt [32]byte) (*types.Transaction, error) {
	return _ContractFactory.Contract.Deploy(&_ContractFactory.TransactOpts, code, salt)
}

// ContractFactoryDeployedIterator is returned from FilterDeployed and is used to iterate over the raw logs and unpacked data for Deployed events raised by the ContractFactory contract.
type ContractFactoryDeployedIterator struct {
	Event *ContractFactoryDeployed // Event containing the contract specifics and raw log

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
func (it *ContractFactoryDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFactoryDeployed)
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
		it.Event = new(ContractFactoryDeployed)
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
func (it *ContractFactoryDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractFactoryDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractFactoryDeployed represents a Deployed event raised by the ContractFactory contract.
type ContractFactoryDeployed struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDeployed is a free log retrieval operation binding the contract event 0xf40fcec21964ffb566044d083b4073f29f7f7929110ea19e1b3ebe375d89055e.
//
// Solidity: event Deployed(address addr)
func (_ContractFactory *ContractFactoryFilterer) FilterDeployed(opts *bind.FilterOpts) (*ContractFactoryDeployedIterator, error) {

	logs, sub, err := _ContractFactory.contract.FilterLogs(opts, "Deployed")
	if err != nil {
		return nil, err
	}
	return &ContractFactoryDeployedIterator{contract: _ContractFactory.contract, event: "Deployed", logs: logs, sub: sub}, nil
}

// WatchDeployed is a free log subscription operation binding the contract event 0xf40fcec21964ffb566044d083b4073f29f7f7929110ea19e1b3ebe375d89055e.
//
// Solidity: event Deployed(address addr)
func (_ContractFactory *ContractFactoryFilterer) WatchDeployed(opts *bind.WatchOpts, sink chan<- *ContractFactoryDeployed) (event.Subscription, error) {

	logs, sub, err := _ContractFactory.contract.WatchLogs(opts, "Deployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractFactoryDeployed)
				if err := _ContractFactory.contract.UnpackLog(event, "Deployed", log); err != nil {
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

// ParseDeployed is a log parse operation binding the contract event 0xf40fcec21964ffb566044d083b4073f29f7f7929110ea19e1b3ebe375d89055e.
//
// Solidity: event Deployed(address addr)
func (_ContractFactory *ContractFactoryFilterer) ParseDeployed(log types.Log) (*ContractFactoryDeployed, error) {
	event := new(ContractFactoryDeployed)
	if err := _ContractFactory.contract.UnpackLog(event, "Deployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
