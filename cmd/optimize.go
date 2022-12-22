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

	"github.com/h2non/bimg"
	"github.com/spf13/cobra"
)

// optimizeCmd represents the optimize command
var optimizeCmd = &cobra.Command{
	Use:   "optimize",
	Short: "Optimize the images in this images directory",
	Run: func(cmd *cobra.Command, args []string) {

		folder, _ := cmd.Flags().GetString("folder")
		quality, _ := cmd.Flags().GetInt("quality")

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

		println(folder)

		filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
			if err == nil {

				if !info.IsDir() {
					if !strings.Contains(info.Name(), "_COMPRESSED") {
						buffer, _ := os.ReadFile(path)
						finalPath := path[:strings.Index(path, info.Name())]
						_, err := imageProcessing(buffer, info.Name(), quality, finalPath)

						if err != nil {
							panic(err)
						}
					}
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
func imageProcessing(buffer []byte, fileName string, quality int, path string) (string, error) {

	fileType := bimg.DetermineImageType(buffer)
	fileTypeString := bimg.DetermineImageTypeName(buffer)

	filename := fileName + "." + fileTypeString

	converted, err := bimg.NewImage(buffer).Convert(fileType)
	if err != nil {
		return filename, err
	}

	processed, err := bimg.NewImage(converted).Process(bimg.Options{Quality: quality})
	if err != nil {
		return filename, err
	}

	errorCreate := createFolder(path + "/COMPRESSED")

	if errorCreate != nil {
		return fileName, errorCreate
	}

	writeError := bimg.Write(fmt.Sprintf(path+"/COMPRESSED/%s", filename), processed)
	if writeError != nil {
		return filename, writeError
	}

	return filename, nil
}

func init() {
	rootCmd.AddCommand(optimizeCmd)

	optimizeCmd.Flags().StringP("folder", "f", ".", "specify the folder that you want to optimize")
	optimizeCmd.Flags().IntP("quality", "q", 50, "quality of final image")
}
