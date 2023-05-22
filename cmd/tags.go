/*
Copyright © 2023 Thiago P. Martinez <thiago.martinez@nuvem.net>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// tagsCmd represents the tags command
var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "Tools for EPC tags",
}

func init() {
	rootCmd.AddCommand(tagsCmd)
}
