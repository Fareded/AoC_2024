package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	// "math"
	// "strconv"
)

func main() {

	// Reads from file
	data, err := os.ReadFile("wordSearch.txt")
	check(err)
	sData := string(data)

	rows := strings.Split(sData, "\n")

	wordSearch := make([][]string, len(rows))
	xmasReg := regexp.MustCompile(`\w{1}`)

	for i := 0; i < len(rows); i++ {
		rows[i] = strings.Trim(rows[i], " ")
		wordSearch[i] = xmasReg.FindAllString(rows[i], -1)
	}


	fmt.Print("Number found: ", findX(wordSearch))
	fmt.Println()
	// for i := 0; i < len(wordSearch); i++ {
	// 	fmt.Println(wordSearch[i])
	// }
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findX(wordSearch [][]string) int {
	xmas := []string{"X", "M", "A", "S"}
	xmasCount := 0

	for i := 0; i < len(wordSearch); i++ {
		for j := 0; j < len(wordSearch[i]); j++ {
			if wordSearch[i][j] == xmas[0] {
				found, wrdCount := dfs(wordSearch, i, j, 0, 0)

				if found {
					xmasCount += wrdCount
				}
			}
		}
	}

	return xmasCount
}

func dfs(wordSearch [][]string, row, column, index int, direction int) (bool, int) {
	//return true if the word is found
	if index == 4 {
		return true, 1
	}

	//makes sure we dont search out of bounds
	r_outbound := row < 0 || row >= len(wordSearch)
	c_outbound := column < 0 || column >= len(wordSearch[0])

	if r_outbound || c_outbound {
		return false, 0
	}

	// if the letter is wrong return false
	wrongLetter := wordSearch[row][column] != string("XMAS"[index])
	if wrongLetter {
		return false, 0
	}

	//mark the letter as visited
	letter := wordSearch[row][column]
	wordSearch[row][column] = "\033[32m" + letter + "\033[0m"
	wordFound := false

	switch direction {
	case 0:
		f1, x1 := dfs(wordSearch, row+1, column, index+1, 1)
		f2, x2 := dfs(wordSearch, row-1, column, index+1, 2)
		f3, x3 := dfs(wordSearch, row, column+1, index+1, 3)
		f4, x4 := dfs(wordSearch, row, column-1, index+1, 4)
		f5, x5 := dfs(wordSearch, row+1, column+1, index+1, 5)
		f6, x6 := dfs(wordSearch, row+1, column-1, index+1, 6)
		f7, x7 := dfs(wordSearch, row-1, column+1, index+1, 7)
		f8, x8 := dfs(wordSearch, row-1, column-1, index+1, 8)

		xSum := x1 + x2 + x3 + x4 + x5 + x6 + x7 + x8
		if f1 || f2 || f3 || f4 || f5 || f6 || f7 || f8 {
			return true, xSum
		}
	case 1:
		found, _ := dfs(wordSearch, row+1, column, index+1, 1)
		if found {
			wordFound = true
		}
	case 2:
		found, _ := dfs(wordSearch, row-1, column, index+1, 2)
		if found {
			wordFound = true
		}
	case 3:
		found, _ := dfs(wordSearch, row, column+1, index+1, 3)
		if found {
			wordFound = true
		}
	case 4:
		found, _ := dfs(wordSearch, row, column-1, index+1, 4)
		if found {
			wordFound = true
		}
	case 5:
		found, _ := dfs(wordSearch, row+1, column+1, index+1, 5)
		if found {
			wordFound = true
		}
	case 6:
		found, _ := dfs(wordSearch, row+1, column-1, index+1, 6)
		if found {
			wordFound = true
		}
	case 7:
		found, _ := dfs(wordSearch, row-1, column+1, index+1, 7)
		if found {
			wordFound = true
		}
	case 8:
		found, _ := dfs(wordSearch, row-1, column-1, index+1, 8)
		if found {
			wordFound = true
		}
	}

	wordSearch[row][column] = letter

	if !wordFound {
		return false, 0
	} else {
		return true, 1
	}
}
