package count

import (
	"fmt"

	"github.com/akshaym-3255/family-tree/cmd"
	"github.com/akshaym-3255/family-tree/internal/repositories"
	"github.com/akshaym-3255/family-tree/internal/service"
	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "counts the diff relationships of a person",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("relation is required")
			return
		}
		relation := args[0]

		member, _ := cmd.Flags().GetString("of")

		relationshipRepo := repositories.NewRelationshipRepository()
		relationshipService := service.NewRelationshipService(relationshipRepo)
		memberRepo := repositories.NewMemberRepository()
		memberService := service.NewMemberService(memberRepo, relationshipService)
		count, err := memberService.GetCountOfRelationShip(member, relation)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(count)

	},
}

func init() {
	cmd.RootCmd.AddCommand(countCmd)
	countCmd.Flags().String("of", "", "name of person for which you need count")
	countCmd.MarkFlagRequired("of")
}
