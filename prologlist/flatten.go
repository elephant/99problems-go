package prologlist

// Flatten the given list of Items, including nested lists
func Flatten(l []interface{}) (flattened []interface{}) {

	flattened = flat(l, flattened, 0)

	return flattened
}

func flat(l []interface{}, flattened []interface{}, i int) []interface{} {
	n := len(l)
	if i >= n {
		return flattened
	}

	switch l[i].(type) {
	case []interface{}:
		k := l[i].([]interface{})
		flattened = append(flattened, k...)
	default:
		k := l[i].(Item)
		flattened = append(flattened, k)
	}

	return flat(l, flattened, i+1)
}
