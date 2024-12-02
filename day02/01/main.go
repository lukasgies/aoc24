package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
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
		numberCache := -1 // init -1 du to no negative values in the reports
		var valuesIncreasing bool
		var safeReport bool

		if report[0] < report[1] {
			valuesIncreasing = true
		} else {
			valuesIncreasing = false
		}

		// entry by entry
		for _, number := range report {
			if numberCache == -1 {
				numberCache = number
				continue
			}

			diff := math.Abs(float64(number - numberCache))
			if diff < 1 || diff > 3 {
				safeReport = false
				break
			}
			// check increasing
			if valuesIncreasing && number < numberCache {
				safeReport = false
				break
			}
			// check decreasing
			if !valuesIncreasing && number > numberCache {
				safeReport = false
				break
			}
			numberCache = number
			safeReport = true
		}
		if safeReport {
			safeReports++
		}
	}

	// output result
	fmt.Println("The number of safe reports is: ", safeReports)
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
