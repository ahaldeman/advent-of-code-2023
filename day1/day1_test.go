package day1

import (
	"fmt"
	"testing"
)

func TestDay1(t *testing.T) {
	t.Run("should calculate the calibration values for a simple input", func(t *testing.T) {
		got := SumCalibrationValues("day1_input_simple.txt")
		want := 142

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("calculate for the days puzzle input", func(t *testing.T) {
		fmt.Println(SumCalibrationValues("day1_input.txt"))
	})
}
