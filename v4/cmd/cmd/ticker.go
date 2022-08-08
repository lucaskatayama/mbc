/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"

	"github.com/lucaskatayama/mbc/v4"
)

type TickerPrintable struct {
	mbc.Ticker
}

func (tp TickerPrintable) ToJSON() string {
	b, _ := json.Marshal(tp.Ticker)
	return string(b)
}

func (tp TickerPrintable) ToTable() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)

	t.AppendRow(table.Row{"Open", tp.Open})
	t.AppendRow(table.Row{"High", tp.High})
	t.AppendRow(table.Row{"Low", tp.Low})
	t.AppendRow(table.Row{"Close", tp.Last})
	t.AppendRow(table.Row{"Volume", tp.Volume})
	t.AppendSeparator()
	t.AppendRow(table.Row{"Buy", tp.Buy})
	t.AppendRow(table.Row{"Sell", tp.Sell})
	t.AppendSeparator()
	t.AppendRow(table.Row{"Timestamp", time.Time(tp.Ts).Format(time.RFC822Z)})

	return t.Render()
}

// tickerCmd represents the ticker command
var tickerCmd = &cobra.Command{
	Use:   "ticker",
	Short: "Get ticker",
	Run: func(cmd *cobra.Command, args []string) {

		params := mbc.TickerParams{
			Symbols: []mbc.InstrumentSymbol{mbc.InstrumentSymbol(strings.ToUpper(instrument))},
		}
		client, err := mbc.NewPublicOnlyClient()
		if err != nil {
			panic(err)
		}
		tickers, _, err := client.PublicData.ListTickers(cmd.Context(), params)
		if err != nil {
			return
		}

		writer.Print(TickerPrintable{tickers[0]})
	},
}

func init() {
	getCmd.AddCommand(tickerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tickerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tickerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
