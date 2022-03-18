/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"

	"github.com/lucaskatayama/mbc/v4"
)

var limit int64
var instrument string

type OrderbookPrintable struct {
	mbc.Orderbook
}

func (o OrderbookPrintable) ToJSON() string {
	b, _ := json.Marshal(o.Orderbook)
	return string(b)
}

func (o OrderbookPrintable) ToTable() string {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"#", "Price", "Volume"})
	bestAsk := o.Asks[len(o.Asks)-1]
	bestBid := o.Bids[0]

	var asks []table.Row
	for i := len(o.Asks) - 1; i >= 0; i-- {
		asks = append(asks, table.Row{i, o.Asks[i][0], o.Asks[i][1]})
	}
	t.AppendRows(asks)

	t.AppendSeparator()
	t.AppendRow(table.Row{
		"Spread", bestAsk[0].Sub(bestBid[0]), bestAsk[1].Sub(bestBid[1]),
	})
	t.AppendSeparator()

	var bids []table.Row
	for i, d := range o.Bids {
		bids = append(bids, table.Row{i, d[0], d[1]})
	}
	t.AppendRows(bids)

	return t.Render()
}

// orderbookCmd represents the orderbook command
var orderbookCmd = &cobra.Command{
	Use:   "orderbook",
	Short: "Get orderbook",
	Run: func(cmd *cobra.Command, args []string) {

		if instrument == "" {
			fmt.Println("Instrument symbol should not be empty")
		}

		params := mbc.OrderbookParams{
			Symbol: mbc.InstrumentSymbol(strings.ToUpper(instrument)),
			Limit:  limit,
		}
		client, err := mbc.NewPublicOnlyClient()
		if err != nil {
			panic(err)
		}
		orderbook, _, err := client.PublicData.GetOrderbook(cmd.Context(), params)
		if err != nil {
			return
		}

		writer.Print(OrderbookPrintable{orderbook})
	},
}

func init() {
	getCmd.AddCommand(orderbookCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// orderbookCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	orderbookCmd.Flags().Int64VarP(&limit, "limit", "l", 10, "Orderbook size limit")

	orderbookCmd.Flags().StringVarP(&instrument, "instrument", "i", "", "Instrument symbol in the format BASE-QUOTE (e.g. BTC-BRL)")
}
