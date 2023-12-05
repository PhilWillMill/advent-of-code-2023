package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		minRed := 0
		minGreen := 0
		minBlue := 0
		line := scanner.Text()
		split := strings.Split(line, ":")
		// splitID := strings.Split(split[0], " ")
		// id, _ := strconv.Atoi(strings.ReplaceAll(splitID[1], ":", ""))
		splitGames := strings.Split(split[1], ";")
		for _, game := range splitGames {
			cubes := strings.Split(game, ",")
			for _, cube := range cubes {
				cube = strings.TrimSpace(cube)
				count, _ := strconv.Atoi(strings.Split(cube, " ")[0])
				if strings.Contains(cube, "red") {
					if count > minRed {
						minRed = count
					}
				}
				if strings.Contains(cube, "green") {
					if count > minGreen {
						minGreen = count
					}
				}
				if strings.Contains(cube, "blue") {
					if count > minBlue {
						minBlue = count
					}
				}
			}
		}
		fmt.Println(line)
		fmt.Println("Red:", minRed, "Green:", minGreen, "Blue:", minBlue)
		power := minRed * minGreen * minBlue
		fmt.Println("Power:", power)
		sum = sum + power
	}
	fmt.Println("Total Power:", sum)
}
