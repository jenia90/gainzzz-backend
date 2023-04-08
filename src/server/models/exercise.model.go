package models

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

type Exercise struct {
	Name string
	ExerciseType ExerciseType 
	Description string
}
type Exercises []Exercise

func (e Exercises) AddExercise(ex Exercise) Exercises {
	return append(e, ex)
}

func (e Exercises) UpdateExercise(name string, ex Exercise) {
	// implement
}

func (e Exercises) DeleteExercise(name string) {
	// implement
}
