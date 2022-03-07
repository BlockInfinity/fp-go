package fp

// Determines whether all the members of an array satisfy the specified test.
func Every[T any](predicate func(T) bool) func([]T) bool {
	return func(xs []T) bool {

		for _, x := range xs {
			if !predicate(x) {
				return false
			}
		}

		return true
	}
}

func EveryWithIndex[T any](predicate func(T, int) bool) func([]T) bool {
	return func(xs []T) bool {

		for i, x := range xs {
			if !predicate(x, i) {
				return false
			}
		}

		return true
	}
}

func EveryWithSlice[T any](predicate func(T, int, []T) bool) func([]T) bool {
	return func(xs []T) bool {

		for i, x := range xs {
			if !predicate(x, i, xs) {
				return false
			}
		}

		return true
	}
}
