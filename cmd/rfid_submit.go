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
	Run: func(cmd *cobra.Command, args []string) {

		host, _ := rootCmd.PersistentFlags().GetString("hostname")
		port, _ := rootCmd.PersistentFlags().GetInt("port")
		if host == "" {
			fmt.Println("hostname must be provided")
			return
		}
		if port <= 0 {
			fmt.Println("port number must be provided")
			return
		}
		url := fmt.Sprintf("http://%s:%d/api/v1/commands/rfid/submit", host, port)
		if p, _ := cmd.Flags().GetBool("print-endpoint"); p {
			fmt.Println(url)
			return
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
				fmt.Println(err)
			}
		}

	},
}

func init() {
	rfidCmd.AddCommand(submitCmd)

	submitCmd.Flags().Bool("print-endpoint", false, "Print the endpoint URL and exit")
	submitCmd.Flags().BoolP("tee", "t", false, "Print the payload to stdout that be sent to the endpoint")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// submitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// submitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
