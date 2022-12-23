/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// deleteCompressedCmd represents the deleteCompressed command
var deleteCompressedCmd = &cobra.Command{
	Use:   "delete-compressed",
	Short: "Delete all previously compressed images",
	Run: func(cmd *cobra.Command, args []string) {

		folder, _ := cmd.Flags().GetString("folder")

		dir, err := os.Getwd()

		if err != nil {
			log.Fatal("cant get current directory")
		}

		if folder == "." {
			folder = ""
		} else {
			folder = "/" + folder
		}

		folder = dir + folder

		filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
			if err == nil {

				if info.Name() == "COMPRESSED" && info.IsDir() {
					err := os.RemoveAll(path)

					if err != nil {
						panic(err)
					}
				}

			}
			return nil
		})

	},
}

func init() {
	rootCmd.AddCommand(deleteCompressedCmd)

	deleteCompressedCmd.Flags().StringP("folder", "f", ".", "Specify the folder that you want to delete")
}
