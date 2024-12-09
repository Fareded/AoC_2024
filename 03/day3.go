package main

import (
	"fmt"
	"os"
	// "strings"
	"regexp"
	// "math"
	"strconv"
)

func main() {

	// Reads from file
	data, err := os.ReadFile("corruptedMemory.txt")
	check(err)
	sData := string(data)

	validReg := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|don't\(\)|do\(\)`)

	muls := validReg.FindAllString(sData, -1)


	var mulsTotal int64 = 0
	mulsReg := regexp.MustCompile(`\d{1,3}`)

	contMuls := true
	for i := 0; i < len(muls); i++ {
		if muls[i] == "don't()" {
			contMuls = false
		} else if muls[i] == "do()" {
			contMuls = true
		} else if contMuls {
			mulsTemp := mulsReg.FindAllString(muls[i], -1)

			x, err := strconv.ParseInt(mulsTemp[0], 10, 64)
			check(err)
			y, err := strconv.ParseInt(mulsTemp[1], 10, 64)
			check(err)

			z := x * y

			mulsTotal += z
		}

	}

	fmt.Println(mulsTotal)
}

// Error handling function
func check(e error) {
	if e != nil {
		panic(e)
	}
}
