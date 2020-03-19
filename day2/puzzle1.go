package main

import (
	"fmt"
	"os"
)

func main() {
	var intCodes []int
	pointer := 0
	calcOnePointer := 0
	calcTwoPointer := 0
	resultPointer := 0

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	for {
	    if n, _ := fmt.Fscanf(file, "%d,", &pointer); n != 1 {break}
	    intCodes = append(intCodes,pointer)
	}

	pointer = 0

	for intCodes[pointer] != 99 {
		if intCodes[pointer] == 1 {
			resultPointer = intCodes[pointer+3]
			if resultPointer <= len(intCodes) {
				calcOnePointer = intCodes[pointer+1]
				calcTwoPointer = intCodes[pointer+2]
				intCodes[resultPointer] = intCodes[calcOnePointer] + intCodes[calcTwoPointer]
			}
			pointer += 4
		} else if intCodes[pointer] == 2 {
			resultPointer = intCodes[pointer+3]
			if resultPointer <= len(intCodes) {
				calcOnePointer = intCodes[pointer+1]
				calcTwoPointer = intCodes[pointer+2]
				intCodes[resultPointer] = intCodes[calcOnePointer] * intCodes[calcTwoPointer]
			}
			pointer += 4
		} else {
			pointer++
		}
	}
	fmt.Println(intCodes)
}
