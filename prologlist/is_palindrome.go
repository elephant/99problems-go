package prologlist

// IsPalindrome returns true if the given list Items are a palindrome
func IsPalindrome(l []Item) bool {
	n := len(l)
	if n == 0 {
		return true
	}
	if n == 1 {
		return true
	}

	c := 0
	for i := n - 1; i != n && i != c; i-- {
		if l[c] != l[i] {
			return false
		}
		c++
	}

	return true
}
