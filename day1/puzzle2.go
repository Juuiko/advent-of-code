package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	mass, fuel, modFuel := 0, 0, 0

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
		modFuel = (mass / 3) - 2

		fuelMass := modFuel

		for {
			fuelMass = (fuelMass / 3) - 2
			if fuelMass <= 0 {break}
			modFuel += fuelMass
		}

		fuel += modFuel
	}

	fmt.Println(fmt.Sprintf("Total fuel required -> %d", fuel))
}
