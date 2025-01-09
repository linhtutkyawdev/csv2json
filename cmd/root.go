/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/FerdinaKusumah/excel2json"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "csv2json",
	Short: "Convert CSV files to JSON format",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a CSV file to convert.")
			return
		}

		var (
			result []*map[string]interface{}
			// relative path

			path, err = filepath.Abs(args[0])
			// select only selected field
			// headers = []string{"humidity", "sound"}
			// if you want to show all headers just passing nil or empty list
			headers   = []string{}
			delimited = ","
		)
		if err != nil {
			log.Fatalf(`unable to parse file, error: %s`, err)
		}
		if result, err = excel2json.GetCsvFilePath(path, delimited, headers); err != nil {
			log.Fatalf(`unable to parse file, error: %s`, err)
		}
		for _, val := range result {
			result, _ := json.Marshal(val)
			fmt.Println(string(result))
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Define flags and configuration settings here.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.csv2json.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
