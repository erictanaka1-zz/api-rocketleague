package main

import (
	"github.com/erictanaka1/api-rocketleague/routes"
	"net/http"
)

func main() {
	router := routes.Routes()
	http.ListenAndServe(":8000", router)
}
