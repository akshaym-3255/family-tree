package connect

import (
	"fmt"

	"github.com/akshaym-3255/family-tree/cmd"
	"github.com/akshaym-3255/family-tree/internal/repositories"
	"github.com/akshaym-3255/family-tree/internal/service"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "used to define between relationships between persons",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("first_member name required")
		}

		first_member := args[0]
		second_member, _ := cmd.Flags().GetString("of")
		relation, _ := cmd.Flags().GetString("as")

		memberRepo := repositories.NewMemberRepository()
		relationshipRepo := repositories.NewRelationshipRepository()
		relationshipService := service.NewRelationshipService(relationshipRepo)
		memberService := service.NewMemberService(memberRepo, relationshipService)

		err := memberService.AddConnection(first_member, second_member, relation)
		if err != nil {
			fmt.Println(err)
			return
		}

	},
}

func init() {
	cmd.RootCmd.AddCommand(connectCmd)
	connectCmd.Flags().String("as", "", "relation with person 1")
	connectCmd.Flags().String("of", "", "name of person with which relation present")

	connectCmd.MarkFlagRequired("of")
	connectCmd.MarkFlagRequired("as")

}
