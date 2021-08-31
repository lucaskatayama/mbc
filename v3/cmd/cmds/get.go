package cmds

import (
	"github.com/lucaskatayama/mbc/v3/cmd/api"
	"github.com/lucaskatayama/mbc/v3/cmd/utils"

	"github.com/spf13/cobra"
)

func NewCmdGet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get command",
		Long:  "Get Command",
		Run:   utils.DefaultCmd(),
	}

	cmd.AddCommand(api.NewPublicCmd()...)

	return cmd
}
