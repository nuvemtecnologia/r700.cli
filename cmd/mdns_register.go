/*
Copyright © 2023 Thiago P. Martinez <thiago.martinez@nuvem.net>
*/
package cmd

import (
	"fmt"
	"github.com/grandcat/zeroconf"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a mDNS service",
	RunE: func(cmd *cobra.Command, args []string) error {

		instance, _ := cmd.Flags().GetString("instance")
		if instance == "" {
			instance = "r700-emulator"
		}
		port, _ := rootCmd.PersistentFlags().GetInt("port")
		if port <= 0 {
			return fmt.Errorf("port number is required")
		}

		srv, err := zeroconf.Register(instance, "_piranha._tcp", "local.", port, nil, nil)
		if err != nil {
			return err
		}

		fmt.Printf("Registered service on _http._tcp (instance=%s port=%d), press Ctrl-C to exit...\n", instance, port)

		c := make(chan os.Signal, 1)
		done := make(chan bool, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			<-c
			srv.Shutdown()
			done <- true
		}()

		<-done

		return nil
	},
}

func init() {
	mdnsCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringP("instance", "i", "r700-emulator", "service instance name")
}
