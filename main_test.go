package main

import "testing"

func TestMinCuts(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"aab", 1},     // aa | b
		{"racecar", 0}, // already a palindrome
		{"a", 0},
		{"ab", 1},
		{"", 0},
		{"noonabbad", 2}, // noon | abba | d
		{"abcde", 4},     // no palindromes longer than 1
		{"aaaa", 0},
	}
	for _, c := range cases {
		if got := MinCuts(c.in); got != c.want {
			t.Errorf("MinCuts(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}
