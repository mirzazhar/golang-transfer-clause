package erc20

import (
	"errors"
	"math/big"

	"github.com/mirzazhar/golang-transfer-clause/utils"
)

// ERC20Transform is a custom type. It only accepts values that have
// GetToAddress and GetValue methods.
type ERC20Transform interface {
	GetToAddress() string
	GetValue() string
}

// ERC20Body holds the necessary transfer information to be used by a
// transaction that will interact with the ERC20-based token standard.
type ERC20Body struct {
	to, value, data string
	tokenAddress    string
}

// New creates and returns an empty instance of ERC20Body.
func New() *ERC20Body {
	return &ERC20Body{}
}

// AddToAddress method adds the recipient address to its instance.
func (eb *ERC20Body) AddToAddress(to string) *ERC20Body {
	eb.to = to
	return eb
}

// AddValue method adds the "amount to be transferred" to its instance.
func (eb *ERC20Body) AddValue(value string) *ERC20Body {
	eb.value = value
	return eb
}

// Init creates an instance of ERC20Body using any type that implements
// ERC20Transform interface.
func Init(erc20 ERC20Transform) *ERC20Body {
	return &ERC20Body{
		to:    erc20.GetToAddress(),
		value: erc20.GetValue(),
	}
}

// AddTokenAddress adds the contract address of the ERC20-based standard
// token.
func (eb *ERC20Body) AddTokenAddress(tokenAddr string) *ERC20Body {
	eb.tokenAddress = tokenAddr
	return eb
}

// AddData adds an account address. Later on, this address will use as a
// parameter for ERC20-based token methods: approve and tokentransferfrom.
func (eb *ERC20Body) AddData(data string) *ERC20Body {
	eb.data = data
	return eb
}

// Build validates its underlying instance and then creates the
// new instance of ERC20Clause.
func (b *ERC20Body) Build() (*ERC20Clause, error) {
	if !utils.IsValidAddress(b.tokenAddress) {
		return nil, utils.ErrTokenAddress
	} else if b.tokenAddress == b.to {
		return nil, utils.ErrSameEOAContractAddr
	} else if !utils.IsValidDecimalValue(b.value) {
		return nil, utils.ErrValue
	}
	return &ERC20Clause{ERC20Body: *b}, nil
}

// ERC20Clause represents the transfer information for the ERC-20 standard used by
// a transaction within an ethereum or ethereum-based fork.
type ERC20Clause struct {
	ERC20Body
}

// GetTokenAddress returns the contract address of the ERC-20 standard token.
func (erc *ERC20Clause) GetTokenAddress() string {
	return erc.tokenAddress
}

// GetToAddress returns the receiver address for the ERC-20 standard token.
func (erc *ERC20Clause) GetToAddress() string {
	return erc.to
}

// TokenName returns the payload of the token name for the ERC-20-based getter.
func (erc *ERC20Clause) TokenName() []byte {
	data := erc20methodIDs[name]
	return data[:]
}

// TokenSymbol returns the payload of the token symbol for the ERC-20-based getter.
func (erc *ERC20Clause) TokenSymbol() []byte {
	data := erc20methodIDs[symbol]
	return data[:]
}

// TokenDecimals returns the payload of the token decimals for the ERC-20-based getter.
func (erc *ERC20Clause) TokenDecimals() []byte {
	data := erc20methodIDs[decimals]
	return data[:]
}

// TokenTotalSupply returns the payload of total supply for the ERC-20-based getter.
func (erc *ERC20Clause) TokenTotalSupply() []byte {
	data := erc20methodIDs[totalSupply]
	return data[:]
}

// TokenBalance returns the payload of token balance for the ERC-20-based getter.
func (erc *ERC20Clause) TokenBalance() ([]byte, error) {
	return erc.payload(balance)
}

// TokenTranfer returns the payload of token transfer for the ERC-20-based method.
func (erc *ERC20Clause) TokenTranfer() ([]byte, error) {
	payload, err := erc.payload(transfer)
	if err != nil {
		return nil, err
	}
	return erc.extendPayload(payload)
}

// TokenApprove returns the payload of tokens to approve for the ERC-20-based method.
func (erc *ERC20Clause) TokenApprove() ([]byte, error) {
	payload, err := erc.payload(approve)
	if err != nil {
		return nil, err
	}

	return erc.extendPayload(payload)
}

// TokenTransferFrom returns the payload of "token transfer from the address of token
// approved to another address" for the ERC-20-based method.
func (erc *ERC20Clause) TokenTransferFrom(from string) ([]byte, error) {
	payload, err := erc.payload(transferFrom)
	if err != nil {
		return nil, err
	}

	toaddress, err := utils.AddresstoBytes(from)
	if err != nil {
		return nil, err
	}

	paddedAddress := utils.LeftPadBytes(toaddress, 32)
	payload = append(payload, paddedAddress...)
	return erc.extendPayload(payload)
}

// TokenAllowance returns the payload to find the remaining number of allowed tokens for
// the ERC-20-based getters.
func (erc *ERC20Clause) TokenAllowance(owner string) ([]byte, error) {
	payload, err := erc.payload(allowance)
	if err != nil {
		return nil, err
	}

	toaddress, err := utils.AddresstoBytes(owner)
	if err != nil {
		return nil, err
	}

	paddedAddress := utils.LeftPadBytes(toaddress, 32)
	return append(payload, paddedAddress...), nil
}

// payload creates the actual data array to be used to interact with the ERC-20-based
// token standard.
func (erc *ERC20Clause) payload(method string) ([]byte, error) {
	address, err := utils.AddresstoBytes(erc.to)
	if err != nil {
		return nil, err
	}

	paddedAddress := utils.LeftPadBytes(address, 32)
	methodID := erc20methodIDs[method]

	var data []byte
	data = append(data, methodID[:]...)
	data = append(data, paddedAddress...)
	return data, nil
}

// extendPayload extends the functionality of the payload method, particularly when three
// or more values are required to complete the payload for the ERC-20-based token standard.
func (erc *ERC20Clause) extendPayload(payload []byte) ([]byte, error) {
	amount, ok := new(big.Int).SetString(erc.value, 10)
	if !ok {
		return nil, errors.New("error in converting string based value to big integers")
	}

	paddedAmount := utils.LeftPadBytes(amount.Bytes(), 32)
	payload = append(payload, paddedAmount...)
	return payload, nil
}

// GetERCPayloadData returns the payload of the given method in a byte array. Moreover, purposely,
// it can be called from outside this package or through the interface.
func (erc *ERC20Clause) GetERCPayloadData(method string) ([]byte, error) {
	var data []byte
	var err error

	switch method {
	case "name":
		data = erc.TokenName()
	case "symbol":
		data = erc.TokenSymbol()
	case "totalSupply":
		data = erc.TokenTotalSupply()
	case "decimals":
		data = erc.TokenDecimals()
	case "balanceOf":
		data, err = erc.TokenBalance()
		if err != nil {
			return nil, err
		}
	case "transfer":
		data, err = erc.TokenTranfer()
		if err != nil {
			return nil, err
		}
	case "approve":
		data, err = erc.TokenApprove()
		if err != nil {
			return nil, err
		}
	case "transferFrom":
		data, err = erc.TokenTransferFrom(erc.data)
		if err != nil {
			return nil, err
		}
	case "allowance":
		data, err = erc.TokenAllowance(erc.data)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("this method is not defined :" + method)
	}
	return data, nil
}
