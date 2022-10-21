package GetIPInfo

import (
	"bufio"
	"fmt"
	"os"
)

func GetIPList(filename string) ([]string, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, 0, err
	}
	defer file.Close()
	var lines []string
	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		lineCount++
	}
	return lines, lineCount, scanner.Err()
}
