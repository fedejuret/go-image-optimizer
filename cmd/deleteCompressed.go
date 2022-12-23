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
	Short: "A brief description of your command",
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCompressedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCompressedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	deleteCompressedCmd.Flags().StringP("folder", "f", ".", "Specify the folder that you want to delete")
}
