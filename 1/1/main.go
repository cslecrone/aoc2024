package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("../input")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Arrays would be more efficient than slices
	first := []int{}
	second := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		fs := strings.Fields(line)
		c1, _ := strconv.Atoi(fs[0])
		c2, _ := strconv.Atoi(fs[1])
		// TODO: Insert in order
		first = append(first, c1)
		second = append(second, c2)
	}

	// It would be more efficient to insert in order, but I'm being lazy
	sort.Ints(first)
	sort.Ints(second)

	tot := 0

	for idx := range first {
		diff := first[idx] - second[idx]
		// math.Abs() does not support integers AFAICT
		if diff < 0 {
			diff = -diff
		}
		tot += diff
	}

	fmt.Println(tot)
}
