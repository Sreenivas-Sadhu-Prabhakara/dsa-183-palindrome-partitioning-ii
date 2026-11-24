package main

import "fmt"

// MinCuts returns the fewest cuts splitting s into palindromic pieces
// (0 when s is already a palindrome). A precomputed palindrome table turns
// the cut DP's inner check into O(1).
func MinCuts(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	// isPal[i][j]: s[i..j] is a palindrome. Filled by increasing length.
	isPal := make([][]bool, n)
	for i := range isPal {
		isPal[i] = make([]bool, n)
		isPal[i][i] = true
	}
	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			if s[i] == s[j] && (length == 2 || isPal[i+1][j-1]) {
				isPal[i][j] = true
			}
		}
	}
	// cuts[i]: fewest cuts for the prefix s[:i].
	cuts := make([]int, n+1)
	cuts[0] = -1 // so a whole-prefix palindrome costs 0
	for i := 1; i <= n; i++ {
		cuts[i] = i - 1 // worst case: cut between every character
		for j := 0; j < i; j++ {
			if isPal[j][i-1] && cuts[j]+1 < cuts[i] {
				cuts[i] = cuts[j] + 1
			}
		}
	}
	return cuts[n]
}

func main() {
	fmt.Println(MinCuts("aab"))
	fmt.Println(MinCuts("racecar"))
}
