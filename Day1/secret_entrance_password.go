package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var (
	dialPosition int = 50
)

func dialRotation(input string) (int, int, error) {

	if input == "" {
		return -1, -1, errors.New("input is empty")
	}

	// get the first character of the input string
	direction := string(input[0])

	// separate the second and third character to a variable
	numClicksStr := input[1:]
	numClicks, err := strconv.Atoi(numClicksStr)
	if err != nil {
		return -1, -1, errors.New("invalid number of clicks")
	}
	var numOfZeroPointing int
	if numClicks > 100 {
		numOfZeroPointing = numClicks / 100
		numClicks = numClicks % 100
	}

	var updatedPosition int
	if direction == "L" {
		updatedPosition = dialPosition - numClicks
		if updatedPosition < 0 {
			if dialPosition != 0 && updatedPosition != 0 {
				numOfZeroPointing = numOfZeroPointing + 1
			}

			updatedPosition = 100 + updatedPosition
		}
	} else if direction == "R" {
		updatedPosition = dialPosition + numClicks
		if updatedPosition >= 100 {

			updatedPosition = updatedPosition - 100
			if dialPosition != 0 && updatedPosition != 0 {
				numOfZeroPointing = numOfZeroPointing + 1
			}
		}
	}

	dialPosition = updatedPosition

	fmt.Println("Dial moved", direction, numClicks, "to position", dialPosition, "and zero pointing", numOfZeroPointing)
	return dialPosition, numOfZeroPointing, nil

}

func main() {

	totalZeros := 0
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		position, numOfZeroPointing, err := dialRotation(input)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if position == 0 {
			totalZeros = totalZeros + 1
		}

		totalZeros = totalZeros + numOfZeroPointing
	}

	fmt.Println("password", totalZeros)
	fmt.Println("updated_position", dialPosition)
}
