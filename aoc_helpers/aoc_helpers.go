// this package just contains functions I use regularly between puzzles

package aoc_helpers

import (
	"os"
)

	
func ReadFile(fileName string) string {
	data, err := os.ReadFile(fileName)
    Check(err)
    sData := string(data)
	return sData
}

// Error handling function
func Check(e error) {
    if e != nil {
        panic(e)
    }
}