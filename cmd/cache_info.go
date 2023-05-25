/*
Copyright © 2023 Thiago P. Martinez <thiago.martinez@nuvem.net>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show R700 emulator cache info",
	RunE: func(cmd *cobra.Command, args []string) error {
		base, err := apiBaseURL()
		if err != nil {
			return err
		}
		url := fmt.Sprintf("%s/commands/cache/info", base)
		fmt.Println(url)

		req, _ := http.NewRequest("GET", url, nil)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("status code: %d", resp.StatusCode)
		}

		j := make(map[string]interface{})
		_ = json.NewDecoder(resp.Body).Decode(&j)
		info, _ := json.MarshalIndent(j, "", "  ")

		fmt.Println(string(info))

		return nil
	},
	SilenceUsage: true,
}

func init() {
	cacheCmd.AddCommand(infoCmd)
}
