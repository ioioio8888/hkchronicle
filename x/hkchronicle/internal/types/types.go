package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MinEventPrice is Initial Starting Price for an event that was never previously owned
var MinEventPrice = sdk.Coins{sdk.NewInt64Coin("hkctoken", 1)}

// WhoseEvent is a struct that contains all the metadata of an event
type WhoseEvent struct {
	Value   string           `json:"value"`
	Time    int64            `json:"time"`
	Owner   sdk.AccAddress   `json:"owner"`
	Price   sdk.Coins        `json:"price"`
	Staker  []sdk.AccAddress `json:"staker"`
	Stake   sdk.Coins        `json:"stake"`
	Content string           `json:"content"`
}

// NewWhoseEvent returns a new WhoseEvent with the minprice as the price
func NewWhoseEvent() WhoseEvent {
	return WhoseEvent{
		Price: MinEventPrice,
	}
}

// implement fmt.Stringer
func (w WhoseEvent) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Value: %s
Price: %s`, w.Owner, w.Value, w.Price))
}
