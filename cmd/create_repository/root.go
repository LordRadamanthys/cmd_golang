package createrepository

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "createRepo",
	Short: "createRepo - a simple CLI to create repository",
	Long: `createRepo is a easy way to create repository in gitlab using cli
			`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing CLI '%s'", err)
		os.Exit(1)
	}
}
