package add

import (
	"fmt"

	"github.com/akshaym-3255/family-tree/internal/repositories"
	"github.com/akshaym-3255/family-tree/internal/service"
	"github.com/spf13/cobra"
)

var personCmd = &cobra.Command{
	Use:   "person",
	Short: "adds a person in a family tree",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("name is required")
			return
		}
		name := args[0]

		relationshipRepo := repositories.NewRelationshipRepository()
		relationshipService := service.NewRelationshipService(relationshipRepo)
		memberRepo := repositories.NewMemberRepository()
		memberService := service.NewMemberService(memberRepo, relationshipService)

		err := memberService.AddMember(name)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	addCmd.AddCommand(personCmd)
}
