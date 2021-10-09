package api

import "github.com/spf13/cobra"

func NewPublicCmd() []*cobra.Command {
	cmds := []*cobra.Command{
		NewGetOrderbookCmd(),
		NewGetTickerCmd(),
		NewGetTradesCmd(),
	}

	for _, cmd := range cmds {
		cmd.Flags().StringP("base", "b", "BTC", "Base symbol")
		cmd.Flags().StringP("quote", "q", "BRL", "Quote symbol")
	}

	return cmds
}
