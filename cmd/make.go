package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var makeCmd = &cobra.Command{
	Use:   "make [directory name]",
	Short: "Create a directory with the given name",
	Long: `This command creates a directory with the specified name. 
For example:

pralex make mydirectory

This will create a directory named 'mydirectory'.`,
	Run: func(cmd *cobra.Command, args []string) {
		dirName := args[0]
		if err := os.Mkdir(dirName, 0755); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dirName, err)
		} else {
			fmt.Printf("Directory %s created successfully\n", dirName)
		}
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)
}
