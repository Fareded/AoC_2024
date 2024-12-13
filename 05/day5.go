package main

import (
	"fmt"
	// "regexp"
	"strings"

	"github.com/Fareded/AoC_2024/aoc_helpers"
	// "math"
	"strconv"
)

func main() {
	rules, pages := parseData()
	validPages := 0
	var midTotal int64 = 0

	for i := 0; i < len(pages); i++ {
		valid, mid := ruleEnforcer(rules, strings.Split(pages[i], ","))
		if valid{
			validPages++
			midTotal += mid
		}
	}

	fmt.Println("Valid Pages: ", validPages)
	fmt.Println("Mid Total: ", midTotal)
}

func parseData() ([][]string, []string) {
	sData := aoc_helpers.ReadFile("pagesAndRules.txt")
	rulesAndPages := strings.Split(sData, "\r\n\r\n")
	tempRules := strings.Split(rulesAndPages[0], "\n")
	for i := 0; i < len(tempRules); i++ {
		tempRules[i] = strings.Trim(tempRules[i], "\r")
	}
	rules := make([][]string, len(tempRules))

	for i := 0; i < len(tempRules); i++ {
		rules[i] = strings.Split(tempRules[i], "|")
	}

	pages := strings.Split(rulesAndPages[1], "\n")
	for i := 0; i < len(pages); i++ {
		pages[i] = strings.Trim(pages[i], "\r")
	}

	return rules, pages
}

// find current value
// find the rules that matches the current value
// find the corresponding value in each rule within the page
// record their positions
// compare to the position of the current value

func ruleEnforcer(rules [][]string, page []string) (bool, int64){
	

	for i := 0; i < len(page); i++ {
		match := make([]string, 0)

		for j := 0; j < len(rules); j++ {
			//finding relevant rules
			if page[i] == rules[j][0] {
				match = append(match, rules[j][1])
			}
			
		}
		for k := 0; k < len(page); k++ {
			if  contains(match, page[k]) {
				if k < i {
					return false, 0
				}
			}
		}
	}

	midPoint := len(page) / 2
	mid, err := strconv.ParseInt(page[midPoint], 10, 64)
	aoc_helpers.Check(err)

	return true, mid
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
