package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("modulemass.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	accumulator := 0
	accumulatorRecursive := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		accumulator += fuelRequirementByMass(mass)
		accumulatorRecursive += fuelRequirementByMassRecursive(mass)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Print("Fuel needed without taking account of fuel's fuel: ")
	fmt.Println(accumulator)

	fmt.Print("Fuel needed taking account of fuel's fuel: ")
	fmt.Println(accumulatorRecursive)
}

func fuelRequirementByMass(mass int) int {
	return int(math.Floor(float64(mass)/3.0) - 2.0)
}

func fuelRequirementByMassRecursive(mass int) int {
	if fuelRequirementByMass(mass) <= 0 {
		return 0
	}
	return fuelRequirementByMass(mass) + fuelRequirementByMassRecursive(fuelRequirementByMass(mass))
}
