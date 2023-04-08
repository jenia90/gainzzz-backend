package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"server/data"
	"server/models"
)

var exercises models.Exercises

func InitDemoExercises() {
	exercises = data.InitDemoExercises()
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
			exBytes, err := io.ReadAll(r.Body)
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