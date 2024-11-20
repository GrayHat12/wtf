package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "wtf",
	Short:   "wtf is a cli tool for quick search.",
	Long:    "wtf is a cli tool for quick search. wtf is <x>",
	Example: "wtf is meow golang",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stderr, "Oops ! No command provided. Check `wtf --help`")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops, An error while executing wtf '%s'\n", err)
		os.Exit(1)
	}
}
