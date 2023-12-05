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
	loadedRed := 12
	loadedGreen := 13
	loadedBlue := 14
	sum := 0

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		validGame := true
		line := scanner.Text()
		split := strings.Split(line, ":")
		splitID := strings.Split(split[0], " ")
		id, _ := strconv.Atoi(strings.ReplaceAll(splitID[1], ":", ""))
		splitGames := strings.Split(split[1], ";")
		for _, game := range splitGames {
			cubes := strings.Split(game, ",")
			for _, cube := range cubes {
				cube = strings.TrimSpace(cube)
				count, _ := strconv.Atoi(strings.Split(cube, " ")[0])
				if strings.Contains(cube, "red") {
					if count > loadedRed {
						validGame = false
					}
				}
				if strings.Contains(cube, "green") {
					if count > loadedGreen {
						validGame = false
					}
				}
				if strings.Contains(cube, "blue") {
					if count > loadedBlue {
						validGame = false
					}
				}
			}
		}
		fmt.Println(id, validGame)
		if validGame {
			sum = sum + id
		}
	}
	fmt.Println(sum)
}
