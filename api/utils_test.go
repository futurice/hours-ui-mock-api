package api

import "testing"

func TestRoundToHalf(t *testing.T) {
	type testCase struct {
		initial float64
		expected float64
	}

	testCases := []testCase{
		testCase{1.0, 1.0},
		testCase{1.2, 1.0},
		testCase{1.3, 1.5},
		testCase{1.6, 1.5},
		testCase{1.85, 2.0},
		testCase{-1.0, -1.0},
		testCase{-1.2, -1.0},
		testCase{-1.3, -1.5},
		testCase{-1.6, -1.5},
		testCase{-1.85, -2.0},
	}

	for _, test := range testCases {
		actual := RoundToHalf(test.initial)
		if actual != test.expected {
			t.Error("Expected", actual, "to be", test.expected)
		}
	}
}
