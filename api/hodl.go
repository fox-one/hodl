// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"MixinEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"MixinTransaction\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NONCE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PID\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"name\":\"custodian\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"members\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"raw\",\"type\":\"bytes\"}],\"name\":\"mixin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vats\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"asset\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052600080546001600160401b031916815560045534801561002357600080fd5b50611800806100336000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806308ae4b0c146100675780632d459b6b1461009057806340a4e253146100be5780635cae8005146101665780635eaec0e41461017b578063e091dd1a146101aa575b600080fd5b61007a61007536600461150d565b6101d5565b604051610087919061153d565b60405180910390f35b6100b061009e366004611592565b60016020526000908152604090205481565b604051908152602001610087565b61011d6100cc3660046115bb565b600360208190526000918252604090912080546001820154600283015492909301546001600160801b0390911692916001600160a01b038116916001600160401b03600160a01b9092048216911685565b604080516001600160801b03909616865260208601949094526001600160a01b03909216928401929092526001600160401b03918216606084015216608082015260a001610087565b6101796101743660046115d4565b61026f565b005b6101926fd402f2c61ba04d82a75c5674fae1acca81565b6040516001600160801b039091168152602001610087565b6000546101bd906001600160401b031681565b6040516001600160401b039091168152602001610087565b600260205260009081526040902080546101ee90611645565b80601f016020809104026020016040519081016040528092919081815260200182805461021a90611645565b80156102675780601f1061023c57610100808354040283529160200191610267565b820191906000526020600020905b81548152906001019060200180831161024a57829003601f168201915b505050505081565b608d8110156102bc5760405162461bcd60e51b8152602060048201526014602482015273195d995b9d0819185d18481d1bdbc81cdb585b1b60621b60448201526064015b60405180910390fd5b60008060006103048286868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092939250506109479050565b90506001600160801b0381166fd402f2c61ba04d82a75c5674fae1acca146103605760405162461bcd60e51b815260206004820152600f60248201526e696e76616c69642070726f6365737360881b60448201526064016102b3565b61036b826010611696565b915060006103b28387878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092939250506109a59050565b6000549091506001600160401b038083169116146104025760405162461bcd60e51b815260206004820152600d60248201526c696e76616c6964206e6f6e636560981b60448201526064016102b3565b600054610419906001600160401b031660016116ae565b6000805467ffffffffffffffff19166001600160401b0392909216919091179055610445836008611696565b9250600061048c8488888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092939250506109479050565b9050610499846010611696565b93506104de8488888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509293925050610a029050565b61ffff1694506104ef846002611696565b9350602085111561053a5760405162461bcd60e51b8152602060048201526015602482015274696e7465676572206f7574206f6620626f756e647360581b60448201526064016102b3565b60006105e260006105dc610589888a8d8d8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929493925050610a5f9050565b6105948a60206116d9565b6001600160401b038111156105ab576105ab6116f0565b6040519080825280601f01601f1916602001820160405280156105d5576020820181803683370190505b5090610b6c565b90610be9565b90506105ee8686611696565b94506106338589898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509293925050610a029050565b61ffff169550610644856002611696565b9450600061068d86888b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929493925050610a5f9050565b90506106998787611696565b955060006106e0878b8b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092939250506109a59050565b90506106ed876008611696565b9650610732878b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509293925050610a029050565b61ffff169750610743886010611706565b61074e906002611696565b610759906002611696565b975060006107a2888a8d8d8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929493925050610a5f9050565b90506107ae8989611696565b97506107bb886002611696565b975060006108058960408e8e8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929493925050610a5f9050565b9050610812896040611696565b98508a89146108635760405162461bcd60e51b815260206004820152601860248201527f6d616c666f726d6564206576656e7420656e636f64696e67000000000000000060448201526064016102b3565b7fd81b55010c63687df072fb24ad086c9e75c10adfe0526499b87c43ab75e257428c8c604051610894929190611725565b60405180910390a16001600160801b0386166000908152600160205260409020546108c0908690611696565b6001600160801b03871660009081526001602052604081209190915582906002906108f083805160209091012090565b6001600160a01b03166001600160a01b031681526020019081526020016000209080519060200190610923929190611474565b5081516020830120610939908888888789610c47565b505050505050505050505050565b6000610954826010611696565b8351101561099c5760405162461bcd60e51b8152602060048201526015602482015274746f55696e743132385f6f75744f66426f756e647360581b60448201526064016102b3565b50016010015190565b60006109b2826008611696565b835110156109f95760405162461bcd60e51b8152602060048201526014602482015273746f55696e7436345f6f75744f66426f756e647360601b60448201526064016102b3565b50016008015190565b6000610a0f826002611696565b83511015610a565760405162461bcd60e51b8152602060048201526014602482015273746f55696e7431365f6f75744f66426f756e647360601b60448201526064016102b3565b50016002015190565b606081610a6d81601f611696565b1015610aac5760405162461bcd60e51b815260206004820152600e60248201526d736c6963655f6f766572666c6f7760901b60448201526064016102b3565b610ab68284611696565b84511015610afa5760405162461bcd60e51b8152602060048201526011602482015270736c6963655f6f75744f66426f756e647360781b60448201526064016102b3565b606082158015610b195760405191506000825260208201604052610b63565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015610b52578051835260209283019201610b3a565b5050858452601f01601f1916604052505b50949350505050565b6060806040519050835180825260208201818101602087015b81831015610b9d578051835260209283019201610b85565b50855184518101855292509050808201602086015b81831015610bca578051835260209283019201610bb2565b508651929092011591909101601f01601f191660405250905092915050565b6000610bf6826020611696565b83511015610c3e5760405162461bcd60e51b8152602060048201526015602482015274746f55696e743235365f6f75744f66426f756e647360581b60448201526064016102b3565b50016020015190565b600080610c548382610be9565b9050610c61826020611696565b915080610c93576000610c7484846109a5565b9050610c8c87878b610c86858a6116ae565b8c610ca8565b5050610c9e565b610c9e818986610dcd565b5050505050505050565b60006001600160a01b038416610cee5760405162461bcd60e51b815260206004820152600b60248201526a19dd5e4b5b9bdd0b5cd95d60aa1b60448201526064016102b3565b60001960045410610d2c5760405162461bcd60e51b81526020600482015260086024820152676f766572666c6f7760c01b60448201526064016102b3565b600460008154610d3b90611754565b9182905550600081815260036020819052604090912080546001600160801b0319166001600160801b039990991698909817885560018801969096556002870180546001600160a01b03969096166001600160e01b031990961695909517600160a01b6001600160401b0395861602179094555093909201805467ffffffffffffffff19169390921692909217905590565b6000838152600360205260409020600201546001600160a01b03838116911614610e2a5760405162461bcd60e51b815260206004820152600e60248201526d1b9bdd0b585d5d1a1bdc9a5e995960921b60448201526064016102b3565b6000838152600360205260409020600201546001600160401b03808316600160a01b9092041610610e875760405162461bcd60e51b81526020600482015260076024820152661b9bdd0b595b9960ca1b60448201526064016102b3565b6000610e9284610ffb565b5060008581526003602081815260408084209283015483546001909401546001600160a01b038a168652600290935290842080549596509394610f74946001600160401b03909216936001600160801b031692918791610ef190611645565b80601f0160208091040260200160405190810160405280929190818152602001828054610f1d90611645565b8015610f6a5780601f10610f3f57610100808354040283529160200191610f6a565b820191906000526020600020905b815481529060010190602001808311610f4d57829003601f168201915b5050505050611109565b90507fdb53e751d28ed0d6e3682814bf8d23f7dd7b29c94f74a56fbb7f88e9dca9f39b81604051610fa5919061153d565b60405180910390a1505050600091825250600360208190526040822080546001600160801b031916815560018101929092556002820180546001600160e01b031916905501805467ffffffffffffffff19169055565b6040805160208082528183019092526060916000918291602082018180368337019050509050836000805b60208161ffff1610156110ce57828161ffff16602081106110495761104961176f565b1a60f81b848261ffff16815181106110635761106361176f565b60200101906001600160f81b031916908160001a905350600060f81b848261ffff16815181106110955761109561176f565b01602001516001600160f81b0319161180156110b3575061ffff8216155b156110bc578091505b806110c681611785565b915050611026565b5060006110dc8260206117a7565b90506110fd61ffff83166110f18460206117a7565b86919061ffff16610a5f565b97909650945050505050565b6060608083511061114e5760405162461bcd60e51b815260206004820152600f60248201526e657874726120746f6f206c6172676560881b60448201526064016102b3565b6001600160801b0385166000908152600160205260409020548411156111af5760405162461bcd60e51b815260206004820152601660248201527534b739bab33334b1b4b2b73a1031bab9ba37b234b0b760511b60448201526064016102b3565b6001600160801b0385166000908152600160205260409020546111d39085906116d9565b6001600160801b0386166000908152600160205260408120919091556112086fd402f2c61ba04d82a75c5674fae1acca6112d6565b905061121d61121688611366565b8290610b6c565b905061122b611216876112d6565b905060008061123987610ffb565b91509150611250611249826113ed565b8490610b6c565b925061125c8383610b6c565b925061126b61124987516113ed565b92506112778387610b6c565b92506112b160085b6040519080825280601f01601f1916602001820160405280156112a9576020820181803683370190505b508490610b6c565b92506112bd8386610b6c565b92506112c9600261127f565b9998505050505050505050565b60408051601080825281830190925260609160009190602082018180368337019050509050608083901b60005b601081101561135d5781816010811061131e5761131e61176f565b1a60f81b8382815181106113345761133461176f565b60200101906001600160f81b031916908160001a9053508061135581611754565b915050611303565b50909392505050565b6040805160088082528183019092526060916000919060208201818036833701905050905060c083901b60005b600881101561135d578181600881106113ae576113ae61176f565b1a60f81b8382815181106113c4576113c461176f565b60200101906001600160f81b031916908160001a905350806113e581611754565b915050611393565b6040805160028082528183019092526060916000919060208201818036833701905050905060f083901b60005b600281101561135d578181600281106114355761143561176f565b1a60f81b83828151811061144b5761144b61176f565b60200101906001600160f81b031916908160001a9053508061146c81611754565b91505061141a565b82805461148090611645565b90600052602060002090601f0160209004810192826114a257600085556114e8565b82601f106114bb57805160ff19168380011785556114e8565b828001600101855582156114e8579182015b828111156114e85782518255916020019190600101906114cd565b506114f49291506114f8565b5090565b5b808211156114f457600081556001016114f9565b60006020828403121561151f57600080fd5b81356001600160a01b038116811461153657600080fd5b9392505050565b600060208083528351808285015260005b8181101561156a5785810183015185820160400152820161154e565b8181111561157c576000604083870101525b50601f01601f1916929092016040019392505050565b6000602082840312156115a457600080fd5b81356001600160801b038116811461153657600080fd5b6000602082840312156115cd57600080fd5b5035919050565b600080602083850312156115e757600080fd5b82356001600160401b03808211156115fe57600080fd5b818501915085601f83011261161257600080fd5b81358181111561162157600080fd5b86602082850101111561163357600080fd5b60209290920196919550909350505050565b600181811c9082168061165957607f821691505b6020821081141561167a57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b600082198211156116a9576116a9611680565b500190565b60006001600160401b038083168185168083038211156116d0576116d0611680565b01949350505050565b6000828210156116eb576116eb611680565b500390565b634e487b7160e01b600052604160045260246000fd5b600081600019048311821515161561172057611720611680565b500290565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b600060001982141561176857611768611680565b5060010190565b634e487b7160e01b600052603260045260246000fd5b600061ffff8083168181141561179d5761179d611680565b6001019392505050565b600061ffff838116908316818110156117c2576117c2611680565b03939250505056fea26469706673582212203ea2dba7f351856c3fdd295db53e7b95d2635fe588f1a27465a355bae71790d364736f6c634300080a0033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// NONCE is a free data retrieval call binding the contract method 0xe091dd1a.
//
// Solidity: function NONCE() view returns(uint64)
func (_Api *ApiCaller) NONCE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "NONCE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NONCE is a free data retrieval call binding the contract method 0xe091dd1a.
//
// Solidity: function NONCE() view returns(uint64)
func (_Api *ApiSession) NONCE() (uint64, error) {
	return _Api.Contract.NONCE(&_Api.CallOpts)
}

// NONCE is a free data retrieval call binding the contract method 0xe091dd1a.
//
// Solidity: function NONCE() view returns(uint64)
func (_Api *ApiCallerSession) NONCE() (uint64, error) {
	return _Api.Contract.NONCE(&_Api.CallOpts)
}

// PID is a free data retrieval call binding the contract method 0x5eaec0e4.
//
// Solidity: function PID() view returns(uint128)
func (_Api *ApiCaller) PID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "PID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PID is a free data retrieval call binding the contract method 0x5eaec0e4.
//
// Solidity: function PID() view returns(uint128)
func (_Api *ApiSession) PID() (*big.Int, error) {
	return _Api.Contract.PID(&_Api.CallOpts)
}

// PID is a free data retrieval call binding the contract method 0x5eaec0e4.
//
// Solidity: function PID() view returns(uint128)
func (_Api *ApiCallerSession) PID() (*big.Int, error) {
	return _Api.Contract.PID(&_Api.CallOpts)
}

// Custodian is a free data retrieval call binding the contract method 0x2d459b6b.
//
// Solidity: function custodian(uint128 ) view returns(uint256)
func (_Api *ApiCaller) Custodian(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "custodian", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Custodian is a free data retrieval call binding the contract method 0x2d459b6b.
//
// Solidity: function custodian(uint128 ) view returns(uint256)
func (_Api *ApiSession) Custodian(arg0 *big.Int) (*big.Int, error) {
	return _Api.Contract.Custodian(&_Api.CallOpts, arg0)
}

// Custodian is a free data retrieval call binding the contract method 0x2d459b6b.
//
// Solidity: function custodian(uint128 ) view returns(uint256)
func (_Api *ApiCallerSession) Custodian(arg0 *big.Int) (*big.Int, error) {
	return _Api.Contract.Custodian(&_Api.CallOpts, arg0)
}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members(address ) view returns(bytes)
func (_Api *ApiCaller) Members(opts *bind.CallOpts, arg0 common.Address) ([]byte, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "members", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members(address ) view returns(bytes)
func (_Api *ApiSession) Members(arg0 common.Address) ([]byte, error) {
	return _Api.Contract.Members(&_Api.CallOpts, arg0)
}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members(address ) view returns(bytes)
func (_Api *ApiCallerSession) Members(arg0 common.Address) ([]byte, error) {
	return _Api.Contract.Members(&_Api.CallOpts, arg0)
}

// Vats is a free data retrieval call binding the contract method 0x40a4e253.
//
// Solidity: function vats(uint256 ) view returns(uint128 asset, uint256 amount, address guy, uint64 end, uint64 nonce)
func (_Api *ApiCaller) Vats(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Asset  *big.Int
	Amount *big.Int
	Guy    common.Address
	End    uint64
	Nonce  uint64
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "vats", arg0)

	outstruct := new(struct {
		Asset  *big.Int
		Amount *big.Int
		Guy    common.Address
		End    uint64
		Nonce  uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Asset = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Guy = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.End = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.Nonce = *abi.ConvertType(out[4], new(uint64)).(*uint64)

	return *outstruct, err

}

// Vats is a free data retrieval call binding the contract method 0x40a4e253.
//
// Solidity: function vats(uint256 ) view returns(uint128 asset, uint256 amount, address guy, uint64 end, uint64 nonce)
func (_Api *ApiSession) Vats(arg0 *big.Int) (struct {
	Asset  *big.Int
	Amount *big.Int
	Guy    common.Address
	End    uint64
	Nonce  uint64
}, error) {
	return _Api.Contract.Vats(&_Api.CallOpts, arg0)
}

// Vats is a free data retrieval call binding the contract method 0x40a4e253.
//
// Solidity: function vats(uint256 ) view returns(uint128 asset, uint256 amount, address guy, uint64 end, uint64 nonce)
func (_Api *ApiCallerSession) Vats(arg0 *big.Int) (struct {
	Asset  *big.Int
	Amount *big.Int
	Guy    common.Address
	End    uint64
	Nonce  uint64
}, error) {
	return _Api.Contract.Vats(&_Api.CallOpts, arg0)
}

// Mixin is a paid mutator transaction binding the contract method 0x5cae8005.
//
// Solidity: function mixin(bytes raw) returns()
func (_Api *ApiTransactor) Mixin(opts *bind.TransactOpts, raw []byte) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "mixin", raw)
}

// Mixin is a paid mutator transaction binding the contract method 0x5cae8005.
//
// Solidity: function mixin(bytes raw) returns()
func (_Api *ApiSession) Mixin(raw []byte) (*types.Transaction, error) {
	return _Api.Contract.Mixin(&_Api.TransactOpts, raw)
}

// Mixin is a paid mutator transaction binding the contract method 0x5cae8005.
//
// Solidity: function mixin(bytes raw) returns()
func (_Api *ApiTransactorSession) Mixin(raw []byte) (*types.Transaction, error) {
	return _Api.Contract.Mixin(&_Api.TransactOpts, raw)
}

// ApiMixinEventIterator is returned from FilterMixinEvent and is used to iterate over the raw logs and unpacked data for MixinEvent events raised by the Api contract.
type ApiMixinEventIterator struct {
	Event *ApiMixinEvent // Event containing the contract specifics and raw log

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
func (it *ApiMixinEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiMixinEvent)
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
		it.Event = new(ApiMixinEvent)
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
func (it *ApiMixinEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiMixinEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiMixinEvent represents a MixinEvent event raised by the Api contract.
type ApiMixinEvent struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMixinEvent is a free log retrieval operation binding the contract event 0xd81b55010c63687df072fb24ad086c9e75c10adfe0526499b87c43ab75e25742.
//
// Solidity: event MixinEvent(bytes arg0)
func (_Api *ApiFilterer) FilterMixinEvent(opts *bind.FilterOpts) (*ApiMixinEventIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "MixinEvent")
	if err != nil {
		return nil, err
	}
	return &ApiMixinEventIterator{contract: _Api.contract, event: "MixinEvent", logs: logs, sub: sub}, nil
}

// WatchMixinEvent is a free log subscription operation binding the contract event 0xd81b55010c63687df072fb24ad086c9e75c10adfe0526499b87c43ab75e25742.
//
// Solidity: event MixinEvent(bytes arg0)
func (_Api *ApiFilterer) WatchMixinEvent(opts *bind.WatchOpts, sink chan<- *ApiMixinEvent) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "MixinEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiMixinEvent)
				if err := _Api.contract.UnpackLog(event, "MixinEvent", log); err != nil {
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

// ParseMixinEvent is a log parse operation binding the contract event 0xd81b55010c63687df072fb24ad086c9e75c10adfe0526499b87c43ab75e25742.
//
// Solidity: event MixinEvent(bytes arg0)
func (_Api *ApiFilterer) ParseMixinEvent(log types.Log) (*ApiMixinEvent, error) {
	event := new(ApiMixinEvent)
	if err := _Api.contract.UnpackLog(event, "MixinEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiMixinTransactionIterator is returned from FilterMixinTransaction and is used to iterate over the raw logs and unpacked data for MixinTransaction events raised by the Api contract.
type ApiMixinTransactionIterator struct {
	Event *ApiMixinTransaction // Event containing the contract specifics and raw log

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
func (it *ApiMixinTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiMixinTransaction)
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
		it.Event = new(ApiMixinTransaction)
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
func (it *ApiMixinTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiMixinTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiMixinTransaction represents a MixinTransaction event raised by the Api contract.
type ApiMixinTransaction struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMixinTransaction is a free log retrieval operation binding the contract event 0xdb53e751d28ed0d6e3682814bf8d23f7dd7b29c94f74a56fbb7f88e9dca9f39b.
//
// Solidity: event MixinTransaction(bytes arg0)
func (_Api *ApiFilterer) FilterMixinTransaction(opts *bind.FilterOpts) (*ApiMixinTransactionIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "MixinTransaction")
	if err != nil {
		return nil, err
	}
	return &ApiMixinTransactionIterator{contract: _Api.contract, event: "MixinTransaction", logs: logs, sub: sub}, nil
}

// WatchMixinTransaction is a free log subscription operation binding the contract event 0xdb53e751d28ed0d6e3682814bf8d23f7dd7b29c94f74a56fbb7f88e9dca9f39b.
//
// Solidity: event MixinTransaction(bytes arg0)
func (_Api *ApiFilterer) WatchMixinTransaction(opts *bind.WatchOpts, sink chan<- *ApiMixinTransaction) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "MixinTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiMixinTransaction)
				if err := _Api.contract.UnpackLog(event, "MixinTransaction", log); err != nil {
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

// ParseMixinTransaction is a log parse operation binding the contract event 0xdb53e751d28ed0d6e3682814bf8d23f7dd7b29c94f74a56fbb7f88e9dca9f39b.
//
// Solidity: event MixinTransaction(bytes arg0)
func (_Api *ApiFilterer) ParseMixinTransaction(log types.Log) (*ApiMixinTransaction, error) {
	event := new(ApiMixinTransaction)
	if err := _Api.contract.UnpackLog(event, "MixinTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
