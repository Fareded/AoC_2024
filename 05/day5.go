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
	reorderValid := 0
	var reMidTotal int64 = 0

	for i := 0; i < len(pages); i++ {
		valid, mid := ruleEnforcer(rules, strings.Split(pages[i], ","))
		if valid {
			validPages++
			midTotal += mid
		}else{
			reordered := reorderPage(rules, strings.Split(pages[i], ","))
			valid, mid := ruleEnforcer(rules, reordered)
			if valid {
				reorderValid++
				reMidTotal += mid
			}
		}
	}

	fmt.Println("Valid Pages: ", validPages)
	fmt.Println("Mid Total: ", midTotal)
	fmt.Println("Reordered Pages: ", reorderValid)
	fmt.Println("Reordered Mid Total: ", reMidTotal)
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

//Checks if a page follows the rules given
func ruleEnforcer(rules [][]string, page []string) (bool, int64) {

	for i := 0; i < len(page); i++ {
		match := findRules(rules, page, i)
		for k := 0; k < len(page); k++ {
			if contains(match, page[k]) {
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

//helper function to find the rules that determine the values the input comes before
func findRules(rules [][]string, page []string, targetValue int) []string {
	match := make([]string, 0)

	for i := 0; i < len(rules); i++ {
		//finding relevant rules
		if page[targetValue] == rules[i][0] {
			match = append(match, rules[i][1])
		}
	}

	return match
}

//helper function to check if a string is in a slice (aka a value in a rules list)
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func reorderPage(allRules [][]string, page []string) []string {
	
	rules := make([][]string, len(page))

	for i := 0; i < len(page); i++ {
		rules[i] = append(rules[i], page[i])
		rules[i] = append(rules[i], findRules(allRules, page, i)...)
	}

	reordered := mergeSort(page, rules)

	return reordered
}

//merge sort for reordering a page
func mergeSort(arr []string, rules [][]string) []string {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2
	left := mergeSort(arr[:middle], rules)
	right := mergeSort(arr[middle:], rules)

	return merge(left, right, rules)
}

func merge(left, right []string, rules [][]string) []string {
	size, i, j := len(left)+len(right), 0, 0
	sorted := make([]string, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			sorted[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			sorted[k] = left[i]
			i++
		} else if contains(getMyRules(left[i], rules), right[j]) {
			sorted[k] = left[i]
			i++
		} else {
			sorted[k] = right[j]
			j++
		}
	}

	return sorted
}

func getMyRules(x string, y [][]string) []string {

	for i := 0; i < len(y); i++ {
		if y[i][0] == x {
			return y[i][1:]
		}
	}
	return nil
}
