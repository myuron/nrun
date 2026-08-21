/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"slices"

	"github.com/spf13/cobra"
)

// PackageJSON represents the fields of package.json that nrun uses.
type PackageJSON struct {
	Scripts map[string]string `json:"scripts"`
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(_ *cobra.Command, _ []string) {
		if err := ScriptsList(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// ScriptsList prints the scripts in package.json sorted by name.
func ScriptsList() error {
	data, err := os.ReadFile("package.json")
	if err != nil {
		return err
	}

	var pkgJSON PackageJSON
	if err := json.Unmarshal(data, &pkgJSON); err != nil {
		return err
	}

	keys := slices.Sorted(maps.Keys(pkgJSON.Scripts))
	for _, key := range keys {
		fmt.Printf("%s: %s\n", key, pkgJSON.Scripts[key])
	}

	return nil
}
