package cli

import (
	"context"

	"github.com/code0xff/interchange/x/ibcdex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListSellOrderBook() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-sell-order-book",
		Short: "list all sell-order-book",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllSellOrderBookRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.SellOrderBookAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowSellOrderBook() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-sell-order-book [index]",
		Short: "shows a sell-order-book",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetSellOrderBookRequest{
				Index: args[0],
			}

			res, err := queryClient.SellOrderBook(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
