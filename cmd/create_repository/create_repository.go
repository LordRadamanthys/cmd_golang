package createrepository

import (
	"github/LordRadamanthys/cmd_golang/cmd/service"

	"github.com/spf13/cobra"
)

var reverseCmd = &cobra.Command{
	Use:     "createRepo",
	Aliases: []string{"cr"},
	Short:   "Reverses a string",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		service.CreateRepositoryGitlabService(args[0], args[1])
		// fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(reverseCmd)
}
