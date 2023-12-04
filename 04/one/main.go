package main

import (
	"bufio"
	"lib/file"
	"log/slog"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var CardIDRegex = regexp.MustCompile(`(?m)Card\s+\d{1,3}:`)
var NumbersRegex = regexp.MustCompile(`(?m)(\d{2}|\d{1})`)

func main() {
	f, err := file.ReadFromFile("../input.txt")
	if err != nil {
		slog.Error("Error reading from file", "error", err)
	}

	scanner := bufio.NewScanner(f)

	var points int

	for scanner.Scan() {

		cardIDMatches := CardIDRegex.FindAllStringSubmatch(scanner.Text(), -1)
		prefixCut, _ := strings.CutPrefix(scanner.Text(), cardIDMatches[0][0])

		values := strings.Split(prefixCut, "|")

		var winningNumbers []int
		for _, num := range NumbersRegex.FindAllString(values[0], -1) {
			n, err := strconv.Atoi(num)
			if err != nil {
				slog.Error("Error convering string to numbers", "error", err)
			}

			winningNumbers = append(winningNumbers, n)
		}

		var yourNumbers []int
		for _, num := range NumbersRegex.FindAllString(values[1], -1) {
			n, err := strconv.Atoi(num)
			if err != nil {
				slog.Error("Error convering string to numbers", "error", err)
			}

			yourNumbers = append(yourNumbers, n)
		}

		var yourWinningNumbers []int
		for _, v := range yourNumbers {
			ok := slices.Contains(winningNumbers, v)
			if ok {
				yourWinningNumbers = append(yourWinningNumbers, v)
			}
		}

		if len(winningNumbers) > 0 {
			var gamePoints int

			for i := range yourWinningNumbers {
				if i == 0 {
					gamePoints = 1
				} else {
					gamePoints = gamePoints * 2
				}
			}

			points += gamePoints
		}

	}

	slog.Info("Total Points Earned", "POINTS", points) //21821
}
