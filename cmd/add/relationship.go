package add

import (
	"fmt"

	"github.com/akshaym-3255/family-tree/internal/repositories"
	"github.com/akshaym-3255/family-tree/internal/service"
	"github.com/spf13/cobra"
)

var relationshipCmd = &cobra.Command{
	Use:   "relationship",
	Short: "add relationship to family tree",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("name is required")
			return
		}
		name := args[0]

		relationshipRepo := repositories.NewRelationshipRepository()
		relationshipService := service.NewRelationshipService(relationshipRepo)
		
		err := relationshipService.AddRelationship(name)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	addCmd.AddCommand(relationshipCmd)
}
