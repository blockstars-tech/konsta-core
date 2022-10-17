package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"konsta.live/command/backup"
	"konsta.live/command/genesis"
	"konsta.live/command/helper"
	"konsta.live/command/ibft"
	"konsta.live/command/license"
	"konsta.live/command/loadbot"
	"konsta.live/command/monitor"
	"konsta.live/command/peers"
	"konsta.live/command/secrets"
	"konsta.live/command/server"
	"konsta.live/command/status"
	"konsta.live/command/txpool"
	"konsta.live/command/version"
	"konsta.live/command/whitelist"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Konsta is a framework for building Ethereum-compatible Blockchain networks",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		monitor.GetCommand(),
		loadbot.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		whitelist.GetCommand(),
		license.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
