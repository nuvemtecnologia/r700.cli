/*
Copyright © 2023 Thiago P. Martinez <thiago.martinez@nuvem.net>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// mdnsCmd represents the mdns command
var mdnsCmd = &cobra.Command{
	Use:   "mdns",
	Short: "Multicast DNS commands",
}

func init() {
	rootCmd.AddCommand(mdnsCmd)
}
