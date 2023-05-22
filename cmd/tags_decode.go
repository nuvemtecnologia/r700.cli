/*
Copyright © 2023 Thiago P. Martinez <thiago.martinez@nuvem.net>
*/
package cmd

import (
	"cli/epc"
	"fmt"
	"github.com/spf13/cobra"
)

const (
	HexFormat string = "hex"
	B64Format string = "b64"
)

var quiet bool
var fields = []string{"header", "manager", "class", "serial"}

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode EPC tag",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("no EPC provided")
		}

		var format string
		switch len(args[0]) {
		case 16:
			format = B64Format
		case 24:
			format = HexFormat
		default:
			return fmt.Errorf("invalid format. Please, use 'hex' or 'b64'")
		}

		var tag *epc.EPC
		var err error

		if format == HexFormat {
			tag, err = epc.DecodeHex(args[0])
			if err != nil {
				return err
			}
		}
		if format == B64Format {
			tag, err = epc.DecodeB64(args[0])
			if err != nil {
				return err
			}
		}

		for _, f := range fields {
			var format string
			switch f {
			case "header":
				format = "Header: %d\n"
				if quiet {
					format = "%d\n"
				}
				fmt.Printf(format, tag.Header)
			case "manager":
				format = "Manager: %d\n"
				if quiet {
					format = "%d\n"
				}
				fmt.Printf(format, tag.Manager)
			case "class":
				format = "Class: %d\n"
				if quiet {
					format = "%d\n"
				}
				fmt.Printf(format, tag.Class)
			case "serial":
				format = "Serial Number: %d\n"
				if quiet {
					format = "%d\n"
				}
				fmt.Printf(format, tag.Serial)
			default:
				return fmt.Errorf("invalid field '%s'", f)
			}
		}

		return nil
	},
	SilenceUsage: true,
}

func init() {
	tagsCmd.AddCommand(decodeCmd)

	decodeCmd.Flags().StringSliceVar(&fields, "field", fields, "fields to be printed")
	decodeCmd.Flags().BoolVarP(&quiet, "quiet", "q", quiet, "no print labels")
}
