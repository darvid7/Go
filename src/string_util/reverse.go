// Package string_util contains utlity functions for working with strings.

package string_util

// Reverse returns a string reversed left to right.
func Reverse(s string) string {
	r := []rune(s) // Array of runes from string s.
	// A rune in an integer value mapped to unicode value i.e. 97 = a
	// for i = 0, j = len(r) -1
	// 	while i < len(r)/2 # Go half way
	//		i += 1, j -= 1
	// Reverse string by swapping first element with last element and so 
	// forth until half the len of array of chars.
	for i, j :=0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

