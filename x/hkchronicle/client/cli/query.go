package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/ioioio8888/hkchronicle/x/hkchronicle/internal/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	hkchronicleQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the hkchronicle module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	hkchronicleQueryCmd.AddCommand(client.GetCommands(
		GetCmdResolveEvent(storeKey, cdc),
		GetCmdWhoseEvent(storeKey, cdc),
		GetCmdAllEvents(storeKey, cdc),
	)...)
	return hkchronicleQueryCmd
}

// GetCmdResolveEvent queries information about an event
func GetCmdResolveEvent(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "eresolve [event]",
		Short: "resolve event",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			event := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/eresolve/%s", queryRoute, event), nil)
			if err != nil {
				fmt.Printf("could not resolve event - %s \n", event)
				return nil
			}

			var out types.QueryResEventResolve
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdWhoseEvents queries information about an event
func GetCmdWhoseEvent(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "whoseevent [name]",
		Short: "Query whose event",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			name := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/whoseevent/%s", queryRoute, name), nil)
			if err != nil {
				fmt.Printf("could not resolve whose event - %s \n", name)
				return nil
			}

			var out types.WhoseEvent
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdAllEvent queries a list of all events
func GetCmdAllEvents(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "allevents",
		Short: "show all events",
		// Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/allevents", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not get query names\n")
				return nil
			}

			var out types.QueryResAllEvents
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
