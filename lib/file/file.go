package file

import (
	"log/slog"
	"os"
)

func ReadFromFile(filename string) (*os.File, error) {
	file, err := os.Open("../input.txt")
	if err != nil {
		slog.Error("Error reading from file", "error", err)
		return nil, err
	}
	defer file.Close()

	return file, nil
}
