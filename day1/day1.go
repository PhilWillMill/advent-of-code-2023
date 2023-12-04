package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func reverseString(str string) (reversed string) {

	for _, v := range str {
		reversed = string(v) + reversed
	}
	return
}

func main() {
	result := 0
	r, err := regexp.Compile("[1-9]")

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		first := r.FindString(scanner.Text())
		last := r.FindString(reverseString(scanner.Text()))
		combined, _ := strconv.Atoi(first + last)
		result = result + combined
	}

	fmt.Println(result)
}
