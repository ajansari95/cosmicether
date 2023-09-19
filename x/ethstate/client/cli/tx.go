package cli

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ajansari95/cosmicether/x/ethstate/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(GetSubmitSlotDataTxCmd())
	cmd.AddCommand(GetGetSlotDataFromEthTxCmd())

	// this line is used by starport scaffolding # 1

	return cmd
}

func GetSubmitSlotDataTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-slot-data [slotData]",
		Short: "Submit slot data",
		Example: strings.TrimSpace(
			fmt.Sprintf(` tx ethstate submit-slot-data [slotData jsonObject]`),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			slotDataRaw := args[0]
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var slotData types.SlotData

			err = json.Unmarshal([]byte(slotDataRaw), &slotData)
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitSlotData(slotData, clientCtx.GetFromAddress())

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd

}

func GetGetSlotDataFromEthTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-slot-data-from-eth [contractAddress] [slot] [height]",
		Short: "Get slot data from eth",
		Example: strings.TrimSpace(
			fmt.Sprintf(`$ %s tx ethstate get-slot-data-from-eth [contractAddress] [slot] [height]`),
		),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			contractAddress := args[0]
			slot := args[1]

			height, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgGetSlotDataFromEth(contractAddress, clientCtx.FromAddress, slot, height)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
