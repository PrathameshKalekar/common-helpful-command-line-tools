/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open current directory in VS Code",
	Run: func(cmd *cobra.Command, args []string) {
		uzo, _ := cmd.Flags().GetString("unzipandopen")
		if uzo != "" {
			if err := unzipAndOpen(uzo); err != nil {
				fmt.Println("Error : ", err)
			}
		} else {
			if err := openDirectoryInVSCode(); err != nil {
				fmt.Println("Error : ", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(openCmd)

	openCmd.Flags().StringP("unzipandopen", "u", "", "Unzip and open file")
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

func unzipAndOpen(zipFile string) error {
	dest := strings.TrimSuffix(zipFile, ".zip")
	if err := unzip(zipFile, dest); err != nil {
		return err
	}
	fmt.Printf("Unzipping %s to %s\n", zipFile, dest)
	openvscode := exec.Command("code", dest)
	if err := openvscode.Run(); err != nil {
		return fmt.Errorf("command not executable : %w", err)
	}
	fmt.Println("Opening unzipped file in VS Code")
	return nil
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	if err := os.MkdirAll(dest, 0755); err != nil {
		return err
	}

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, f.Mode())
		} else {
			if err := os.MkdirAll(filepath.Dir(fpath), f.Mode()); err != nil {
				return err
			}

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			rc, err := f.Open()
			if err != nil {
				return err
			}

			_, err = io.Copy(outFile, rc)
			outFile.Close()
			rc.Close()

			if err != nil {
				return err
			}
		}
	}
	return nil
}
