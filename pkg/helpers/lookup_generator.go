package helpers

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	// change size when you need to generate more square lookups
	size = 36
)

func outputLookup() {
	fmt.Println(formatSlice(generateSquareLookup(size)))
}

func generateSquareLookup(n int) [][]int {
	lookup := make([][]int, n)

	root := int(math.Sqrt(float64(n)))
	var idx, baseIdx = 0, 0
	for i := 0; i < n*n; i++ {

		lookup[idx] = append(lookup[idx], i)

		// for every multiple of root, move to the next index
		if (i+1)%root == 0 {
			idx++
		}

		if (i+1)%(n*root) == 0 {
			baseIdx += root
		}

		if (i+1)%n == 0 {
			idx = baseIdx
		}
	}

	return lookup
}

func formatSlice(s [][]int) []string {
	var l []string

	for _, v := range s {
		a := []string{}
		for _, i := range v {
			a = append(a, strconv.Itoa(i))
		}

		b := strings.Join(a, ",")
		l = append(l, fmt.Sprintf("{%s},", b))
	}

	return l
}
