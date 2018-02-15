package prologlist_test

import (
	"testing"

	"github.com/elephant/99problems-go/prologlist"
)

type flattenTest struct {
	Expected []prologlist.Item
	List     []interface{}
}

func TestFlatten(t *testing.T) {
	tests := []flattenTest{
		flattenTest{
			Expected: []prologlist.Item{},
			List:     []interface{}{},
		}, // empty list
		flattenTest{
			Expected: []prologlist.Item{
				prologlist.Item{Value: 5},
			},
			List: []interface{}{
				prologlist.Item{Value: 5},
			},
		}, // single Item
		flattenTest{
			Expected: []prologlist.Item{
				prologlist.Item{Value: 5},
				prologlist.Item{Value: 4},
			},
			List: []interface{}{
				prologlist.Item{Value: 5},
				prologlist.Item{Value: 4},
			},
		}, // already flattened list
		flattenTest{
			Expected: []prologlist.Item{
				prologlist.Item{Value: 5},
				prologlist.Item{Value: 4},
			},
			List: []interface{}{
				[]interface{}{
					prologlist.Item{Value: 5},
					prologlist.Item{Value: 4},
				},
			},
		}, // nested
	}

	for _, test := range tests {
		res := prologlist.Flatten(test.List)

		for i, itm := range res {
			if test.Expected[i] != itm.(prologlist.Item) {
				t.Errorf("Expected: %+v; Got: %+v; List: +%v", test.Expected, res, test.List)
			}
		}
	}
}
