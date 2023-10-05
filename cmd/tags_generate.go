/*
Copyright © 2023 Thiago P. Martinez <thiago.martinez@nuvem.net>
*/
package cmd

import (
	"cli/epc"
	"fmt"
	"github.com/spf13/cobra"
)

var header uint8
var manager uint32
var class uint32
var serialNumber uint64
var printDocs bool
var printHex bool
var printB64 bool
var printTagUri bool
var printPureIdentityUri bool
var times uint32
var inline bool

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:          "generate",
	Short:        "Generate a new EPC tag",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		if printDocs {
			fmt.Println("Please, visit https://www.gs1.org/services/epc-encoderdecoder to learn more about EPC encoding")
			return nil
		}

		for i := uint32(0); i < times; i++ {
			epc, err := epc.NewEPC(header, manager, class, serialNumber)
			if err != nil {
				return err
			}

			if i > 0 {
				if inline {
					fmt.Printf(" ")
				} else {
					fmt.Println()
				}
			}

			if printPureIdentityUri {
				fmt.Printf(epc.PureIdentityURI())
				continue
			}

			if printTagUri {
				fmt.Printf(epc.TagURI())
				continue
			}

			if printB64 {
				b64, err := epc.B64()
				if err != nil {
					return err
				}
				fmt.Printf(b64)
				continue
			}

			if printHex {
				fmt.Printf(epc.Hex())
				continue
			}

			return fmt.Errorf("no output format specified")
		}

		fmt.Println()

		return nil
	},
}

func init() {
	tagsCmd.AddCommand(generateCmd)

	generateCmd.Flags().Uint8Var(&header, "header", header, "header of generated tag")
	generateCmd.Flags().Uint32VarP(&manager, "manager", "m", manager, "company prefix of generated tag")
	generateCmd.Flags().Uint32VarP(&class, "class", "c", class, "item reference of generated tag (randomly generated)")
	generateCmd.Flags().Uint64VarP(&serialNumber, "serial", "s", serialNumber, "serial number of generated tag (randomly generated)")
	generateCmd.Flags().BoolVar(&printHex, "hex", true, "print hex representation of generated tag")
	generateCmd.Flags().BoolVar(&printB64, "b64", false, "print base64 representation of generated tag")
	generateCmd.Flags().BoolVar(&printPureIdentityUri, "identity-uri", false, "print pure identity uri")
	generateCmd.Flags().BoolVar(&printDocs, "docs", false, "print documentation")
	generateCmd.Flags().Uint32VarP(&times, "times", "t", 1, "number of tags to generate")
	generateCmd.Flags().BoolVar(&inline, "inline", false, "print output in a single line")
}
