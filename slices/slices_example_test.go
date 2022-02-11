package slices

import (
	"fmt"
	"strconv"
)

func ExampleMap() {
	ints := []int{1, 2, 3, 4}
	res1 := Map(ints, func(i int, _ int, _ []int) string { return strconv.Itoa(i) })
	fmt.Printf("%#v\n", res1)

	res2 := Map(ints, func(i int, _ int, _ []int) int { return i + i })
	fmt.Printf("%#v\n", res2)

	// Output:
	// []string{"1", "2", "3", "4"}
	// []int{2, 4, 6, 8}
}

func ExampleFilter() {
	ints := []int{1, 2, 3, 4}
	res1 := Filter(ints, func(i int, _ int, _ []int) bool { return i%2 == 0 })
	fmt.Printf("%#v\n", res1)

	strs := []string{"a", "b", "c", "d"}
	res2 := Filter(strs, func(s string, pos int, _ []string) bool { return s == "a" || pos == 3 })
	fmt.Printf("%#v\n", res2)

	// Output:
	// []int{2, 4}
	// []string{"a", "d"}
}

func ExampleReduce() {
	ints := []int{1, 2, 3, 4}
	res := Reduce(ints, func(acc string, v int, _ int) string {
		if len(acc) > 0 {
			acc += ", "
		}
		acc += strconv.Itoa(v)
		return acc
	}, "")
	fmt.Printf("%#v\n", res)

	// Output:
	// "1, 2, 3, 4"
}

func ExampleForEach() {
	in := []string{"a", "b", "c", "d"}

	ForEach(in, func(s string, pos int) {
		fmt.Println(s)
	})

	// Output:
	// a
	// b
	// c
	// d
}

func ExampleCopy() {
	in := []int{1, 2, 3, 4}

	out := Copy(in)
	fmt.Printf("%#v\n", out)

	// Output:
	// []int{1, 2, 3, 4}
}

func ExampleChunk() {
	ints := []int{1, 2, 3, 4}
	res := Chunk(ints, 3)
	fmt.Printf("%#v\n", res)

	strs := []string{"a", "b", "c", "d"}
	res2 := Chunk(strs, 3)
	fmt.Printf("%#v\n", res2)

	// Output:
	// [][]int{[]int{1, 2, 3}, []int{4}}
	// [][]string{[]string{"a", "b", "c"}, []string{"d"}}
}

func ExampleFill() {
	ints := make([]int, 5)
	// Fill from 2 to 4 with 1
	res := Fill(ints, 1, 2, 4)
	fmt.Printf("%#v\n", res)

	strs := make([]string, 5)
	// Fill from 1 to 3 with "bu"
	res2 := Fill(strs, "bu", 1, 3)
	fmt.Printf("%#v\n", res2)

	// Output:
	// []int{0, 0, 1, 1, 0}
	// []string{"", "bu", "bu", "", ""}
}

func ExampleFindIndex() {
	ints := []int{1, 2, 3, 4}
	res := FindIndex(ints, func(i int) bool { return i == 3 })
	fmt.Printf("%#v\n", res)

	strs := []string{"a", "b", "c", "d"}
	res2 := FindIndex(strs, func(s string) bool { return s == "b" })
	fmt.Printf("%#v\n", res2)

	// Output:
	// 2
	// 1
}

func ExampleRemove() {
	strs := []string{"a", "b", "c", "d"}
	out, removed := Remove(strs, func(s string, pos int) bool { return s == "b" || pos == 2 })
	fmt.Printf("%#v\n", out)
	fmt.Printf("%#v\n", removed)

	// Output:
	// []string{"a", "d"}
	// []string{"b", "c"}
}

func ExampleReverseInPlace() {
	strs := []string{"a", "b", "c", "d"}
	ReverseInPlace(strs)
	fmt.Printf("%#v\n", strs)

	// Output:
	// []string{"d", "c", "b", "a"}
}
