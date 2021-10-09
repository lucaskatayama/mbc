package api

import (
	"context"
	"fmt"
	"github.com/lucaskatayama/mbc/v3"
	"github.com/spf13/cobra"
)



func NewGetOrderbookCmd() *cobra.Command {
	limit := 10
	cmd := &cobra.Command{
		Use:   "orderbook",
		Short: "Get orderbook",
		Long:  "Get ordebook",
		Run: func(cmd *cobra.Command, args []string) {
			b := cmd.Flag("base").Value.String()
			q := cmd.Flag("quote").Value.String()
			client := mbc.New()
			t, err := client.Orderbook(context.Background(), b, q, mbc.WithLimit(limit))
			if err != nil {
				panic(err)
			}
			fmt.Printf("%#v", t)
		},
	}


	cmd.Flags().IntVarP(&limit, "limit", "l", 10, "Orderbook size")

	return cmd
}
