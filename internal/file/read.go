package file

import (
	"bufio"
	"os"
)

func ReadLines(path string) ([][]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := make([][]byte, 0)
	for scanner.Scan() {
		result = append(result, scanner.Bytes())
	}

	return result, nil
}
