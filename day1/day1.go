package day1

import (
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func SumCalibrationValues(filename string) (sum int) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		sum += calculateCalibrationValue(line)
	}
	return sum
}

func calculateCalibrationValue(line string) (value int) {
	firstValue := findFirstNumberAsString(line)
	secondValue := findFirstNumberAsString(reverse(line))
	calibrationValueString := firstValue + secondValue
	calibrationValue, err := strconv.Atoi(calibrationValueString)

	if err != nil {
		panic("Unable to convert to number!")
	}

	return calibrationValue
}

func findFirstNumberAsString(s string) string {
	for _, char := range s {
		if unicode.IsDigit(char) {
			return string(char)
		}
	}
	panic("Did not find any numbers!")
}

func reverse(s string) string {
	runes := []rune(s)

	length := len(runes)

	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}

	return string(runes)
}
