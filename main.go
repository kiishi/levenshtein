package levenshtein_distance

import (
	"fmt"
	"math"
)

func GetDistance(firstString, secondString string) int {
	matrix := map[int]map[int]int{}
	for a := 1; a < len(firstString)+1; a++ {
		for b := 1; b < len(secondString)+1; b++ {
			getLevenshteinValue(firstString, secondString, a, b, matrix)
		}
	}
	PrintMatrix(matrix)
	return matrix[len(firstString)][len(secondString)]
}

func PrintMatrix(memo map[int]map[int]int) {
	for _, a := range memo {
		for _, b := range a {
			fmt.Print(b, "-")
		}
		fmt.Println()
	}
}

func GetMinimum(values ...int) int {
	min := -1
	for _, val := range values {
		if min < 0 {
			min = val
			continue
		} else {
			if val < min {
				min = val
			}
		}
	}
	return min
}

func resolveValue(firstString, secondString string, i, j int) int {
	if firstString[i-1] == secondString[j-1] {
		return 0
	} else {
		return 1
	}
}

func getLevenshteinValue(first, second string, i, j int, memo map[int]map[int]int) int {
	//	check if the min is zero
	min := int(math.Min(float64(i), float64(j)))
	if min == 0 {
		max := int(math.Max(float64(i), float64(j)))
		if memo[i] == nil {
			memo[i] = make(map[int]int)
			memo[i][j] = max
		} else {
			memo[i][j] = max
		}
		return max
	} else {
		value := GetMinimum(
			getLevenshteinValue(first, second, i-1, j, memo)+1,
			getLevenshteinValue(first, second, i, j-1, memo)+1,
			getLevenshteinValue(first, second, i-1, j-1, memo)+resolveValue(first, second, i, j),
		)
		if memo[i] == nil {
			memo[i] = make(map[int]int)
			memo[i][j] = value
		} else {
			memo[i][j] = value
		}
		return value
	}
}
