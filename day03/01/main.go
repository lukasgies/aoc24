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
const FIND_MULS = "mul\\(" + EXTRACT_NUMBERS + "," + EXTRACT_NUMBERS + "\\)"

func main() {
	mulSum := 0
	numbers := regexp.MustCompile(EXTRACT_NUMBERS)
	mulInstructions := regexp.MustCompile(FIND_MULS)
	findings := mulInstructions.FindAll([]byte(data), -1) // [[mul(x,y)], [mul(x,y)],...]
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
	fmt.Println("The sum of all multiplikations is: ", mulSum)
}
