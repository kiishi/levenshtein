package levenshtein_test

import (
	"github.com/kiishi/levenshtein"
	"testing"
)

func TestGetDistance(t *testing.T) {
	//extracted from an online tool
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
		{"intern", "interim", 2},
	}


	for _ , testcase  := range testCases {
		value := levenshtein.GetDistance(testcase.First, testcase.Second)
		if value != testcase.Solution {
			t.Errorf("Expected %d got %d", testcase.Solution, value)
		}
	}
}
func BenchmarkGetDistance(b *testing.B) {
	for n:= 0; n < b.N ; n++{
		levenshtein.GetDistance("kitten", "sitting")
	}
}

func TestGetMinimum(t *testing.T) {
	value := levenshtein.GetMinimum(2, 3, 5, 1, 2, 4)
	if value != 1 {
		t.Fail()
	}
}
