/*
Copyright © 2023 Thiago P. Martinez <thiago.martinez@nuvem.net>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// rfidCmd represents the rfid command
var rfidCmd = &cobra.Command{
	Use:   "rfid",
	Short: "RIFD commands",
}

func init() {
	rootCmd.AddCommand(rfidCmd)
}
