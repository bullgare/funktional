# Funktional

Is an implementation of basic functional programming helpers in go.

It is still work in progress.

## Compatibility

**golang 1.18 or above**, as it is using generics.

## Functions

Of course, they work with any values, not just integers or strings.

### For slices

[More detailed examples](./slices/slices_example_test.go).

| Function       | Example                                                                                                                                  | Description                                                                                                                                 |
|----------------|------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| Chunk          | Chunk([]int{1, 2, 3, 4}, 3)                                                                                                              | creates an array of elements splitted into groups the length of size.                                                                       |
| Copy           | Copy([]int{1, 2, 3, 4})                                                                                                                  | creates a shallow copy of the given slice.                                                                                                  |
| Fill           | Fill([]int{1, 2, 3, 4}, 1, 2, 4)                                                                                                         | fills elements of array with value from start up to, but not including, end.                                                                |
| Filter         | Filter([]int{1, 2, 3, 4}, func(i int, _ int, _ []int) bool { return i%2 == 0 })                                                          | iterates through a slice and returns a new slice with values filtered by given function.                                                    |
| FindIndex      | FindIndex([]int{1, 2, 3, 4}, func(i int) bool { return i == 3 })                                                                         | iterates over elements of collection, returning the first index assertion returns truthy for. If no valid was found, return -1.             |
| ForEach        | ForEach([]string{"a", "b", "c", "d"}, func(s string, pos int) {fmt.Println(s)})                                                          | runs given function for each element of a slice.                                                                                            |
| Map            | Map([]int{1, 2, 3, 4}, func(i int, _ int, _ []int) int { return i + i })                                                                 | creates a slice by iterating over a given slice and applying a function to it.                                                              |
| Reduce         | Reduce([]int{1, 2, 3, 4}, func(acc string, v int, _ int) string {if len(acc) > 0 {acc += ", "}; acc += strconv.Itoa(v); return acc}, "") | iterates through a slice and reduces it to a given accumulator.                                                                             |
| Remove         | Remove([]string{"a", "b", "c", "d"}, func(s string, pos int) bool { return s == "b" })                                                   | from the slice given all values assertion returns truthy for. Returns 2 slices: cleaned slice and all removed elements (keeping the order). |
| ReverseInPlace | ReverseInPlace([]string{"a", "b", "c", "d"})                                                                                             | reverses original slice elements order. Mutates original slice.                                                                             |

### For maps

_TBD_
