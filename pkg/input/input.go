package input

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func Load(filePath string) (lines []string, err error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get current path: %w", err)
	}
	file, err := os.Open(filepath.Join(path, filePath))
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read file lines: %w", err)
	}
	return lines, nil
}
