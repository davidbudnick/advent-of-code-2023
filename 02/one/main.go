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

	sumGameIDs := 0

	for scanner.Scan() {
		gameIDMatches := gameIDRegex.FindAllStringSubmatch(scanner.Text(), -1)

		gameID, err := strconv.Atoi(gameIDMatches[0][gameIDRegex.SubexpIndex(GAME_ID)])
		if err != nil {
			slog.Error("Error convert string to a number", "error", err)
		}

		prefixCut, _ := strings.CutPrefix(scanner.Text(), gameIDMatches[0][0])
		groups := groupRegex.FindAllStringSubmatch(prefixCut, -1)

		possible := true

		for _, group := range groups {
			pairs := pairRegex.FindAllStringSubmatch(group[1], -1)
			for _, pair := range pairs {
				var greenCount int
				var blueCount int
				var redCount int

				switch pair[pairRegex.SubexpIndex(CUBE_COLOR)] {
				case GREEN:
					i, err := strconv.Atoi(pair[pairRegex.SubexpIndex(CUBE_COUNT)])
					if err != nil {
						slog.Error("Error convert string to a number", "error", err)
					}
					greenCount = i

				case BLUE:
					i, err := strconv.Atoi(pair[pairRegex.SubexpIndex(CUBE_COUNT)])
					if err != nil {
						slog.Error("Error convert string to a number", "error", err)
					}
					blueCount = i

				case RED:
					i, err := strconv.Atoi(pair[pairRegex.SubexpIndex(CUBE_COUNT)])
					if err != nil {
						slog.Error("Error convert string to a number", "error", err)
					}
					redCount = i
				}

				if redCount > 12 || greenCount > 13 || blueCount > 14 {
					possible = false
				}

			}
		}

		if possible {
			sumGameIDs += gameID
		}

	}

	slog.Info("sum of the IDs", "SUM", sumGameIDs) //2006
}
