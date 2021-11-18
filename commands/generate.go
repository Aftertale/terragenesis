package commands

import (
	"github.com/spf13/cobra"
)

var _ cmder = (*generateCmd)(nil)

type generateCmd struct {
	*baseCmd
}

func newGenerateCmd() *generateCmd {
	cc := &generateCmd{}
	cc.baseCmd = newBaseCmd(&cobra.Command{
		Use:   "generate",
		Short: "generate a terraform provider.",
	})

	cc.cmd.AddCommand(
		newGenerateProviderCmd().getCommand(),
		updateGenerateProviderCmd().getCommand())

	return cc
}
