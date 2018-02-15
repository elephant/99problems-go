package prologlist

// Last Item from list
func Last(l []Item) (last Item) {
	n := len(l)
	if n == 0 {
		return last
	}

	return l[n-1]
}
