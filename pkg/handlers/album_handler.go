package handlers

import (
	"fmt"
	"net/http"

	model "image-store-service/pkg/model"
	services "image-store-service/pkg/services"
	utilities "image-store-service/pkg/utilities"
)

// swagger:route POST /createAlbum Album CreateAlbum
// It creates an Album
// consumes:
// - multipart/form-data
// responses:
// 200:
//    AlbumResponse

//CreateAlbumHandler Creates an Album
func CreateAlbumHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the album nam
		folder := r.FormValue("albumName")
		// Create the album if it doesn't exist
		serviceResponse := services.CreateDirectory(folder)
		response := model.AlbumResponse{
			Response: fmt.Sprintf("%s", serviceResponse),
		}
		//Return the JSON response
		utilities.JSONResponse(w, 200, response)
	})
}

// swagger:route POST /deleteAlbum Album DeleteAlbum
// It deletes an Album
// consumes:
// - multipart/form-data
// responses:
// 200: AlbumResponse

// DeleteAlbumHandler Deletes an Album
func DeleteAlbumHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the album name
		folder := r.FormValue("albumName")
		// Delete the album if it exists
		serviceResponse := services.DeleteDirectory(folder)
		response := model.AlbumResponse{
			Response: fmt.Sprintf("%s", serviceResponse),
		}
		//Return the JSON response
		utilities.JSONResponse(w, 200, response)
	})
}
