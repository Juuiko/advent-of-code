package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	mass, fuel := 0, 0

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err.Error())
		}
		fuel += (mass / 3) - 2
	}

	fmt.Println(fmt.Sprintf("Total fuel required -> %d", fuel))
}
