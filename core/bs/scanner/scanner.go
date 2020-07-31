/*
Copyright Lemon Corp. All Rights Reserved.

Written by hama
*/

package scanner

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	scannerFuncName = "scanner"
	scannerCmdDes   = "Operate a scanner: start|stop."
)

var ConfigPath string

func Cmd() *cobra.Command {

	scannerCmd.AddCommand(startCmd())

	return scannerCmd
}

var scannerCmd = &cobra.Command{
	Use:   scannerFuncName,
	Short: fmt.Sprint(scannerCmdDes),
	Long:  fmt.Sprint(scannerCmdDes),
}
