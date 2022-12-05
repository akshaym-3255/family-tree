package add

import (
	"fmt"

	"github.com/akshaym-3255/family-tree/cmd"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add command used to add resources",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("please add resource type one of person or relationship")
	},
}

func init() {
	cmd.RootCmd.AddCommand(addCmd)
}
