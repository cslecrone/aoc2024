package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("../input")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	simMap := make(map[int]int)
	first := []int{}
	second := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		fs := strings.Fields(line)
		c1, _ := strconv.Atoi(fs[0])
		c2, _ := strconv.Atoi(fs[1])
		simMap[c1] = 0
		first = append(first, c1)
		second = append(second, c2)
	}

	for _, val := range second {
		simMap[val] += 1
	}

	tot := 0

	for _, val := range first {
		tot += val * simMap[val]
	}

	fmt.Println(tot)
}
