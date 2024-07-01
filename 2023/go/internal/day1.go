package internal

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func SumCalibrationValues(filePath string) int {
	//#region day 1.2
	numStrings := make(map[string]int)
	numStrings["one"] = 1
	numStrings["two"] = 2
	numStrings["three"] = 3
	numStrings["four"] = 4
	numStrings["five"] = 5
	numStrings["six"] = 6
	numStrings["seven"] = 7
	numStrings["eight"] = 8
	numStrings["nine"] = 9
	//#endregion

	items, _ := os.ReadFile(filePath)
	lines := strings.Split(string(items), "\n")
	sum := 0
	for _, items := range lines {
		converted_string := string(items)
		first := ""
		second := ""
		for id, char := range converted_string {
			converted_char := string(char)
			isNumber := regexp.MustCompile(`[0-9]`).MatchString(converted_char)

			if isNumber {
				if first == "" {
					first = converted_char
				}
				second = converted_char
			} else {
				//#region day 1.2
				for key := range numStrings {
					val := ""
					if id+len(key) <= len(converted_string) {
						val = converted_string[id : id+len(key)]
					}

					if numStrings[val] == 0 {
						continue
					}

					if first == "" {
						first = strconv.Itoa(numStrings[val])
					}
					second = strconv.Itoa(numStrings[val])
				}
				//#endregion
			}
		}
		n, _ := strconv.Atoi(first + second)
		sum += n
	}

	return sum
}
