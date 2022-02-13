package maps

// Map creates a new map by iterating over a given map and applying a function to it.
func Map[T, Y any, K comparable](in map[K]T, convert func(T, K, map[K]T) Y) map[K]Y {
	if in == nil {
		return nil
	}

	out := make(map[K]Y, len(in))
	for k, v := range in {
		out[k] = convert(v, k, in)
	}

	return out
}

// Filter iterates over a map and returns a new map with values filtered by a given function.
func Filter[T any, K comparable](in map[K]T, filter func(T, K, map[K]T) bool) map[K]T {
	if in == nil {
		return nil
	}

	res := make(map[K]T, len(in))

	for k, v := range in {
		if filter(v, k, in) {
			res[k] = v
		}
	}

	return res
}

// Reduce iterates over a map and reduces it to a given accumulator.
func Reduce[T, Y any, K comparable](in map[K]T, reduce func(Y, T, K) Y, acc Y) Y {
	if in == nil {
		return acc
	}

	for k, v := range in {
		acc = reduce(acc, v, k)
	}

	return acc
}

// ForEach runs given function for each element of a map.
func ForEach[T any, K comparable](in map[K]T, fn func(T, K)) {
	for k, v := range in {
		fn(v, k)
	}
}

// Copy creates a shallow copy of a map.
func Copy[T any, K comparable](in map[K]T) map[K]T {
	if in == nil {
		return nil
	}

	out := make(map[K]T, len(in))
	for k, v := range in {
		out[k] = v
	}

	return out
}

// Keys returns all map keys in random order.
func Keys[T any, K comparable](in map[K]T) []K {
	if in == nil {
		return nil
	}

	out := make([]K, 0, len(in))
	for k := range in {
		out = append(out, k)
	}
	return out
}

// Values returns all map values in random order.
func Values[T any, K comparable](in map[K]T) []T {
	if in == nil {
		return nil
	}

	out := make([]T, 0, len(in))
	for _, v := range in {
		out = append(out, v)
	}
	return out
}

// FindKeyBy iterates over a map, returning a pointer to the first (random) key assertion returns truthy for.
// If no valid value was found, returns nil.
func FindKeyBy[T any, K comparable](in map[K]T, assertion func(T) bool) *K {
	for k, v := range in {
		if assertion(v) {
			return &k
		}
	}

	return nil
}

// FindAllKeysBy iterates over a map, returning a slice of keys assertion returns truthy for.
// If no valid value was found, returns nil.
func FindAllKeysBy[T any, K comparable](in map[K]T, assertion func(T) bool) []K {
	if len(in) == 0 {
		return nil
	}

	var out []K

	for k, v := range in {
		if assertion(v) {
			out = append(out, k)
		}
	}

	return out
}

// Invert creates a new map switching the keys and values from the original map (k->v, v->k).
// If original map has non-unique values, the output map will have one of the original keys as a value (as map keys order is random).
func Invert[K1, K2 comparable](in map[K1]K2) map[K2]K1 {
	if in == nil {
		return nil
	}

	out := make(map[K2]K1, len(in))
	for k, v := range in {
		out[v] = k
	}
	return out
}

// InvertBy creates a new map switching the keys and values from the original map and a function applied to the values (k->v, fn(v)->k).
// If a function given generates non-unique results, the output map will have one of the results as a value.
func InvertBy[T any, K1, K2 comparable](in map[K1]T, fn func(v T) K2) map[K2]K1 {
	if in == nil {
		return nil
	}

	out := make(map[K2]K1, len(in))
	for k, v := range in {
		out[fn(v)] = k
	}
	return out
}

// InvertGrouped creates a new map switching the keys and values from the original map (k->[]v, v->k).
// The value is always a slice of original keys.
func InvertGrouped[K1, K2 comparable](in map[K1]K2) map[K2][]K1 {
	if in == nil {
		return nil
	}

	out := make(map[K2][]K1, len(in))
	for k, v := range in {
		prev := out[v]
		if prev == nil {
			prev = make([]K1, 0, 1)
		}
		out[v] = append(prev, k)
	}
	return out
}
