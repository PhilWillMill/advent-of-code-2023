package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func read(file *os.File) (lines []string) {

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return
}

func checkLine(line string, startPos, endPos int) bool {
	for j, slice := range strings.Split(line, "") {
		if j >= startPos-1 && j <= endPos {
			if slice != "." {
				_, err := strconv.Atoi(slice)
				if err != nil {
					return true
				}
			}
		}
		if j > endPos {
			return false
		}
	}
	return false
}

func findSymbol(startPos, endPos, currentLine int, lines []string) (found bool) {
	found = false
	startPos = max(1, startPos)
	endPos = min(len(lines[currentLine])-1, endPos)
	if currentLine == 0 { // Check current and next
		found = checkLine(lines[currentLine], startPos, endPos)
		if found {
			return true
		}
		found = checkLine(lines[currentLine+1], startPos, endPos)
		if found {
			return true
		}
	} else if currentLine+1 >= len(lines) { // Check previous and current
		found = checkLine(lines[currentLine-1], startPos, endPos)
		if found {
			return true
		}
		found = checkLine(lines[currentLine], startPos, endPos)
		if found {
			return true
		}
	} else { // Check previous, current and next
		found = checkLine(lines[currentLine-1], startPos, endPos)
		if found {
			return true
		}
		found = checkLine(lines[currentLine], startPos, endPos)
		if found {
			return true
		}
		found = checkLine(lines[currentLine+1], startPos, endPos)
		if found {
			return true
		}
	}
	return
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := read(file)

	numberSum := 0

	for i, line := range lines {
		if i <= 1300 {
			numStart := -1
			numEnd := 0
			number := 0
			for j, slice := range strings.Split(line, "") {
				_, err := strconv.Atoi(slice)
				if (j == len(line)-1 || err != nil) && numStart != -1 {
					if j == len(line)-1 && err == nil {
						numEnd = j + 1
					} else {
						numEnd = j
					}

					number, _ = strconv.Atoi(line[numStart:numEnd])
					symbol := findSymbol(numStart, numEnd, i, lines)

					if symbol {
						numberSum = numberSum + number
						fmt.Println(i+1, " ", number, " ", symbol, " ", numberSum)
					}

					numStart = -1
					numEnd = 0
				}
				if err == nil {
					if numStart == -1 {
						numStart = j
					}
				}
			}
		}
	}
}
