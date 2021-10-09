package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


func NewConfigSetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set <key> <value>",
		Short: "Set config",
		Run: func(cmd *cobra.Command, args []string) {
			viper.Set(args[0], args[1])
			if err := viper.SafeWriteConfig(); err != nil {
				if err := viper.WriteConfig(); err != nil {
					panic(err)
				}
			}
		},
		Args: cobra.ExactArgs(2),
	}

	return cmd
}


func NewConfigInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Init config",
		Run: func(cmd *cobra.Command, args []string) {
			if err := viper.SafeWriteConfig(); err != nil {
				fmt.Printf("File %s already exists", viper.ConfigFileUsed())
			}
		},
	}

	return cmd
}
