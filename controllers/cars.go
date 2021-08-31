package controllers

import (
	"encoding/json"
	"github.com/erictanaka1/api-rocketleague/db"
	"github.com/erictanaka1/api-rocketleague/models"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))
var mock = db.MockData()

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Home", nil)
}

func GetAllCarsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mock)
}

func GetCarHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range mock {
		if strconv.Itoa(item.Id) == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Car{})
}

func CreateCarHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var car models.Car
	_ = json.NewDecoder(r.Body).Decode(&car)
	car.Id = len(mock) + 1
	mock = append(mock, car)
	json.NewEncoder(w).Encode(car)
}

func EditCarHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range mock {
		if strconv.Itoa(item.Id) == params["id"] {
			var car models.Car
			car.Id = item.Id
			_ = json.NewDecoder(r.Body).Decode(&car)
			mock[index].Name = car.Name
			mock[index].Desc = car.Desc
			json.NewEncoder(w).Encode(car)
			return
		}
	}
	json.NewEncoder(w).Encode(mock)
}

func DeleteCarHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range mock {
		if strconv.Itoa(item.Id) == params["id"] {
			mock = append(mock[:index], mock[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(mock)
}
