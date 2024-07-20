/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "rename a file on the current location",
	Long: `This command renames a file  with the specified name.
For example:

pralex rename main.go rename.go    # renames main.go into rename.go `,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]
		renameName := args[1]
		if _, err := os.Stat(fileName); err != nil {
			fmt.Printf("The given file with filename %s does not exist or provide the file extention", fileName)
			return
		} else {
			if strings.Contains(renameName, ".") {
				os.Rename(fileName, renameName)
				fmt.Printf("file %s is renamed to %s", fileName, renameName)
				return
			}
			fmt.Printf("%s is not a valid file name provide proper name with extention", renameName)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)

}
