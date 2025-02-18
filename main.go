package main

import (
	"fmt"
	"os"
	"os/exec"
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

// func remove(slice []string, s int) []string {
// 	return append(slice[:s], slice[s+1:]...)
// }

func main() {
	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg)

	newargs := []string{}

	for i := range argsWithoutProg {
		if i == 5 || i == 6 || i == 13 || i == 14 || i == 15 || i == 16 {
			continue
		}
		newargs = append(newargs, argsWithoutProg[i])
	}

	WriteStringToFile(LocalDateStringForFileName(), fmt.Sprintf("%+v %+v", argsWithoutProg, newargs))

	exec.Command("ffmpeg_real.exe", newargs...).Run()
}
