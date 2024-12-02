package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var data string

var delimeter = ","

func main() {
	scanner := bufio.NewScanner(strings.NewReader(data))
	var leftList []int
	var rightList []int

	// scan each line
	for scanner.Scan() {
		currentLine := strings.Split(scanner.Text(), delimeter)
		l, err := strconv.Atoi(currentLine[0])
		if err != nil {
			fmt.Println("Error no numeric value:", currentLine[0])
		}

		r, err := strconv.Atoi(currentLine[1])
		if err != nil {
			fmt.Println("Error no numeric value:", currentLine[1])
		}

		leftList = append(leftList, l)
		rightList = append(rightList, r)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning input:", err)
	}

	// sort input lists
	sort.Ints(leftList)
	sort.Ints(rightList)

	// print results
	fmt.Println("Total distance: ", calculateDistances(leftList, rightList))
	fmt.Println("Similarity Score: ", findSimilarityScore(leftList, rightList))
}

func calculateDistances(leftList []int, rightList []int) int {
	totalDistance := 0
	if len(leftList) == len(rightList) {
		for i := range leftList {
			distance := math.Abs(float64(leftList[i] - rightList[i]))
			totalDistance += int(distance)
		}
	} else {
		fmt.Println("Error input lists have different lenghts!")
	}
	return totalDistance
}

func findSimilarityScore(leftList []int, rightList []int) int {
	totalSimilarity := 0
	if len(leftList) == len(rightList) {
		for i := range leftList {
			similarity := leftList[i] * rightListAppearances(leftList[i], rightList)
			totalSimilarity += similarity
		}
	} else {
		fmt.Println("Error input lists have different lenghts!")
	}
	return totalSimilarity
}

func rightListAppearances(left int, rightList []int) int {
	count := 0
	for i := range rightList {
		if left == rightList[i] {
			count++
		}
	}
	return count
}
