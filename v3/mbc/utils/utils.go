package utils

import (
	"github.com/spf13/cobra"
)

func DefaultCmd() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		cmd.Help()
	}
}
