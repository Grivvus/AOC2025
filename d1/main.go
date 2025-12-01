package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	inputData, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	const startPosition = 50

	sequence := strings.Split(string(inputData), "\n")
	// last element is always empty
	sequence = sequence[:len(sequence)-1]

	// answer := Part1Solution(sequence, startPosition)
	answer := Part2Solution(sequence, startPosition)

	fmt.Println(answer)
}

func Part1Solution(sequence []string, currentPosition int) int {
	pointsAt0Times := 0
	for _, move := range sequence {
		var preparedMove string
		if move[0] == 'L' {
			preparedMove = strings.Replace(move, "L", "-", 1)
		} else {
			preparedMove = move[1:]
		}

		num, err := strconv.Atoi(preparedMove)
		if err != nil {
			panic(err)
		}

		currentPosition += num

		for currentPosition < 0 {
			currentPosition += 100
		}
		for currentPosition >= 100 {
			currentPosition -= 100
		}

		if currentPosition == 0 {
			pointsAt0Times++
		}
	}

	return pointsAt0Times
}

func Part2Solution(sequence []string, currentPosition int) int {
	pointsAt0Times := 0
	for _, move := range sequence {
		var preparedMove string
		if move[0] == 'L' {
			preparedMove = strings.Replace(move, "L", "-", 1)
		} else {
			preparedMove = move[1:]
		}

		num, err := strconv.Atoi(preparedMove)
		if err != nil {
			panic(err)
		}

		fullCircles := Abs(num) / 100
		var reminder int
		if num >= 0 {
			reminder %= 100
		} else {
			reminder = num + fullCircles*100
		}

		pointsAt0Times += fullCircles

		currentPosition += reminder
		if currentPosition < 0 || currentPosition >= 100 {
			reminder++
			currentPosition = ((currentPosition % 100) + 100) % 100
		}
	}

	return pointsAt0Times
}

func Abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}
