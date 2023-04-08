package data

import "server/models"

func InitDemoWorkouts() models.Workouts{
	return parseDataFromFile[models.Workouts]("../tests/workouts.sample.json")
}