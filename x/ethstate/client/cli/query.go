package cli

import (
	"fmt"
	"strconv"

	"github.com/ajansari95/cosmicether/x/ethstate/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	// Group ethstate queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetSlotDataCmd(),
		GetContractDataCmd(),
		GetEthBlockCmd(),
	)
	// this line is used by starport scaffolding # 1

	return cmd
}

// GetSlotDataCmd returns the cli query commands for this module
func GetSlotDataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "slot-data [contractAddress] [slot]",
		Short: fmt.Sprintf("Querying slot data  for the [contract} and [slot}"),
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			//args
			contractAddress := args[0]
			slot := args[1]

			queryClient := types.NewQueryClient(clientCtx)
			req := &types.QuerySlotDataRequest{
				ContractAddress: contractAddress,
				Slot:            slot,
			}

			res, err := queryClient.SlotData(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)

		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetContractDataCmd returns the cli query commands for this module
func GetContractDataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "contract-data [contractAddress]",
		Short: fmt.Sprintf("Querying contract data for the [contract}"),
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			contractAddress := args[0]

			queryClient := types.NewQueryClient(clientCtx)
			req := &types.QueryContractDataRequest{
				ContractAddress: contractAddress,
			}

			res, err := queryClient.ContractData(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetEthBlockCmd returns the cli query commands for this module
func GetEthBlockCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "eth-block [blockNumber]",
		Short: fmt.Sprintf("Querying block data for the [blockNumber}"),
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			blockNumber := args[0]

			queryClient := types.NewQueryClient(clientCtx)

			height, err := strconv.ParseUint(blockNumber, 10, 64)
			if err != nil {
				return err
			}

			req := &types.QueryEthBlockRequest{
				BlockHeight: height,
			}

			res, err := queryClient.EthBlock(cmd.Context(), req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
