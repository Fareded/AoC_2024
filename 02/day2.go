package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"github.com/Fareded/AoC_2024/aoc_helpers"
)


func main() {
    // Reads locationIds from file
    sData := aoc_helpers.ReadFile("reports.txt")

    reports := strings.Split(sData, "\n")

	reportStatus := make([]string, len(reports))

	safe := 0
	unsafe := 0

	for i := 0; i < len(reports); i++ {
		temp := strings.TrimSpace(reports[i])
		report := strings.Split(temp , " ")


		if reportCheck(report) {
			reportStatus[i] = "Safe"
			safe++
		} else if problemDampener(report) {
			reportStatus[i] = "Safe"
			safe++
		} else {
			reportStatus[i] = "Unsafe"
			unsafe++
		}
	}

	fmt.Println("Total Reports: ", len(reports))
	fmt.Println("Safe Reports: ", safe)
	fmt.Println("Unsafe Reports: ", unsafe)

}

func reportCheck(report []string) bool {
	reportData := make([]float64, len(report))
	var err error

	for i := 0; i < len(reportData); i++ {
		reportData[i], err = strconv.ParseFloat(report[i], 64)
		aoc_helpers.Check(err)
	}
	
	isSafe := true
	var asc bool

	// Check if the data starts by ascending or descending
	if reportData[0] -  reportData[1] < 0 {
		asc = true
	} else if reportData[0] -  reportData[1] > 0 {
		asc = false
	}else{
		return false
	}

	for i := 1; i < len(reportData); i++ {

		// Check if the difference between the data points is between 1 and 3
		rawDiff := reportData[i] -  reportData[i-1]
		dataDiff := math.Abs(rawDiff)
		if dataDiff < 1 || dataDiff > 3 {
			isSafe = false
		}

		// Check for a change in the direction of the data
		if !asc && rawDiff > 0 {
			isSafe = false
		} else if asc && rawDiff < 0 {
			isSafe = false
		}

		if !isSafe {
			break
		}
	}
	return isSafe
}

func problemDampener(report []string) bool{
	var valid bool
	tempReport := make([]string, len(report))
	
	for i := 0; i < len(report); i++ {
	copy(tempReport, report)
		if i == 0 {
			valid = reportCheck(report[1:])
		} else if i == len(tempReport) - 1 {
			valid = reportCheck(tempReport[:len(tempReport) - 1])
		} else {
			valid = reportCheck(append(tempReport[:i], tempReport[i+1:]...))
		}
		if valid{
			return true
		}
		
	}
	return false
}