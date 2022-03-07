package fp

// Calls the specified callback function for all the elements in an array. The return value of the callback function is the accumulated result, and is provided as an argument in the next call to the callback function.
func Reduce[T any, R any](callback func(R, T) R, acc R) func([]T) R {
	return func(xs []T) R {

		for _, x := range xs {
			acc = callback(acc, x)
		}

		return acc
	}
}

func ReduceWithIndex[T any, R any](callback func(R, T, int) R, acc R) func([]T) R {
	return func(xs []T) R {

		for i, x := range xs {
			acc = callback(acc, x, i)
		}

		return acc
	}
}

func ReduceWithSlice[T any, R any](callback func(R, T, int, []T) R, acc R) func([]T) R {
	return func(xs []T) R {

		for i, x := range xs {
			acc = callback(acc, x, i, xs)
		}

		return acc
	}
}
