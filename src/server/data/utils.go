package data

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"server/models"
)

func parseDataFromFile[T models.Workouts|models.Exercises](path string) (T) {
	var data T
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatal("Unable to parse abs path of", path)
	}
	readBytes, err := os.ReadFile(absPath)
	if err != nil {
		log.Fatal("Unable to read ", absPath)
		log.Fatal(err)
	}
	
	err = json.Unmarshal(readBytes, &data)
	if err != nil {
		log.Fatal("Failed to parse json from ", absPath)
	}

	return data
}