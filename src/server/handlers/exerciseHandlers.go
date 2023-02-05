package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"server/models"
	"time"
)

var workouts models.Workouts
var exercises models.Exercises
func InitDemoData() {
	exercises = models.Exercises{
		{Name: "Deadlift", ExerciseType: models.Back, Description: "Deadlift exercise"},
		{Name: "Pull Down", ExerciseType: models.Back, Description: "Lats exercise"},
		{Name: "Biceps Curls", ExerciseType: models.Biceps, Description: "Biceps curls"},
	}

	workouts = models.Workouts{
		{ 
			Name: "Back", 
			Exercises: map[string]models.Sets{
				exercises[0].Name: {
					{Id: 0, Reps: 6, Weight: 30, Units: "kg"},
				},
			},
			Start: time.Now().Add(-time.Hour),
			End: time.Now(),
		},
	}
}

func ExerciseHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
		case http.MethodGet:
			exerciseJson, err := json.Marshal(exercises)
			if err != nil {
				log.Fatal(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			log.Printf("Returning all Exercises to %s", r.Host)
			w.Header().Set("Content-Type", "application/json")
			w.Write(exerciseJson) 
		case http.MethodPost:
			var newEx models.Exercise
			exBytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatal("Failed to read request")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			err = json.Unmarshal(exBytes, &newEx)
			if err != nil {
				log.Fatal("Failed to parse request json")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			log.Printf("Adding exercise Name: %s, Type: %s, Description: %s", newEx.Name, newEx.ExerciseType, newEx.Description)
			exercises = exercises.AddExercise(newEx)
			w.WriteHeader(http.StatusCreated)
			return
	}
}

func WorkoutHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			workoutJson, err := json.Marshal(workouts)
			if err != nil {
				log.Fatal(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Conetent-Type", "application/json")
			w.Write(workoutJson)
	}
}