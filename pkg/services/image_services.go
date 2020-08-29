package services

import (
	"fmt"
	utilities "image-store-service/pkg/utilities"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

//Add an Image to the album
func CreateImage(file multipart.File, folder string, fileName string) string {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Print("Error Spotted \n", err)

	}
	// Check whether the album exists or not
	albumpath := filepath.Join(wd, "..", "Albums", folder)
	albumexisting, err := utilities.Exists(albumpath)
	if !(albumexisting) {
		return "Album doesn't exist. Please create the Album"
	}

	path := filepath.Join(albumpath, fileName)
	existing, err := utilities.Exists(path)
	if existing {
		return "file already exists"
	}
	if err != nil {
		return "Error Encountered"
	}
	//create a dummy file to clone the data
	newFile, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return "Unable to create a new file"
	}
	defer newFile.Close()

	// set position back to start.
	file.Seek(0, 0)
	// Copy the file
	io.Copy(newFile, file)

	//Produce a message to Kafka
	AlbumKafkaProducer(fileName + "is Created")
	return "Image Created"
}

//Delete an Image from the album
func DeleteImage(folder string, fileName string) string {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Print("Error Spotted \n", err)
	}
	// Get the path of the file
	path := filepath.Join(wd, "..", "Albums", folder, fileName)

	//Check if the file exists
	fileExists, err := utilities.Exists(path)
	if fileExists {
		err := os.Remove(path)
		fmt.Println(path)
		if err != nil {
			return "Error encountered while deleting the file"
		}
		//Produce a message to Kafka
		AlbumKafkaProducer(fileName + "is Deleted")
	} else {
		return "Unable to delete file, File doesn't exist"
	}

	return "File will be deleted"

}
