package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// CreateFile create file
func CreateFile(path string, onFileCreate func(*os.File) error) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	return onFileCreate(file)
}

// WriteStringToFile
func WriteStringToFile(path, content string) error {
	return CreateFile(path, func(file *os.File) error {
		_, err := file.WriteString(content)
		return err
	})
}

// LocalDateStringForFileName
func LocalDateStringForFileName() string {
	return strings.ReplaceAll(time.Now().Format(time.RFC3339), ":", "-")
}

func main() {
	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg)

	WriteStringToFile(LocalDateStringForFileName(), fmt.Sprintf("%+v", argsWithoutProg))
}
