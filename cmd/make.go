package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var makeCmd = &cobra.Command{
	Use:   "make [directory name]",
	Short: "Create a directory or file with the given name",
	Long: `This command creates a directory or file with the specified name.
For example:

pralex make mydirectory    # Creates a directory named 'mydirectory'
pralex make myfile.go      # Creates a file named 'myfile.go'
pralex make .gitignore     # Creates a file named '.gitignore'
pralex make .env           # Creates a file named '.env'`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		makeFileOrDirectory(name)
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)
}

func makeFileOrDirectory(name string) {
	if _, err := os.Stat(name); err != nil {
		if strings.Contains(name, ".") {

			file, err := os.Create(name)
			if err != nil {
				fmt.Printf("error while creating file : %s", err)
				return
			}
			defer file.Close()
			fmt.Printf("File created named %s \n", name)

		} else {
			if err := os.Mkdir(name, 0755); err != nil {
				fmt.Printf("Error creating directory %s: %v\n", name, err)
			} else {
				fmt.Printf("Directory %s created successfully\n", name)
			}
		}
	} else {
		if strings.Contains(name, ".") {
			fmt.Println("File already exist")
		} else {
			fmt.Println("Directory already exist")
		}
	}
}
