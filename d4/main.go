package main

import (
	"fmt"
	"io"
	"os"
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

	rows := strings.Split(string(inputData), "\n")

	// res := Solution(rows)
	res := Solution2(rows, 0)

	fmt.Println(res)
}

func Solution(rows []string) int {
	cnt := 0

	moves := [...][2]int{{-1, -1}, {-1, 0}, {0, -1}, {1, 0}, {0, 1}, {1, 1}, {1, -1}, {-1, 1}}

	for i := range rows {
		for j := range rows[i] {
			neighbours := 0
			for _, m := range moves {
				ni, nj := i+m[0], j+m[1]
				if !inBounds(rows, ni, nj) {
					continue
				}
				if rows[ni][nj] == '@' {
					neighbours++
				}
			}
			if neighbours < 4 && rows[i][j] == '@' {
				cnt++
			}
		}
	}

	return cnt
}

func Solution2(rows []string, prevCnt int) int {
	cnt := 0
	moves := [...][2]int{{-1, -1}, {-1, 0}, {0, -1}, {1, 0}, {0, 1}, {1, 1}, {1, -1}, {-1, 1}}

	nextRows := make([]string, len(rows))

	for i := range rows {
		var b strings.Builder
		for j := range rows[i] {
			neighbours := 0
			for _, m := range moves {
				ni, nj := i+m[0], j+m[1]
				if !inBounds(rows, ni, nj) {
					continue
				}
				if rows[ni][nj] == '@' {
					neighbours++
				}
			}
			if neighbours < 4 && rows[i][j] == '@' {
				cnt++
				b.WriteByte('.')
			} else {
				b.WriteByte(rows[i][j])
			}
		}
		nextRows[i] = b.String()
	}
	if cnt == 0 {
		return prevCnt
	}
	return Solution2(nextRows, cnt+prevCnt)
}

func inBounds(rows []string, x, y int) bool {
	return (x >= 0 && x < len(rows)) && (y >= 0 && y < len(rows[0]))
}
