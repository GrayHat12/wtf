package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var noOfLinesFlag int

var isCmd = &cobra.Command{
	Use:     "is",
	Aliases: []string{"Is", "IS", "iS"},
	Short:   "wtf is <x>.",
	Long:    "wtf is <x>.",
	Args:    cobra.RangeArgs(1, 50),
	Example: "wtf is meow",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if noOfLinesFlag < 1 || noOfLinesFlag > 10 {
			return fmt.Errorf("lines flag cannot be outside range [1, 10]. Invalid value %d", noOfLinesFlag)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", Is(args))
	},
}

func init() {
	isCmd.Flags().IntVarP(&noOfLinesFlag, "lines", "l", 3, "number of lines in summary. Ex : `wtf is -l 2 <x>`. Must be in range [1, 10]. Defaults to 3")
	rootCmd.AddCommand(isCmd)
}
