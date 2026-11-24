# Palindrome Partitioning II (Min Cuts)

> Striver SDE Sheet 183 · Day 26 DP II — two DPs stacked: a palindrome table answering "is s[i..j] one?" in O(1), and a prefix DP minimizing cuts over last pieces. Go, stdlib only, tested.

## Problem

The fewest cuts that split a string into palindromic pieces — the optimization sibling of #053, which enumerated *all* partitions (and whose output could be exponential; asking only for the minimum collapses the cost to polynomial).

## Approach

Layer one: `isPal[i][j]` by interval DP — a span is a palindrome when its ends match and its interior is one, filled by increasing length so interiors are always ready. Layer two: `cuts[i]`, the fewest cuts for the prefix of length i — choose the final piece `s[j..i-1]`, demand it be a palindrome (now an O(1) lookup), and pay `cuts[j] + 1`. Seeding `cuts[0] = −1` makes a whole-prefix palindrome cost zero cuts without a special case. The pattern — precompute a predicate table so the optimizing DP's guard is free — recurs all over string DP.

**Complexity:** O(n²) time, O(n²) space for the table.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
