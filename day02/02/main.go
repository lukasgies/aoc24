package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var data string

var delimeter = " "

func main() {
	safeReports := 0
	scanner := bufio.NewScanner(strings.NewReader(data))

	// line by line
	for scanner.Scan() {
		report := convertToIntSlice(strings.Split(scanner.Text(), delimeter))

		if ascOrDesc(report) && correctLevelDistance(report) {
			safeReports++
			continue
		}
		for i := 0; i < len(report); i++ {
			permutation := slices.Clone(report)
			permutation = slices.Delete(permutation, i, i+1)
			if ascOrDesc(permutation) && correctLevelDistance(permutation) {
				safeReports++
				break
			}
		}
	}

	// output result
	fmt.Println("The number of safe reports is: ", safeReports)
}

func ascOrDesc(report []int) bool {
	if len(report) < 2 {
		return true // A single element is either ascending or descending
	}

	isAscending := report[0] < report[1]
	compare := func(a, b int) bool {
		if isAscending {
			return a <= b
		}
		return a >= b
	}

	for i := 1; i < len(report); i++ {
		if !compare(report[i-1], report[i]) {
			return false
		}
	}
	return true
}

func correctLevelDistance(report []int) bool {
	cache := -1
	for _, number := range report {
		if cache == -1 {
			cache = number
			continue
		}
		diff := math.Abs(float64(number - cache))
		if diff < 1 || diff > 3 {
			return false
		}
		cache = number
	}
	return true
}

func convertToIntSlice(data []string) []int {
	res := make([]int, len(data))
	for i, s := range data {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Error parsing input data to int: ", err)
		}
		res[i] = n
	}
	return res
}
