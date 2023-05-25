/*
Copyright © 2023 Thiago P. Martinez <thiago.martinez@nuvem.net>
*/
package cmd

import (
	"bytes"
	epc_tools "cli/epc"
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

		base, err := apiBaseURL()
		if err != nil {
			return err
		}
		url := fmt.Sprintf("%s/commands/rfid/submit", base)
		if p, _ := cmd.Flags().GetBool("print-endpoint"); p {
			fmt.Println(url)
			return nil
		}

		for _, epc := range args {

			tag, err := epc_tools.DecodeHex(epc)
			if err != nil {
				return fmt.Errorf("invalid EPC: %s", epc)
			}

			body := model.NewTagEvent(*tag)
			buf := new(bytes.Buffer)
			_ = json.NewEncoder(buf).Encode(body)

			if t, _ := cmd.Flags().GetBool("tee"); t {
				fmt.Printf(buf.String())
			}

			req, _ := http.NewRequest("POST", url, buf)
			client := &http.Client{}
			_, err = client.Do(req)
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
