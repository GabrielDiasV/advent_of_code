package utils

import (
	"fmt"
	"os"
)

func LoadFile(day int) string {
	filePath := fmt.Sprintf("input_files/input_day%d.txt", day)
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return ""
	}
	return string(fileContent)
}
