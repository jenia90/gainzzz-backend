package data

import "server/models"

func InitDemoExercises() models.Exercises{
	return parseDataFromFile[models.Exercises]("../tests/exercises.sample.json")
}