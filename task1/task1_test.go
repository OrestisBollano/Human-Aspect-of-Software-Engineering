package task1

import "testing"

func TestAdd(t *testing.T) {
    // Define all the test cases
    testCases := []struct {
        a    int // input 1
        b    int // input 2
        want int // expected result
    }{
        {1, 2, 3},
        {5, 5, 10},
        {0, 0, 0},
        {-1, 1, 0},
        {-5, -10, -15},
    }

    // Loop through all test cases
    for _, tc := range testCases {
        got := Add(tc.a, tc.b)
        if got != tc.want {
            // The test fails if 'got' is not equal to 'want'
            t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.want)
        }
    }
}