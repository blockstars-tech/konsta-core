package secrets

import (
	"github.com/spf13/cobra"
	"konsta.live/command/helper"
	"konsta.live/command/secrets/generate"
	initCmd "konsta.live/command/secrets/init"
)

func GetCommand() *cobra.Command {
	secretsCmd := &cobra.Command{
		Use:   "secrets",
		Short: "Top level SecretsManager command for interacting with secrets functionality. Only accepts subcommands.",
	}

	helper.RegisterGRPCAddressFlag(secretsCmd)

	registerSubcommands(secretsCmd)

	return secretsCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// secrets init
		initCmd.GetCommand(),
		// secrets generate
		generate.GetCommand(),
	)
}
