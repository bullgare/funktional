package maps

import (
	"fmt"
	"sort"
)

func ExampleMap() {
	in := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	out := Map(in, func(v int, k string, all map[string]int) int {
		return v * 2
	})
	fmt.Printf("%#v", out)

	// Output:
	// map[string]int{"a":2, "b":4, "c":6, "d":8}
}

func ExampleFilter() {
	in := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	out := Filter(in, func(v int, k string, all map[string]int) bool {
		return v < 3
	})
	fmt.Printf("%#v", out)

	// Output:
	// map[string]int{"a":1, "b":2}
}

func ExampleReduce() {
	in := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	out := Reduce(in, func(acc int, v int, k string) int { return acc + v }, 0)
	fmt.Printf("%#v", out)

	// Output:
	// 10
}

func ExampleForEach() {
	in := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	out := make([]string, 0, 4)
	ForEach(in, func(v int, k string) {
		out = append(out, fmt.Sprint(k, v))
	})
	sort.Strings(out)
	fmt.Printf("%#v", out)

	// Output:
	// []string{"a1", "b2", "c3", "d4"}
}

func ExampleCopy() {
	in := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	out := Copy(in)
	fmt.Printf("%#v", out)

	// Output:
	// map[string]int{"a":1, "b":2, "c":3, "d":4}
}

func ExampleKeys() {
	in := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	out := Keys(in)
	sort.Strings(out)
	fmt.Printf("%#v", out)

	// Output:
	// []string{"a", "b", "c", "d"}
}

func ExampleValues() {
	in := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	out := Values(in)
	sort.Ints(out)
	fmt.Printf("%#v", out)

	// Output:
	// []int{1, 2, 3, 4}
}

func ExampleFindKeyBy() {
	in := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	out := FindKeyBy(in, func(v int) bool {
		return v == 2
	})
	fmt.Printf("%#v", *out)

	// Output:
	// "b"
}

func ExampleFindAllKeysBy() {
	in := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	out := FindAllKeysBy(in, func(v int) bool {
		return v < 3
	})
	sort.Strings(out)
	fmt.Printf("%#v", out)

	// Output:
	// []string{"a", "b"}
}

func ExampleInvert() {
	in := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	out := Invert(in)
	fmt.Printf("%#v", out)

	// Output:
	// map[int]string{1:"a", 2:"b", 3:"c", 4:"d"}
}

func ExampleInvertBy() {
	in := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	out := InvertBy(in, func(v int) float64 {
		return float64(v)
	})
	fmt.Printf("%#v", out)

	// Output:
	// map[float64]string{1:"a", 2:"b", 3:"c", 4:"d"}
}

func ExampleInvertGrouped() {
	in := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 1}
	out := InvertGrouped(in)
	for _, v := range out {
		// sorting as map order is undefined, optional
		sort.Strings(v)
	}
	fmt.Printf("%#v", out)

	// Output:
	// map[int][]string{1:[]string{"a", "e"}, 2:[]string{"b"}, 3:[]string{"c"}, 4:[]string{"d"}}
}
