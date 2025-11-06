package task2

import "testing"

func TestReverse(t *testing.T) {
    testCases := []struct {
        input string
        want  string
    }{
        {"hello", "olleh"},
        {"", ""},
        {"Go", "oG"},
        {"a", "a"},
        {"Greetings", "sgniteerG"},
    }

    for _, tc := range testCases {
        got := Reverse(tc.input)
        if got != tc.want {
            t.Errorf("Reverse(%q) = %q; want %q", tc.input, got, tc.want)
        }
    }
}