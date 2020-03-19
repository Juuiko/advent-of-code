package main

import (
	"fmt"
	"os"
)

func main() {
	var intCodesFile []int
	pointer := 0
	calcOnePointer := 0
	calcTwoPointer := 0
	resultPointer := 0

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	for {
	    if n, _ := fmt.Fscanf(file, "%d,", &pointer); n != 1 {break}
	    intCodesFile = append(intCodesFile,pointer)
	}
	file.Close()
	intCodes := make([]int, len(intCodesFile))

	for i:=0; i<=99; i++ {
	for j:=0; j<=99; j++ {
		copy(intCodes, intCodesFile)
		intCodes[1]=i;
		intCodes[2]=j;
		pointer = 0

		for intCodes[pointer] != 99 {
			if intCodes[pointer] == 1 {
				resultPointer = intCodes[pointer+3]
				if resultPointer <= len(intCodes) {
					calcOnePointer = intCodes[pointer+1]
					calcTwoPointer = intCodes[pointer+2]
					if calcOnePointer <= len(intCodes) && calcTwoPointer <= len(intCodes) {
						intCodes[resultPointer] = intCodes[calcOnePointer] + intCodes[calcTwoPointer]
					}
				}
				pointer += 4
			} else if intCodes[pointer] == 2 {
				resultPointer = intCodes[pointer+3]
				if resultPointer <= len(intCodes) {
					calcOnePointer = intCodes[pointer+1]
					calcTwoPointer = intCodes[pointer+2]
					if calcOnePointer <= len(intCodes) && calcTwoPointer <= len(intCodes) && calcTwoPointer > 0 {
						intCodes[resultPointer] = intCodes[calcOnePointer] * intCodes[calcTwoPointer]
					}
				}
				pointer += 4
			} else {
				pointer++
			}
		}
		if intCodes[0]==19690720 {
			fmt.Println(intCodes)
		}
	}
	}
}
