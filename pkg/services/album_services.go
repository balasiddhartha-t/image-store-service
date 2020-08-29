package services

import (
	"fmt"
	utilities "image-store-service/pkg/utilities"
	"os"
	"path/filepath"
)

// Create the album
func CreateDirectory(albumName string) string {
	// Get the current working directory
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	// Create the album if there is no existing album
	error := os.Mkdir(filepath.Join(workingDir, "..", "Albums", albumName), 0755)
	if error != nil {
		fmt.Println(error)
		fmt.Println("File already exists")
		return "File existing error"
	}
	return "Album Created"
}

//Delete the album
func DeleteDirectory(albumName string) string {
	// Get the current working directory
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	// Check if album exists
	fileExists, err := utilities.Exists(filepath.Join(workingDir, "..", "Albums", albumName))

	if fileExists {
		// Remove the album
		removeSucess := os.Remove(filepath.Join(workingDir, "..", "Albums", albumName))
		if removeSucess != nil {
			fmt.Println(removeSucess)
			return "Error occured while deleting directory"
		}
		return "Album is Successfully deleted"
	} else {
		return "Album Doesn't exist"
	}
}
