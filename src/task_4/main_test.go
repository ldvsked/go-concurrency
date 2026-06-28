package main 

import (
	"testing"
)

func getCacheItemsSlice[T comparable](c *Cache[T]) []T {
	if c.root == nil {
		return []T{}
	}

	var res []T = make([]T, 0, len(c.itemsMap))
	var cur *node[T] = c.root
	for cur != nil {
		res = append(res, cur.value)
		cur = cur.next
	}
	return res
}

func compareSlice[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

type testCase[T comparable] struct {
	name string 
	capacity uint 
	actionsSet []T
	actionsGet []actionGet[T]
	expected []T
	clearFlag bool
}

type actionGet[T comparable] struct {
	key T
	expectedFlag bool
}

func runTestCase[T comparable](t *testing.T, tc testCase[T]){
	t.Run(tc.name, func(t *testing.T){
		var c *Cache[T] = NewCache[T](tc.capacity)
		for _, action := range tc.actionsSet {
			c.Set(action)
		}
		for _, action := range tc.actionsGet {
			var _, actualFlag = c.Get(action.key)
			if action.expectedFlag != actualFlag {
				if action.expectedFlag{
					t.Errorf("Error GET for element %v: it must be in cache, but actually it is not in cache", action.key)
				} else {
					t.Errorf("Error GET for element %v: it must not be in cache, but actually it is in cache", action.key)
				}
			}  
		}

		if tc.clearFlag {
			c.Clear()
		}

		var got = getCacheItemsSlice(c)
		if !compareSlice(got, tc.expected) {
			t.Errorf("Expected: %v, got: %v", tc.expected, got)
			return
		}
	})
}

func TestInt(t *testing.T) { //тестовые функции не могут быть дженериками
	var testCases []testCase[int] = []testCase[int]{
		{"intSetEviction", 3, []int{1, 2, 3, 4}, []actionGet[int]{}, []int{4, 3, 2}, false},
		{"intSetDeleteRare", 3, []int{1, 2, 1, 3, 4}, []actionGet[int]{}, []int{4, 3, 1}, false},
		
	}
	for _, tc := range testCases {
		runTestCase(t, tc)
	}
}

func TestString(t *testing.T) {
	var testCases []testCase[string] = []testCase[string] {
		{"stringSet", 4, []string{"ab", "bb", "c", "ab", "d", "ee"}, 
		[]actionGet[string]{{"ab", true}, {"eee", false}},
		[]string{"ab", "ee", "d", "c"}, false},
	}
	for _, tc := range testCases {
		runTestCase(t, tc)
	}
}

func TestFloat64(t *testing.T) {
	var testCases []testCase[float64] = []testCase[float64]{
		{"doubleClear", 3, []float64{8.9, 99.0, 21.0, 1.333}, []actionGet[float64]{}, 
		[]float64{}, true},
	}
	for _, tc := range testCases {
		runTestCase(t, tc)
	}
}