/*
Copyright Lemon Corp. All Rights Reserved.

Written by hama
*/

package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wowlsh93/goscan/core/bs/scanner"
	"github.com/wowlsh93/goscan/core/bs/version"
	"os"
)

var mainCmd = &cobra.Command{

	Use:   "bs",
	Short: "Sample scanner",
	Long:  `This application is simple scanner to learn ethereum`,
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func main() {

	mainCmd.AddCommand(scanner.Cmd())
	mainCmd.AddCommand(version.Cmd())

	if err := mainCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
