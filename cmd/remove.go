/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	Long: `This command removes a file or folder with the specified name.
For example:

pralex remove main.go rename.go    # removes main.go `,
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]
		if _, err := os.Stat(fileName); err != nil {
			fmt.Printf("%s dosent present in current directory", fileName)
			return
		}
		if err := os.Remove(fileName); err != nil {
			fmt.Printf("Unable to remove file or directory named %s", fileName)
			return
		}
		fmt.Printf("%s removed sucessfully ", fileName)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
