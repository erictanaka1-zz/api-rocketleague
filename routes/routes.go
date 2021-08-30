package routes

import (
	"github.com/erictanaka1/api-rocketleague/controllers"
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.HomeHandler)
	r.HandleFunc("/cars", controllers.GetAllCarsHandler).Methods("GET")
	r.HandleFunc("/cars/{id}", controllers.GetCarHandler).Methods("GET")
	r.HandleFunc("/cars", controllers.CreateCarHandler).Methods("POST")
	r.HandleFunc("/cars/{id}", controllers.EditCarHandler).Methods("PUT")
	r.HandleFunc("/cars/{id}", controllers.DeleteCarHandler).Methods("DELETE")

	return r
}
