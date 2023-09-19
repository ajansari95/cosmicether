package cli

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ajansari95/cosmicether/x/ethquery/types"
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

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(GetSubmitQueryResponseTxCmd())
	return cmd
}

func GetSubmitQueryResponseTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-query-response [query_id] [result] [height]",
		Short: "Submit query response",
		Example: strings.TrimSpace(
			fmt.Sprintf(`$ %s tx ethquery submit-query-response [query_id] [result]`,
				types.ModuleName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			queryID := args[0]
			result := args[1]

			height, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitQueryResponse(queryID, height, []byte(result), clientCtx.GetFromAddress().String())

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
