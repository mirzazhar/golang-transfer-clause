package erc20

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
