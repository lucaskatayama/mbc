/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"

	"github.com/lucaskatayama/mbc/v4"
)

type TickersPrintable struct {
	tickers []mbc.Ticker
}

func (tp TickersPrintable) ToJSON() string {
	b, _ := json.Marshal(tp.tickers)
	return string(b)
}

func (tp TickersPrintable) ToTable() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)

	t.AppendHeader(table.Row{"Instrument", "Open", "High", "Low", "Close", "Volume", "Buy", "Sell"})

	for _, ticker := range tp.tickers {
		t.AppendRow(table.Row{
			ticker.Pair,
			ticker.Open,
			ticker.High,
			ticker.Low,
			ticker.Last,
			ticker.Volume,
			ticker.Buy,
			ticker.Sell,
		})
	}

	return t.Render()
}

// tickersCmd represents the tickers command
var tickersCmd = &cobra.Command{
	Use:   "tickers",
	Short: "List tickers",
	Run: func(cmd *cobra.Command, args []string) {
		instruments := strings.Split(instrument, ",")
		var ins []mbc.InstrumentSymbol
		for _, i := range instruments {
			ins = append(ins, mbc.InstrumentSymbol(strings.ToUpper(i)))
		}

		params := mbc.TickerParams{
			Symbols: ins,
		}
		client, err := mbc.NewPublicOnlyClient()
		if err != nil {
			panic(err)
		}
		tickers, _, err := client.PublicData.ListTickers(cmd.Context(), params)
		if err != nil {
			return
		}

		writer.Print(TickersPrintable{tickers: tickers})
	},
}

func init() {
	getCmd.AddCommand(tickersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tickersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tickersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
