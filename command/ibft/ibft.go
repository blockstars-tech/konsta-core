package ibft

import (
	"github.com/spf13/cobra"
	"konsta.live/command/helper"
	"konsta.live/command/ibft/candidates"
	"konsta.live/command/ibft/propose"
	"konsta.live/command/ibft/quorum"
	"konsta.live/command/ibft/snapshot"
	"konsta.live/command/ibft/status"
	_switch "konsta.live/command/ibft/switch"
)

func GetCommand() *cobra.Command {
	ibftCmd := &cobra.Command{
		Use:   "ibft",
		Short: "Top level IBFT command for interacting with the IBFT consensus. Only accepts subcommands.",
	}

	helper.RegisterGRPCAddressFlag(ibftCmd)

	registerSubcommands(ibftCmd)

	return ibftCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// ibft status
		status.GetCommand(),
		// ibft snapshot
		snapshot.GetCommand(),
		// ibft propose
		propose.GetCommand(),
		// ibft candidates
		candidates.GetCommand(),
		// ibft switch
		_switch.GetCommand(),
		// ibft quorum
		quorum.GetCommand(),
	)
}
