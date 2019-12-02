package hkchronicle

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	EventRecords []Event `json:"Event_records"`
}

func NewGenesisState(EventRecords []Event) GenesisState {
	return GenesisState{EventRecords: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.EventRecords {
		if record.Owner == nil {
			return fmt.Errorf("Invalid EventRecord: Value: %s. Error: Missing Owner", record.Value)
		}
		if record.Value == "" {
			return fmt.Errorf("Invalid EventRecord: Owner: %s. Error: Missing Value", record.Owner)
		}
		if record.Time == 0 {
			return fmt.Errorf("Invalid EventRecord: Time: %d. Error: Missing Time", record.Time)
		}
		if record.Price == nil {
			return fmt.Errorf("Invalid EventRecord: Value: %s. Error: Missing Price", record.Value)
		}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		EventRecords: []Event{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.EventRecords {
		keeper.SetEvent(ctx, record.Value, record)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Event
	iterator := k.GetEventsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		name := string(iterator.Key())
		Event := k.GetEvent(ctx, name)
		records = append(records, Event)
	}
	return GenesisState{EventRecords: records}
}
