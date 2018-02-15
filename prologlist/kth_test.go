package prologlist_test

import (
	"testing"

	"github.com/elephant/99problems-go/prologlist"
)

type kthListTest struct {
	Kth      int
	Expected prologlist.Item
	List     []prologlist.Item
}

func TestKth(t *testing.T) {
	var none prologlist.Item
	tests := []kthListTest{
		kthListTest{
			Kth:      0,
			Expected: none,
			List:     []prologlist.Item{},
		}, // empty list
		kthListTest{
			Kth:      -1,
			Expected: none,
			List: []prologlist.Item{
				prologlist.Item{Value: 5},
			},
		}, // out of bounds
		kthListTest{
			Kth:      1,
			Expected: none,
			List: []prologlist.Item{
				prologlist.Item{Value: 5},
			},
		}, // out of bounds
		kthListTest{
			Kth:      0,
			Expected: prologlist.Item{Value: 5},
			List: []prologlist.Item{
				prologlist.Item{Value: 5},
			},
		},
		kthListTest{
			Kth: 3,
			Expected: prologlist.Item{
				Value: 2,
			},
			List: []prologlist.Item{
				prologlist.Item{Value: 5},
				prologlist.Item{Value: 4},
				prologlist.Item{Value: 3},
				prologlist.Item{Value: 2},
				prologlist.Item{Value: 1},
			},
		},
	}

	for _, test := range tests {
		res := prologlist.Kth(test.List, test.Kth)
		if test.Expected != res {
			t.Errorf("Expected: %v; Got: %v", test.Expected, res)
		}
	}
}
