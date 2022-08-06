package erc20

import "github.com/mirzazhar/golang-transfer-clause/utils"

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
