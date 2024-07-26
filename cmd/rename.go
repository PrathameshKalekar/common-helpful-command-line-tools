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
		oldName := args[0]
		newName := args[1]

		if err := validateOldName(oldName); err != nil {
			fmt.Println(err)
			return
		}

		if err := renameItem(oldName, newName); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)

}

// validateOldName checks if the old file or folder exists
func validateOldName(name string) error {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return fmt.Errorf("The given file or folder with name %s does not exist", name)
	}
	return nil
}

// renameItem renames a file or folder based on the new name
func renameItem(oldName, newName string) error {
	if strings.Contains(newName, ".") {
		// Rename the file
		if err := os.Rename(oldName, newName); err != nil {
			return fmt.Errorf("Error renaming file %s to %s: %v", oldName, newName, err)
		}
		fmt.Printf("File %s is renamed to %s\n", oldName, newName)
	} else {
		// Rename the folder
		if err := os.Rename(oldName, newName); err != nil {
			return fmt.Errorf("Error renaming folder %s to %s: %v", oldName, newName, err)
		}
		fmt.Printf("Folder %s is renamed to %s\n", oldName, newName)
	}
	return nil
}
