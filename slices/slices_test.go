package slices

import (
	"reflect"
	"strconv"
	"testing"
)

func init() {
	testing.Init()
}

func Test_Map(t *testing.T) {
	tt := []struct {
		name     string
		in       []int
		convert  func(i int, _ int, _ []int) interface{} // it should have a real value instead of empty interface in real life
		expected []interface{}                           // it should have a real value instead of empty interface in real life
	}{
		{
			name:     "func(i int, _ int, _ []int) string { return strconv.Itoa(i) }",
			in:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			convert:  func(i int, _ int, _ []int) interface{} { return strconv.Itoa(i) },
			expected: []interface{}{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
		},
		{
			name:     "func(i int, _ int, _ []int) float64 { return float64(i) }",
			in:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			convert:  func(i int, _ int, _ []int) interface{} { return float64(i) / 10 },
			expected: []interface{}{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9},
		},
		{
			name:     "func(i int, _ int, _ []int) bool { return i % 2 == 0 }",
			in:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			convert:  func(i int, _ int, _ []int) interface{} { return i%2 == 0 },
			expected: []interface{}{false, true, false, true, false, true, false, true, false},
		},
		{
			name:     "nil in - nil out",
			in:       nil,
			convert:  func(i int, _ int, _ []int) interface{} { return i%2 == 0 },
			expected: nil,
		},
	}

	for _, tc := range tt {
		// t.Run is not working for me in go2 now
		res := Map(tc.in, tc.convert)
		if !reflect.DeepEqual(res, tc.expected) {
			t.Fatalf(`Map %s: expected 
				%#v, got 
				%#v`, tc.name, tc.expected, res)
		}
	}
}

func Test_Filter(t *testing.T) {
	tt := []struct {
		name     string
		in       []int
		filter   func(i int, _ int, _ []int) bool
		expected []int
	}{
		{
			name:     "func(i int, _ int, _ []int) bool { return i % 2 == 0 }",
			in:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			filter:   func(i int, _ int, _ []int) bool { return i%2 == 0 },
			expected: []int{2, 4, 6, 8},
		},
		{
			name:     "func(i int, _ int, _ []int) bool { return false }",
			in:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			filter:   func(i int, _ int, _ []int) bool { return false },
			expected: []int{},
		},
		{
			name:     "func(i int) bool { return true }",
			in:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			filter:   func(i int, _ int, _ []int) bool { return true },
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "nil in - nil out",
			in:       nil,
			filter:   func(i int, _ int, _ []int) bool { return true },
			expected: nil,
		},
	}

	for _, tc := range tt {
		res := Filter(tc.in, tc.filter)
		if !reflect.DeepEqual(res, tc.expected) {
			t.Fatalf(`Map %s: expected
				%#v, got
				%#v`, tc.name, tc.expected, res)
		}
	}
}

func Test_Reduce(t *testing.T) {
	tt := []struct {
		name        string
		in          []int
		reduce      func(acc int, elem int, _ int) int
		accumulator int
		expected    int
	}{
		{
			name: "sum: func(int, int, int) int",
			in:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			reduce: func(acc int, elem int, _ int) int {
				return acc + elem
			},
			accumulator: 0,
			expected:    45,
		},
		{
			name: "count: func(int, int, int) int",
			in:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			reduce: func(acc int, _ int, _ int) int {
				return acc + 1
			},
			accumulator: 0,
			expected:    9,
		},
		{
			name: "count: func(int, int, int) int",
			in:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			reduce: func(acc int, _ int, index int) int {
				return acc + index
			},
			accumulator: 0,
			expected:    36,
		},
		{
			name: "nil in - initial accumulator",
			in:   nil,
			reduce: func(acc int, _ int, index int) int {
				return acc + index
			},
			accumulator: 0,
			expected:    0,
		},
	}

	for _, tc := range tt {
		res := Reduce(tc.in, tc.reduce, tc.accumulator)
		if !reflect.DeepEqual(res, tc.expected) {
			t.Fatalf(`Reduce %s: expected
				%#v, got
				%#v`, tc.name, tc.expected, res)
		}
	}
}

func Test_ForEach(t *testing.T) {
	in := []string{"a", "b", "c", "d"}
	expectedRes := "a0,b1,c2,d3,"

	acc := ""
	ForEach(in, func(s string, pos int) {
		acc += s + strconv.Itoa(pos) + ","
	})

	if !reflect.DeepEqual(acc, expectedRes) {
		t.Fatalf(`ForEach: expected
				%#v, got
				%#v`, expectedRes, acc)
	}
}

func Test_Copy_Integers(t *testing.T) {
	tt := []struct {
		name        string
		in          []int
		expectedOut []int
	}{
		{
			name:        "happy path",
			in:          []int{1, 2, 3},
			expectedOut: []int{1, 2, 3},
		},
		{
			name:        "nil - nil",
			in:          nil,
			expectedOut: nil,
		},
		{
			name:        "empty slice - empty slice",
			in:          []int{},
			expectedOut: []int{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := Copy(tc.in)
			if tc.in != nil {
				// altering initial slice should not affect the result
				tc.in = append(tc.in, 4)
			}
			if !reflect.DeepEqual(res, tc.expectedOut) {
				t.Fatalf(`Copy %s: expected
				%#v, got
				%#v`, tc.name, tc.expectedOut, res)
			}
		})
	}
}

func Test_Copy_DeepCopyShouldNotWork(t *testing.T) {
	type my struct {
		a string
	}

	in := []*my{
		{a: "a"},
		{a: "b"},
	}
	res := Copy(in)
	in[0].a = "c"
	expectedRes := []*my{
		{a: "c"}, // we are expecting it to change here as we do not handle deep copy
		{a: "b"},
	}

	if !reflect.DeepEqual(res, expectedRes) {
		t.Fatalf(`Copy should not care about deep equality: expected
				%#v, got
				%#v`, expectedRes, res)
	}
}

func Test_Chunk(t *testing.T) {
	tt := []struct {
		name     string
		in       []int
		size     int
		expected [][]int
	}{
		{
			name:     "by 3",
			in:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			size:     3,
			expected: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
		{
			name:     "by 2",
			in:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			size:     2,
			expected: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9}},
		},
		{
			name:     "by 1",
			in:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			size:     1,
			expected: [][]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}},
		},
		{
			name:     "by 0",
			in:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			size:     0,
			expected: [][]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}},
		},
		{
			name:     "by -1",
			in:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			size:     -1,
			expected: [][]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}},
		},
		{
			name:     "by 10",
			in:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			size:     10,
			expected: [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},
		{
			name:     "by 100",
			in:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			size:     100,
			expected: [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},
		{
			name:     "by 9",
			in:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			size:     9,
			expected: [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},
		{
			name:     "by 8",
			in:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			size:     8,
			expected: [][]int{{1, 2, 3, 4, 5, 6, 7, 8}, {9}},
		},
		{
			name:     "nil in - nil out",
			in:       nil,
			size:     8,
			expected: nil,
		},
	}

	for _, tc := range tt {
		res := Chunk(tc.in, tc.size)
		if !reflect.DeepEqual(res, tc.expected) {
			t.Fatalf(`Chunk %s: expected 
				%#v, got 
				%#v`, tc.name, tc.expected, res)
		}
	}
}

func Test_Fill(t *testing.T) {
	tt := []struct {
		name     string
		in       []int
		value    int
		from     int
		to       int
		expected []int
	}{
		{
			name:     "3 to 6 by 1",
			in:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			value:    1,
			from:     3,
			to:       6,
			expected: []int{1, 2, 3, 1, 1, 1, 7, 8, 9},
		},
		{
			name:     "0 to 9 by 1",
			in:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			value:    1,
			from:     0,
			to:       9,
			expected: []int{1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name:     "-1 to 3 by 1",
			in:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			value:    1,
			from:     -1,
			to:       3,
			expected: []int{1, 1, 1, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "6 to 15 by 1",
			in:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			value:    1,
			from:     6,
			to:       15,
			expected: []int{1, 2, 3, 4, 5, 6, 1, 1, 1},
		},
		{
			name:     "6 to 3 by 1",
			in:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			value:    1,
			from:     6,
			to:       3,
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "10 to 15 by 1",
			in:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			value:    1,
			from:     6,
			to:       3,
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "-3 to -1 by 1",
			in:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			value:    1,
			from:     -3,
			to:       -1,
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "nil in - nil out",
			in:       nil,
			value:    1,
			from:     0,
			to:       0,
			expected: nil,
		},
	}

	for _, tc := range tt {
		res := Fill(tc.in, tc.value, tc.from, tc.to)
		if !reflect.DeepEqual(res, tc.expected) {
			t.Fatalf(`Fill %s: expected 
				%#v, got 
				%#v`, tc.name, tc.expected, res)
		}
	}
}

func Test_FindIndex(t *testing.T) {
	tt := []struct {
		name      string
		in        []int
		predicate func(i int) bool
		fromIndex *int
		expected  int
	}{
		{
			name:      "func(i int) bool { return i == 3 }",
			in:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			predicate: func(i int) bool { return i == 3 },
			fromIndex: nil,
			expected:  2,
		},
		{
			name:      "func(i int) bool { return i == 30 }",
			in:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			predicate: func(i int) bool { return i == 30 },
			fromIndex: nil,
			expected:  -1,
		},
		{
			name:      "func(i int) bool { return i == 3 }, from index=1",
			in:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			predicate: func(i int) bool { return i == 3 },
			fromIndex: func() *int { i := 1; return &i }(),
			expected:  2,
		},
		{
			name:      "func(i int) bool { return i == 3 }, from index=5",
			in:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			predicate: func(i int) bool { return i == 3 },
			fromIndex: func() *int { i := 5; return &i }(),
			expected:  -1,
		},
		{
			name:      "func(i int) bool { return i == 3 }, from index=15",
			in:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			predicate: func(i int) bool { return i == 3 },
			fromIndex: func() *int { i := 15; return &i }(),
			expected:  -1,
		},
		{
			name:      "func(i int) bool { return i == 3 }, from index=-1",
			in:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			predicate: func(i int) bool { return i == 3 },
			fromIndex: func() *int { i := -1; return &i }(),
			expected:  2,
		},
		{
			name:      "nil in - -1",
			in:        nil,
			predicate: func(i int) bool { return i == 3 },
			fromIndex: nil,
			expected:  -1,
		},
	}

	for _, tc := range tt {
		var res int
		if tc.fromIndex != nil {
			res = FindIndex(tc.in, tc.predicate, *tc.fromIndex)
		} else {
			res = FindIndex(tc.in, tc.predicate)
		}
		if !reflect.DeepEqual(res, tc.expected) {
			t.Fatalf(`Fill %s: expected
				%#v, got
				%#v`, tc.name, tc.expected, res)
		}
	}
}

func Test_Remove(t *testing.T) {
	tt := []struct {
		name            string
		in              []string
		assertion       func(string, int) bool
		expectedOut     []string
		expectedRemoved []string
	}{
		{
			name: "happy path",
			in:   []string{"a", "b", "c", "d"},
			assertion: func(s string, pos int) bool {
				return s == "b" || pos == 2
			},
			expectedOut:     []string{"a", "d"},
			expectedRemoved: []string{"b", "c"},
		},
		{
			name: "nil in - nil res",
			in:   nil,
			assertion: func(s string, pos int) bool {
				return s == "b" || pos == 2
			},
			expectedOut:     nil,
			expectedRemoved: nil,
		},
		{
			name: "empty in - empty res",
			in:   []string{},
			assertion: func(s string, pos int) bool {
				return s == "b" || pos == 2
			},
			expectedOut:     []string{},
			expectedRemoved: []string{},
		},
		{
			name: "remove none",
			in:   []string{"a", "b", "c", "d"},
			assertion: func(s string, pos int) bool {
				return false
			},
			expectedOut:     []string{"a", "b", "c", "d"},
			expectedRemoved: []string{},
		},
		{
			name: "remove all",
			in:   []string{"a", "b", "c", "d"},
			assertion: func(s string, pos int) bool {
				return true
			},
			expectedOut:     []string{},
			expectedRemoved: []string{"a", "b", "c", "d"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res, removed := Remove(tc.in, tc.assertion)
			if !reflect.DeepEqual(res, tc.expectedOut) || !reflect.DeepEqual(removed, tc.expectedRemoved) {
				t.Fatalf(`Fill %s: expected 
				%#v and %#v, got 
				%#v and %#v`, tc.name, tc.expectedOut, tc.expectedRemoved, res, removed)
			}
		})
	}
}

func Test_ReverseInPlace(t *testing.T) {
	tt := []struct {
		name        string
		in          []int
		expectedRes []int
	}{
		{
			name:        "happy path odd",
			in:          []int{1, 2, 3},
			expectedRes: []int{3, 2, 1},
		},
		{
			name:        "happy path even",
			in:          []int{1, 2, 3, 4},
			expectedRes: []int{4, 3, 2, 1},
		},
		{
			name:        "happy path - values do not matter",
			in:          []int{100, 2, 30, 4},
			expectedRes: []int{4, 30, 2, 100},
		},
		{
			name:        "nil value is also valid",
			in:          nil,
			expectedRes: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ReverseInPlace(tc.in)

			if !reflect.DeepEqual(tc.in, tc.expectedRes) {
				t.Fatalf(`Fill %s: expected
				%#v, got
				%#v`, tc.name, tc.expectedRes, tc.in)
			}
		})
	}
}
