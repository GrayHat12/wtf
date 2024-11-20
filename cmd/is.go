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
	Args:    cobra.RangeArgs(1, 50),
	Example: "wtf is meow",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", Is(args))
	},
}

func init() {
	rootCmd.AddCommand(isCmd)
}
