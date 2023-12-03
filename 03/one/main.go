package main

import (
	"bufio"
	"lib/file"
	"log/slog"
	"strconv"
	"strings"
)

func main() {
	f, err := file.ReadFromFile("../input.txt")
	if err != nil {
		slog.Error("Error reading from file", "error", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var parts [][]string
	for scanner.Scan() {
		var row []string
		for _, item := range scanner.Text() {
			row = append(row, string(item))
		}
		parts = append(parts, row)
	}

	var sum int

	for rowIndex, row := range parts {
		var activeCountSum bool
		var activeNums []string
		for itemIndex, item := range row {

			ok := isInteger(item)
			if ok {

				if scan(rowIndex, itemIndex-1, parts) {
					activeCountSum = true
				} else if scan(rowIndex, itemIndex+1, parts) {
					activeCountSum = true
				} else if scan(rowIndex-1, itemIndex, parts) {
					activeCountSum = true
				} else if scan(rowIndex-1, itemIndex-1, parts) {
					activeCountSum = true
				} else if scan(rowIndex-1, itemIndex+1, parts) {
					activeCountSum = true
				} else if scan(rowIndex+1, itemIndex, parts) {
					activeCountSum = true
				} else if scan(rowIndex+1, itemIndex-1, parts) {
					activeCountSum = true
				} else if scan(rowIndex+1, itemIndex+1, parts) {
					activeCountSum = true
				}

				activeNums = append(activeNums, item)

			}

			//reset if num does not find symbol
			if !activeCountSum && len(activeNums) != 0 && stopCheck(item, itemIndex) {
				activeNums = []string{}
			}

			if activeCountSum && len(activeNums) != 0 && stopCheck(item, itemIndex) {
				cn, err := strconv.Atoi(strings.Join(activeNums, ""))
				if err != nil {
					slog.Error("Error converting string to number", "error", err)
				}

				activeCountSum = false
				activeNums = []string{}

				sum += cn
			}

		}
	}

	slog.Info("Total part number ", "SUM", sum) //527446

}

func isInteger(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func scan(x int, y int, parts [][]string) bool {
	if x < 0 || y < 0 || x == 140 || y == 140 {
		return false
	}

	return isSymbol(parts[x][y])

}

func isSymbol(s string) bool {
	symbols := []string{
		"*",
		"-",
		"+",
		"$",
		"@",
		"#",
		"%",
		"&",
		"/",
		"=",
	}

	for _, v := range symbols {
		if s == v {
			return true
		}
	}

	return false
}

func stopCheck(s string, index int) bool {
	return s == "." || index == 139 || isSymbol(s)
}
