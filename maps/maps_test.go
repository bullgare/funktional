package maps

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func Test_Map(t *testing.T) {
	tt := []struct {
		name     string
		in       map[interface{}]int
		convert  func(int, interface{}, map[interface{}]int) interface{}
		expected map[interface{}]interface{}
	}{
		{
			name: "happy path",
			in: map[interface{}]int{
				"a": 1,
				2:   2,
			},
			convert: func(v int, k interface{}, in map[interface{}]int) interface{} {
				return v * 2
			},
			expected: map[interface{}]interface{}{
				"a": 2,
				2:   4,
			},
		},
		{
			name: "nil in - nil out",
			in:   nil,
			convert: func(v int, k interface{}, in map[interface{}]int) interface{} {
				return v * 2
			},
			expected: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := Map(tc.in, tc.convert)

			if !reflect.DeepEqual(res, tc.expected) {
				t.Fatalf(`Map %s: expected 
				%#v, got 
				%#v`, tc.name, tc.expected, res)
			}
		})
	}
}

func Test_Filter(t *testing.T) {
	tt := []struct {
		name     string
		in       map[string]int
		filter   func(int, string, map[string]int) bool
		expected map[string]int
	}{
		{
			name: "only even values",
			in:   map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			filter: func(v int, k string, in map[string]int) bool {
				return v%2 == 0
			},
			expected: map[string]int{"b": 2, "d": 4},
		},
		{
			name: "only odd values",
			in:   map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			filter: func(v int, k string, in map[string]int) bool {
				return v%2 == 1
			},
			expected: map[string]int{"a": 1, "c": 3},
		},
		{
			name: "nil in - nil out",
			in:   nil,
			filter: func(v int, k string, in map[string]int) bool {
				return v%2 == 1
			},
			expected: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := Filter(tc.in, tc.filter)

			if !reflect.DeepEqual(res, tc.expected) {
				t.Fatalf(`Filter %s: expected 
				%#v, got 
				%#v`, tc.name, tc.expected, res)
			}
		})
	}
}

func Test_Reduce(t *testing.T) {
	tt := []struct {
		name        string
		in          map[string]int
		reduce      func(acc map[string]bool, v int, k string) map[string]bool
		accumulator map[string]bool
		expectedRes map[string]bool
	}{
		{
			name: "concat all values",
			in:   map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			reduce: func(acc map[string]bool, v int, k string) map[string]bool {
				acc[strconv.Itoa(v)] = true
				return acc
			},
			accumulator: map[string]bool{},
			expectedRes: map[string]bool{"1": true, "2": true, "3": true, "4": true},
		},
		{
			name: "concat all keys",
			in:   map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			reduce: func(acc map[string]bool, v int, k string) map[string]bool {
				acc[k] = true
				return acc
			},
			accumulator: map[string]bool{},
			expectedRes: map[string]bool{"a": true, "b": true, "c": true, "d": true},
		},
		{
			name: "concat all keys and values",
			in:   map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			reduce: func(acc map[string]bool, v int, k string) map[string]bool {
				acc[k+strconv.Itoa(v)] = true
				return acc
			},
			accumulator: map[string]bool{},
			expectedRes: map[string]bool{"a1": true, "b2": true, "c3": true, "d4": true},
		},
		{
			name: "nil in - initial accumulator (no panic)",
			in:   nil,
			reduce: func(acc map[string]bool, v int, k string) map[string]bool {
				acc[k+strconv.Itoa(v)] = true
				return acc
			},
			accumulator: nil,
			expectedRes: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := Reduce(tc.in, tc.reduce, tc.accumulator)

			if !reflect.DeepEqual(res, tc.expectedRes) {
				t.Fatalf(`Reduce %s: expected 
				%#v, got 
				%#v`, tc.name, tc.expectedRes, res)
			}
		})
	}
}

func Test_ForEach(t *testing.T) {
	in := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	expectedRes := map[string]bool{"a1": true, "b2": true, "c3": true, "d4": true}

	acc := map[string]bool{}
	ForEach(in, func(v int, k string) {
		acc[k+strconv.Itoa(v)] = true
	})

	if !reflect.DeepEqual(acc, expectedRes) {
		t.Fatalf(`ForEach: expected
				%#v, got
				%#v`, expectedRes, acc)
	}
}

func Test_Copy(t *testing.T) {
	tt := []struct {
		name     string
		in       map[string]int
		expected map[string]int
	}{
		{
			name:     "happy path",
			in:       map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			expected: map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
		},
		{
			name:     "nil in - nil out",
			in:       nil,
			expected: nil,
		},
		{
			name:     "empty in - empty out",
			in:       map[string]int{},
			expected: map[string]int{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := Copy(tc.in)

			if !reflect.DeepEqual(res, tc.expected) {
				t.Fatalf(`Copy: expected
				%#v, got
				%#v`, tc.expected, res)
			}
		})
	}
}

func Test_Keys(t *testing.T) {
	tt := []struct {
		name     string
		in       map[string]int
		expected []string
	}{
		{
			name:     "happy path",
			in:       map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			expected: []string{"a", "b", "c", "d"},
		},
		{
			name:     "nil in - nil out",
			in:       nil,
			expected: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := Keys(tc.in)
			// sorting as map order is undefined
			sort.Strings(res)
			sort.Strings(tc.expected)

			if !reflect.DeepEqual(res, tc.expected) {
				t.Fatalf(`Keys: expected
				%#v, got
				%#v`, tc.expected, res)
			}
		})
	}
}

func Test_Values(t *testing.T) {
	tt := []struct {
		name     string
		in       map[string]int
		expected []int
	}{
		{
			name:     "happy path",
			in:       map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "nil in - nil out",
			in:       nil,
			expected: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := Values(tc.in)
			// sorting as map order is undefined
			sort.Ints(res)
			sort.Ints(tc.expected)

			if !reflect.DeepEqual(res, tc.expected) {
				t.Fatalf(`Values: expected
				%#v, got
				%#v`, tc.expected, res)
			}
		})
	}
}

func Test_FindKeyBy(t *testing.T) {
	tt := []struct {
		name      string
		in        map[string]int
		assertion func(int) bool
		expected  *string
	}{
		{
			name: "happy path",
			in:   map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			assertion: func(v int) bool {
				return v == 1
			},
			expected: func() *string { s := "a"; return &s }(),
		},
		{
			name: "value not found - nil returned",
			in:   map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			assertion: func(v int) bool {
				return v == -1
			},
			expected: nil,
		},
		// hard to test value collision case as map range order is random
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := FindKeyBy(tc.in, tc.assertion)

			if !reflect.DeepEqual(res, tc.expected) {
				t.Fatalf(`FindKeyBy: expected
				%#v, got
				%#v`, tc.expected, res)
			}
		})
	}
}

func Test_FindAllKeysBy(t *testing.T) {
	tt := []struct {
		name      string
		in        map[string]int
		assertion func(int) bool
		expected  []string
	}{
		{
			name: "happy path",
			in:   map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			assertion: func(v int) bool {
				return v < 3
			},
			expected: []string{"a", "b"},
		},
		{
			name: "values collision",
			in:   map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 1},
			assertion: func(v int) bool {
				return v < 3
			},
			expected: []string{"a", "b", "e"},
		},
		{
			name: "value not found - nil returned",
			in:   map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			assertion: func(v int) bool {
				return v == -1
			},
			expected: nil,
		},
		// hard to test value collision case as map range order is random
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := FindAllKeysBy(tc.in, tc.assertion)
			// sorting as map order is undefined
			sort.Strings(res)
			sort.Strings(tc.expected)

			if !reflect.DeepEqual(res, tc.expected) {
				t.Fatalf(`FindKeyBy: expected
				%#v, got
				%#v`, tc.expected, res)
			}
		})
	}
}

func Test_Invert(t *testing.T) {
	tt := []struct {
		name     string
		in       map[string]int
		expected map[int]string
	}{
		{
			name:     "happy path",
			in:       map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			expected: map[int]string{1: "a", 2: "b", 3: "c", 4: "d"},
		},
		{
			name:     "nil in - nil out",
			in:       nil,
			expected: nil,
		},
		// hard to test value collision case as map range order is random
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := Invert(tc.in)

			if !reflect.DeepEqual(res, tc.expected) {
				t.Fatalf(`Invert: expected
				%#v, got
				%#v`, tc.expected, res)
			}
		})
	}
}

func Test_InvertBy(t *testing.T) {
	tt := []struct {
		name     string
		in       map[string]int
		fn       func(int) float64
		expected map[float64]string
	}{
		{
			name: "happy path",
			in:   map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			fn: func(v int) float64 {
				return float64(v)
			},
			expected: map[float64]string{1.: "a", 2.: "b", 3.: "c", 4.: "d"},
		},
		{
			name:     "nil in - nil out",
			in:       nil,
			expected: nil,
		},
		// hard to test value collision case as map range order is random
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := InvertBy(tc.in, tc.fn)

			if !reflect.DeepEqual(res, tc.expected) {
				t.Fatalf(`InvertBy: expected
				%#v, got
				%#v`, tc.expected, res)
			}
		})
	}
}

func Test_InvertGrouped(t *testing.T) {
	tt := []struct {
		name     string
		in       map[string]int
		expected map[int][]string
	}{
		{
			name:     "happy path",
			in:       map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			expected: map[int][]string{1: {"a"}, 2: {"b"}, 3: {"c"}, 4: {"d"}},
		},
		{
			name:     "nil in - nil out",
			in:       nil,
			expected: nil,
		},
		{
			name:     "value collision case - multiple values in a slice",
			in:       map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 1, "f": 1, "g": 3},
			expected: map[int][]string{1: {"a", "e", "f"}, 2: {"b"}, 3: {"c", "g"}, 4: {"d"}},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := InvertGrouped(tc.in)
			for _, v := range res {
				// sorting as map order is undefined
				sort.Strings(v)
			}

			if !reflect.DeepEqual(res, tc.expected) {
				t.Fatalf(`InvertGrouped: expected
				%#v, got
				%#v`, tc.expected, res)
			}
		})
	}
}
