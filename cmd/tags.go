/*
Copyright © 2023 Thiago P. Martinez <thiago.martinez@nuvem.net>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// tagsCmd represents the tags command
var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "A brief description of your command",
}

func init() {
	rootCmd.AddCommand(tagsCmd)

	rootCmd.PersistentFlags().VisitAll(func(flag *pflag.Flag) {
		flag.Hidden = true
	})
}
