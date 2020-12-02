package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	values := readfile()
	fmt.Println(calculateWithTwoValue(values, 2020))
	fmt.Println(calculateWithThreeValue(values, 2020))
}

func readfile() []int {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []int
	for scanner.Scan() {
		intValue, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, intValue)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func calculateWithTwoValue(array []int, target int) int {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if target-array[i] == array[j] {
				return array[i] * array[j]
			}
		}
	}
	return 0
}

func calculateWithThreeValue(array []int, target int) int {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			for h := 0; h < len(array); h++ {
				if target-array[i]-array[j] == array[h] {
					return array[i] * array[j] * array[h]
				}
			}
		}
	}
	return 0
}
