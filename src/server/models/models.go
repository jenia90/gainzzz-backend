package models

import (
	"time"
)

type ExerciseType string
const (
	Abs 		= "abs"
	Legs 		= "legs"
	Back 		= "back"
	Chest 		= "chest"
	Biceps 		= "biceps"
	Triceps 	= "triceps"
	Shoulders 	= "shoulders"
)

type Set struct {
	Id uint
	Reps uint
	Weight uint
	Units string
}
type Sets []Set

type Exercise struct {
	Name string
	ExerciseType ExerciseType 
	Description string
}
type Exercises []Exercise
type Workout struct {
	Name string
	Exercises map[string]Sets
	Start time.Time
	End time.Time
}
type Workouts []Workout

func (e Exercises) AddExercise(ex Exercise) Exercises {
	return append(e, ex)
}

func (e Exercises) UpdateExercise(name string, ex Exercise) {
	// implement
}

func (e Exercises) DeleteExercise(name string) {
	// implement
}

func (ws Workouts) AddWorkout(w Workout) {
	// implement
}

func (ws Workouts) UpdateWorkout(name string, w Workout) {
	// implement
}

func (ws Workouts) DeleteWorkout(name string) {
	// implement
}

func (w Workout) AddExercise(e Exercise) {
	// implement
}

func (w Workout) UpdateExercise(exerciseName string, e Exercise) {
	// implement
}

func (w Workout) AddExerciseSet(exerciseName string, newSet Set) {
	//implement
}

func (w Workout) DeleteExerciseSet(exerciseName string, setId uint, newSet Set) {
	// implement
}

func (w Workout) UpdateExerciseSet(exerciseName string, setId uint, updatedSet Set) {
	// implement
}