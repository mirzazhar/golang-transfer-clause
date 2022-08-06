package erc20

import "github.com/mirzazhar/golang-transfer-clause/utils"

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

// AddTokenAddress adds the contract address of the ERC20-based standard
// token.
func (b *ERC20Body) AddTokenAddress(tokenAddr string) *ERC20Body {
	b.tokenAddress = tokenAddr
	return b
}

// AddData adds an account address. Later on, this address will use as a
// parameter for ERC20-based token methods: approve and tokentransferfrom.
func (b *ERC20Body) AddData(data string) *ERC20Body {
	b.data = data
	return b
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
