/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/h2non/bimg"
	"github.com/spf13/cobra"
)

type File struct {
	name     string
	location string
}

// optimizeCmd represents the optimize command
var optimizeCmd = &cobra.Command{
	Use:   "optimize",
	Short: "Optimize the images in this images directory",
	Run: func(cmd *cobra.Command, args []string) {

		folder, errFolder := cmd.Flags().GetString("folder")

		if errFolder != nil {
			folder = "."
		}

		dir, err := os.Getwd()

		if err != nil {
			log.Fatal("cant get current directory")
		}

		if folder == "." {
			folder = ""
		} else {
			folder = "\\" + folder
		}

		folder = dir + folder

		filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
			if err == nil {

				if !info.IsDir() {
					buffer, _ := os.ReadFile(info.Name())
					imageProcessing(buffer, 50, path)

				}

			}
			return nil
		})
	},
}

// A new folder is created at the root of the project.
func createFolder(dirname string) error {
	_, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirname, 0755)
		if errDir != nil {
			return errDir
		}
	}
	return nil
}

// The mime type of the image is changed, it is compressed and then saved in the specified folder.
func imageProcessing(buffer []byte, quality int, dirname string) (string, error) {
	filename := strings.Replace(uuid.New().String(), "-", "", -1) + ".webp"

	converted, err := bimg.NewImage(buffer).Convert(bimg.WEBP)
	if err != nil {
		return filename, err
	}

	processed, err := bimg.NewImage(converted).Process(bimg.Options{Quality: quality})
	if err != nil {
		return filename, err
	}

	writeError := bimg.Write(fmt.Sprintf("./"+dirname+"/%s", filename), processed)
	if writeError != nil {
		return filename, writeError
	}

	return filename, nil
}

func init() {
	rootCmd.AddCommand(optimizeCmd)

	optimizeCmd.Flags().StringP("folder", "f", "", "specify the folder that you want to optimize")
	optimizeCmd.Flags().StringP("exclude-folders", "", "", "exclude folders")
}
