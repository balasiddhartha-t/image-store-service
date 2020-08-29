package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	model "image-store-service/pkg/model"
	services "image-store-service/pkg/services"
	utilities "image-store-service/pkg/utilities"
)

// swagger:route POST /createImage Image CreateImage
// Creates an image in the album

// consumes:
// - multipart/form-data
// responses:
// 200: AlbumResponse
// CreateImageHandler Creates the image into an Album
// Response back to the system
//
// CreateImage has albumName and file
func CreateImageHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var imageResponse string

		//read the album name from the form
		folder := r.FormValue("albumName")
		if folder == "" {
			imageResponse = "Folder name cannot be Empty"
			return
		}

		//read the file name from the form
		file, handler, error := r.FormFile("file")
		if error != nil {
			imageResponse = "File cannot be Empty"
			fmt.Println(error)
			return
		}

		//Get the type of the file
		fileType, err := utilities.GetFileContentType(file)
		if err != nil {
			imageResponse = "Error Resolving the type of file uploaded"
		}

		defer file.Close()

		//Persist if the type of the file is image
		if strings.Contains(fileType, "image") {
			imageResponse = services.CreateImage(file, folder, handler.Filename)
		} else {
			imageResponse = "Provided file type is not an Image File"
		}
		response := model.AlbumResponse{
			Response: fmt.Sprintf("%s", imageResponse),
		}

		//return the response back
		utilities.JSONResponse(w, 200, response)

	})
}

//swagger:route POST /deleteImage Image DeleteImage
//Deletes an image in the album
// responses:
// 200:AlbumResponse
// consumes:
// - multipart/form-data

//DeleteImageHandler Deletes the image into an Album

func DeleteImageHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var imageResponse string

		//Get the album name from the form
		folder := r.FormValue("albumName")
		if folder == "" {
			imageResponse = "Folder name cannot be Empty"
			return
		}

		//Get the file name from the form
		fileName := r.FormValue("fileName")
		if folder == "" {
			imageResponse = "File name cannot be Empty"
			return
		}
		//Delete the Image if it is present in the directory
		imageResponse = services.DeleteImage(folder, fileName)
		response := model.AlbumResponse{
			Response: fmt.Sprintf("%s", imageResponse),
		}
		//return the response back
		utilities.JSONResponse(w, 200, response)
	})
}

// swagger:route GET /getImage Image GetImage
// Get an image in the album

//DeleteImageHandler Deletes the image into an Album

func GetImageHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var imageResponse string
		//Read the Album name from the form
		folder := r.FormValue("albumName")
		if folder == "" {
			imageResponse = "Folder name cannot be Empty"
			fmt.Printf(imageResponse)
			return
		}
		//Read the Image name from the form
		fileName := r.FormValue("fileName")
		if folder == "" {
			imageResponse = "File name cannot be Empty"
			fmt.Printf(imageResponse)
			return
		}
		//Get the current working directory
		workingDir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		//Serve the file back to the server
		http.ServeFile(w, r, filepath.Join(workingDir, "..", "Albums/"+folder+"/"+fileName))
	})
}
