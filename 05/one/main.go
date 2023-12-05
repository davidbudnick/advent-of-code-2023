package main

import (
	"bufio"
	"fmt"
	"lib/file"
	"log/slog"
)

//? List all of the seeds which need to be planted

//? seed-to-soil
//? soil-to-fertilizer
//? fertilizer-to-water
//? water-to-light
//? light-to-temperature
//? temperature-to-humidity
//? humidity-to-location

//? Every seed is idenified with a number
//? Numbers CAN repeat for reach area

//? First list what seeds need to be planted
//? list of maps which describe how to convert numbers from a source category into numbers in a destination category

//? range of numbers

//? seed-to-soil map:
//? [50] (destination range start) [98] (source range start) [2] (range length)

//? Range [2]
//? Source 98 -> Destination -> 50
//? Source 99 -> Destination -> 51w

func main() {
	f, err := file.ReadFromFile("../input.txt")
	if err != nil {
		slog.Error("Error reading from file", "error", err)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
