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

func GetByteArrSliceFromFile(path string) [][]byte {
	var slice [][]byte
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
		row := scanner.Text()
		var rb []byte
		for _, c := range row {
			rb = append(rb, byte(c))
		}
		slice = append(slice, rb)
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

func RemoveIndex[T comparable](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}

func InsertAtIndex[T comparable](s *[]T, index int, value T) {
	if len(*s) == index {
		*s = append(*s, value)
		return
	}
	*s = append((*s)[:index+1], (*s)[index:]...)
	(*s)[index] = value
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Reverse[T comparable](t []T) []T {
	for i, j := 0, len(t)-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
	return t
}
