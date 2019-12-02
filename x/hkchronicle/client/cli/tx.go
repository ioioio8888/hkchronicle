package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/ioioio8888/hkchronicle/x/hkchronicle/internal/types"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	hkchronicleTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "hkchronicle transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	hkchronicleTxCmd.AddCommand(client.PostCommands(
		GetCmdBuyEvent(cdc),
		GetCmdSetEvent(cdc),
		GetCmdStakeEvent(cdc),
		GetCmdDeleteEvent(cdc),
	)...)

	return hkchronicleTxCmd
}

// GetCmdBuyEvent is the CLI command for sending a BuyEvent transaction
func GetCmdBuyEvent(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "buy-event [name] [amount]",
		Short: "bid for existing event or claim new event",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			coins, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgBuyEvent(args[0], coins, cliCtx.GetFromAddress())
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdSetEvent is the CLI command for sending a SetName transaction
func GetCmdSetEvent(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-event [name] [time] [value]",
		Short: "set the value and time associated with an event that you own",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			// if err := cliCtx.EnsureAccountExists(); err != nil {
			// 	return err
			// }
			time, terr := strconv.ParseInt(args[1], 10, 64)
			if terr != nil {
				return terr
			}
			msg := types.NewMsgSetEvent(args[0], time, args[2], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			// return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, msgs)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdBuyEvent is the CLI command for sending a BuyEvent transaction
func GetCmdStakeEvent(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "stake-event [name] [amount]",
		Short: "Stake a coin on an event",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			coins, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgStakeEvent(args[0], coins, cliCtx.GetFromAddress())
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdDeleteEvent is the CLI command for sending a DeleteEvent transaction
func GetCmdDeleteEvent(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-event [name]",
		Short: "delete the event that you own along with it's associated fields",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteEvent(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			// return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, msgs)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
