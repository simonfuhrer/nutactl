package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func newCompletionCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "completion",
		Short: "Generates bash completion scripts",
		Long: `To load completion run
	
	. <(bitbucket completion)
	
	To configure your bash shell to load completions for each session add to your bashrc
	
	# ~/.bashrc or ~/.profile
	. <(nutactl completion)
	`,
		Run: func(cmd *cobra.Command, args []string) {
			_ = rootCmd.GenBashCompletion(os.Stdout)
		},
	}

	return cmd
}
