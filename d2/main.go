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
	inputData = inputData[:len(inputData)-1]

	ranges := strings.Split(string(inputData), ",")
	res := Solution(ranges)

	fmt.Println(res)
}

func Solution(ranges []string) uint64 {
	var sum uint64 = 0
	for _, r := range ranges {
		s := strings.Split(r, "-")
		startStr := s[0]
		endStr := s[1]

		start, err := strconv.Atoi(startStr)
		if err != nil {
			panic(err)
		}

		end, err := strconv.Atoi(endStr)
		if err != nil {
			panic(err)
		}
		for i := start; i <= end; i++ {
			/* first part */
			// if IsDoubled(strconv.Itoa(i)) {
			// 	sum += uint64(i)
			// }

			/* second part */
			if IsRepeated(strconv.Itoa(i)) {
				sum += uint64(i)
			}
		}
	}

	return sum
}

func IsDoubled(id string) bool {
	n := len(id)
	if n%2 == 1 {
		return false
	}
	halfN := n / 2
	return id[:halfN] == id[halfN:]
}

func IsRepeated(id string) bool {
	n := len(id)

nextStep:
	for step := 1; step <= n/2; step++ {
		if n%step != 0 {
			continue
		}
		firstPart := id[:step]
		for i := 1; (i+1)*step <= n; i++ {
			startBound := i * step
			endBound := (i + 1) * step
			if firstPart != id[startBound:endBound] {
				continue nextStep
			}
		}
		return true
	}
	return false
}
