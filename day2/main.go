package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type line struct {
	Min      int
	Max      int
	Letter   string
	Password string
}

func main() {
	values := readfile()
	countValidPasswordPart1 := 0
	countValidPasswordPart2 := 0
	for i := 0; i < len(values); i++ {
		if isValidPasswordPart1(values[i]) == true {
			countValidPasswordPart1++
		}
		if (isValidPasswordPart2(values[i]) == true) {
			countValidPasswordPart2++
		}
	}
	fmt.Println(countValidPasswordPart1)
	fmt.Println(countValidPasswordPart2)
}

func readfile() []line {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []line
	for scanner.Scan() {
		regex := regexp.MustCompile(`^(?P<min>[0-9]+)\-(?P<max>[0-9]+) (?P<letter>[a-zA-Z]+): (?P<password>[a-zA-Z]+)`)
		match := regex.FindStringSubmatch(scanner.Text())

		result := make(map[string]string)
		for i, name := range regex.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}
		line := line{}
		line.Min, _ = strconv.Atoi(result["min"])
		line.Max, _ = strconv.Atoi(result["max"])
		line.Letter = result["letter"]
		line.Password = result["password"]

		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func isValidPasswordPart1(value line) bool {
	countLetter := strings.Count(value.Password, value.Letter)
	if countLetter >= value.Min && countLetter <= value.Max {
		return true
	}
	return false
}

func isValidPasswordPart2(value line) bool {
	if (string(value.Password[value.Min-1]) == value.Letter || string(value.Password[value.Max-1]) == value.Letter) && 
	string(value.Password[value.Min-1]) != string(value.Password[value.Max-1]) {
		return true
	}
	return false
}
