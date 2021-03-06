// Album Management for the System
//
// Documentation for the Image Management System
//
// Schemees :http
// BasePath: /
// Version: 1.0.0
// Consumes:
//  -multipart/form-data;
// Produces:
// -application/json
// swagger:meta

package main

import (
	"log"
	"net/http"

	handler "image-store-service/pkg/handlers"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := mux.NewRouter()
	log.Println("Server has been started")

	// handlers for API
	r.Handle("/createAlbum", handler.CreateAlbumHandler()).Methods("POST")
	r.Handle("/deleteAlbum", handler.DeleteAlbumHandler()).Methods("POST")
	r.Handle("/createImage", handler.CreateImageHandler()).Methods("POST")
	r.Handle("/deleteImage", handler.DeleteImageHandler()).Methods("POST")
	r.Handle("/getImage", handler.GetImageHandler()).Methods("GET")

	r.Handle("/swagger.yaml", http.FileServer(http.Dir("../"))).Methods("GET")
	http.Handle("/metrics", promhttp.Handler())
	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	r.Handle("/docs", sh).Methods("GET")

	http.ListenAndServe(":8080", r)

}
