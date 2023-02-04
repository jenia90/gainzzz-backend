package models

import (
	"strings"
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
	id uint
	reps uint
	weight uint
	units string
}
type Sets []Set

type Exercise struct {
	name string
	exerciseType ExerciseType 
	description string
}
type Exercises []Exercise
var exercises = Exercises{
	{name: "deadlift", exerciseType: Back, description: "keep your back straight!!" },
}

type Workout struct {
	name string
	exercises map[Exercise]Sets
	start time.Time
	end time.Time
}
type Workouts []Workout

func (this Exercise) compare(other Exercise) int {
	return strings.Compare(this.name, other.name)
}

func (e Exercises) AddExercise(ex Exercise) {
	e = append(e, ex)
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