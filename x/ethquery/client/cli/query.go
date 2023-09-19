package cli

import (
	"fmt"
	"strings"

	"github.com/ajansari95/cosmicether/x/ethquery/types"
	// "strings"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group ethquery queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetAllQueriesCmd())

	return cmd
}

func GetAllQueriesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "eth-queries",
		Short: "Querying ethqueries pending to be processed",
		Example: strings.TrimSpace(
			fmt.Sprintf(`$ %s query eth-queries`),
		),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryRequestsRequest{}

			res, err := queryClient.Queries(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)

		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
