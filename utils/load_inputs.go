package utils

import (
	"fmt"
	"os"
)

func LoadFile(day, year int) string {
	filePath := fmt.Sprintf("%d/input_files/input_day%d.txt", year, day)
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return ""
	}
	return string(fileContent)
}
