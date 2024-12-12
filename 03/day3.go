package main

import (
	"fmt"
	// "strings"
	"regexp"
	// "math"
	"strconv"

	"github.com/Fareded/AoC_2024/aoc_helpers"
)

func main() {

	// Reads from file
	
	sData := aoc_helpers.ReadFile("corruptedMemory.txt")

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
			aoc_helpers.Check(err)
			y, err := strconv.ParseInt(mulsTemp[1], 10, 64)
			aoc_helpers.Check(err)

			z := x * y

			mulsTotal += z
		}

	}

	fmt.Println(mulsTotal)
}

