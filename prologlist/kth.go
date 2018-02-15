package prologlist

// Kth Item from list
func Kth(l []Item, kth int) (k Item) {
	n := len(l)
	if n == 0 || kth > n-1 || kth < 0 {
		return k
	}

	return l[kth]
}
