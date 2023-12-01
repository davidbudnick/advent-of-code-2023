package main

import (
	"bufio"
	"fmt"
	"lib/file"
	"log/slog"
	"regexp"
	"strconv"
)

var regex = regexp.MustCompile(`[1-9]`)

func main() {
	f, err := file.ReadFromFile("../input.txt")
	if err != nil {
		slog.Error("Error reading from file", "error", err)
	}

	var sum int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		matches := regex.FindAllString(scanner.Text(), -1)

		if len(matches) == 0 {
			slog.Error("Matches length is zero")
			return
		}

		s, err := strconv.Atoi(fmt.Sprintf("%s%s", matches[0], matches[len(matches)-1]))
		if err != nil {
			slog.Error("Error converting string to a number", "error", err)
		}

		sum += s
	}

	slog.Info("Sum of all of the valibration values", "sum", sum)
}
