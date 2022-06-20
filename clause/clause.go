package clause

import "github.com/mirzazhar/golang-transfer-clause/utils"

// Clause represents the transfer information used by a transaction
// within an ethereum or ethereum-based fork.
type Clause struct {
	ClauseBody
}

// GetToAddress method returns the receiver address.
func (cl *Clause) GetToAddress() string {
	return cl.to
}

// GetValue method returns the value.
func (cl *Clause) GetValue() string {
	return cl.value
}

// GetData method returns the arbitrary data.
func (cl *Clause) GetData() string {
	return cl.data
}

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

// Build validates its underlying instance and then creates the
// new instance of Clause.
func (cb *ClauseBody) Build() (*Clause, error) {
	if !utils.IsValidAddress(cb.to) {
		return nil, utils.ErrToAddress
	} else if !utils.IsValidValue(cb.value) {
		return nil, utils.ErrValue
	}
	return &Clause{ClauseBody: *cb}, nil
}
