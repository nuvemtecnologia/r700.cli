/*
Copyright © 2023 Thiago P. Martinez <thiago.martinez@nuvem.net>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "r700_cli",
	Short: "client for R700 emulator",
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	initConfig()

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.r700_cli.yaml)")
	rootCmd.PersistentFlags().String("hostname", viper.GetString("hostname"), "hostname")
	rootCmd.PersistentFlags().Int("port", viper.GetInt("port"), "port")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".r700_cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	_ = viper.ReadInConfig()
}

func apiBaseURL() (string, error) {
	host, _ := rootCmd.PersistentFlags().GetString("hostname")
	port, _ := rootCmd.PersistentFlags().GetInt("port")
	if host == "" {
		return "", fmt.Errorf("hostname must be provided")
	}
	if port <= 0 {
		return "", fmt.Errorf("port number must be provided")
	}
	return fmt.Sprintf("http://%s:%d/api/v1", host, port), nil
}
