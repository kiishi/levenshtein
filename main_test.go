package levenshtein_distance_test

import (
	levenshtein_distance "github.com/kiishi/levenshtein"
	"testing"
)

func TestGetDistance(t *testing.T) {
	testCases := []struct {
		First    string
		Second   string
		Solution int
	}{
		{"book", "boot", 1},
		{"kitten", "sitting", 3},
		{"testing", "tester", 3},
		{"parrallelogram", "ellipse", 12},
		{"international", "interim", 7},
	}


	for _ , testcase  := range testCases {
		value := levenshtein_distance.GetDistance(testcase.First, testcase.Second)
		if value != testcase.Solution {
			t.Errorf("Expected %d got %d", testcase.Solution, value)
		}
	}
}

func TestGetMinimum(t *testing.T) {
	value := levenshtein_distance.GetMinimum(2, 3, 5, 1, 2, 4)
	if value != 1 {
		t.Fail()
	}
}
