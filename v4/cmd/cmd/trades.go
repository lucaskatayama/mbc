/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"

	"github.com/lucaskatayama/mbc/v4"
)

type TradesPrintable struct {
	trades []mbc.Trade
}

func (tp TradesPrintable) ToJSON() string {
	// TODO implement me
	panic("implement me")
}

func (tp TradesPrintable) ToTable() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{"Datetime", "Id", "Side", "Price", "Volume"})

	for i := len(tp.trades) - 1; i >= 0; i-- {
		if limit > 0 && int64(len(tp.trades)-i) > limit {
			break
		}
		trade := tp.trades[i]
		color := text.FgGreen
		if trade.Type == "sell" {
			color = text.FgRed
		}
		t.AppendRow(table.Row{time.Time(trade.Ts).Format(time.RFC3339), trade.Tid, color.Sprint(trade.Type), trade.Price, trade.Amount})
	}
	return t.Render()
}

var (
	tid  int64
	from int64
	to   int64
)

// tradesCmd represents the trades command
var tradesCmd = &cobra.Command{
	Use:   "trades",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := mbc.NewPublicOnlyClient()
		if err != nil {
			panic(err)
		}

		params := mbc.TradeParams{
			Symbol: mbc.InstrumentSymbol(instrument),
			Tid:    tid,
			From:   time.Unix(from, 0),
			To:     time.Unix(to, 0),
		}
		trades, _, err := client.PublicData.ListTrades(cmd.Context(), params)
		writer.Print(TradesPrintable{trades: trades})
	},
}

func init() {
	getCmd.AddCommand(tradesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tradesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	tradesCmd.Flags().Int64VarP(&limit, "limit", "l", -1, "Size limit")
	tradesCmd.Flags().Int64Var(&tid, "tid", 0, "From tid")
	tradesCmd.Flags().Int64Var(&from, "from", 0, "From unix time")
	tradesCmd.Flags().Int64Var(&to, "to", 0, "To unix time")
}
