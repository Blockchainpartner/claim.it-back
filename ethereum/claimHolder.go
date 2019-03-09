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

// ClaimHolderABI is the input ABI used to generate the binding from.
const ClaimHolderABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimsByTopic\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_claimId\",\"type\":\"bytes32\"}],\"name\":\"removeClaim\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_claimId\",\"type\":\"bytes32\"},{\"name\":\"_topic\",\"type\":\"bytes32\"},{\"name\":\"_scheme\",\"type\":\"uint256\"},{\"name\":\"_issuer\",\"type\":\"address\"},{\"name\":\"_signature\",\"type\":\"bytes32\"},{\"name\":\"_data\",\"type\":\"bytes32\"},{\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"changeClaim\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_claimId\",\"type\":\"bytes32\"}],\"name\":\"toggleApprovaleClaim\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_claimId\",\"type\":\"bytes32\"}],\"name\":\"getClaim\",\"outputs\":[{\"name\":\"topic\",\"type\":\"bytes32\"},{\"name\":\"scheme\",\"type\":\"uint256\"},{\"name\":\"issuer\",\"type\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes32\"},{\"name\":\"data\",\"type\":\"bytes32\"},{\"name\":\"uri\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_topic\",\"type\":\"bytes32\"},{\"name\":\"_scheme\",\"type\":\"uint256\"},{\"name\":\"_signature\",\"type\":\"bytes32\"},{\"name\":\"_data\",\"type\":\"bytes32\"},{\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"addClaim\",\"outputs\":[{\"name\":\"claimId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"claims\",\"outputs\":[{\"name\":\"topic\",\"type\":\"bytes32\"},{\"name\":\"scheme\",\"type\":\"uint256\"},{\"name\":\"issuer\",\"type\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes32\"},{\"name\":\"data\",\"type\":\"bytes32\"},{\"name\":\"uri\",\"type\":\"string\"},{\"name\":\"recipientApproval\",\"type\":\"bool\"},{\"name\":\"parentClaimId\",\"type\":\"bytes32\"},{\"name\":\"isValid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_topic\",\"type\":\"bytes32\"}],\"name\":\"getClaimIdsByTopic\",\"outputs\":[{\"name\":\"claimIds\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"claimId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"topic\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"scheme\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"signature\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"ClaimApprovalToggled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"claimId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"topic\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"scheme\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"signature\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"ClaimAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"claimId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"topic\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"scheme\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"signature\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"ClaimRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"claimId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"topic\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"scheme\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"signature\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"ClaimChanged\",\"type\":\"event\"}]"

// ClaimHolderBin is the compiled bytecode used for deploying new contracts.
const ClaimHolderBin = `0x608060405260018054600160a060020a03191673b1dc4198dc4633105fe04cb1cf2505972498e44317905534801561003657600080fd5b506040516020806110118339810180604052602081101561005657600080fd5b505160008054600160a060020a03909216600160a060020a0319909216919091179055610f89806100886000396000f3fe608060405234801561001057600080fd5b50600436106100b0576000357c010000000000000000000000000000000000000000000000000000000090048063b23bbc7a11610083578063b23bbc7a146101dc578063c9100bcb146101f9578063df64dde2146102c3578063eff0f5921461034b578063f9e9b61814610432576100b0565b80630e4e4dfa146100b55780634eee424a146100ea5780638da5cb5b1461011b578063a85a9bd61461013f575b600080fd5b6100d8600480360360408110156100cb57600080fd5b508035906020013561049f565b60408051918252519081900360200190f35b6101076004803603602081101561010057600080fd5b50356104cf565b604080519115158252519081900360200190f35b610123610704565b60408051600160a060020a039092168252519081900360200190f35b610107600480360360e081101561015557600080fd5b813591602081013591604082013591600160a060020a036060820135169160808201359160a08101359181019060e0810160c082013564010000000081111561019d57600080fd5b8201836020820111156101af57600080fd5b803590602001918460018302840111640100000000831117156101d157600080fd5b509092509050610713565b610107600480360360208110156101f257600080fd5b50356108b0565b6102166004803603602081101561020f57600080fd5b5035610a49565b6040518087815260200186815260200185600160a060020a0316600160a060020a0316815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561028357818101518382015260200161026b565b50505050905090810190601f1680156102b05780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b6100d8600480360360a08110156102d957600080fd5b81359160208101359160408201359160608101359181019060a08101608082013564010000000081111561030c57600080fd5b82018360208201111561031e57600080fd5b8035906020019184600183028401116401000000008311171561034057600080fd5b509092509050610b2c565b6103686004803603602081101561036157600080fd5b5035610d5b565b604051808a815260200189815260200188600160a060020a0316600160a060020a03168152602001878152602001868152602001806020018515151515815260200184815260200183151515158152602001828103825286818151815260200191508051906020019080838360005b838110156103ef5781810151838201526020016103d7565b50505050905090810190601f16801561041c5780820380516001836020036101000a031916815260200191505b509a505050505050505050505060405180910390f35b61044f6004803603602081101561044857600080fd5b5035610e3b565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561048b578181015183820152602001610473565b505050509050019250505060405180910390f35b6003602052816000526040600020818154811015156104ba57fe5b90600052602060002001600091509150505481565b600081815260026020819052604082200154600160a060020a0316331415610542576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180610f396025913960400191505060405180910390fd5b60008281526002602081815260409283902080546001808301548386015460038501546004860154600590960180548a516101009682161596909602600019011698909804601f810188900488028501880190995288845293979196600160a060020a039091169593949360609392908301828280156106035780601f106105d857610100808354040283529160200191610603565b820191906000526020600020905b8154815290600101906020018083116105e657829003601f168201915b505050505090506000600260008a815260200190815260200160002060080160006101000a81548160ff02191690831515021790555083600160a060020a031686897fd8926d7f59d697f02d719bfaf05a3a564b8b17dcea1d82f69a54d6f13c5e3548888787876040518085815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156106b95781810151838201526020016106a1565b50505050905090810190601f1680156106e65780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a4506001979650505050505050565b600054600160a060020a031681565b600088815260026020819052604082200154600160a060020a0316331415610786576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180610f396025913960400191505060405180910390fd5b600089815260026020819052604090912089815560018101899055908101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03891617905560038101869055600481018590556107e6906005018484610e9d565b5060008981526002602090815260409182902060068101805460ff199081169091556008909101805490911660011790558151898152908101879052908101859052608060608201818152908201849052600160a060020a038816918a918c917f8e3848bbbb6124bac6e78d7d6f351516db9e27b626b756b5b8ada2034c4336ed918c918b918b918b918b919060a08201848480828437600083820152604051601f909101601f19169092018290039850909650505050505050a450600198975050505050505050565b60008054600160a060020a0316331461092a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f6d73672e73656e6465722073686f756c6420626520746865206f776e65720000604482015290519081900360640190fd5b60008281526002602081815260409283902060068101805460ff19811660ff909116151790558083015481546001808401546003850154600486015489518381529788018290529887018990526080606088018181526005909701805460001995811615610100029590950190941698909804978701889052600160a060020a039094169792968a967f9d524b960fb88ad472994586b295c3227def57acdcbca7804e91bea2b18aa0af969295949392909160a083019084908015610a305780601f10610a0557610100808354040283529160200191610a30565b820191906000526020600020905b815481529060010190602001808311610a1357829003601f168201915b50509550505050505060405180910390a4506001919050565b6000818152600260208181526040808420805460018083015483870154600385015460048601546005909601805488516101009682161596909602600019011699909904601f8101899004890285018901909752868452899889988998899860609897600160a060020a039096169590939290918391830182828015610b105780601f10610ae557610100808354040283529160200191610b10565b820191906000526020600020905b815481529060010190602001808311610af357829003601f168201915b5050505050905095509550955095509550955091939550919395565b6000805460408051336c01000000000000000000000000818102602080850191909152600160a060020a0390951602603483015260488083018c905283518084039091018152606890920183528151918401919091208085526002938490529190932089815560018101899055918201805473ffffffffffffffffffffffffffffffffffffffff19169093179092556003810186905560048101859055610bd7906005018484610e9d565b506000818152600260209081526040808320600801805460ff191660019081179091558a8452600383528184208054918201815584529282902090920183905581518881529081018790529081018590526080606082018181529082018490523391899184917f80095c7241bc123695ff70400336540c6e8c10a85a94249ed3eb16a51ea80304918b918b918b918b918b919060a08201848480828437600083820152604051601f909101601f19169092018290039850909650505050505050a46000805460408051600160a060020a039283166024820152604481018b9052606480820186905282518083039091018152608490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f9918925d00000000000000000000000000000000000000000000000000000000178152600154825193519295941693828483876187965a03f4808015610d425760018114610d4657610d4a565b8484fd5b8484f35b505050505050509695505050505050565b600260208181526000928352604092839020805460018083015483860154600385015460048601546005870180548b516101009782161597909702600019011699909904601f8101899004890286018901909a5289855294989297600160a060020a03909216969095939290830182828015610e185780601f10610ded57610100808354040283529160200191610e18565b820191906000526020600020905b815481529060010190602001808311610dfb57829003601f168201915b5050505060068301546007840154600890940154929360ff918216939092501689565b600081815260036020908152604091829020805483518184028101840190945280845260609392830182828015610e9157602002820191906000526020600020905b815481526020019060010190808311610e7d575b50505050509050919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610ede5782800160ff19823516178555610f0b565b82800160010185558215610f0b579182015b82811115610f0b578235825591602001919060010190610ef0565b50610f17929150610f1b565b5090565b610f3591905b80821115610f175760008155600101610f21565b9056fe6d73672e73656e6465722073686f756c642062652074686520636c61696d20697373756572a165627a7a723058208c41eaaf05b002fc00d4ff074348f646989acf0a1b6035fbc653fd41328394380029`

// DeployClaimHolder deploys a new Ethereum contract, binding an instance of ClaimHolder to it.
func DeployClaimHolder(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *ClaimHolder, error) {
	parsed, err := abi.JSON(strings.NewReader(ClaimHolderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ClaimHolderBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ClaimHolder{ClaimHolderCaller: ClaimHolderCaller{contract: contract}, ClaimHolderTransactor: ClaimHolderTransactor{contract: contract}, ClaimHolderFilterer: ClaimHolderFilterer{contract: contract}}, nil
}

// ClaimHolder is an auto generated Go binding around an Ethereum contract.
type ClaimHolder struct {
	ClaimHolderCaller     // Read-only binding to the contract
	ClaimHolderTransactor // Write-only binding to the contract
	ClaimHolderFilterer   // Log filterer for contract events
}

// ClaimHolderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClaimHolderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimHolderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClaimHolderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimHolderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClaimHolderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimHolderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClaimHolderSession struct {
	Contract     *ClaimHolder      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClaimHolderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClaimHolderCallerSession struct {
	Contract *ClaimHolderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ClaimHolderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClaimHolderTransactorSession struct {
	Contract     *ClaimHolderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ClaimHolderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClaimHolderRaw struct {
	Contract *ClaimHolder // Generic contract binding to access the raw methods on
}

// ClaimHolderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClaimHolderCallerRaw struct {
	Contract *ClaimHolderCaller // Generic read-only contract binding to access the raw methods on
}

// ClaimHolderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClaimHolderTransactorRaw struct {
	Contract *ClaimHolderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClaimHolder creates a new instance of ClaimHolder, bound to a specific deployed contract.
func NewClaimHolder(address common.Address, backend bind.ContractBackend) (*ClaimHolder, error) {
	contract, err := bindClaimHolder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClaimHolder{ClaimHolderCaller: ClaimHolderCaller{contract: contract}, ClaimHolderTransactor: ClaimHolderTransactor{contract: contract}, ClaimHolderFilterer: ClaimHolderFilterer{contract: contract}}, nil
}

// NewClaimHolderCaller creates a new read-only instance of ClaimHolder, bound to a specific deployed contract.
func NewClaimHolderCaller(address common.Address, caller bind.ContractCaller) (*ClaimHolderCaller, error) {
	contract, err := bindClaimHolder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimHolderCaller{contract: contract}, nil
}

// NewClaimHolderTransactor creates a new write-only instance of ClaimHolder, bound to a specific deployed contract.
func NewClaimHolderTransactor(address common.Address, transactor bind.ContractTransactor) (*ClaimHolderTransactor, error) {
	contract, err := bindClaimHolder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimHolderTransactor{contract: contract}, nil
}

// NewClaimHolderFilterer creates a new log filterer instance of ClaimHolder, bound to a specific deployed contract.
func NewClaimHolderFilterer(address common.Address, filterer bind.ContractFilterer) (*ClaimHolderFilterer, error) {
	contract, err := bindClaimHolder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClaimHolderFilterer{contract: contract}, nil
}

// bindClaimHolder binds a generic wrapper to an already deployed contract.
func bindClaimHolder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClaimHolderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimHolder *ClaimHolderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ClaimHolder.Contract.ClaimHolderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimHolder *ClaimHolderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimHolder.Contract.ClaimHolderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimHolder *ClaimHolderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimHolder.Contract.ClaimHolderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimHolder *ClaimHolderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ClaimHolder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimHolder *ClaimHolderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimHolder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimHolder *ClaimHolderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimHolder.Contract.contract.Transact(opts, method, params...)
}

// Claims is a free data retrieval call binding the contract method 0xeff0f592.
//
// Solidity: function claims(bytes32 ) constant returns(bytes32 topic, uint256 scheme, address issuer, bytes32 signature, bytes32 data, string uri, bool recipientApproval, bytes32 parentClaimId, bool isValid)
func (_ClaimHolder *ClaimHolderCaller) Claims(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Topic             [32]byte
	Scheme            *big.Int
	Issuer            common.Address
	Signature         [32]byte
	Data              [32]byte
	Uri               string
	RecipientApproval bool
	ParentClaimId     [32]byte
	IsValid           bool
}, error) {
	ret := new(struct {
		Topic             [32]byte
		Scheme            *big.Int
		Issuer            common.Address
		Signature         [32]byte
		Data              [32]byte
		Uri               string
		RecipientApproval bool
		ParentClaimId     [32]byte
		IsValid           bool
	})
	out := ret
	err := _ClaimHolder.contract.Call(opts, out, "claims", arg0)
	return *ret, err
}

// Claims is a free data retrieval call binding the contract method 0xeff0f592.
//
// Solidity: function claims(bytes32 ) constant returns(bytes32 topic, uint256 scheme, address issuer, bytes32 signature, bytes32 data, string uri, bool recipientApproval, bytes32 parentClaimId, bool isValid)
func (_ClaimHolder *ClaimHolderSession) Claims(arg0 [32]byte) (struct {
	Topic             [32]byte
	Scheme            *big.Int
	Issuer            common.Address
	Signature         [32]byte
	Data              [32]byte
	Uri               string
	RecipientApproval bool
	ParentClaimId     [32]byte
	IsValid           bool
}, error) {
	return _ClaimHolder.Contract.Claims(&_ClaimHolder.CallOpts, arg0)
}

// Claims is a free data retrieval call binding the contract method 0xeff0f592.
//
// Solidity: function claims(bytes32 ) constant returns(bytes32 topic, uint256 scheme, address issuer, bytes32 signature, bytes32 data, string uri, bool recipientApproval, bytes32 parentClaimId, bool isValid)
func (_ClaimHolder *ClaimHolderCallerSession) Claims(arg0 [32]byte) (struct {
	Topic             [32]byte
	Scheme            *big.Int
	Issuer            common.Address
	Signature         [32]byte
	Data              [32]byte
	Uri               string
	RecipientApproval bool
	ParentClaimId     [32]byte
	IsValid           bool
}, error) {
	return _ClaimHolder.Contract.Claims(&_ClaimHolder.CallOpts, arg0)
}

// ClaimsByTopic is a free data retrieval call binding the contract method 0x0e4e4dfa.
//
// Solidity: function claimsByTopic(bytes32 , uint256 ) constant returns(bytes32)
func (_ClaimHolder *ClaimHolderCaller) ClaimsByTopic(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ClaimHolder.contract.Call(opts, out, "claimsByTopic", arg0, arg1)
	return *ret0, err
}

// ClaimsByTopic is a free data retrieval call binding the contract method 0x0e4e4dfa.
//
// Solidity: function claimsByTopic(bytes32 , uint256 ) constant returns(bytes32)
func (_ClaimHolder *ClaimHolderSession) ClaimsByTopic(arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _ClaimHolder.Contract.ClaimsByTopic(&_ClaimHolder.CallOpts, arg0, arg1)
}

// ClaimsByTopic is a free data retrieval call binding the contract method 0x0e4e4dfa.
//
// Solidity: function claimsByTopic(bytes32 , uint256 ) constant returns(bytes32)
func (_ClaimHolder *ClaimHolderCallerSession) ClaimsByTopic(arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _ClaimHolder.Contract.ClaimsByTopic(&_ClaimHolder.CallOpts, arg0, arg1)
}

// GetClaim is a free data retrieval call binding the contract method 0xc9100bcb.
//
// Solidity: function getClaim(bytes32 _claimId) constant returns(bytes32 topic, uint256 scheme, address issuer, bytes32 signature, bytes32 data, string uri)
func (_ClaimHolder *ClaimHolderCaller) GetClaim(opts *bind.CallOpts, _claimId [32]byte) (struct {
	Topic     [32]byte
	Scheme    *big.Int
	Issuer    common.Address
	Signature [32]byte
	Data      [32]byte
	Uri       string
}, error) {
	ret := new(struct {
		Topic     [32]byte
		Scheme    *big.Int
		Issuer    common.Address
		Signature [32]byte
		Data      [32]byte
		Uri       string
	})
	out := ret
	err := _ClaimHolder.contract.Call(opts, out, "getClaim", _claimId)
	return *ret, err
}

// GetClaim is a free data retrieval call binding the contract method 0xc9100bcb.
//
// Solidity: function getClaim(bytes32 _claimId) constant returns(bytes32 topic, uint256 scheme, address issuer, bytes32 signature, bytes32 data, string uri)
func (_ClaimHolder *ClaimHolderSession) GetClaim(_claimId [32]byte) (struct {
	Topic     [32]byte
	Scheme    *big.Int
	Issuer    common.Address
	Signature [32]byte
	Data      [32]byte
	Uri       string
}, error) {
	return _ClaimHolder.Contract.GetClaim(&_ClaimHolder.CallOpts, _claimId)
}

// GetClaim is a free data retrieval call binding the contract method 0xc9100bcb.
//
// Solidity: function getClaim(bytes32 _claimId) constant returns(bytes32 topic, uint256 scheme, address issuer, bytes32 signature, bytes32 data, string uri)
func (_ClaimHolder *ClaimHolderCallerSession) GetClaim(_claimId [32]byte) (struct {
	Topic     [32]byte
	Scheme    *big.Int
	Issuer    common.Address
	Signature [32]byte
	Data      [32]byte
	Uri       string
}, error) {
	return _ClaimHolder.Contract.GetClaim(&_ClaimHolder.CallOpts, _claimId)
}

// GetClaimIdsByTopic is a free data retrieval call binding the contract method 0xf9e9b618.
//
// Solidity: function getClaimIdsByTopic(bytes32 _topic) constant returns(bytes32[] claimIds)
func (_ClaimHolder *ClaimHolderCaller) GetClaimIdsByTopic(opts *bind.CallOpts, _topic [32]byte) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _ClaimHolder.contract.Call(opts, out, "getClaimIdsByTopic", _topic)
	return *ret0, err
}

// GetClaimIdsByTopic is a free data retrieval call binding the contract method 0xf9e9b618.
//
// Solidity: function getClaimIdsByTopic(bytes32 _topic) constant returns(bytes32[] claimIds)
func (_ClaimHolder *ClaimHolderSession) GetClaimIdsByTopic(_topic [32]byte) ([][32]byte, error) {
	return _ClaimHolder.Contract.GetClaimIdsByTopic(&_ClaimHolder.CallOpts, _topic)
}

// GetClaimIdsByTopic is a free data retrieval call binding the contract method 0xf9e9b618.
//
// Solidity: function getClaimIdsByTopic(bytes32 _topic) constant returns(bytes32[] claimIds)
func (_ClaimHolder *ClaimHolderCallerSession) GetClaimIdsByTopic(_topic [32]byte) ([][32]byte, error) {
	return _ClaimHolder.Contract.GetClaimIdsByTopic(&_ClaimHolder.CallOpts, _topic)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ClaimHolder *ClaimHolderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ClaimHolder.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ClaimHolder *ClaimHolderSession) Owner() (common.Address, error) {
	return _ClaimHolder.Contract.Owner(&_ClaimHolder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ClaimHolder *ClaimHolderCallerSession) Owner() (common.Address, error) {
	return _ClaimHolder.Contract.Owner(&_ClaimHolder.CallOpts)
}

// AddClaim is a paid mutator transaction binding the contract method 0xdf64dde2.
//
// Solidity: function addClaim(bytes32 _topic, uint256 _scheme, bytes32 _signature, bytes32 _data, string _uri) returns(bytes32 claimId)
func (_ClaimHolder *ClaimHolderTransactor) AddClaim(opts *bind.TransactOpts, _topic [32]byte, _scheme *big.Int, _signature [32]byte, _data [32]byte, _uri string) (*types.Transaction, error) {
	return _ClaimHolder.contract.Transact(opts, "addClaim", _topic, _scheme, _signature, _data, _uri)
}

// AddClaim is a paid mutator transaction binding the contract method 0xdf64dde2.
//
// Solidity: function addClaim(bytes32 _topic, uint256 _scheme, bytes32 _signature, bytes32 _data, string _uri) returns(bytes32 claimId)
func (_ClaimHolder *ClaimHolderSession) AddClaim(_topic [32]byte, _scheme *big.Int, _signature [32]byte, _data [32]byte, _uri string) (*types.Transaction, error) {
	return _ClaimHolder.Contract.AddClaim(&_ClaimHolder.TransactOpts, _topic, _scheme, _signature, _data, _uri)
}

// AddClaim is a paid mutator transaction binding the contract method 0xdf64dde2.
//
// Solidity: function addClaim(bytes32 _topic, uint256 _scheme, bytes32 _signature, bytes32 _data, string _uri) returns(bytes32 claimId)
func (_ClaimHolder *ClaimHolderTransactorSession) AddClaim(_topic [32]byte, _scheme *big.Int, _signature [32]byte, _data [32]byte, _uri string) (*types.Transaction, error) {
	return _ClaimHolder.Contract.AddClaim(&_ClaimHolder.TransactOpts, _topic, _scheme, _signature, _data, _uri)
}

// ChangeClaim is a paid mutator transaction binding the contract method 0xa85a9bd6.
//
// Solidity: function changeClaim(bytes32 _claimId, bytes32 _topic, uint256 _scheme, address _issuer, bytes32 _signature, bytes32 _data, string _uri) returns(bool success)
func (_ClaimHolder *ClaimHolderTransactor) ChangeClaim(opts *bind.TransactOpts, _claimId [32]byte, _topic [32]byte, _scheme *big.Int, _issuer common.Address, _signature [32]byte, _data [32]byte, _uri string) (*types.Transaction, error) {
	return _ClaimHolder.contract.Transact(opts, "changeClaim", _claimId, _topic, _scheme, _issuer, _signature, _data, _uri)
}

// ChangeClaim is a paid mutator transaction binding the contract method 0xa85a9bd6.
//
// Solidity: function changeClaim(bytes32 _claimId, bytes32 _topic, uint256 _scheme, address _issuer, bytes32 _signature, bytes32 _data, string _uri) returns(bool success)
func (_ClaimHolder *ClaimHolderSession) ChangeClaim(_claimId [32]byte, _topic [32]byte, _scheme *big.Int, _issuer common.Address, _signature [32]byte, _data [32]byte, _uri string) (*types.Transaction, error) {
	return _ClaimHolder.Contract.ChangeClaim(&_ClaimHolder.TransactOpts, _claimId, _topic, _scheme, _issuer, _signature, _data, _uri)
}

// ChangeClaim is a paid mutator transaction binding the contract method 0xa85a9bd6.
//
// Solidity: function changeClaim(bytes32 _claimId, bytes32 _topic, uint256 _scheme, address _issuer, bytes32 _signature, bytes32 _data, string _uri) returns(bool success)
func (_ClaimHolder *ClaimHolderTransactorSession) ChangeClaim(_claimId [32]byte, _topic [32]byte, _scheme *big.Int, _issuer common.Address, _signature [32]byte, _data [32]byte, _uri string) (*types.Transaction, error) {
	return _ClaimHolder.Contract.ChangeClaim(&_ClaimHolder.TransactOpts, _claimId, _topic, _scheme, _issuer, _signature, _data, _uri)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0x4eee424a.
//
// Solidity: function removeClaim(bytes32 _claimId) returns(bool success)
func (_ClaimHolder *ClaimHolderTransactor) RemoveClaim(opts *bind.TransactOpts, _claimId [32]byte) (*types.Transaction, error) {
	return _ClaimHolder.contract.Transact(opts, "removeClaim", _claimId)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0x4eee424a.
//
// Solidity: function removeClaim(bytes32 _claimId) returns(bool success)
func (_ClaimHolder *ClaimHolderSession) RemoveClaim(_claimId [32]byte) (*types.Transaction, error) {
	return _ClaimHolder.Contract.RemoveClaim(&_ClaimHolder.TransactOpts, _claimId)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0x4eee424a.
//
// Solidity: function removeClaim(bytes32 _claimId) returns(bool success)
func (_ClaimHolder *ClaimHolderTransactorSession) RemoveClaim(_claimId [32]byte) (*types.Transaction, error) {
	return _ClaimHolder.Contract.RemoveClaim(&_ClaimHolder.TransactOpts, _claimId)
}

// ToggleApprovaleClaim is a paid mutator transaction binding the contract method 0xb23bbc7a.
//
// Solidity: function toggleApprovaleClaim(bytes32 _claimId) returns(bool success)
func (_ClaimHolder *ClaimHolderTransactor) ToggleApprovaleClaim(opts *bind.TransactOpts, _claimId [32]byte) (*types.Transaction, error) {
	return _ClaimHolder.contract.Transact(opts, "toggleApprovaleClaim", _claimId)
}

// ToggleApprovaleClaim is a paid mutator transaction binding the contract method 0xb23bbc7a.
//
// Solidity: function toggleApprovaleClaim(bytes32 _claimId) returns(bool success)
func (_ClaimHolder *ClaimHolderSession) ToggleApprovaleClaim(_claimId [32]byte) (*types.Transaction, error) {
	return _ClaimHolder.Contract.ToggleApprovaleClaim(&_ClaimHolder.TransactOpts, _claimId)
}

// ToggleApprovaleClaim is a paid mutator transaction binding the contract method 0xb23bbc7a.
//
// Solidity: function toggleApprovaleClaim(bytes32 _claimId) returns(bool success)
func (_ClaimHolder *ClaimHolderTransactorSession) ToggleApprovaleClaim(_claimId [32]byte) (*types.Transaction, error) {
	return _ClaimHolder.Contract.ToggleApprovaleClaim(&_ClaimHolder.TransactOpts, _claimId)
}

// ClaimHolderClaimAddedIterator is returned from FilterClaimAdded and is used to iterate over the raw logs and unpacked data for ClaimAdded events raised by the ClaimHolder contract.
type ClaimHolderClaimAddedIterator struct {
	Event *ClaimHolderClaimAdded // Event containing the contract specifics and raw log

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
func (it *ClaimHolderClaimAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimHolderClaimAdded)
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
		it.Event = new(ClaimHolderClaimAdded)
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
func (it *ClaimHolderClaimAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimHolderClaimAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimHolderClaimAdded represents a ClaimAdded event raised by the ClaimHolder contract.
type ClaimHolderClaimAdded struct {
	ClaimId   [32]byte
	Topic     [32]byte
	Scheme    *big.Int
	Issuer    common.Address
	Signature [32]byte
	Data      [32]byte
	Uri       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimAdded is a free log retrieval operation binding the contract event 0x80095c7241bc123695ff70400336540c6e8c10a85a94249ed3eb16a51ea80304.
//
// Solidity: event ClaimAdded(bytes32 indexed claimId, bytes32 indexed topic, uint256 scheme, address indexed issuer, bytes32 signature, bytes32 data, string uri)
func (_ClaimHolder *ClaimHolderFilterer) FilterClaimAdded(opts *bind.FilterOpts, claimId [][32]byte, topic [][32]byte, issuer []common.Address) (*ClaimHolderClaimAddedIterator, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ClaimHolder.contract.FilterLogs(opts, "ClaimAdded", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &ClaimHolderClaimAddedIterator{contract: _ClaimHolder.contract, event: "ClaimAdded", logs: logs, sub: sub}, nil
}

// WatchClaimAdded is a free log subscription operation binding the contract event 0x80095c7241bc123695ff70400336540c6e8c10a85a94249ed3eb16a51ea80304.
//
// Solidity: event ClaimAdded(bytes32 indexed claimId, bytes32 indexed topic, uint256 scheme, address indexed issuer, bytes32 signature, bytes32 data, string uri)
func (_ClaimHolder *ClaimHolderFilterer) WatchClaimAdded(opts *bind.WatchOpts, sink chan<- *ClaimHolderClaimAdded, claimId [][32]byte, topic [][32]byte, issuer []common.Address) (event.Subscription, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ClaimHolder.contract.WatchLogs(opts, "ClaimAdded", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimHolderClaimAdded)
				if err := _ClaimHolder.contract.UnpackLog(event, "ClaimAdded", log); err != nil {
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

// ClaimHolderClaimApprovalToggledIterator is returned from FilterClaimApprovalToggled and is used to iterate over the raw logs and unpacked data for ClaimApprovalToggled events raised by the ClaimHolder contract.
type ClaimHolderClaimApprovalToggledIterator struct {
	Event *ClaimHolderClaimApprovalToggled // Event containing the contract specifics and raw log

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
func (it *ClaimHolderClaimApprovalToggledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimHolderClaimApprovalToggled)
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
		it.Event = new(ClaimHolderClaimApprovalToggled)
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
func (it *ClaimHolderClaimApprovalToggledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimHolderClaimApprovalToggledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimHolderClaimApprovalToggled represents a ClaimApprovalToggled event raised by the ClaimHolder contract.
type ClaimHolderClaimApprovalToggled struct {
	ClaimId   [32]byte
	Topic     [32]byte
	Scheme    *big.Int
	Issuer    common.Address
	Signature [32]byte
	Data      [32]byte
	Uri       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimApprovalToggled is a free log retrieval operation binding the contract event 0x9d524b960fb88ad472994586b295c3227def57acdcbca7804e91bea2b18aa0af.
//
// Solidity: event ClaimApprovalToggled(bytes32 indexed claimId, bytes32 indexed topic, uint256 scheme, address indexed issuer, bytes32 signature, bytes32 data, string uri)
func (_ClaimHolder *ClaimHolderFilterer) FilterClaimApprovalToggled(opts *bind.FilterOpts, claimId [][32]byte, topic [][32]byte, issuer []common.Address) (*ClaimHolderClaimApprovalToggledIterator, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ClaimHolder.contract.FilterLogs(opts, "ClaimApprovalToggled", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &ClaimHolderClaimApprovalToggledIterator{contract: _ClaimHolder.contract, event: "ClaimApprovalToggled", logs: logs, sub: sub}, nil
}

// WatchClaimApprovalToggled is a free log subscription operation binding the contract event 0x9d524b960fb88ad472994586b295c3227def57acdcbca7804e91bea2b18aa0af.
//
// Solidity: event ClaimApprovalToggled(bytes32 indexed claimId, bytes32 indexed topic, uint256 scheme, address indexed issuer, bytes32 signature, bytes32 data, string uri)
func (_ClaimHolder *ClaimHolderFilterer) WatchClaimApprovalToggled(opts *bind.WatchOpts, sink chan<- *ClaimHolderClaimApprovalToggled, claimId [][32]byte, topic [][32]byte, issuer []common.Address) (event.Subscription, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ClaimHolder.contract.WatchLogs(opts, "ClaimApprovalToggled", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimHolderClaimApprovalToggled)
				if err := _ClaimHolder.contract.UnpackLog(event, "ClaimApprovalToggled", log); err != nil {
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

// ClaimHolderClaimChangedIterator is returned from FilterClaimChanged and is used to iterate over the raw logs and unpacked data for ClaimChanged events raised by the ClaimHolder contract.
type ClaimHolderClaimChangedIterator struct {
	Event *ClaimHolderClaimChanged // Event containing the contract specifics and raw log

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
func (it *ClaimHolderClaimChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimHolderClaimChanged)
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
		it.Event = new(ClaimHolderClaimChanged)
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
func (it *ClaimHolderClaimChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimHolderClaimChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimHolderClaimChanged represents a ClaimChanged event raised by the ClaimHolder contract.
type ClaimHolderClaimChanged struct {
	ClaimId   [32]byte
	Topic     [32]byte
	Scheme    *big.Int
	Issuer    common.Address
	Signature [32]byte
	Data      [32]byte
	Uri       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimChanged is a free log retrieval operation binding the contract event 0x8e3848bbbb6124bac6e78d7d6f351516db9e27b626b756b5b8ada2034c4336ed.
//
// Solidity: event ClaimChanged(bytes32 indexed claimId, bytes32 indexed topic, uint256 scheme, address indexed issuer, bytes32 signature, bytes32 data, string uri)
func (_ClaimHolder *ClaimHolderFilterer) FilterClaimChanged(opts *bind.FilterOpts, claimId [][32]byte, topic [][32]byte, issuer []common.Address) (*ClaimHolderClaimChangedIterator, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ClaimHolder.contract.FilterLogs(opts, "ClaimChanged", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &ClaimHolderClaimChangedIterator{contract: _ClaimHolder.contract, event: "ClaimChanged", logs: logs, sub: sub}, nil
}

// WatchClaimChanged is a free log subscription operation binding the contract event 0x8e3848bbbb6124bac6e78d7d6f351516db9e27b626b756b5b8ada2034c4336ed.
//
// Solidity: event ClaimChanged(bytes32 indexed claimId, bytes32 indexed topic, uint256 scheme, address indexed issuer, bytes32 signature, bytes32 data, string uri)
func (_ClaimHolder *ClaimHolderFilterer) WatchClaimChanged(opts *bind.WatchOpts, sink chan<- *ClaimHolderClaimChanged, claimId [][32]byte, topic [][32]byte, issuer []common.Address) (event.Subscription, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ClaimHolder.contract.WatchLogs(opts, "ClaimChanged", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimHolderClaimChanged)
				if err := _ClaimHolder.contract.UnpackLog(event, "ClaimChanged", log); err != nil {
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

// ClaimHolderClaimRemovedIterator is returned from FilterClaimRemoved and is used to iterate over the raw logs and unpacked data for ClaimRemoved events raised by the ClaimHolder contract.
type ClaimHolderClaimRemovedIterator struct {
	Event *ClaimHolderClaimRemoved // Event containing the contract specifics and raw log

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
func (it *ClaimHolderClaimRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimHolderClaimRemoved)
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
		it.Event = new(ClaimHolderClaimRemoved)
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
func (it *ClaimHolderClaimRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimHolderClaimRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimHolderClaimRemoved represents a ClaimRemoved event raised by the ClaimHolder contract.
type ClaimHolderClaimRemoved struct {
	ClaimId   [32]byte
	Topic     [32]byte
	Scheme    *big.Int
	Issuer    common.Address
	Signature [32]byte
	Data      [32]byte
	Uri       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimRemoved is a free log retrieval operation binding the contract event 0xd8926d7f59d697f02d719bfaf05a3a564b8b17dcea1d82f69a54d6f13c5e3548.
//
// Solidity: event ClaimRemoved(bytes32 indexed claimId, bytes32 indexed topic, uint256 scheme, address indexed issuer, bytes32 signature, bytes32 data, string uri)
func (_ClaimHolder *ClaimHolderFilterer) FilterClaimRemoved(opts *bind.FilterOpts, claimId [][32]byte, topic [][32]byte, issuer []common.Address) (*ClaimHolderClaimRemovedIterator, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ClaimHolder.contract.FilterLogs(opts, "ClaimRemoved", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &ClaimHolderClaimRemovedIterator{contract: _ClaimHolder.contract, event: "ClaimRemoved", logs: logs, sub: sub}, nil
}

// WatchClaimRemoved is a free log subscription operation binding the contract event 0xd8926d7f59d697f02d719bfaf05a3a564b8b17dcea1d82f69a54d6f13c5e3548.
//
// Solidity: event ClaimRemoved(bytes32 indexed claimId, bytes32 indexed topic, uint256 scheme, address indexed issuer, bytes32 signature, bytes32 data, string uri)
func (_ClaimHolder *ClaimHolderFilterer) WatchClaimRemoved(opts *bind.WatchOpts, sink chan<- *ClaimHolderClaimRemoved, claimId [][32]byte, topic [][32]byte, issuer []common.Address) (event.Subscription, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ClaimHolder.contract.WatchLogs(opts, "ClaimRemoved", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimHolderClaimRemoved)
				if err := _ClaimHolder.contract.UnpackLog(event, "ClaimRemoved", log); err != nil {
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

// ERC735ABI is the input ABI used to generate the binding from.
const ERC735ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_claimId\",\"type\":\"bytes32\"}],\"name\":\"removeClaim\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_claimId\",\"type\":\"bytes32\"},{\"name\":\"_topic\",\"type\":\"bytes32\"},{\"name\":\"_scheme\",\"type\":\"uint256\"},{\"name\":\"_issuer\",\"type\":\"address\"},{\"name\":\"_signature\",\"type\":\"bytes32\"},{\"name\":\"_data\",\"type\":\"bytes32\"},{\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"changeClaim\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_claimId\",\"type\":\"bytes32\"}],\"name\":\"toggleApprovaleClaim\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_claimId\",\"type\":\"bytes32\"}],\"name\":\"getClaim\",\"outputs\":[{\"name\":\"topic\",\"type\":\"bytes32\"},{\"name\":\"scheme\",\"type\":\"uint256\"},{\"name\":\"issuer\",\"type\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes32\"},{\"name\":\"data\",\"type\":\"bytes32\"},{\"name\":\"uri\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_topic\",\"type\":\"bytes32\"},{\"name\":\"_scheme\",\"type\":\"uint256\"},{\"name\":\"_signature\",\"type\":\"bytes32\"},{\"name\":\"_data\",\"type\":\"bytes32\"},{\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"addClaim\",\"outputs\":[{\"name\":\"claimRequestId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_topic\",\"type\":\"bytes32\"}],\"name\":\"getClaimIdsByTopic\",\"outputs\":[{\"name\":\"claimIds\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"claimId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"topic\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"scheme\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"signature\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"ClaimApprovalToggled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"claimId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"topic\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"scheme\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"signature\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"ClaimAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"claimId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"topic\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"scheme\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"signature\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"ClaimRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"claimId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"topic\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"scheme\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"signature\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"ClaimChanged\",\"type\":\"event\"}]"

// ERC735Bin is the compiled bytecode used for deploying new contracts.
const ERC735Bin = `0x`

// DeployERC735 deploys a new Ethereum contract, binding an instance of ERC735 to it.
func DeployERC735(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC735, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC735ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC735Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC735{ERC735Caller: ERC735Caller{contract: contract}, ERC735Transactor: ERC735Transactor{contract: contract}, ERC735Filterer: ERC735Filterer{contract: contract}}, nil
}

// ERC735 is an auto generated Go binding around an Ethereum contract.
type ERC735 struct {
	ERC735Caller     // Read-only binding to the contract
	ERC735Transactor // Write-only binding to the contract
	ERC735Filterer   // Log filterer for contract events
}

// ERC735Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC735Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC735Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC735Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC735Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC735Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC735Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC735Session struct {
	Contract     *ERC735           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC735CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC735CallerSession struct {
	Contract *ERC735Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC735TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC735TransactorSession struct {
	Contract     *ERC735Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC735Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC735Raw struct {
	Contract *ERC735 // Generic contract binding to access the raw methods on
}

// ERC735CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC735CallerRaw struct {
	Contract *ERC735Caller // Generic read-only contract binding to access the raw methods on
}

// ERC735TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC735TransactorRaw struct {
	Contract *ERC735Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC735 creates a new instance of ERC735, bound to a specific deployed contract.
func NewERC735(address common.Address, backend bind.ContractBackend) (*ERC735, error) {
	contract, err := bindERC735(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC735{ERC735Caller: ERC735Caller{contract: contract}, ERC735Transactor: ERC735Transactor{contract: contract}, ERC735Filterer: ERC735Filterer{contract: contract}}, nil
}

// NewERC735Caller creates a new read-only instance of ERC735, bound to a specific deployed contract.
func NewERC735Caller(address common.Address, caller bind.ContractCaller) (*ERC735Caller, error) {
	contract, err := bindERC735(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC735Caller{contract: contract}, nil
}

// NewERC735Transactor creates a new write-only instance of ERC735, bound to a specific deployed contract.
func NewERC735Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC735Transactor, error) {
	contract, err := bindERC735(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC735Transactor{contract: contract}, nil
}

// NewERC735Filterer creates a new log filterer instance of ERC735, bound to a specific deployed contract.
func NewERC735Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC735Filterer, error) {
	contract, err := bindERC735(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC735Filterer{contract: contract}, nil
}

// bindERC735 binds a generic wrapper to an already deployed contract.
func bindERC735(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC735ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC735 *ERC735Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC735.Contract.ERC735Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC735 *ERC735Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC735.Contract.ERC735Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC735 *ERC735Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC735.Contract.ERC735Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC735 *ERC735CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC735.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC735 *ERC735TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC735.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC735 *ERC735TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC735.Contract.contract.Transact(opts, method, params...)
}

// GetClaim is a free data retrieval call binding the contract method 0xc9100bcb.
//
// Solidity: function getClaim(bytes32 _claimId) constant returns(bytes32 topic, uint256 scheme, address issuer, bytes32 signature, bytes32 data, string uri)
func (_ERC735 *ERC735Caller) GetClaim(opts *bind.CallOpts, _claimId [32]byte) (struct {
	Topic     [32]byte
	Scheme    *big.Int
	Issuer    common.Address
	Signature [32]byte
	Data      [32]byte
	Uri       string
}, error) {
	ret := new(struct {
		Topic     [32]byte
		Scheme    *big.Int
		Issuer    common.Address
		Signature [32]byte
		Data      [32]byte
		Uri       string
	})
	out := ret
	err := _ERC735.contract.Call(opts, out, "getClaim", _claimId)
	return *ret, err
}

// GetClaim is a free data retrieval call binding the contract method 0xc9100bcb.
//
// Solidity: function getClaim(bytes32 _claimId) constant returns(bytes32 topic, uint256 scheme, address issuer, bytes32 signature, bytes32 data, string uri)
func (_ERC735 *ERC735Session) GetClaim(_claimId [32]byte) (struct {
	Topic     [32]byte
	Scheme    *big.Int
	Issuer    common.Address
	Signature [32]byte
	Data      [32]byte
	Uri       string
}, error) {
	return _ERC735.Contract.GetClaim(&_ERC735.CallOpts, _claimId)
}

// GetClaim is a free data retrieval call binding the contract method 0xc9100bcb.
//
// Solidity: function getClaim(bytes32 _claimId) constant returns(bytes32 topic, uint256 scheme, address issuer, bytes32 signature, bytes32 data, string uri)
func (_ERC735 *ERC735CallerSession) GetClaim(_claimId [32]byte) (struct {
	Topic     [32]byte
	Scheme    *big.Int
	Issuer    common.Address
	Signature [32]byte
	Data      [32]byte
	Uri       string
}, error) {
	return _ERC735.Contract.GetClaim(&_ERC735.CallOpts, _claimId)
}

// GetClaimIdsByTopic is a free data retrieval call binding the contract method 0xf9e9b618.
//
// Solidity: function getClaimIdsByTopic(bytes32 _topic) constant returns(bytes32[] claimIds)
func (_ERC735 *ERC735Caller) GetClaimIdsByTopic(opts *bind.CallOpts, _topic [32]byte) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _ERC735.contract.Call(opts, out, "getClaimIdsByTopic", _topic)
	return *ret0, err
}

// GetClaimIdsByTopic is a free data retrieval call binding the contract method 0xf9e9b618.
//
// Solidity: function getClaimIdsByTopic(bytes32 _topic) constant returns(bytes32[] claimIds)
func (_ERC735 *ERC735Session) GetClaimIdsByTopic(_topic [32]byte) ([][32]byte, error) {
	return _ERC735.Contract.GetClaimIdsByTopic(&_ERC735.CallOpts, _topic)
}

// GetClaimIdsByTopic is a free data retrieval call binding the contract method 0xf9e9b618.
//
// Solidity: function getClaimIdsByTopic(bytes32 _topic) constant returns(bytes32[] claimIds)
func (_ERC735 *ERC735CallerSession) GetClaimIdsByTopic(_topic [32]byte) ([][32]byte, error) {
	return _ERC735.Contract.GetClaimIdsByTopic(&_ERC735.CallOpts, _topic)
}

// AddClaim is a paid mutator transaction binding the contract method 0xdf64dde2.
//
// Solidity: function addClaim(bytes32 _topic, uint256 _scheme, bytes32 _signature, bytes32 _data, string _uri) returns(bytes32 claimRequestId)
func (_ERC735 *ERC735Transactor) AddClaim(opts *bind.TransactOpts, _topic [32]byte, _scheme *big.Int, _signature [32]byte, _data [32]byte, _uri string) (*types.Transaction, error) {
	return _ERC735.contract.Transact(opts, "addClaim", _topic, _scheme, _signature, _data, _uri)
}

// AddClaim is a paid mutator transaction binding the contract method 0xdf64dde2.
//
// Solidity: function addClaim(bytes32 _topic, uint256 _scheme, bytes32 _signature, bytes32 _data, string _uri) returns(bytes32 claimRequestId)
func (_ERC735 *ERC735Session) AddClaim(_topic [32]byte, _scheme *big.Int, _signature [32]byte, _data [32]byte, _uri string) (*types.Transaction, error) {
	return _ERC735.Contract.AddClaim(&_ERC735.TransactOpts, _topic, _scheme, _signature, _data, _uri)
}

// AddClaim is a paid mutator transaction binding the contract method 0xdf64dde2.
//
// Solidity: function addClaim(bytes32 _topic, uint256 _scheme, bytes32 _signature, bytes32 _data, string _uri) returns(bytes32 claimRequestId)
func (_ERC735 *ERC735TransactorSession) AddClaim(_topic [32]byte, _scheme *big.Int, _signature [32]byte, _data [32]byte, _uri string) (*types.Transaction, error) {
	return _ERC735.Contract.AddClaim(&_ERC735.TransactOpts, _topic, _scheme, _signature, _data, _uri)
}

// ChangeClaim is a paid mutator transaction binding the contract method 0xa85a9bd6.
//
// Solidity: function changeClaim(bytes32 _claimId, bytes32 _topic, uint256 _scheme, address _issuer, bytes32 _signature, bytes32 _data, string _uri) returns(bool success)
func (_ERC735 *ERC735Transactor) ChangeClaim(opts *bind.TransactOpts, _claimId [32]byte, _topic [32]byte, _scheme *big.Int, _issuer common.Address, _signature [32]byte, _data [32]byte, _uri string) (*types.Transaction, error) {
	return _ERC735.contract.Transact(opts, "changeClaim", _claimId, _topic, _scheme, _issuer, _signature, _data, _uri)
}

// ChangeClaim is a paid mutator transaction binding the contract method 0xa85a9bd6.
//
// Solidity: function changeClaim(bytes32 _claimId, bytes32 _topic, uint256 _scheme, address _issuer, bytes32 _signature, bytes32 _data, string _uri) returns(bool success)
func (_ERC735 *ERC735Session) ChangeClaim(_claimId [32]byte, _topic [32]byte, _scheme *big.Int, _issuer common.Address, _signature [32]byte, _data [32]byte, _uri string) (*types.Transaction, error) {
	return _ERC735.Contract.ChangeClaim(&_ERC735.TransactOpts, _claimId, _topic, _scheme, _issuer, _signature, _data, _uri)
}

// ChangeClaim is a paid mutator transaction binding the contract method 0xa85a9bd6.
//
// Solidity: function changeClaim(bytes32 _claimId, bytes32 _topic, uint256 _scheme, address _issuer, bytes32 _signature, bytes32 _data, string _uri) returns(bool success)
func (_ERC735 *ERC735TransactorSession) ChangeClaim(_claimId [32]byte, _topic [32]byte, _scheme *big.Int, _issuer common.Address, _signature [32]byte, _data [32]byte, _uri string) (*types.Transaction, error) {
	return _ERC735.Contract.ChangeClaim(&_ERC735.TransactOpts, _claimId, _topic, _scheme, _issuer, _signature, _data, _uri)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0x4eee424a.
//
// Solidity: function removeClaim(bytes32 _claimId) returns(bool success)
func (_ERC735 *ERC735Transactor) RemoveClaim(opts *bind.TransactOpts, _claimId [32]byte) (*types.Transaction, error) {
	return _ERC735.contract.Transact(opts, "removeClaim", _claimId)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0x4eee424a.
//
// Solidity: function removeClaim(bytes32 _claimId) returns(bool success)
func (_ERC735 *ERC735Session) RemoveClaim(_claimId [32]byte) (*types.Transaction, error) {
	return _ERC735.Contract.RemoveClaim(&_ERC735.TransactOpts, _claimId)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0x4eee424a.
//
// Solidity: function removeClaim(bytes32 _claimId) returns(bool success)
func (_ERC735 *ERC735TransactorSession) RemoveClaim(_claimId [32]byte) (*types.Transaction, error) {
	return _ERC735.Contract.RemoveClaim(&_ERC735.TransactOpts, _claimId)
}

// ToggleApprovaleClaim is a paid mutator transaction binding the contract method 0xb23bbc7a.
//
// Solidity: function toggleApprovaleClaim(bytes32 _claimId) returns(bool success)
func (_ERC735 *ERC735Transactor) ToggleApprovaleClaim(opts *bind.TransactOpts, _claimId [32]byte) (*types.Transaction, error) {
	return _ERC735.contract.Transact(opts, "toggleApprovaleClaim", _claimId)
}

// ToggleApprovaleClaim is a paid mutator transaction binding the contract method 0xb23bbc7a.
//
// Solidity: function toggleApprovaleClaim(bytes32 _claimId) returns(bool success)
func (_ERC735 *ERC735Session) ToggleApprovaleClaim(_claimId [32]byte) (*types.Transaction, error) {
	return _ERC735.Contract.ToggleApprovaleClaim(&_ERC735.TransactOpts, _claimId)
}

// ToggleApprovaleClaim is a paid mutator transaction binding the contract method 0xb23bbc7a.
//
// Solidity: function toggleApprovaleClaim(bytes32 _claimId) returns(bool success)
func (_ERC735 *ERC735TransactorSession) ToggleApprovaleClaim(_claimId [32]byte) (*types.Transaction, error) {
	return _ERC735.Contract.ToggleApprovaleClaim(&_ERC735.TransactOpts, _claimId)
}

// ERC735ClaimAddedIterator is returned from FilterClaimAdded and is used to iterate over the raw logs and unpacked data for ClaimAdded events raised by the ERC735 contract.
type ERC735ClaimAddedIterator struct {
	Event *ERC735ClaimAdded // Event containing the contract specifics and raw log

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
func (it *ERC735ClaimAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC735ClaimAdded)
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
		it.Event = new(ERC735ClaimAdded)
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
func (it *ERC735ClaimAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC735ClaimAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC735ClaimAdded represents a ClaimAdded event raised by the ERC735 contract.
type ERC735ClaimAdded struct {
	ClaimId   [32]byte
	Topic     [32]byte
	Scheme    *big.Int
	Issuer    common.Address
	Signature [32]byte
	Data      [32]byte
	Uri       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimAdded is a free log retrieval operation binding the contract event 0x80095c7241bc123695ff70400336540c6e8c10a85a94249ed3eb16a51ea80304.
//
// Solidity: event ClaimAdded(bytes32 indexed claimId, bytes32 indexed topic, uint256 scheme, address indexed issuer, bytes32 signature, bytes32 data, string uri)
func (_ERC735 *ERC735Filterer) FilterClaimAdded(opts *bind.FilterOpts, claimId [][32]byte, topic [][32]byte, issuer []common.Address) (*ERC735ClaimAddedIterator, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ERC735.contract.FilterLogs(opts, "ClaimAdded", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &ERC735ClaimAddedIterator{contract: _ERC735.contract, event: "ClaimAdded", logs: logs, sub: sub}, nil
}

// WatchClaimAdded is a free log subscription operation binding the contract event 0x80095c7241bc123695ff70400336540c6e8c10a85a94249ed3eb16a51ea80304.
//
// Solidity: event ClaimAdded(bytes32 indexed claimId, bytes32 indexed topic, uint256 scheme, address indexed issuer, bytes32 signature, bytes32 data, string uri)
func (_ERC735 *ERC735Filterer) WatchClaimAdded(opts *bind.WatchOpts, sink chan<- *ERC735ClaimAdded, claimId [][32]byte, topic [][32]byte, issuer []common.Address) (event.Subscription, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ERC735.contract.WatchLogs(opts, "ClaimAdded", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC735ClaimAdded)
				if err := _ERC735.contract.UnpackLog(event, "ClaimAdded", log); err != nil {
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

// ERC735ClaimApprovalToggledIterator is returned from FilterClaimApprovalToggled and is used to iterate over the raw logs and unpacked data for ClaimApprovalToggled events raised by the ERC735 contract.
type ERC735ClaimApprovalToggledIterator struct {
	Event *ERC735ClaimApprovalToggled // Event containing the contract specifics and raw log

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
func (it *ERC735ClaimApprovalToggledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC735ClaimApprovalToggled)
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
		it.Event = new(ERC735ClaimApprovalToggled)
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
func (it *ERC735ClaimApprovalToggledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC735ClaimApprovalToggledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC735ClaimApprovalToggled represents a ClaimApprovalToggled event raised by the ERC735 contract.
type ERC735ClaimApprovalToggled struct {
	ClaimId   [32]byte
	Topic     [32]byte
	Scheme    *big.Int
	Issuer    common.Address
	Signature [32]byte
	Data      [32]byte
	Uri       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimApprovalToggled is a free log retrieval operation binding the contract event 0x9d524b960fb88ad472994586b295c3227def57acdcbca7804e91bea2b18aa0af.
//
// Solidity: event ClaimApprovalToggled(bytes32 indexed claimId, bytes32 indexed topic, uint256 scheme, address indexed issuer, bytes32 signature, bytes32 data, string uri)
func (_ERC735 *ERC735Filterer) FilterClaimApprovalToggled(opts *bind.FilterOpts, claimId [][32]byte, topic [][32]byte, issuer []common.Address) (*ERC735ClaimApprovalToggledIterator, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ERC735.contract.FilterLogs(opts, "ClaimApprovalToggled", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &ERC735ClaimApprovalToggledIterator{contract: _ERC735.contract, event: "ClaimApprovalToggled", logs: logs, sub: sub}, nil
}

// WatchClaimApprovalToggled is a free log subscription operation binding the contract event 0x9d524b960fb88ad472994586b295c3227def57acdcbca7804e91bea2b18aa0af.
//
// Solidity: event ClaimApprovalToggled(bytes32 indexed claimId, bytes32 indexed topic, uint256 scheme, address indexed issuer, bytes32 signature, bytes32 data, string uri)
func (_ERC735 *ERC735Filterer) WatchClaimApprovalToggled(opts *bind.WatchOpts, sink chan<- *ERC735ClaimApprovalToggled, claimId [][32]byte, topic [][32]byte, issuer []common.Address) (event.Subscription, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ERC735.contract.WatchLogs(opts, "ClaimApprovalToggled", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC735ClaimApprovalToggled)
				if err := _ERC735.contract.UnpackLog(event, "ClaimApprovalToggled", log); err != nil {
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

// ERC735ClaimChangedIterator is returned from FilterClaimChanged and is used to iterate over the raw logs and unpacked data for ClaimChanged events raised by the ERC735 contract.
type ERC735ClaimChangedIterator struct {
	Event *ERC735ClaimChanged // Event containing the contract specifics and raw log

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
func (it *ERC735ClaimChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC735ClaimChanged)
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
		it.Event = new(ERC735ClaimChanged)
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
func (it *ERC735ClaimChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC735ClaimChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC735ClaimChanged represents a ClaimChanged event raised by the ERC735 contract.
type ERC735ClaimChanged struct {
	ClaimId   [32]byte
	Topic     [32]byte
	Scheme    *big.Int
	Issuer    common.Address
	Signature [32]byte
	Data      [32]byte
	Uri       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimChanged is a free log retrieval operation binding the contract event 0x8e3848bbbb6124bac6e78d7d6f351516db9e27b626b756b5b8ada2034c4336ed.
//
// Solidity: event ClaimChanged(bytes32 indexed claimId, bytes32 indexed topic, uint256 scheme, address indexed issuer, bytes32 signature, bytes32 data, string uri)
func (_ERC735 *ERC735Filterer) FilterClaimChanged(opts *bind.FilterOpts, claimId [][32]byte, topic [][32]byte, issuer []common.Address) (*ERC735ClaimChangedIterator, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ERC735.contract.FilterLogs(opts, "ClaimChanged", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &ERC735ClaimChangedIterator{contract: _ERC735.contract, event: "ClaimChanged", logs: logs, sub: sub}, nil
}

// WatchClaimChanged is a free log subscription operation binding the contract event 0x8e3848bbbb6124bac6e78d7d6f351516db9e27b626b756b5b8ada2034c4336ed.
//
// Solidity: event ClaimChanged(bytes32 indexed claimId, bytes32 indexed topic, uint256 scheme, address indexed issuer, bytes32 signature, bytes32 data, string uri)
func (_ERC735 *ERC735Filterer) WatchClaimChanged(opts *bind.WatchOpts, sink chan<- *ERC735ClaimChanged, claimId [][32]byte, topic [][32]byte, issuer []common.Address) (event.Subscription, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ERC735.contract.WatchLogs(opts, "ClaimChanged", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC735ClaimChanged)
				if err := _ERC735.contract.UnpackLog(event, "ClaimChanged", log); err != nil {
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

// ERC735ClaimRemovedIterator is returned from FilterClaimRemoved and is used to iterate over the raw logs and unpacked data for ClaimRemoved events raised by the ERC735 contract.
type ERC735ClaimRemovedIterator struct {
	Event *ERC735ClaimRemoved // Event containing the contract specifics and raw log

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
func (it *ERC735ClaimRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC735ClaimRemoved)
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
		it.Event = new(ERC735ClaimRemoved)
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
func (it *ERC735ClaimRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC735ClaimRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC735ClaimRemoved represents a ClaimRemoved event raised by the ERC735 contract.
type ERC735ClaimRemoved struct {
	ClaimId   [32]byte
	Topic     [32]byte
	Scheme    *big.Int
	Issuer    common.Address
	Signature [32]byte
	Data      [32]byte
	Uri       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimRemoved is a free log retrieval operation binding the contract event 0xd8926d7f59d697f02d719bfaf05a3a564b8b17dcea1d82f69a54d6f13c5e3548.
//
// Solidity: event ClaimRemoved(bytes32 indexed claimId, bytes32 indexed topic, uint256 scheme, address indexed issuer, bytes32 signature, bytes32 data, string uri)
func (_ERC735 *ERC735Filterer) FilterClaimRemoved(opts *bind.FilterOpts, claimId [][32]byte, topic [][32]byte, issuer []common.Address) (*ERC735ClaimRemovedIterator, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ERC735.contract.FilterLogs(opts, "ClaimRemoved", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &ERC735ClaimRemovedIterator{contract: _ERC735.contract, event: "ClaimRemoved", logs: logs, sub: sub}, nil
}

// WatchClaimRemoved is a free log subscription operation binding the contract event 0xd8926d7f59d697f02d719bfaf05a3a564b8b17dcea1d82f69a54d6f13c5e3548.
//
// Solidity: event ClaimRemoved(bytes32 indexed claimId, bytes32 indexed topic, uint256 scheme, address indexed issuer, bytes32 signature, bytes32 data, string uri)
func (_ERC735 *ERC735Filterer) WatchClaimRemoved(opts *bind.WatchOpts, sink chan<- *ERC735ClaimRemoved, claimId [][32]byte, topic [][32]byte, issuer []common.Address) (event.Subscription, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ERC735.contract.WatchLogs(opts, "ClaimRemoved", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC735ClaimRemoved)
				if err := _ERC735.contract.UnpackLog(event, "ClaimRemoved", log); err != nil {
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

// ERC780ABI is the input ABI used to generate the binding from.
const ERC780ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"setSelfClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"subject\",\"type\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"setClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"issuer\",\"type\":\"address\"},{\"name\":\"subject\",\"type\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"removeClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"issuer\",\"type\":\"address\"},{\"name\":\"subject\",\"type\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getClaim\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"subject\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"ClaimSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"subject\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"ClaimRemoved\",\"type\":\"event\"}]"

// ERC780Bin is the compiled bytecode used for deploying new contracts.
const ERC780Bin = `0x`

// DeployERC780 deploys a new Ethereum contract, binding an instance of ERC780 to it.
func DeployERC780(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC780, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC780ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC780Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC780{ERC780Caller: ERC780Caller{contract: contract}, ERC780Transactor: ERC780Transactor{contract: contract}, ERC780Filterer: ERC780Filterer{contract: contract}}, nil
}

// ERC780 is an auto generated Go binding around an Ethereum contract.
type ERC780 struct {
	ERC780Caller     // Read-only binding to the contract
	ERC780Transactor // Write-only binding to the contract
	ERC780Filterer   // Log filterer for contract events
}

// ERC780Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC780Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC780Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC780Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC780Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC780Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC780Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC780Session struct {
	Contract     *ERC780           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC780CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC780CallerSession struct {
	Contract *ERC780Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC780TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC780TransactorSession struct {
	Contract     *ERC780Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC780Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC780Raw struct {
	Contract *ERC780 // Generic contract binding to access the raw methods on
}

// ERC780CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC780CallerRaw struct {
	Contract *ERC780Caller // Generic read-only contract binding to access the raw methods on
}

// ERC780TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC780TransactorRaw struct {
	Contract *ERC780Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC780 creates a new instance of ERC780, bound to a specific deployed contract.
func NewERC780(address common.Address, backend bind.ContractBackend) (*ERC780, error) {
	contract, err := bindERC780(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC780{ERC780Caller: ERC780Caller{contract: contract}, ERC780Transactor: ERC780Transactor{contract: contract}, ERC780Filterer: ERC780Filterer{contract: contract}}, nil
}

// NewERC780Caller creates a new read-only instance of ERC780, bound to a specific deployed contract.
func NewERC780Caller(address common.Address, caller bind.ContractCaller) (*ERC780Caller, error) {
	contract, err := bindERC780(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC780Caller{contract: contract}, nil
}

// NewERC780Transactor creates a new write-only instance of ERC780, bound to a specific deployed contract.
func NewERC780Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC780Transactor, error) {
	contract, err := bindERC780(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC780Transactor{contract: contract}, nil
}

// NewERC780Filterer creates a new log filterer instance of ERC780, bound to a specific deployed contract.
func NewERC780Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC780Filterer, error) {
	contract, err := bindERC780(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC780Filterer{contract: contract}, nil
}

// bindERC780 binds a generic wrapper to an already deployed contract.
func bindERC780(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC780ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC780 *ERC780Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC780.Contract.ERC780Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC780 *ERC780Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC780.Contract.ERC780Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC780 *ERC780Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC780.Contract.ERC780Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC780 *ERC780CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC780.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC780 *ERC780TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC780.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC780 *ERC780TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC780.Contract.contract.Transact(opts, method, params...)
}

// GetClaim is a free data retrieval call binding the contract method 0xe1661eff.
//
// Solidity: function getClaim(address issuer, address subject, bytes32 key) constant returns(bytes32)
func (_ERC780 *ERC780Caller) GetClaim(opts *bind.CallOpts, issuer common.Address, subject common.Address, key [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ERC780.contract.Call(opts, out, "getClaim", issuer, subject, key)
	return *ret0, err
}

// GetClaim is a free data retrieval call binding the contract method 0xe1661eff.
//
// Solidity: function getClaim(address issuer, address subject, bytes32 key) constant returns(bytes32)
func (_ERC780 *ERC780Session) GetClaim(issuer common.Address, subject common.Address, key [32]byte) ([32]byte, error) {
	return _ERC780.Contract.GetClaim(&_ERC780.CallOpts, issuer, subject, key)
}

// GetClaim is a free data retrieval call binding the contract method 0xe1661eff.
//
// Solidity: function getClaim(address issuer, address subject, bytes32 key) constant returns(bytes32)
func (_ERC780 *ERC780CallerSession) GetClaim(issuer common.Address, subject common.Address, key [32]byte) ([32]byte, error) {
	return _ERC780.Contract.GetClaim(&_ERC780.CallOpts, issuer, subject, key)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0xc7508ec7.
//
// Solidity: function removeClaim(address issuer, address subject, bytes32 key) returns()
func (_ERC780 *ERC780Transactor) RemoveClaim(opts *bind.TransactOpts, issuer common.Address, subject common.Address, key [32]byte) (*types.Transaction, error) {
	return _ERC780.contract.Transact(opts, "removeClaim", issuer, subject, key)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0xc7508ec7.
//
// Solidity: function removeClaim(address issuer, address subject, bytes32 key) returns()
func (_ERC780 *ERC780Session) RemoveClaim(issuer common.Address, subject common.Address, key [32]byte) (*types.Transaction, error) {
	return _ERC780.Contract.RemoveClaim(&_ERC780.TransactOpts, issuer, subject, key)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0xc7508ec7.
//
// Solidity: function removeClaim(address issuer, address subject, bytes32 key) returns()
func (_ERC780 *ERC780TransactorSession) RemoveClaim(issuer common.Address, subject common.Address, key [32]byte) (*types.Transaction, error) {
	return _ERC780.Contract.RemoveClaim(&_ERC780.TransactOpts, issuer, subject, key)
}

// SetClaim is a paid mutator transaction binding the contract method 0x9918925d.
//
// Solidity: function setClaim(address subject, bytes32 key, bytes32 value) returns()
func (_ERC780 *ERC780Transactor) SetClaim(opts *bind.TransactOpts, subject common.Address, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _ERC780.contract.Transact(opts, "setClaim", subject, key, value)
}

// SetClaim is a paid mutator transaction binding the contract method 0x9918925d.
//
// Solidity: function setClaim(address subject, bytes32 key, bytes32 value) returns()
func (_ERC780 *ERC780Session) SetClaim(subject common.Address, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _ERC780.Contract.SetClaim(&_ERC780.TransactOpts, subject, key, value)
}

// SetClaim is a paid mutator transaction binding the contract method 0x9918925d.
//
// Solidity: function setClaim(address subject, bytes32 key, bytes32 value) returns()
func (_ERC780 *ERC780TransactorSession) SetClaim(subject common.Address, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _ERC780.Contract.SetClaim(&_ERC780.TransactOpts, subject, key, value)
}

// SetSelfClaim is a paid mutator transaction binding the contract method 0x9155b01a.
//
// Solidity: function setSelfClaim(bytes32 key, bytes32 value) returns()
func (_ERC780 *ERC780Transactor) SetSelfClaim(opts *bind.TransactOpts, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _ERC780.contract.Transact(opts, "setSelfClaim", key, value)
}

// SetSelfClaim is a paid mutator transaction binding the contract method 0x9155b01a.
//
// Solidity: function setSelfClaim(bytes32 key, bytes32 value) returns()
func (_ERC780 *ERC780Session) SetSelfClaim(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _ERC780.Contract.SetSelfClaim(&_ERC780.TransactOpts, key, value)
}

// SetSelfClaim is a paid mutator transaction binding the contract method 0x9155b01a.
//
// Solidity: function setSelfClaim(bytes32 key, bytes32 value) returns()
func (_ERC780 *ERC780TransactorSession) SetSelfClaim(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _ERC780.Contract.SetSelfClaim(&_ERC780.TransactOpts, key, value)
}

// ERC780ClaimRemovedIterator is returned from FilterClaimRemoved and is used to iterate over the raw logs and unpacked data for ClaimRemoved events raised by the ERC780 contract.
type ERC780ClaimRemovedIterator struct {
	Event *ERC780ClaimRemoved // Event containing the contract specifics and raw log

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
func (it *ERC780ClaimRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC780ClaimRemoved)
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
		it.Event = new(ERC780ClaimRemoved)
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
func (it *ERC780ClaimRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC780ClaimRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC780ClaimRemoved represents a ClaimRemoved event raised by the ERC780 contract.
type ERC780ClaimRemoved struct {
	Issuer  common.Address
	Subject common.Address
	Key     [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimRemoved is a free log retrieval operation binding the contract event 0x895b934151a219f5c7dd6868b1397f44e9435ef5969693eed4bdc1c74dff2df1.
//
// Solidity: event ClaimRemoved(address indexed issuer, address indexed subject, bytes32 indexed key)
func (_ERC780 *ERC780Filterer) FilterClaimRemoved(opts *bind.FilterOpts, issuer []common.Address, subject []common.Address, key [][32]byte) (*ERC780ClaimRemovedIterator, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _ERC780.contract.FilterLogs(opts, "ClaimRemoved", issuerRule, subjectRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &ERC780ClaimRemovedIterator{contract: _ERC780.contract, event: "ClaimRemoved", logs: logs, sub: sub}, nil
}

// WatchClaimRemoved is a free log subscription operation binding the contract event 0x895b934151a219f5c7dd6868b1397f44e9435ef5969693eed4bdc1c74dff2df1.
//
// Solidity: event ClaimRemoved(address indexed issuer, address indexed subject, bytes32 indexed key)
func (_ERC780 *ERC780Filterer) WatchClaimRemoved(opts *bind.WatchOpts, sink chan<- *ERC780ClaimRemoved, issuer []common.Address, subject []common.Address, key [][32]byte) (event.Subscription, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _ERC780.contract.WatchLogs(opts, "ClaimRemoved", issuerRule, subjectRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC780ClaimRemoved)
				if err := _ERC780.contract.UnpackLog(event, "ClaimRemoved", log); err != nil {
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

// ERC780ClaimSetIterator is returned from FilterClaimSet and is used to iterate over the raw logs and unpacked data for ClaimSet events raised by the ERC780 contract.
type ERC780ClaimSetIterator struct {
	Event *ERC780ClaimSet // Event containing the contract specifics and raw log

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
func (it *ERC780ClaimSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC780ClaimSet)
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
		it.Event = new(ERC780ClaimSet)
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
func (it *ERC780ClaimSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC780ClaimSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC780ClaimSet represents a ClaimSet event raised by the ERC780 contract.
type ERC780ClaimSet struct {
	Issuer  common.Address
	Subject common.Address
	Key     [32]byte
	Value   [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimSet is a free log retrieval operation binding the contract event 0xdad2a5befe683e1646aae561032e8841b2c4f3db59d2fdedd9985163cb7be33e.
//
// Solidity: event ClaimSet(address indexed issuer, address indexed subject, bytes32 indexed key, bytes32 value)
func (_ERC780 *ERC780Filterer) FilterClaimSet(opts *bind.FilterOpts, issuer []common.Address, subject []common.Address, key [][32]byte) (*ERC780ClaimSetIterator, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _ERC780.contract.FilterLogs(opts, "ClaimSet", issuerRule, subjectRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &ERC780ClaimSetIterator{contract: _ERC780.contract, event: "ClaimSet", logs: logs, sub: sub}, nil
}

// WatchClaimSet is a free log subscription operation binding the contract event 0xdad2a5befe683e1646aae561032e8841b2c4f3db59d2fdedd9985163cb7be33e.
//
// Solidity: event ClaimSet(address indexed issuer, address indexed subject, bytes32 indexed key, bytes32 value)
func (_ERC780 *ERC780Filterer) WatchClaimSet(opts *bind.WatchOpts, sink chan<- *ERC780ClaimSet, issuer []common.Address, subject []common.Address, key [][32]byte) (event.Subscription, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _ERC780.contract.WatchLogs(opts, "ClaimSet", issuerRule, subjectRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC780ClaimSet)
				if err := _ERC780.contract.UnpackLog(event, "ClaimSet", log); err != nil {
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
