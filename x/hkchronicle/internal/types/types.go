package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MinEventPrice is Initial Starting Price for an event that was never previously owned
var MinEventPrice = sdk.Coins{sdk.NewInt64Coin("hkctoken", 1)}

// Event is a struct that contains all the metadata of an event
type Event struct {
	Value   string           `json:"value"`
	Time    int64            `json:"time"`
	Owner   sdk.AccAddress   `json:"owner"`
	Price   sdk.Coins        `json:"price"`
	Stakers []sdk.AccAddress `json:"stakers"`
	Stake   sdk.Coins        `json:"stake"`
	Content string           `json:"content"`
}

// NewEvent returns a new Event with the minprice as the price
func NewEvent() Event {
	return Event{
		Price: MinEventPrice,
	}
}

// implement fmt.Stringer
func (w Event) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Value: %s
Price: %s`, w.Owner, w.Value, w.Price))
}

//-------fact------//
// MinPrice is 1 facttoken
var MinPrice = sdk.Coins{sdk.NewInt64Coin("facttoken", 1)}

//Fact struct
type Fact struct {
	Title       string           `json:"title"`
	Time        int64            `json:"time"`
	Place       string           `json:"place"`
	Description string           `json:"description"`
	Creator     sdk.AccAddress   `json:"creator"`
	Price       sdk.Coins        `json:"price"`
	Stakers     []sdk.AccAddress `json:"stakers"`
}

// NewFact returns a new Fact
func NewFact() Fact {
	return Fact{
		Price: MinPrice,
	}
}

//Fact Check Delegation struct
type FactCheckDelegation struct {
	DelegatorAddress sdk.AccAddress `json:"delegator_address"`
	Title            string         `json:"title"`
	Shares           sdk.Dec        `json:"shares"`
}

// NewDelegation creates a new delegation object
func NewDelegation(delegatorAddr sdk.AccAddress, title string, shares sdk.Dec) FactCheckDelegation {
	return FactCheckDelegation{
		DelegatorAddress: delegatorAddr,
		Title:            title,
		Shares:           shares,
	}
}
