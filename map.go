package fp

// Calls a defined callback function on each element of an array, and returns an array that contains the results.
func Map[T any, R any](callback func(T) R) func([]T) []R {
	return func(xs []T) []R {

		result := make([]R, 0, len(xs))

		for _, x := range xs {
			result = append(result, callback(x))
		}

		return result
	}
}

func MapWithError[T any, R any](callback func(T) (R, error)) func([]T) ([]R, error) {
	return func(xs []T) ([]R, error) {

		result := make([]R, 0, len(xs))

		for _, x := range xs {
			resultElement, err := callback(x)
			if err != nil {
				return make([]R, 0), err
			}

			result = append(result, resultElement)
		}

		return result, nil
	}
}

// See Map but callback receives index of element.
func MapWithIndex[T any, R any](callback func(T, int) R) func([]T) []R {
	return func(xs []T) []R {

		result := make([]R, 0, len(xs))

		for i, x := range xs {
			result = append(result, callback(x, i))
		}

		return result
	}
}

// Like Map but callback receives index of element and the whole array.
func MapWithSlice[T any, R any](callback func(T, int, []T) R) func([]T) []R {
	return func(xs []T) []R {

		result := make([]R, 0, len(xs))

		for i, x := range xs {
			result = append(result, callback(x, i, xs))
		}

		return result
	}
}
