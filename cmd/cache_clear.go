/*
Copyright © 2023 Thiago P. Martinez <thiago.martinez@nuvem.net>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear R700 emulator cache",
	RunE: func(cmd *cobra.Command, args []string) error {
		base, err := apiBaseURL()
		if err != nil {
			return err
		}
		url := fmt.Sprintf("%s/commands/cache/info", base)
		req, _ := http.NewRequest("DELETE", url, nil)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("status code: %d", resp.StatusCode)
		}

		fmt.Println("cache cleared")

		return nil
	},
	SilenceUsage: true,
}

func init() {
	cacheCmd.AddCommand(clearCmd)
}
