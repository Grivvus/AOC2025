package main

import (
	"fmt"
	"io"
	"math"
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
	inputData = inputData[:len(inputData)-1]

	banks := strings.Split(string(inputData), "\n")
	// res := Solution1(banks)
	res := Solution2(banks)

	fmt.Println(res)
}

func Solution1(banks []string) int {
	sum := 0
	for _, bank := range banks {
		firstMax := 0
		secondMax := 0
		for i, digitStr := range bank {
			digit, err := strconv.Atoi(string(digitStr))
			if err != nil {
				panic(err)
			}
			if i == len(bank)-1 {
				if digit > secondMax {
					secondMax = digit
				}
				break
			}

			if digit > firstMax {
				firstMax = digit
				secondMax = 0
			} else if digit > secondMax {
				secondMax = digit
			}
		}
		localSum := (firstMax * 10) + secondMax
		sum += localSum
	}

	return sum
}

func Solution2(banks []string) uint64 {
	var sum uint64 = 0
	const getLimit = 12
	for _, bank := range banks {
		var prevMax [getLimit]int
		for i, digitStr := range bank {
			digit, err := strconv.Atoi(string(digitStr))
			if err != nil {
				panic(err)
			}
			leftSpace := len(bank) - i
			startIndex := max(0, getLimit-leftSpace)

			for j := startIndex; j < len(prevMax); j++ {
				if digit > prevMax[j] {
					prevMax[j] = digit
					for k := j + 1; k < len(prevMax); k++ {
						prevMax[k] = 0
					}
					break
				}
			}
		}
		localSum := 0
		for i := range getLimit {
			localSum += (prevMax[getLimit-i-1] * int(math.Pow(10.0, float64(i))))
		}
		sum += uint64(localSum)
	}

	return sum
}
