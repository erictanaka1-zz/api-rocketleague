package db

import "github.com/erictanaka1/api-rocketleague/models"

// Mock data for testing

func MockData() []models.Car {
	var cars = []models.Car{}

	cars = append(cars, models.Car{Id: 1, Name: "Octane", Desc: "Descricao teste"})
	cars = append(cars, models.Car{Id: 2, Name: "Dominus", Desc: "Descricao teste 2"})
	cars = append(cars, models.Car{Id: 3, Name: "Merc", Desc: "Descricao teste 3"})
	cars = append(cars, models.Car{Id: 3, Name: "Breakout", Desc: "Descricao teste 4"})

	return cars
}
