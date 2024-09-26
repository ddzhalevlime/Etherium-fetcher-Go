// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package web3

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

// SimplePersonInfoContractMetaData contains all meta data concerning the SimplePersonInfoContract contract.
var SimplePersonInfoContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"personIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAge\",\"type\":\"uint256\"}],\"name\":\"PersonInfoUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_personIndex\",\"type\":\"uint256\"}],\"name\":\"getPersonInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPersonsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"persons\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"age\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_age\",\"type\":\"uint256\"}],\"name\":\"setPersonInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SimplePersonInfoContractABI is the input ABI used to generate the binding from.
// Deprecated: Use SimplePersonInfoContractMetaData.ABI instead.
var SimplePersonInfoContractABI = SimplePersonInfoContractMetaData.ABI

// SimplePersonInfoContract is an auto generated Go binding around an Ethereum contract.
type SimplePersonInfoContract struct {
	SimplePersonInfoContractCaller     // Read-only binding to the contract
	SimplePersonInfoContractTransactor // Write-only binding to the contract
	SimplePersonInfoContractFilterer   // Log filterer for contract events
}

// SimplePersonInfoContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimplePersonInfoContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplePersonInfoContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimplePersonInfoContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplePersonInfoContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimplePersonInfoContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplePersonInfoContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimplePersonInfoContractSession struct {
	Contract     *SimplePersonInfoContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SimplePersonInfoContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimplePersonInfoContractCallerSession struct {
	Contract *SimplePersonInfoContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// SimplePersonInfoContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimplePersonInfoContractTransactorSession struct {
	Contract     *SimplePersonInfoContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// SimplePersonInfoContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimplePersonInfoContractRaw struct {
	Contract *SimplePersonInfoContract // Generic contract binding to access the raw methods on
}

// SimplePersonInfoContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimplePersonInfoContractCallerRaw struct {
	Contract *SimplePersonInfoContractCaller // Generic read-only contract binding to access the raw methods on
}

// SimplePersonInfoContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimplePersonInfoContractTransactorRaw struct {
	Contract *SimplePersonInfoContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimplePersonInfoContract creates a new instance of SimplePersonInfoContract, bound to a specific deployed contract.
func NewSimplePersonInfoContract(address common.Address, backend bind.ContractBackend) (*SimplePersonInfoContract, error) {
	contract, err := bindSimplePersonInfoContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimplePersonInfoContract{SimplePersonInfoContractCaller: SimplePersonInfoContractCaller{contract: contract}, SimplePersonInfoContractTransactor: SimplePersonInfoContractTransactor{contract: contract}, SimplePersonInfoContractFilterer: SimplePersonInfoContractFilterer{contract: contract}}, nil
}

// NewSimplePersonInfoContractCaller creates a new read-only instance of SimplePersonInfoContract, bound to a specific deployed contract.
func NewSimplePersonInfoContractCaller(address common.Address, caller bind.ContractCaller) (*SimplePersonInfoContractCaller, error) {
	contract, err := bindSimplePersonInfoContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimplePersonInfoContractCaller{contract: contract}, nil
}

// NewSimplePersonInfoContractTransactor creates a new write-only instance of SimplePersonInfoContract, bound to a specific deployed contract.
func NewSimplePersonInfoContractTransactor(address common.Address, transactor bind.ContractTransactor) (*SimplePersonInfoContractTransactor, error) {
	contract, err := bindSimplePersonInfoContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimplePersonInfoContractTransactor{contract: contract}, nil
}

// NewSimplePersonInfoContractFilterer creates a new log filterer instance of SimplePersonInfoContract, bound to a specific deployed contract.
func NewSimplePersonInfoContractFilterer(address common.Address, filterer bind.ContractFilterer) (*SimplePersonInfoContractFilterer, error) {
	contract, err := bindSimplePersonInfoContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimplePersonInfoContractFilterer{contract: contract}, nil
}

// bindSimplePersonInfoContract binds a generic wrapper to an already deployed contract.
func bindSimplePersonInfoContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimplePersonInfoContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimplePersonInfoContract *SimplePersonInfoContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimplePersonInfoContract.Contract.SimplePersonInfoContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimplePersonInfoContract *SimplePersonInfoContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimplePersonInfoContract.Contract.SimplePersonInfoContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimplePersonInfoContract *SimplePersonInfoContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimplePersonInfoContract.Contract.SimplePersonInfoContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimplePersonInfoContract *SimplePersonInfoContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimplePersonInfoContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimplePersonInfoContract *SimplePersonInfoContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimplePersonInfoContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimplePersonInfoContract *SimplePersonInfoContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimplePersonInfoContract.Contract.contract.Transact(opts, method, params...)
}

// GetPersonInfo is a free data retrieval call binding the contract method 0xd336ac80.
//
// Solidity: function getPersonInfo(uint256 _personIndex) view returns(string, uint256)
func (_SimplePersonInfoContract *SimplePersonInfoContractCaller) GetPersonInfo(opts *bind.CallOpts, _personIndex *big.Int) (string, *big.Int, error) {
	var out []interface{}
	err := _SimplePersonInfoContract.contract.Call(opts, &out, "getPersonInfo", _personIndex)

	if err != nil {
		return *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPersonInfo is a free data retrieval call binding the contract method 0xd336ac80.
//
// Solidity: function getPersonInfo(uint256 _personIndex) view returns(string, uint256)
func (_SimplePersonInfoContract *SimplePersonInfoContractSession) GetPersonInfo(_personIndex *big.Int) (string, *big.Int, error) {
	return _SimplePersonInfoContract.Contract.GetPersonInfo(&_SimplePersonInfoContract.CallOpts, _personIndex)
}

// GetPersonInfo is a free data retrieval call binding the contract method 0xd336ac80.
//
// Solidity: function getPersonInfo(uint256 _personIndex) view returns(string, uint256)
func (_SimplePersonInfoContract *SimplePersonInfoContractCallerSession) GetPersonInfo(_personIndex *big.Int) (string, *big.Int, error) {
	return _SimplePersonInfoContract.Contract.GetPersonInfo(&_SimplePersonInfoContract.CallOpts, _personIndex)
}

// GetPersonsCount is a free data retrieval call binding the contract method 0x8f97cff0.
//
// Solidity: function getPersonsCount() view returns(uint256)
func (_SimplePersonInfoContract *SimplePersonInfoContractCaller) GetPersonsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SimplePersonInfoContract.contract.Call(opts, &out, "getPersonsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPersonsCount is a free data retrieval call binding the contract method 0x8f97cff0.
//
// Solidity: function getPersonsCount() view returns(uint256)
func (_SimplePersonInfoContract *SimplePersonInfoContractSession) GetPersonsCount() (*big.Int, error) {
	return _SimplePersonInfoContract.Contract.GetPersonsCount(&_SimplePersonInfoContract.CallOpts)
}

// GetPersonsCount is a free data retrieval call binding the contract method 0x8f97cff0.
//
// Solidity: function getPersonsCount() view returns(uint256)
func (_SimplePersonInfoContract *SimplePersonInfoContractCallerSession) GetPersonsCount() (*big.Int, error) {
	return _SimplePersonInfoContract.Contract.GetPersonsCount(&_SimplePersonInfoContract.CallOpts)
}

// Persons is a free data retrieval call binding the contract method 0xa2f9eac6.
//
// Solidity: function persons(uint256 ) view returns(string name, uint256 age)
func (_SimplePersonInfoContract *SimplePersonInfoContractCaller) Persons(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name string
	Age  *big.Int
}, error) {
	var out []interface{}
	err := _SimplePersonInfoContract.contract.Call(opts, &out, "persons", arg0)

	outstruct := new(struct {
		Name string
		Age  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Age = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Persons is a free data retrieval call binding the contract method 0xa2f9eac6.
//
// Solidity: function persons(uint256 ) view returns(string name, uint256 age)
func (_SimplePersonInfoContract *SimplePersonInfoContractSession) Persons(arg0 *big.Int) (struct {
	Name string
	Age  *big.Int
}, error) {
	return _SimplePersonInfoContract.Contract.Persons(&_SimplePersonInfoContract.CallOpts, arg0)
}

// Persons is a free data retrieval call binding the contract method 0xa2f9eac6.
//
// Solidity: function persons(uint256 ) view returns(string name, uint256 age)
func (_SimplePersonInfoContract *SimplePersonInfoContractCallerSession) Persons(arg0 *big.Int) (struct {
	Name string
	Age  *big.Int
}, error) {
	return _SimplePersonInfoContract.Contract.Persons(&_SimplePersonInfoContract.CallOpts, arg0)
}

// SetPersonInfo is a paid mutator transaction binding the contract method 0x33f3b2a4.
//
// Solidity: function setPersonInfo(string _name, uint256 _age) returns()
func (_SimplePersonInfoContract *SimplePersonInfoContractTransactor) SetPersonInfo(opts *bind.TransactOpts, _name string, _age *big.Int) (*types.Transaction, error) {
	return _SimplePersonInfoContract.contract.Transact(opts, "setPersonInfo", _name, _age)
}

// SetPersonInfo is a paid mutator transaction binding the contract method 0x33f3b2a4.
//
// Solidity: function setPersonInfo(string _name, uint256 _age) returns()
func (_SimplePersonInfoContract *SimplePersonInfoContractSession) SetPersonInfo(_name string, _age *big.Int) (*types.Transaction, error) {
	return _SimplePersonInfoContract.Contract.SetPersonInfo(&_SimplePersonInfoContract.TransactOpts, _name, _age)
}

// SetPersonInfo is a paid mutator transaction binding the contract method 0x33f3b2a4.
//
// Solidity: function setPersonInfo(string _name, uint256 _age) returns()
func (_SimplePersonInfoContract *SimplePersonInfoContractTransactorSession) SetPersonInfo(_name string, _age *big.Int) (*types.Transaction, error) {
	return _SimplePersonInfoContract.Contract.SetPersonInfo(&_SimplePersonInfoContract.TransactOpts, _name, _age)
}

// SimplePersonInfoContractPersonInfoUpdatedIterator is returned from FilterPersonInfoUpdated and is used to iterate over the raw logs and unpacked data for PersonInfoUpdated events raised by the SimplePersonInfoContract contract.
type SimplePersonInfoContractPersonInfoUpdatedIterator struct {
	Event *SimplePersonInfoContractPersonInfoUpdated // Event containing the contract specifics and raw log

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
func (it *SimplePersonInfoContractPersonInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimplePersonInfoContractPersonInfoUpdated)
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
		it.Event = new(SimplePersonInfoContractPersonInfoUpdated)
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
func (it *SimplePersonInfoContractPersonInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimplePersonInfoContractPersonInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimplePersonInfoContractPersonInfoUpdated represents a PersonInfoUpdated event raised by the SimplePersonInfoContract contract.
type SimplePersonInfoContractPersonInfoUpdated struct {
	PersonIndex *big.Int
	NewName     string
	NewAge      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPersonInfoUpdated is a free log retrieval operation binding the contract event 0x96fb71ab58332a1b713976cc33c58781380e987f4cf4f8b2ef62be13218fec32.
//
// Solidity: event PersonInfoUpdated(uint256 indexed personIndex, string newName, uint256 newAge)
func (_SimplePersonInfoContract *SimplePersonInfoContractFilterer) FilterPersonInfoUpdated(opts *bind.FilterOpts, personIndex []*big.Int) (*SimplePersonInfoContractPersonInfoUpdatedIterator, error) {

	var personIndexRule []interface{}
	for _, personIndexItem := range personIndex {
		personIndexRule = append(personIndexRule, personIndexItem)
	}

	logs, sub, err := _SimplePersonInfoContract.contract.FilterLogs(opts, "PersonInfoUpdated", personIndexRule)
	if err != nil {
		return nil, err
	}
	return &SimplePersonInfoContractPersonInfoUpdatedIterator{contract: _SimplePersonInfoContract.contract, event: "PersonInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchPersonInfoUpdated is a free log subscription operation binding the contract event 0x96fb71ab58332a1b713976cc33c58781380e987f4cf4f8b2ef62be13218fec32.
//
// Solidity: event PersonInfoUpdated(uint256 indexed personIndex, string newName, uint256 newAge)
func (_SimplePersonInfoContract *SimplePersonInfoContractFilterer) WatchPersonInfoUpdated(opts *bind.WatchOpts, sink chan<- *SimplePersonInfoContractPersonInfoUpdated, personIndex []*big.Int) (event.Subscription, error) {

	var personIndexRule []interface{}
	for _, personIndexItem := range personIndex {
		personIndexRule = append(personIndexRule, personIndexItem)
	}

	logs, sub, err := _SimplePersonInfoContract.contract.WatchLogs(opts, "PersonInfoUpdated", personIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimplePersonInfoContractPersonInfoUpdated)
				if err := _SimplePersonInfoContract.contract.UnpackLog(event, "PersonInfoUpdated", log); err != nil {
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

// ParsePersonInfoUpdated is a log parse operation binding the contract event 0x96fb71ab58332a1b713976cc33c58781380e987f4cf4f8b2ef62be13218fec32.
//
// Solidity: event PersonInfoUpdated(uint256 indexed personIndex, string newName, uint256 newAge)
func (_SimplePersonInfoContract *SimplePersonInfoContractFilterer) ParsePersonInfoUpdated(log types.Log) (*SimplePersonInfoContractPersonInfoUpdated, error) {
	event := new(SimplePersonInfoContractPersonInfoUpdated)
	if err := _SimplePersonInfoContract.contract.UnpackLog(event, "PersonInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
