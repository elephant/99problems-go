package prologlist_test

import (
	"testing"

	"github.com/elephant/99problems-go/prologlist"
)

type isPalindromeTest struct {
	Expected bool
	List     []prologlist.Item
}

func TestIsPalindrome(t *testing.T) {
	tests := []isPalindromeTest{
		isPalindromeTest{
			Expected: true,
			List:     []prologlist.Item{},
		}, // empty list
		isPalindromeTest{
			Expected: true,
			List: []prologlist.Item{
				prologlist.Item{},
			},
		}, // single item list
		isPalindromeTest{
			Expected: true,
			List: []prologlist.Item{
				prologlist.Item{Value: 3},
				prologlist.Item{Value: 1},
				prologlist.Item{Value: 2},
				prologlist.Item{Value: 1},
				prologlist.Item{Value: 3},
			},
		},
		isPalindromeTest{
			Expected: false,
			List: []prologlist.Item{
				prologlist.Item{Value: 3},
				prologlist.Item{Value: 1},
			},
		},
	}

	for _, test := range tests {
		res := prologlist.IsPalindrome(test.List)

		if test.Expected != res {
			t.Errorf("Expected: %v; Got: %v. List: %+v", test.Expected, res, test.List)
		}
	}
}
