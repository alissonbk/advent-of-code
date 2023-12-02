package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func GetSliceFromFile(path string) []string {
	var slice []string
	dir, _ := os.Getwd()
	if !strings.Contains(dir, "2023") {
		dir += "/2023/"
	}
	file, err := os.Open(dir + path)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("Could not close fd, msg: " + err.Error())
		}
	}(file)
	if err != nil {
		log.Fatal(err.Error())
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return slice
}

func Contains[T comparable](slice []T, n T) bool {
	for _, v := range slice {
		if v == n {
			return true
		}
	}
	return false
}
