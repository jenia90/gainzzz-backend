package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"server/data"
	"server/models"
)

var workouts models.Workouts

func InitDemoWorkouts(){
	workouts = data.InitDemoWorkouts()
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
		case http.MethodPost:
			var newWrkt models.Workout
			wrktBytes, err := io.ReadAll(r.Body)
			if err != nil {
				log.Fatal("Failed to read workouts bytes")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			err = json.Unmarshal(wrktBytes, &newWrkt)
			if err != nil {
				log.Fatal("Failed to parse request json")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			log.Printf("Adding new workout: %s", newWrkt.Name)
			workouts = workouts.AddWorkout(newWrkt)
			w.WriteHeader(http.StatusCreated)
	}
}