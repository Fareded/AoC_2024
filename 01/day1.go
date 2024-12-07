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
    data, err := os.ReadFile("locationIDs.txt")
    check(err)
    sData := string(data)

    locationIDs := strings.Split(sData, "\n")

    // fmt.Println(len(locationIDs) - 1)
    n := len(locationIDs)

    // Create two slices the length of locationIDs
    idList1 := make([]float64, n)
    idList2 := make([]float64, n)
    
    for i := 0; i < len(locationIDs); i++ {
        x := strings.Split(locationIDs[i], "   ")

        // fmt.Println(x[1])

        // should have error handling here but yk
        idList1[i], err = strconv.ParseFloat(x[0],64)
        check(err)
        idList2[i], err = strconv.ParseFloat(strings.TrimSpace(x[1]),64)
        check(err)

    }

    fmt.Println("Location ID lists created")

    

    idList1 = mergeSort(idList1)
    idList2 = mergeSort(idList2)

    fmt.Println("Location ID lists sorted")

    distances := make([]float64, n)

    for i := 0; i < n; i++ {
        distances[i] = math.Abs(idList1[i] - idList2[i])
    }

    fmt.Println("Distances calculated")

    totalDistance := sumList(distances)

    //using prinf to print the full number
    fmt.Printf("Total Distance: %f\n", totalDistance)

    similarityScore := calcSimilarityScore(idList1, idList2)

    fmt.Printf("Similarity Score: %f\n", similarityScore)
}

// Error handling function
func check(e error) {
    if e != nil {
        panic(e)
    }
}


func mergeSort(arr []float64) []float64 {
    if len(arr) <= 1 {
        return arr
    }

    middle := len(arr) / 2
    left := mergeSort(arr[:middle])
    right := mergeSort(arr[middle:])

    return merge(left, right)
}

func merge(left, right []float64) []float64 {
    size, i, j := len(left)+len(right), 0, 0
    sorted := make([]float64, size)

    for k := 0; k < size; k++ {
        if i > len(left)-1 && j <= len(right)-1 {
            sorted[k] = right[j]
            j++
        } else if j > len(right)-1 && i <= len(left)-1 {
            sorted[k] = left[i]
            i++
        } else if left[i] < right[j] {
            sorted[k] = left[i]
            i++
        } else {
            sorted[k] = right[j]
            j++
        }
    }

    return sorted
}

func sumList(arr []float64) float64 {
    sum := 0.0
    for i := 0; i < len(arr); i++ {
        sum += arr[i]
    }
    return sum
}

func calcSimilarityScore(list1, list2 []float64) float64 {
    simScores := make([]float64, len(list1))

    for i := 0 ; i < len(list1); i++ {
        noSims := 0.0

        for j := 0; j < len(list2); j++ {
            if list1[i] == list2[j] {
                noSims++
            }
        }

        simScores[i] = list1[i] * noSims
    } 

    return sumList(simScores)
}