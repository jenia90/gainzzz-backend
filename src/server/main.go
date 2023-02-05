package main

import (
	"fmt"
	"log"
	"net/http"
	"server/handlers"
)

func main() {
	handlers.InitDemoData()
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/exercises", handlers.ExerciseHandler)
	mux.HandleFunc("/api/v1/workouts", handlers.WorkoutHandler)

	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}