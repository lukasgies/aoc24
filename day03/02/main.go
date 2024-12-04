package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var data string

const EXTRACT_NUMBERS = "\\d{1,3}"
const EXTRACT_DO = "do\\(\\)"
const EXTRACT_DO_NOT = "don't\\(\\)"
const FIND_MULS = "mul\\(" + EXTRACT_NUMBERS + "," + EXTRACT_NUMBERS + "\\)"

func main() {
	mulSum := 0

	// cut out don't()-sections
	dos := regexp.MustCompile(EXTRACT_DO)
	doIndexes := dos.FindAllIndex([]byte(data), -1)
	donts := regexp.MustCompile(EXTRACT_DO_NOT)
	dontIndexes := donts.FindAllIndex([]byte(data), -1)

	instructionSet := data[0:dontIndexes[0][0]]
	to := dontIndexes[0][0]
	for i := range doIndexes {
		from := doIndexes[i][0]
		if from < to {
			continue
		}
		for _, currentDontIndex := range dontIndexes {
			if currentDontIndex[0] > from {
				to = currentDontIndex[0]
				break
			}
			if dontIndexes[len(dontIndexes)-1][1] < from {
				to = len(data) - 1
				break
			}
		}
		if from < to {
			instructionSet = instructionSet + data[from:to]
		}
	}

	// find and addup valid multiplications
	numbers := regexp.MustCompile(EXTRACT_NUMBERS)
	mulInstructions := regexp.MustCompile(FIND_MULS)

	findings := mulInstructions.FindAll([]byte(instructionSet), -1) // [[mul(x,y)], [mul(x,y)],...]
	if findings == nil {
		fmt.Println("No Instruction found! Exiting!")
		return
	}

	for _, mul := range findings {
		factors := numbers.FindAll(mul, -1) // [[x],[y]]
		if len(factors) != 2 {
			fmt.Println("Error multiplikation must be binary!")
			return
		}
		factorX, err := strconv.Atoi(string(factors[0]))
		if err != nil {
			fmt.Println("Error converting factor X to int! ", err)
			return
		}
		factorY, err := strconv.Atoi(string(factors[1]))
		if err != nil {
			fmt.Println("Error converting factor Y to int! ", err)
			return
		}

		mulSum += factorX * factorY
	}
	fmt.Println("The sum of all multiplications is: ", mulSum)
}
