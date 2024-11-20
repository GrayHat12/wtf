package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var isCmd = &cobra.Command{
	Use:     "is",
	Aliases: []string{"Is", "IS", "iS"},
	Short:   "wtf is <x>.",
	Long:    "wtf is <x>.",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", Is(args))
	},
}

func init() {
	rootCmd.AddCommand(isCmd)
}
