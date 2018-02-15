package prologlist_test

import (
	"testing"

	"github.com/elephant/99problems-go/prologlist"
)

type reverseTest struct {
	Expected []prologlist.Item
	List     []prologlist.Item
}

func TestReverse(t *testing.T) {
	var none []prologlist.Item
	tests := []reverseTest{
		reverseTest{
			Expected: none,
			List:     []prologlist.Item{},
		}, //empty list
		reverseTest{
			Expected: []prologlist.Item{
				prologlist.Item{Value: 5},
			}, List: []prologlist.Item{
				prologlist.Item{Value: 5},
			},
		},
		reverseTest{
			Expected: []prologlist.Item{
				prologlist.Item{Value: 5},
				prologlist.Item{Value: 4},
				prologlist.Item{Value: 3},
				prologlist.Item{Value: 2},
				prologlist.Item{Value: 1},
			},
			List: []prologlist.Item{
				prologlist.Item{Value: 1},
				prologlist.Item{Value: 2},
				prologlist.Item{Value: 3},
				prologlist.Item{Value: 4},
				prologlist.Item{Value: 5},
			},
		},
	}

	for _, test := range tests {
		res := prologlist.Reverse(test.List)
		switch {
		case len(res) != len(test.Expected):
			t.Errorf("Expected: %v; Got: %v", test.Expected, res)
		default:
			n := len(res)
			for i, v := range res {
				i++
				if test.List[n-i] != v {
					t.Errorf("Not reversed. Expected list[%d] = %d; Got %d", i, test.List[n-i], v)
				}
			}
		}
	}
}
