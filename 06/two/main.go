package main

import (
	"bufio"
	"lib/file"
	"log/slog"
	"regexp"
	"strconv"
	"strings"
)

var NumbersRegex = regexp.MustCompile(`(?m)(\d{4}|\d{3}|\d{2}|\d{1})`)

func main() {

	f, err := file.ReadFromFile("../input.txt")
	if err != nil {
		slog.Error("Error reading from file", "error", err)
	}

	scanner := bufio.NewScanner(f)

	var time int
	var distance int

	for scanner.Scan() {
		var text string

		timeText, okTime := strings.CutPrefix(scanner.Text(), "Time:")
		if okTime {
			text = timeText
		}

		distanceText, okDistance := strings.CutPrefix(scanner.Text(), "Distance:")
		if okDistance {
			text = distanceText
		}

		if okTime {
			time, err = strconv.Atoi(strings.ReplaceAll(text, " ", ""))
			if err != nil {
				slog.Error("Error converting string to a number", "error", err)
			}
		}

		if okDistance {
			distance, err = strconv.Atoi(strings.ReplaceAll(text, " ", ""))
			if err != nil {
				slog.Error("Error converting string to a number", "error", err)
			}
		}

	}

	var count int
	for hold := 1; hold < time; hold++ {
		if hold*(time-hold) > distance {
			count++
		}
	}

	slog.Info("What is the count", "COUNT", count)
}
