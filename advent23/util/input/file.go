package input

import (
	"advent23/util/assert"
	"bufio"
	"os"
)

func LineByLine(fname string) (*os.File, []string) {
	file, err := os.Open(fname)
	assert.Empty(err)
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	assert.Empty(scanner.Err())
	return file, lines
}
