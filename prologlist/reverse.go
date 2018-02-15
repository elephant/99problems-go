package prologlist

// Reverse an Item list
func Reverse(l []Item) (rev []Item) {
	n := len(l)
	if n == 0 {
		return rev
	}
	if n == 1 {
		return l
	}

	for i := n - 1; i >= 0; i-- {
		rev = append(rev, l[i])
	}

	return rev
}
