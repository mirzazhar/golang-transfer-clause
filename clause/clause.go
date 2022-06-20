package clause

// ClauseBody holds the necessary transfer information to be used by
// a transaction like a receiver address, amount, and arbitrary data.
type ClauseBody struct {
	to, value, data string
}

// New creates and returns an empty instance of the ClauseBody.
func New() *ClauseBody {
	return &ClauseBody{}
}

// AddToAddress method adds the recipient address to its instance.
func (cb *ClauseBody) AddToAddress(to string) *ClauseBody {
	cb.to = to
	return cb
}

// AddValue method adds the "amount to be transferred" to its
// instance.
func (cb *ClauseBody) AddValue(value string) *ClauseBody {
	cb.value = value
	return cb
}

// AddData method adds the arbitrary data to its object. Moreover,
// this data will store within a transaction on the ledger.
func (cb *ClauseBody) AddData(data string) *ClauseBody {
	cb.data = data
	return cb
}
