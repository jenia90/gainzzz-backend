package models

import (
	"errors"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

type Workout struct {
	Name string
	Exercises map[string]Sets
	Start time.Time
	End time.Time
}

type Workouts []Workout

func (ws Workouts) AddWorkout(w Workout) Workouts {
	ws = append(ws, w)
	return ws
}

func (ws Workouts) StartWorkout(w Workout) Workout {
	w.Start = time.Now()
	if !slices.ContainsFunc(ws, func(b Workout) bool {
		return strings.Compare(w.Name, b.Name) == 0
	}) {
		ws.AddWorkout(w)
	}

	return w
}

func (ws Workouts) EndWorkout(w Workout) (Workout, error) {
	if !slices.ContainsFunc(ws, func(b Workout) bool {
		return strings.Compare(w.Name, b.Name) == 0
	}) {
		return w, errors.New("unknown workout")
	} else if w.Start.IsZero() {
		return w, errors.New("can't end workout that hasn't been started")
	} else {
		w.End = time.Now()
	}

	return w, nil
}

func (ws Workouts) UpdateWorkout(name string, w Workout) {
	// implemenC
}

func (ws Workouts) DeleteWorkout(name string) {
	// implement
}

func (w Workout) AddExercise(name string) Workout {
	w.Exercises[name] = Sets{}
	return w
}

func (w Workout) DeleteExercise(exerciseName string) Workout {
	delete(w.Exercises, exerciseName)
	return w
}

func (w Workout) AddExerciseSet(exerciseName string, newSet Set) Workout {
	ex := w.Exercises[exerciseName]
	ex = append(ex, newSet)
	w.Exercises[exerciseName] = ex
	return w
}

func (w Workout) DeleteExerciseSet(exerciseName string, setId uint, newSet Set) {
	// implement
}

func (w Workout) UpdateExerciseSet(exerciseName string, setId uint, updatedSet Set) {
	// implement
}