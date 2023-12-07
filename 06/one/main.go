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

	var time []int
	var distance []int

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

		for _, v := range strings.Split(text, " ") {
			for _, num := range NumbersRegex.FindAllString(strings.TrimSpace(v), -1) {
				i, err := strconv.Atoi(num)
				if err != nil {
					slog.Error("Error converting string to a number", "error", err)
				}

				if okTime {
					time = append(time, i)
				}

				if okDistance {
					distance = append(distance, i)
				}
			}
		}

	}

	score := 1
	for timeIndex, t := range time {
		count := 0
		for hold := 1; hold < t; hold++ {
			if hold*(t-hold) > distance[timeIndex] {
				count++
			}
		}
		score = score * count
	}

	slog.Info("What is the score", "SCORE", score) //512295
}
