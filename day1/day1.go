package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func stringToNumber(str string) (value string) {
	numbers := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	if len(str) > 1 {
		value = numbers[str]
		return
	}
	return str
}

func main() {
	result := 0
	r, err := regexp.Compile("one|two|three|four|five|six|seven|eight|nine|[1-9]")

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		first := r.FindString(scanner.Text())

		i := len(line) - 1
		last := ""
		for last == "" {
			last = r.FindString(line[i:])
			i--
		}
		first = stringToNumber(first)
		last = stringToNumber(last)
		combined, _ := strconv.Atoi(first + last)
		result = result + combined
	}

	fmt.Println(result)
}
