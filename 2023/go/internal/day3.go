package internal

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numChecker = regexp.MustCompile(`\d+`)
var symbolChecker = regexp.MustCompile(`[^\w\s.]`)

func SumAllPartNumbers(filePath string) int {
	items, _ := os.ReadFile(filePath)
	lines := strings.Split(string(items), "\n")
	lines = lines[:len(lines)-1]
	sum := 0

	for row := range lines {
		line := lines[row]
		for col := 0; col < len(line); {
			if line[col] == '.' {
				col++
				continue
			}
			isNumber := numChecker.MatchString(string(line[col]))
			if isNumber {
				num := numChecker.FindString(line[col:])
				foundConnection := CheckNumberForConnection(
					col,
					row,
					col+len(num),
					lines,
				)
				col += len(num)
				if foundConnection {
					n, _ := strconv.Atoi(num)
					sum += n
				}
				continue
			}
			col++
		}
	}
	return sum
}

func CheckNumberForConnection(
	col, row, end int,
	lines []string,
) bool {
	start := col
	for ; col < end; col++ {

		if col == start && col > 0 {
			startChecks := (row > 0 &&
				(symbolChecker.MatchString(string(lines[row-1][col-1])))) ||
				(row < len(lines)-1 &&
					symbolChecker.MatchString(string(lines[row+1][col-1]))) ||
				symbolChecker.MatchString(string(lines[row][col-1]))
			if startChecks {
				return true
			}
		}
		middleChecks := (row > 0 &&
			(symbolChecker.MatchString(string(lines[row-1][col])))) ||
			(row < len(lines)-1 &&
				symbolChecker.MatchString(string(lines[row+1][col])))
		if middleChecks {
			return true
		}

		if col == end-1 && col < len(lines[row])-1 {
			endChecks := (row > 0 &&
				(symbolChecker.MatchString(string(lines[row-1][col+1])))) ||
				(row < len(lines)-1 &&
					symbolChecker.MatchString(string(lines[row+1][col+1]))) ||
				symbolChecker.MatchString(string(lines[row][col+1]))
			if endChecks {
				return true
			}
		}
	}
	return false
}
