package api

import (
	"context"
	"fmt"
	"github.com/lucaskatayama/mbc/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewGetBalancesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "balances",
		Short: "Get balances",
		Long:  "Get balances",
		Run: func(cmd *cobra.Command, args []string) {
			p := cmd.Flag("profile").Value.String()
			id := viper.GetString("profiles." + p + ".id")
			secret := viper.GetString("profiles." + p + ".secret")
			client := mbc.New(mbc.WithIdSecret(id, secret))
			t, err := client.GetBalances(context.Background())
			if err != nil {
				panic(err)
			}
			fmt.Printf("%#v", t)
		},
	}

	return cmd
}
