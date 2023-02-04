package handlers

import (
	"io"
	"net/http"
)

func ExerciseHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world from exercise Handler\n")
}