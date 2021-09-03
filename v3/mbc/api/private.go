package api

import "github.com/spf13/cobra"

func NewPrivateCmd() []*cobra.Command {
	cmds := []*cobra.Command{
		NewGetBalancesCmd(),
	}

	for _, cmd := range cmds {
		cmd.Flags().StringP("profile", "p", "default", "Profile")
	}

	return cmds
}

