package controllers

import (
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Home", nil)
}

func GetAllCarsHandler(w http.ResponseWriter, r *http.Request) {

}

func GetCarHandler(w http.ResponseWriter, r *http.Request) {

}

func CreateCarHandler(w http.ResponseWriter, r *http.Request) {

}

func EditCarHandler(w http.ResponseWriter, r *http.Request) {

}

func DeleteCarHandler(w http.ResponseWriter, r *http.Request) {

}
