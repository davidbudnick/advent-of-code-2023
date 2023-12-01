package main

import (
	"bufio"
	"fmt"
	"lib/file"
	"log/slog"
	"strconv"

	"github.com/dlclark/regexp2"
)

var re = regexp2.MustCompile(`(?=(one|two|three|four|five|six|seven|eight|nine|ten|[1-9]))`, 0)

func main() {
	fullNumbers := make(map[string]string)
	fullNumbers["one"] = "1"
	fullNumbers["two"] = "2"
	fullNumbers["three"] = "3"
	fullNumbers["four"] = "4"
	fullNumbers["five"] = "5"
	fullNumbers["six"] = "6"
	fullNumbers["seven"] = "7"
	fullNumbers["eight"] = "8"
	fullNumbers["nine"] = "9"

	f, err := file.ReadFromFile("../input.txt")
	if err != nil {
		slog.Error("Error reading from file", "error", err)
	}

	var sum int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		var matches []string
		match, _ := re.FindStringMatch(scanner.Text())
		for match != nil {
			for _, group := range match.Groups()[1:] {
				if group.Length != 0 {
					matches = append(matches, group.String())
				}
			}
			match, _ = re.FindNextMatch(match)
		}

		if len(matches) == 0 {
			slog.Error("Matches length is zero")
			return
		}

		var firstValue string
		_, ok := fullNumbers[matches[0]]
		if ok {
			firstValue = fullNumbers[matches[0]]
		} else {
			firstValue = matches[0]
		}

		var lastValue string
		_, ok = fullNumbers[matches[len(matches)-1]]
		if ok {
			lastValue = fullNumbers[matches[len(matches)-1]]
		} else {
			lastValue = matches[len(matches)-1]
		}

		s, err := strconv.Atoi(fmt.Sprintf("%s%s", firstValue, lastValue))
		if err != nil {
			slog.Error("Error converting string to a number", "error", err)
		}

		sum += s
	}

	slog.Info("Sum of all of the valibration values", "sum", sum)
}
