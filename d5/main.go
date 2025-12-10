package main

import (
	"cmp"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(f)
	b = b[:len(b)-1]
	splited := strings.Split(string(b), "\n\n")
	rangesStr := strings.Split(splited[0], "\n")
	idsStr := strings.Split(splited[1], "\n")

	ranges := make([][2]int, len(rangesStr))
	ids := make([]int, len(idsStr))

	for i, r := range rangesStr {
		rangeSplited := strings.Split(r, "-")
		rangeStart, err := strconv.Atoi(rangeSplited[0])
		if err != nil {
			panic(err)
		}
		rangeEnd, err := strconv.Atoi(rangeSplited[1])
		if err != nil {
			panic(err)
		}
		ranges[i] = [2]int{rangeStart, rangeEnd}
	}

	for i, id := range idsStr {
		idNum, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}

		ids[i] = idNum
	}

	// res := Solution(ranges, ids)
	res := Solution2(ranges)

	fmt.Println(res)
}

func Solution(ranges [][2]int, ids []int) int {
	cnt := 0

	for _, id := range ids {
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				cnt++
				break
			}
		}
	}
	return cnt
}

func Solution2(ranges [][2]int) int {
	cnt := 0
	nonOverlapingRagnes := make([][2]int, 0)
	slices.SortFunc(ranges, func(a, b [2]int) int {
		return cmp.Compare(a[0], b[0])
	})

	lb, rb := ranges[0][0], ranges[0][1]

	for _, r := range ranges {
		if r[0] <= rb {
			rb = max(rb, r[1])
		} else {
			nonOverlapingRagnes = append(nonOverlapingRagnes, [2]int{lb, rb})
			lb, rb = r[0], r[1]
		}
	}

	nonOverlapingRagnes = append(nonOverlapingRagnes, [2]int{lb, rb})

	for _, r := range nonOverlapingRagnes {
		cnt += (r[1] - r[0] + 1)
	}

	return cnt
}
