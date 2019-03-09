// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum

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

// ERC725ABI is the input ABI used to generate the binding from.
const ERC725ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_operationType\",\"type\":\"uint256\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getData\",\"outputs\":[{\"name\":\"_value\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"bytes32\"}],\"name\":\"setData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"DataChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"ContractCreated\",\"type\":\"event\"}]"

// ERC725Bin is the compiled bytecode used for deploying new contracts.
const ERC725Bin = `0x`

// DeployERC725 deploys a new Ethereum contract, binding an instance of ERC725 to it.
func DeployERC725(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC725, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC725ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC725Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC725{ERC725Caller: ERC725Caller{contract: contract}, ERC725Transactor: ERC725Transactor{contract: contract}, ERC725Filterer: ERC725Filterer{contract: contract}}, nil
}

// ERC725 is an auto generated Go binding around an Ethereum contract.
type ERC725 struct {
	ERC725Caller     // Read-only binding to the contract
	ERC725Transactor // Write-only binding to the contract
	ERC725Filterer   // Log filterer for contract events
}

// ERC725Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC725Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC725Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC725Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC725Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC725Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC725Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC725Session struct {
	Contract     *ERC725           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC725CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC725CallerSession struct {
	Contract *ERC725Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC725TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC725TransactorSession struct {
	Contract     *ERC725Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC725Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC725Raw struct {
	Contract *ERC725 // Generic contract binding to access the raw methods on
}

// ERC725CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC725CallerRaw struct {
	Contract *ERC725Caller // Generic read-only contract binding to access the raw methods on
}

// ERC725TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC725TransactorRaw struct {
	Contract *ERC725Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC725 creates a new instance of ERC725, bound to a specific deployed contract.
func NewERC725(address common.Address, backend bind.ContractBackend) (*ERC725, error) {
	contract, err := bindERC725(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC725{ERC725Caller: ERC725Caller{contract: contract}, ERC725Transactor: ERC725Transactor{contract: contract}, ERC725Filterer: ERC725Filterer{contract: contract}}, nil
}

// NewERC725Caller creates a new read-only instance of ERC725, bound to a specific deployed contract.
func NewERC725Caller(address common.Address, caller bind.ContractCaller) (*ERC725Caller, error) {
	contract, err := bindERC725(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC725Caller{contract: contract}, nil
}

// NewERC725Transactor creates a new write-only instance of ERC725, bound to a specific deployed contract.
func NewERC725Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC725Transactor, error) {
	contract, err := bindERC725(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC725Transactor{contract: contract}, nil
}

// NewERC725Filterer creates a new log filterer instance of ERC725, bound to a specific deployed contract.
func NewERC725Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC725Filterer, error) {
	contract, err := bindERC725(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC725Filterer{contract: contract}, nil
}

// bindERC725 binds a generic wrapper to an already deployed contract.
func bindERC725(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC725ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC725 *ERC725Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC725.Contract.ERC725Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC725 *ERC725Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC725.Contract.ERC725Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC725 *ERC725Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC725.Contract.ERC725Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC725 *ERC725CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC725.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC725 *ERC725TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC725.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC725 *ERC725TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC725.Contract.contract.Transact(opts, method, params...)
}

// GetData is a free data retrieval call binding the contract method 0x54f6127f.
//
// Solidity: function getData(bytes32 _key) constant returns(bytes32 _value)
func (_ERC725 *ERC725Caller) GetData(opts *bind.CallOpts, _key [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ERC725.contract.Call(opts, out, "getData", _key)
	return *ret0, err
}

// GetData is a free data retrieval call binding the contract method 0x54f6127f.
//
// Solidity: function getData(bytes32 _key) constant returns(bytes32 _value)
func (_ERC725 *ERC725Session) GetData(_key [32]byte) ([32]byte, error) {
	return _ERC725.Contract.GetData(&_ERC725.CallOpts, _key)
}

// GetData is a free data retrieval call binding the contract method 0x54f6127f.
//
// Solidity: function getData(bytes32 _key) constant returns(bytes32 _value)
func (_ERC725 *ERC725CallerSession) GetData(_key [32]byte) ([32]byte, error) {
	return _ERC725.Contract.GetData(&_ERC725.CallOpts, _key)
}

// Execute is a paid mutator transaction binding the contract method 0x44c028fe.
//
// Solidity: function execute(uint256 _operationType, address _to, uint256 _value, bytes _data) returns()
func (_ERC725 *ERC725Transactor) Execute(opts *bind.TransactOpts, _operationType *big.Int, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC725.contract.Transact(opts, "execute", _operationType, _to, _value, _data)
}

// Execute is a paid mutator transaction binding the contract method 0x44c028fe.
//
// Solidity: function execute(uint256 _operationType, address _to, uint256 _value, bytes _data) returns()
func (_ERC725 *ERC725Session) Execute(_operationType *big.Int, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC725.Contract.Execute(&_ERC725.TransactOpts, _operationType, _to, _value, _data)
}

// Execute is a paid mutator transaction binding the contract method 0x44c028fe.
//
// Solidity: function execute(uint256 _operationType, address _to, uint256 _value, bytes _data) returns()
func (_ERC725 *ERC725TransactorSession) Execute(_operationType *big.Int, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC725.Contract.Execute(&_ERC725.TransactOpts, _operationType, _to, _value, _data)
}

// SetData is a paid mutator transaction binding the contract method 0x749ebfb8.
//
// Solidity: function setData(bytes32 _key, bytes32 _value) returns()
func (_ERC725 *ERC725Transactor) SetData(opts *bind.TransactOpts, _key [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _ERC725.contract.Transact(opts, "setData", _key, _value)
}

// SetData is a paid mutator transaction binding the contract method 0x749ebfb8.
//
// Solidity: function setData(bytes32 _key, bytes32 _value) returns()
func (_ERC725 *ERC725Session) SetData(_key [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _ERC725.Contract.SetData(&_ERC725.TransactOpts, _key, _value)
}

// SetData is a paid mutator transaction binding the contract method 0x749ebfb8.
//
// Solidity: function setData(bytes32 _key, bytes32 _value) returns()
func (_ERC725 *ERC725TransactorSession) SetData(_key [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _ERC725.Contract.SetData(&_ERC725.TransactOpts, _key, _value)
}

// ERC725ContractCreatedIterator is returned from FilterContractCreated and is used to iterate over the raw logs and unpacked data for ContractCreated events raised by the ERC725 contract.
type ERC725ContractCreatedIterator struct {
	Event *ERC725ContractCreated // Event containing the contract specifics and raw log

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
func (it *ERC725ContractCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC725ContractCreated)
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
		it.Event = new(ERC725ContractCreated)
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
func (it *ERC725ContractCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC725ContractCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC725ContractCreated represents a ContractCreated event raised by the ERC725 contract.
type ERC725ContractCreated struct {
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterContractCreated is a free log retrieval operation binding the contract event 0xcf78cf0d6f3d8371e1075c69c492ab4ec5d8cf23a1a239b6a51a1d00be7ca312.
//
// Solidity: event ContractCreated(address indexed contractAddress)
func (_ERC725 *ERC725Filterer) FilterContractCreated(opts *bind.FilterOpts, contractAddress []common.Address) (*ERC725ContractCreatedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _ERC725.contract.FilterLogs(opts, "ContractCreated", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC725ContractCreatedIterator{contract: _ERC725.contract, event: "ContractCreated", logs: logs, sub: sub}, nil
}

// WatchContractCreated is a free log subscription operation binding the contract event 0xcf78cf0d6f3d8371e1075c69c492ab4ec5d8cf23a1a239b6a51a1d00be7ca312.
//
// Solidity: event ContractCreated(address indexed contractAddress)
func (_ERC725 *ERC725Filterer) WatchContractCreated(opts *bind.WatchOpts, sink chan<- *ERC725ContractCreated, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _ERC725.contract.WatchLogs(opts, "ContractCreated", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC725ContractCreated)
				if err := _ERC725.contract.UnpackLog(event, "ContractCreated", log); err != nil {
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

// ERC725DataChangedIterator is returned from FilterDataChanged and is used to iterate over the raw logs and unpacked data for DataChanged events raised by the ERC725 contract.
type ERC725DataChangedIterator struct {
	Event *ERC725DataChanged // Event containing the contract specifics and raw log

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
func (it *ERC725DataChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC725DataChanged)
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
		it.Event = new(ERC725DataChanged)
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
func (it *ERC725DataChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC725DataChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC725DataChanged represents a DataChanged event raised by the ERC725 contract.
type ERC725DataChanged struct {
	Key   [32]byte
	Value [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDataChanged is a free log retrieval operation binding the contract event 0x35553580e4553c909abeb5764e842ce1f93c45f9f614bde2a2ca5f5b7b7dc0fb.
//
// Solidity: event DataChanged(bytes32 indexed key, bytes32 indexed value)
func (_ERC725 *ERC725Filterer) FilterDataChanged(opts *bind.FilterOpts, key [][32]byte, value [][32]byte) (*ERC725DataChangedIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _ERC725.contract.FilterLogs(opts, "DataChanged", keyRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &ERC725DataChangedIterator{contract: _ERC725.contract, event: "DataChanged", logs: logs, sub: sub}, nil
}

// WatchDataChanged is a free log subscription operation binding the contract event 0x35553580e4553c909abeb5764e842ce1f93c45f9f614bde2a2ca5f5b7b7dc0fb.
//
// Solidity: event DataChanged(bytes32 indexed key, bytes32 indexed value)
func (_ERC725 *ERC725Filterer) WatchDataChanged(opts *bind.WatchOpts, sink chan<- *ERC725DataChanged, key [][32]byte, value [][32]byte) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _ERC725.contract.WatchLogs(opts, "DataChanged", keyRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC725DataChanged)
				if err := _ERC725.contract.UnpackLog(event, "DataChanged", log); err != nil {
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

// IdentityABI is the input ABI used to generate the binding from.
const IdentityABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_operationType\",\"type\":\"uint256\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getData\",\"outputs\":[{\"name\":\"_value\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"store\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"bytes32\"}],\"name\":\"setData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"DataChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"ContractCreated\",\"type\":\"event\"}]"

// IdentityBin is the compiled bytecode used for deploying new contracts.
const IdentityBin = `0x608060405234801561001057600080fd5b506040516020806105128339810180604052602081101561003057600080fd5b505161004481640100000000610075810204565b6000805260016020527fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb49555061007d565b600081905290565b6104868061008c6000396000f3fe608060405234801561001057600080fd5b5060043610610068577c0100000000000000000000000000000000000000000000000000000000600035046344c028fe811461006d57806354f6127f14610108578063654cf88c14610137578063749ebfb814610154575b600080fd5b6101066004803603608081101561008357600080fd5b81359173ffffffffffffffffffffffffffffffffffffffff60208201351691604082013591908101906080810160608201356401000000008111156100c757600080fd5b8201836020820111156100d957600080fd5b803590602001918460018302840111640100000000831117156100fb57600080fd5b509092509050610177565b005b6101256004803603602081101561011e57600080fd5b50356102fc565b60408051918252519081900360200190f35b6101256004803603602081101561014d57600080fd5b503561030e565b6101066004803603604081101561016a57600080fd5b5080359060200135610320565b6000805260016020527fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb49546101ab336103ff565b1461021757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f6f6e6c792d6f776e65722d616c6c6f7765640000000000000000000000000000604482015290519081900360640190fd5b8415156102655761025f848484848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061040792505050565b506102f5565b60018514156100685760006102af83838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061044992505050565b60405190915073ffffffffffffffffffffffffffffffffffffffff8216907fcf78cf0d6f3d8371e1075c69c492ab4ec5d8cf23a1a239b6a51a1d00be7ca31290600090a2505b5050505050565b60009081526001602052604090205490565b60016020526000908152604090205481565b6000805260016020527fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb4954610354336103ff565b146103c057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f6f6e6c792d6f776e65722d616c6c6f7765640000000000000000000000000000604482015290519081900360640190fd5b60008281526001602052604080822083905551829184917f35553580e4553c909abeb5764e842ce1f93c45f9f614bde2a2ca5f5b7b7dc0fb9190a35050565b600081905290565b8051604051600091906020840183828483898b6187965a03f19350838015610436576001811461043a5761043e565b8383fd5b8383f35b505050509392505050565b60008151602083016000f09291505056fea165627a7a723058204623339372eb0e1e02aedada5b45c119166d37f580672be6500a9b4bc2bdf8610029`

// DeployIdentity deploys a new Ethereum contract, binding an instance of Identity to it.
func DeployIdentity(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *Identity, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IdentityBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Identity{IdentityCaller: IdentityCaller{contract: contract}, IdentityTransactor: IdentityTransactor{contract: contract}, IdentityFilterer: IdentityFilterer{contract: contract}}, nil
}

// Identity is an auto generated Go binding around an Ethereum contract.
type Identity struct {
	IdentityCaller     // Read-only binding to the contract
	IdentityTransactor // Write-only binding to the contract
	IdentityFilterer   // Log filterer for contract events
}

// IdentityCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentitySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentitySession struct {
	Contract     *Identity         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityCallerSession struct {
	Contract *IdentityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IdentityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityTransactorSession struct {
	Contract     *IdentityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IdentityRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityRaw struct {
	Contract *Identity // Generic contract binding to access the raw methods on
}

// IdentityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityCallerRaw struct {
	Contract *IdentityCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityTransactorRaw struct {
	Contract *IdentityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentity creates a new instance of Identity, bound to a specific deployed contract.
func NewIdentity(address common.Address, backend bind.ContractBackend) (*Identity, error) {
	contract, err := bindIdentity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Identity{IdentityCaller: IdentityCaller{contract: contract}, IdentityTransactor: IdentityTransactor{contract: contract}, IdentityFilterer: IdentityFilterer{contract: contract}}, nil
}

// NewIdentityCaller creates a new read-only instance of Identity, bound to a specific deployed contract.
func NewIdentityCaller(address common.Address, caller bind.ContractCaller) (*IdentityCaller, error) {
	contract, err := bindIdentity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityCaller{contract: contract}, nil
}

// NewIdentityTransactor creates a new write-only instance of Identity, bound to a specific deployed contract.
func NewIdentityTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityTransactor, error) {
	contract, err := bindIdentity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityTransactor{contract: contract}, nil
}

// NewIdentityFilterer creates a new log filterer instance of Identity, bound to a specific deployed contract.
func NewIdentityFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityFilterer, error) {
	contract, err := bindIdentity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityFilterer{contract: contract}, nil
}

// bindIdentity binds a generic wrapper to an already deployed contract.
func bindIdentity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identity *IdentityRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Identity.Contract.IdentityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identity *IdentityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.Contract.IdentityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identity *IdentityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identity.Contract.IdentityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identity *IdentityCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Identity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identity *IdentityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identity *IdentityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identity.Contract.contract.Transact(opts, method, params...)
}

// GetData is a free data retrieval call binding the contract method 0x54f6127f.
//
// Solidity: function getData(bytes32 _key) constant returns(bytes32 _value)
func (_Identity *IdentityCaller) GetData(opts *bind.CallOpts, _key [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Identity.contract.Call(opts, out, "getData", _key)
	return *ret0, err
}

// GetData is a free data retrieval call binding the contract method 0x54f6127f.
//
// Solidity: function getData(bytes32 _key) constant returns(bytes32 _value)
func (_Identity *IdentitySession) GetData(_key [32]byte) ([32]byte, error) {
	return _Identity.Contract.GetData(&_Identity.CallOpts, _key)
}

// GetData is a free data retrieval call binding the contract method 0x54f6127f.
//
// Solidity: function getData(bytes32 _key) constant returns(bytes32 _value)
func (_Identity *IdentityCallerSession) GetData(_key [32]byte) ([32]byte, error) {
	return _Identity.Contract.GetData(&_Identity.CallOpts, _key)
}

// Store is a free data retrieval call binding the contract method 0x654cf88c.
//
// Solidity: function store(bytes32 ) constant returns(bytes32)
func (_Identity *IdentityCaller) Store(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Identity.contract.Call(opts, out, "store", arg0)
	return *ret0, err
}

// Store is a free data retrieval call binding the contract method 0x654cf88c.
//
// Solidity: function store(bytes32 ) constant returns(bytes32)
func (_Identity *IdentitySession) Store(arg0 [32]byte) ([32]byte, error) {
	return _Identity.Contract.Store(&_Identity.CallOpts, arg0)
}

// Store is a free data retrieval call binding the contract method 0x654cf88c.
//
// Solidity: function store(bytes32 ) constant returns(bytes32)
func (_Identity *IdentityCallerSession) Store(arg0 [32]byte) ([32]byte, error) {
	return _Identity.Contract.Store(&_Identity.CallOpts, arg0)
}

// Execute is a paid mutator transaction binding the contract method 0x44c028fe.
//
// Solidity: function execute(uint256 _operationType, address _to, uint256 _value, bytes _data) returns()
func (_Identity *IdentityTransactor) Execute(opts *bind.TransactOpts, _operationType *big.Int, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "execute", _operationType, _to, _value, _data)
}

// Execute is a paid mutator transaction binding the contract method 0x44c028fe.
//
// Solidity: function execute(uint256 _operationType, address _to, uint256 _value, bytes _data) returns()
func (_Identity *IdentitySession) Execute(_operationType *big.Int, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Identity.Contract.Execute(&_Identity.TransactOpts, _operationType, _to, _value, _data)
}

// Execute is a paid mutator transaction binding the contract method 0x44c028fe.
//
// Solidity: function execute(uint256 _operationType, address _to, uint256 _value, bytes _data) returns()
func (_Identity *IdentityTransactorSession) Execute(_operationType *big.Int, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Identity.Contract.Execute(&_Identity.TransactOpts, _operationType, _to, _value, _data)
}

// SetData is a paid mutator transaction binding the contract method 0x749ebfb8.
//
// Solidity: function setData(bytes32 _key, bytes32 _value) returns()
func (_Identity *IdentityTransactor) SetData(opts *bind.TransactOpts, _key [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "setData", _key, _value)
}

// SetData is a paid mutator transaction binding the contract method 0x749ebfb8.
//
// Solidity: function setData(bytes32 _key, bytes32 _value) returns()
func (_Identity *IdentitySession) SetData(_key [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _Identity.Contract.SetData(&_Identity.TransactOpts, _key, _value)
}

// SetData is a paid mutator transaction binding the contract method 0x749ebfb8.
//
// Solidity: function setData(bytes32 _key, bytes32 _value) returns()
func (_Identity *IdentityTransactorSession) SetData(_key [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _Identity.Contract.SetData(&_Identity.TransactOpts, _key, _value)
}

// IdentityContractCreatedIterator is returned from FilterContractCreated and is used to iterate over the raw logs and unpacked data for ContractCreated events raised by the Identity contract.
type IdentityContractCreatedIterator struct {
	Event *IdentityContractCreated // Event containing the contract specifics and raw log

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
func (it *IdentityContractCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityContractCreated)
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
		it.Event = new(IdentityContractCreated)
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
func (it *IdentityContractCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityContractCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityContractCreated represents a ContractCreated event raised by the Identity contract.
type IdentityContractCreated struct {
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterContractCreated is a free log retrieval operation binding the contract event 0xcf78cf0d6f3d8371e1075c69c492ab4ec5d8cf23a1a239b6a51a1d00be7ca312.
//
// Solidity: event ContractCreated(address indexed contractAddress)
func (_Identity *IdentityFilterer) FilterContractCreated(opts *bind.FilterOpts, contractAddress []common.Address) (*IdentityContractCreatedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _Identity.contract.FilterLogs(opts, "ContractCreated", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &IdentityContractCreatedIterator{contract: _Identity.contract, event: "ContractCreated", logs: logs, sub: sub}, nil
}

// WatchContractCreated is a free log subscription operation binding the contract event 0xcf78cf0d6f3d8371e1075c69c492ab4ec5d8cf23a1a239b6a51a1d00be7ca312.
//
// Solidity: event ContractCreated(address indexed contractAddress)
func (_Identity *IdentityFilterer) WatchContractCreated(opts *bind.WatchOpts, sink chan<- *IdentityContractCreated, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _Identity.contract.WatchLogs(opts, "ContractCreated", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityContractCreated)
				if err := _Identity.contract.UnpackLog(event, "ContractCreated", log); err != nil {
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

// IdentityDataChangedIterator is returned from FilterDataChanged and is used to iterate over the raw logs and unpacked data for DataChanged events raised by the Identity contract.
type IdentityDataChangedIterator struct {
	Event *IdentityDataChanged // Event containing the contract specifics and raw log

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
func (it *IdentityDataChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityDataChanged)
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
		it.Event = new(IdentityDataChanged)
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
func (it *IdentityDataChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityDataChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityDataChanged represents a DataChanged event raised by the Identity contract.
type IdentityDataChanged struct {
	Key   [32]byte
	Value [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDataChanged is a free log retrieval operation binding the contract event 0x35553580e4553c909abeb5764e842ce1f93c45f9f614bde2a2ca5f5b7b7dc0fb.
//
// Solidity: event DataChanged(bytes32 indexed key, bytes32 indexed value)
func (_Identity *IdentityFilterer) FilterDataChanged(opts *bind.FilterOpts, key [][32]byte, value [][32]byte) (*IdentityDataChangedIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Identity.contract.FilterLogs(opts, "DataChanged", keyRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &IdentityDataChangedIterator{contract: _Identity.contract, event: "DataChanged", logs: logs, sub: sub}, nil
}

// WatchDataChanged is a free log subscription operation binding the contract event 0x35553580e4553c909abeb5764e842ce1f93c45f9f614bde2a2ca5f5b7b7dc0fb.
//
// Solidity: event DataChanged(bytes32 indexed key, bytes32 indexed value)
func (_Identity *IdentityFilterer) WatchDataChanged(opts *bind.WatchOpts, sink chan<- *IdentityDataChanged, key [][32]byte, value [][32]byte) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Identity.contract.WatchLogs(opts, "DataChanged", keyRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityDataChanged)
				if err := _Identity.contract.UnpackLog(event, "DataChanged", log); err != nil {
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
