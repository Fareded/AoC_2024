package main

import (
    "fmt"
    "os"
    "strings"
    "math"
    "strconv"
)


func main() {
    // Reads locationIds from file
    data, err := os.ReadFile("reports.txt")
    check(err)
    sData := string(data)

    reports := strings.Split(sData, "\n")

	reportStatus := make([]string, len(reports))

	safe := 0
	unsafe := 0

	for i := 0; i < len(reports); i++ {

		if reportCheck(reports[i]) {
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

// Error handling function
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func reportCheck(report string) bool {
	report = strings.TrimSpace(report)
	reportTemp := strings.Split(report, " ")
	reportData := make([]float64, len(reportTemp))
	var err error

	for i := 0; i < len(reportData); i++ {
		reportData[i], err = strconv.ParseFloat(reportTemp[i],64)
		check(err)
	}
	
	status := true
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
			status = false
		}

		// Check for a change in the direction of the data
		if asc == false && rawDiff > 0 {
			status = false
		} else if asc == true && rawDiff < 0 {
			status = false
		}

		if status == false {
			break
		}
	}
	return status
}