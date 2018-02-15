package prologlist_test

import (
	"testing"

	"github.com/elephant/99problems-go/prologlist"
)

type listTest struct {
	Expected prologlist.Item
	List     []prologlist.Item
}

func TestLastInt(t *testing.T) {
	var none prologlist.Item
	tests := []listTest{
		listTest{
			Expected: none,
			List:     []prologlist.Item{},
		}, // empty list
		listTest{
			Expected: prologlist.Item{Value: 5},
			List: []prologlist.Item{
				prologlist.Item{Value: 5},
			},
		},
		listTest{
			Expected: prologlist.Item{Value: 1},
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
		res := prologlist.Last(test.List)
		if test.Expected != res {
			t.Errorf("Expected: %v, Got: %v", test.Expected, res)
		}
	}
}
