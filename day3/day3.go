package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type numberLoc struct {
	starI  int
	starJ  int
	number int
}

func read(file *os.File) (lines []string) {

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return
}

func checkLine(line string, startPos, endPos int) (bool, int) {
	for j, slice := range strings.Split(line, "") {
		if j >= startPos-1 && j <= endPos {
			if slice == "*" {
				return true, j
			}
		}
		if j > endPos {
			return false, -1
		}
	}
	return false, -1
}

func findSymbol(startPos, endPos, currentLine int, lines []string) (posI, posJ int, found bool) {
	found = false
	startPos = max(1, startPos)
	endPos = min(len(lines[currentLine])-1, endPos)
	if currentLine == 0 { // Check current and next
		found, posJ = checkLine(lines[currentLine], startPos, endPos)
		if found {
			return currentLine, posJ, found
		}
		found, posJ = checkLine(lines[currentLine+1], startPos, endPos)
		if found {
			return currentLine + 1, posJ, true
		}
	} else if currentLine+1 >= len(lines) { // Check previous and current
		found, posJ = checkLine(lines[currentLine-1], startPos, endPos)
		if found {
			return currentLine - 1, posJ, true
		}
		found, posJ = checkLine(lines[currentLine], startPos, endPos)
		if found {
			return currentLine, posJ, true
		}
	} else { // Check previous, current and next
		found, posJ = checkLine(lines[currentLine-1], startPos, endPos)
		if found {
			return currentLine - 1, posJ, true
		}
		found, posJ = checkLine(lines[currentLine], startPos, endPos)
		if found {
			return currentLine, posJ, true
		}
		found, posJ = checkLine(lines[currentLine+1], startPos, endPos)
		if found {
			return currentLine + 1, posJ, true
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

	// For part 2, when we find a number adjacent to a * symbol add it to a structure of (i,j,number) where i,j are the coords of the *
	// We then run through the data structure following the logic of,
	// if same * then mult numbers, increment count
	// if count > 2 then quit and start on the next item
	// we'll be double adding so divide by 2

	var numbers []numberLoc

	lines := read(file)

	for i, line := range lines {
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
				starI, starJ, found := findSymbol(numStart, numEnd, i, lines)

				if found {
					numbers = append(numbers, numberLoc{starI, starJ, number})
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
	// fmt.Println(numbers)
	gearMultSum := 0
	for number := range numbers {
		matchCount := 0
		gearMult := 0
		for i := number + 1; i < len(numbers); i++ {
			if numbers[number].starI == numbers[i].starI && numbers[number].starJ == numbers[i].starJ {
				matchCount++
				gearMult = gearMult + numbers[number].number*numbers[i].number
				fmt.Println(numbers[number])
				fmt.Println(numbers[i])
			}
		}
		if matchCount == 1 {
			gearMultSum = gearMultSum + gearMult
		}
	}
	fmt.Println(gearMultSum)
}
