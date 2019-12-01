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
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		accumulator += fuelRequirementByMass(mass)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(accumulator)
}

func fuelRequirementByMass(mass int) int {
	return int(math.Floor(float64(mass)/3.0) - 2.0)
}
