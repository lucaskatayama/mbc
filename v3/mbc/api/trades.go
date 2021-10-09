package api

import (
	"context"
	"fmt"
	"github.com/lucaskatayama/mbc/v3"
	"github.com/spf13/cobra"
)

func NewGetTradesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trades",
		Short: "Get trades",
		Long:  "Get trades",
		Run: func(cmd *cobra.Command, args []string) {
			b := cmd.Flag("base").Value.String()
			q := cmd.Flag("quote").Value.String()
			client := mbc.New()
			t, err := client.Trades(context.Background(), b, q)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%#v", t)
		},
	}

	return cmd
}
