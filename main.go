package levenshtein

import (
	"math"
)

func GetDistance(firstString, secondString string) int {
	memory := map[int]map[int]int{}
	return getLevenshteinValue(firstString, secondString, len(firstString), len(secondString), memory)
}


func GetMinimum(values ...int) int {
	// assuming can never be negative
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

func getConstant(firstString, secondString string, i, j int) int {
	if firstString[i-1] == secondString[j-1] {
		return 0
	} else {
		return 1
	}
}

func getLevenshteinValue(first, second string, i, j int, memo map[int]map[int]int) int {

	if memo[i] != nil{
		if val , ok := memo[i][j]; ok{
			return val
		}
	}
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
			getLevenshteinValue(first, second, i-1, j-1, memo)+getConstant(first, second, i, j),
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
