/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"embed"

	"github.com/akshaym-3255/family-tree/cmd"
	_ "github.com/akshaym-3255/family-tree/cmd/add"
	_ "github.com/akshaym-3255/family-tree/cmd/connect"
	_ "github.com/akshaym-3255/family-tree/cmd/count"
	"github.com/akshaym-3255/family-tree/internal/repositories"
)

//go:embed internal/database/*
var f embed.FS

func main() {
	repositories.MemberFile = f
	repositories.RelationFile = f
	cmd.Execute()
}
