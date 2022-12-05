/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/akshaym-3255/family-tree/internal/repositories"
	"github.com/akshaym-3255/family-tree/internal/service"
	"github.com/spf13/cobra"
)

// fatherCmd represents the father command
var fatherCmd = &cobra.Command{
	Use:   "father",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		member, _ := cmd.Flags().GetString("of")

		relationshipRepo := repositories.NewRelationshipRepository()
		relationshipService := service.NewRelationshipService(relationshipRepo)
		memberRepo := repositories.NewMemberRepository()
		memberService := service.NewMemberService(memberRepo, relationshipService)

		father, err := memberService.GetFather(member)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(father)
	},
}

func init() {
	RootCmd.AddCommand(fatherCmd)
	fatherCmd.Flags().String("of", "", "name of person for which need to find father")
	fatherCmd.MarkFlagRequired("of")
}
