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

// NON WORKING DRAFT
func main() {
	safeReports := 0
	scanner := bufio.NewScanner(strings.NewReader(data))

	// line by line
	for scanner.Scan() {
		report := convertToIntSlice(strings.Split(scanner.Text(), delimeter))
		numberCache := -1 // init -1 du to no negative values in the reports
		var safeReport bool
		levelRemoved := false

		// entry by entry
		for i := 0; i < len(report); i++ {
			number := report[i]
			if numberCache == -1 {
				numberCache = number
				continue
			}
			diff := math.Abs(float64(number - numberCache))
			if diff < 1 || diff > 3 {
				if levelRemoved {
					safeReport = false
					break
				}
				// remove last elemt before error detection
				report = slices.Delete(report, i, i+1)
				i-- // Important! Decrement index!
				levelRemoved = true
				numberCache = -1 // invalidate cache
				continue
			}
			numberCache = number
			safeReport = true
		}
		if reportNotAscOrDesc(report) {
			safeReport = false
		}
		if safeReport {
			safeReports++
		}
	}

	// output result
	fmt.Println("The number of safe reports is: ", safeReports)
}

func reportNotAscOrDesc(report []int) bool {
	if len(report) < 2 {
		return false // A single element is either ascending or descending
	}

	isAscending := report[0] < report[1]
	compare := func(a, b int) bool {
		if isAscending {
			return b <= a
		}
		return b >= a
	}

	for i := 1; i < len(report); i++ {
		if compare(report[i-1], report[i]) {
			return true
		}
	}
	return false
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
