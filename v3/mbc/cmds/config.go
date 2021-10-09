package cmds

import (
	"github.com/lucaskatayama/mbc/v3/mbc/config"
	"github.com/lucaskatayama/mbc/v3/mbc/utils"
	"github.com/spf13/cobra"
)

func NewCmdConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Config command",
		Run:   utils.DefaultCmd(),
	}

	cmd.AddCommand(config.NewConfigSetCmd())
	cmd.AddCommand(config.NewConfigInitCmd())

	return cmd
}
