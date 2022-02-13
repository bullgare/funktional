# Funktional

Is an implementation of basic functional programming helpers in go.

It does not have any external dependency.

## Compatibility

**golang 1.18 or above**, as it is using generics.

## Functions

Of course, they work with any type, not just integers or strings.

### For slices

[More detailed examples](./slices/slices_example_test.go).

| Function       | Example                                                                                                                                      | Description                                                                                                                                 |
|----------------|----------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| Chunk          | `Chunk([]int{1, 2, 3, 4}, 3)`                                                                                                                | creates an array of elements splitted into groups the length of size.                                                                       |
| Copy           | `Copy([]int{1, 2, 3, 4})`                                                                                                                    | creates a shallow copy of the given slice.                                                                                                  |
| Fill           | `Fill([]int{1, 2, 3, 4}, 1, 2, 4)`                                                                                                           | fills elements of array with value from start up to, but not including, end.                                                                |
| Filter         | `Filter([]int{1, 2, 3, 4}, func(i int, _ int, _ []int) bool { return i%2 == 0 })`                                                            | iterates over a slice and returns a new slice with values filtered by given function.                                                       |
| FindIndex      | `FindIndex([]int{1, 2, 3, 4}, func(i int) bool { return i == 3 })`                                                                           | iterates over elements of collection, returning the first index assertion returns truthy for. If no valid was found, return -1.             |
| ForEach        | `ForEach([]string{"a", "b", "c", "d"}, func(s string, pos int) { fmt.Println(s) })`                                                          | runs given function for each element of a slice.                                                                                            |
| Map            | `Map([]int{1, 2, 3, 4}, func(i int, _ int, _ []int) int { return i + i })`                                                                   | creates a slice by iterating over a given slice and applying a function to it.                                                              |
| Reduce         | `Reduce([]int{1, 2, 3, 4}, func(acc string, v int, _ int) string { if len(acc) > 0 {acc += ", "}; acc += strconv.Itoa(v); return acc }, "")` | iterates over a slice and reduces it to a given accumulator.                                                                                |
| Remove         | `Remove([]string{"a", "b", "c", "d"}, func(s string, pos int) bool { return s == "b" })`                                                     | from the slice given all values assertion returns truthy for. Returns 2 slices: cleaned slice and all removed elements (keeping the order). |
| ReverseInPlace | `ReverseInPlace([]string{"a", "b", "c", "d"})`                                                                                               | reverses original slice elements order. Mutates original slice.                                                                             |

### For maps

[More detailed examples](./maps/maps_example_test.go)

| Function      | Example                                                                                                                   | Description                                                                                                                                |
|---------------|---------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------|
| Copy          | `Copy(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})`                                                                    | creates a shallow copy of a map.                                                                                                           |
| Filter        | `Filter(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}, func(v int, _ string, _ map[string]int) bool { { return v < 3 })` | iterates over a map and returns a new map with values filtered by a given function.                                                        |
| FindKeyBy     | `FindKeyBy(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}, func(v int) bool { return v == 2 })`                           | iterates over a map, returning a pointer to the first (random) key assertion returns truthy for. If no valid value was found, returns nil. |
| FindAllKeysBy | `FindAllKeysBy(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}, func(v int) bool { return v < 3 })`                        | iterates over a map, returning a slice of keys assertion returns truthy for. If no valid value was found, returns nil.                     |
| ForEach       | `ForEach(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}, func(v int, k string) { fmt.Println(k, v) })`                    | runs given function for each element of a map.                                                                                             |
| Invert        | `Invert(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})`                                                                  | creates a new map switching the keys and values from the original map (k->v, v->k)                                                         |                                                                                                                                            |
| InvertBy      | `InvertBy(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}, func(v int) float64 { return float64(v) }) `                    | creates a new map switching the keys and values from the original map and a function applied to the values (k->v, fn(v)->k).               |                                                                                                                                            |
| InvertGrouped | `InvertGrouped(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 1})`                                                   | creates a new map switching the keys and values from the original map (k->[]v, v->k).                                                      |
| Keys          | `Keys(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})`                                                                    | returns all map keys in random order.                                                                                                      |
| Map           | `Map(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}, func(v int, k string, all map[string]int) int { return v * 2 })`     | creates a new map by iterating over a given map and applying a function to it.                                                             |
| Reduce        | `Reduce(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}, func(acc int, v int, k string) int { return acc + v }, 0)`        | iterates over a map and reduces it to a given accumulator.                                                                                 |
| Values        | `Values(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})`                                                                  | returns all map values in random order.                                                                                                    |
