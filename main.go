/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/akshaym-3255/family-tree/cmd"
	_ "github.com/akshaym-3255/family-tree/cmd/add"
	_ "github.com/akshaym-3255/family-tree/cmd/connect"
	_ "github.com/akshaym-3255/family-tree/cmd/count"
)

func main() {
	cmd.Execute()
}
