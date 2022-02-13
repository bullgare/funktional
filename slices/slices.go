package slices

// Map creates a slice by iterating over a given slice and applying a function to it.
func Map[T, Y any](in []T, convert func(T, int, []T) Y) []Y {
	if in == nil {
		return nil
	}

	res := make([]Y, 0, len(in))

	for i, elem := range in {
		res = append(res, convert(elem, i, in))
	}

	return res
}

// Filter iterates over a slice and returns a new slice with values filtered by a given function.
func Filter[T any](in []T, filter func(T, int, []T) bool) []T {
	if in == nil {
		return nil
	}

	res := make([]T, 0, len(in))

	for i, elem := range in {
		if filter(elem, i, in) {
			res = append(res, elem)
		}
	}

	return res
}

// Reduce iterates over a slice and reduces it to a given accumulator.
func Reduce[T, Y any](in []T, reduce func(Y, T, int) Y, acc Y) Y {
	if in == nil {
		return acc
	}

	for i, elem := range in {
		acc = reduce(acc, elem, i)
	}

	return acc
}

// ForEach runs given function for each element of a slice.
func ForEach[T any](in []T, fn func(T, int)) {
	for pos, elem := range in {
		fn(elem, pos)
	}
}

// Copy creates a shallow copy of the given slice.
func Copy[T any](in []T) []T {
	if in == nil {
		return nil
	}

	out := make([]T, len(in))
	copy(out, in)

	return out
}

// Chunk creates an array of elements splitted into groups the length of size.
func Chunk[T any](in []T, size int) [][]T {
	if in == nil {
		return nil
	}

	if size < 1 {
		size = 1
	}
	res := make([][]T, 0, 1+len(in)/size)
	var current []T

	for i := 0; i < len(in); i++ {
		if (i % size) == 0 {
			if current != nil {
				res = append(res, current)
			}
			current = make([]T, 0, size)
		}
		current = append(current, in[i])
	}
	res = append(res, current)

	return res
}

// Fill fills elements of array with value from start up to, but not including, end.
// Mutates original slice.
func Fill[T any](in []T, value T, start, end int) []T {
	if start > end || start > len(in) || end < 1 {
		return in
	}

	if start < 0 {
		start = 0
	}
	if end > len(in) {
		end = len(in)
	}

	for i := start; i < end; i++ {
		in[i] = value
	}

	return in
}

// FindIndex iterates over elements of collection, returning the first index assertion returns truthy for.
// If no valid value was found, return -1.
func FindIndex[T any](in []T, assertion func(T) bool, fromIndex ...int) int {
	startFrom := 0
	if len(fromIndex) > 0 {
		startFrom = fromIndex[0]
	}
	if startFrom >= len(in) {
		return -1
	}
	if startFrom < 0 {
		startFrom = 0
	}

	for i := startFrom; i < len(in); i++ {
		if assertion(in[i]) {
			return i
		}
	}

	return -1
}

// Remove from the slice given all values assertion returns truthy for. Original slice stays untouched.
// Returns 2 slices: cleaned slice and all removed elements (keeping the order).
// It is not optimal in terms of space used - it pre-allocates 1.5x of initial slice length.
func Remove[T any](in []T, assertion func(T, int) bool) ([]T, []T) {
	if in == nil {
		return nil, nil
	}

	out := make([]T, 0, len(in))
	removed := make([]T, 0, len(in)/2)
	for i, v := range in {
		if assertion(v, i) {
			removed = append(removed, v)
		} else {
			out = append(out, v)
		}
	}

	return out, removed
}

// ReverseInPlace reverses original slice elements order. Mutates original slice.
func ReverseInPlace[T any](in []T) {
	if in == nil {
		return
	}

	l := len(in)
	for i := 0; i < l/2; i++ {
		in[i], in[l-1-i] = in[l-1-i], in[i]
	}
}
