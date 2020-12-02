package main

import (
	"fmt"
	"testing"
)

func TestLogging(t *testing.T) {
	line := line{}
	line.Min = 2
	line.Max = 9
	line.Letter = "c"
	line.Password = "ccccccccc"
	fmt.Println(isValidPasswordPart2(line))
}
