package main

import (
	"bufio"
	"fmt"
	"lib/file"
	"log/slog"
	"regexp"
	"strconv"
	"strings"
)

const (
	GAME_ID = "game_id"

	CUBE_COUNT = "cube_count"
	CUBE_COLOR = "cube_color"

	GREEN = "green"
	BLUE  = "blue"
	RED   = "red"
)

var gameIDRegex = regexp.MustCompile(fmt.Sprintf("Game (?P<%s>[0-9]+):", GAME_ID))
var groupRegex = regexp.MustCompile(`([^;]+);?`)
var pairRegex = regexp.MustCompile(fmt.Sprintf(`(?P<%s>\d+)\s+(?P<%s>\w+)`, CUBE_COUNT, CUBE_COLOR))

func main() {
	f, err := file.ReadFromFile("../input.txt")
	if err != nil {
		slog.Error("Error reading from file", "error", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	powerOfSets := 0

	for scanner.Scan() {
		gameIDMatches := gameIDRegex.FindAllStringSubmatch(scanner.Text(), -1)

		prefixCut, _ := strings.CutPrefix(scanner.Text(), gameIDMatches[0][0])
		groups := groupRegex.FindAllStringSubmatch(prefixCut, -1)

		var greenMaxCount int
		var blueMaxCount int
		var redMaxCount int

		for _, group := range groups {
			pairs := pairRegex.FindAllStringSubmatch(group[1], -1)
			for _, pair := range pairs {
				switch pair[pairRegex.SubexpIndex(CUBE_COLOR)] {
				case GREEN:
					greenCount, err := strconv.Atoi(pair[pairRegex.SubexpIndex(CUBE_COUNT)])
					if err != nil {
						slog.Error("Error convert string to a number", "error", err)
					}

					if greenMaxCount == 0 {
						greenMaxCount = greenCount
					} else if greenCount > greenMaxCount {
						greenMaxCount = greenCount
					}
				case BLUE:
					blueCount, err := strconv.Atoi(pair[pairRegex.SubexpIndex(CUBE_COUNT)])
					if err != nil {
						slog.Error("Error convert string to a number", "error", err)
					}

					if blueMaxCount == 0 {
						blueMaxCount = blueCount
					} else if blueCount > blueMaxCount {
						blueMaxCount = blueCount
					}

				case RED:
					redCount, err := strconv.Atoi(pair[pairRegex.SubexpIndex(CUBE_COUNT)])
					if err != nil {
						slog.Error("Error convert string to a number", "error", err)
					}

					if redMaxCount == 0 {
						redMaxCount = redCount
					} else if redCount > redMaxCount {
						redMaxCount = redCount
					}
				}

			}
		}

		if greenMaxCount != 0 && redMaxCount != 0 && blueMaxCount != 0 {
			powerOfSets += (greenMaxCount * redMaxCount * blueMaxCount)
		}

	}

	slog.Info("Power of sets", "SUM", powerOfSets) //84911
}
