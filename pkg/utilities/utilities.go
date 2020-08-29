package utilities

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
)

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func JSONResponse(w http.ResponseWriter, code int, output interface{}) {
	// Convert our interface to JSON
	response, _ := json.Marshal(output)
	// Set the content type to json for browsers
	w.Header().Set("Content-Type", "application/json")
	// Our response code
	w.WriteHeader(code)

	w.Write(response)
}

func GetFileContentType(out multipart.File) (string, error) {

	// Create a buffer to store the header of the file in
	fileHeader := make([]byte, 512)

	// Copy the headers into the FileHeader buffer
	if _, err := out.Read(fileHeader); err != nil {
		fmt.Println(err)
	}

	// set position back to start.
	if _, err := out.Seek(0, 0); err != nil {
		fmt.Println(err)

	}

	fileType := http.DetectContentType(fileHeader)

	return fileType, nil
}
