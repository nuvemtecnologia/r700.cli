/*
Copyright © 2023 Thiago P. Martinez <thiago.martinez@nuvem.net>
*/
package cmd

import (
	"bytes"
	"cli/model"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

// submitCmd represents the submit command
var submitCmd = &cobra.Command{
	Use:                   "submit epc_hex [...epc_hex]",
	Short:                 "Send a tag event to the rfid emulator",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	SilenceUsage:          true,
	RunE: func(cmd *cobra.Command, args []string) error {

		host, _ := rootCmd.PersistentFlags().GetString("hostname")
		port, _ := rootCmd.PersistentFlags().GetInt("port")
		if host == "" {
			return fmt.Errorf("hostname must be provided")
		}
		if port <= 0 {
			return fmt.Errorf("port number must be provided")
		}
		url := fmt.Sprintf("http://%s:%d/api/v1/commands/rfid/submit", host, port)
		if p, _ := cmd.Flags().GetBool("print-endpoint"); p {
			fmt.Println(url)
			return nil
		}

		for _, epc := range args {
			body := model.NewTagEvent(epc)
			buf := new(bytes.Buffer)
			_ = json.NewEncoder(buf).Encode(body)

			if t, _ := cmd.Flags().GetBool("tee"); t {
				fmt.Printf(buf.String())
			}

			req, _ := http.NewRequest("POST", url, buf)
			client := &http.Client{}
			_, err := client.Do(req)
			if err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	rfidCmd.AddCommand(submitCmd)

	submitCmd.Flags().Bool("print-endpoint", false, "Print the endpoint URL and exit")
	submitCmd.Flags().BoolP("tee", "t", false, "Print the payload to stdout that be sent to the endpoint")
}
