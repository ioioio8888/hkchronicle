package hkchronicle

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	WhoseEventRecords []WhoseEvent `json:"whoseevent_records"`
}

func NewGenesisState(whoseEventRecords []WhoseEvent) GenesisState {
	return GenesisState{WhoseEventRecords: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.WhoseEventRecords {
		if record.Owner == nil {
			return fmt.Errorf("Invalid WhoseEventRecord: Value: %s. Error: Missing Owner", record.Value)
		}
		if record.Value == "" {
			return fmt.Errorf("Invalid WhoseEventRecord: Owner: %s. Error: Missing Value", record.Owner)
		}
		if record.Time == 0 {
			return fmt.Errorf("Invalid WhoseEventRecord: Time: %d. Error: Missing Time", record.Time)
		}
		if record.Price == nil {
			return fmt.Errorf("Invalid WhoseEventRecord: Value: %s. Error: Missing Price", record.Value)
		}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		WhoseEventRecords: []WhoseEvent{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.WhoseEventRecords {
		keeper.SetWhoseEvent(ctx, record.Value, record)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []WhoseEvent
	iterator := k.GetEventsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		name := string(iterator.Key())
		whoseEvent := k.GetWhoseEvent(ctx, name)
		records = append(records, whoseEvent)
	}
	return GenesisState{WhoseEventRecords: records}
}
