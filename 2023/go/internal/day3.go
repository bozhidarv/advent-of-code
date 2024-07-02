package internal

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func SumAllPartNumbers(filePath string) int {
	items, _ := os.ReadFile(filePath)
	lines := strings.Split(string(items), "\n")
	lines = lines[:len(lines)-1]
	sum := 0

	symbolChecker := regexp.MustCompile(`[^\w\s.]`)
	numChecker := regexp.MustCompile(`\d+`)

	for col := range lines {
		line := lines[col]
		for row := 0; row < len(line); {
			if line[row] == '.' {
				row++
				continue
			}
			isNumber := numChecker.MatchString(string(line[row]))
			if isNumber {
				num := regexp.MustCompile(`\d+`).FindString(line[row:])
				foundConnection := CheckNumberForConnection(
					row,
					col,
					row+len(num),
					lines,
					symbolChecker,
				)
				row += len(num)
				if foundConnection {
					n, _ := strconv.Atoi(num)
					sum += n
				}
				continue
			}
			row++
		}
	}
	return sum
}

func CheckNumberForConnection(
	row, col, end int,
	lines []string,
	symbolChecker *regexp.Regexp,
) bool {
	start := row
	for ; row < end; row++ {

		if row == start && row > 0 {
			startChecks := (col > 0 &&
				(symbolChecker.MatchString(string(lines[col-1][row-1])))) ||
				(col < len(lines)-1 &&
					symbolChecker.MatchString(string(lines[col+1][row-1]))) ||
				symbolChecker.MatchString(string(lines[col][row-1]))
			if startChecks {
				return true
			}
		}
		middleChecks := (col > 0 &&
			(symbolChecker.MatchString(string(lines[col-1][row])))) ||
			(col < len(lines)-1 &&
				symbolChecker.MatchString(string(lines[col+1][row])))
		if middleChecks {
			return true
		}

		if row == end-1 && row < len(lines[col])-1 {
			endChecks := (col > 0 &&
				(symbolChecker.MatchString(string(lines[col-1][row+1])))) ||
				(col < len(lines)-1 &&
					symbolChecker.MatchString(string(lines[col+1][row+1]))) ||
				symbolChecker.MatchString(string(lines[col][row+1]))
			if endChecks {
				return true
			}
		}
	}
	return false
}
