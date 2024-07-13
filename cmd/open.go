/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open current directory in VS Code",
	Run: func(cmd *cobra.Command, args []string) {
		if err := openDirectoryInVSCode(); err != nil {
			fmt.Println("Error : ", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(openCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// openCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// openCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	openCmd.Flags().StringP("unzipopen", "uzo", "", "Unzip and open file")
}

func openDirectoryInVSCode() error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("current working directory not found : %w", err)
	}

	openvscode := exec.Command("code", cwd)
	if err = openvscode.Run(); err != nil {
		return fmt.Errorf("command not executable : %w ", err)
	}
	fmt.Println("Opening current directory in VS Code ")
	return nil
}
