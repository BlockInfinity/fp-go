package fp

import "sync"

// Calls a defined callback function on each element of an array. Then, flattens the result into a new array. This is identical to a map followed by flat with depth 1.
func FlatMap[T any, R any](callback func(T) []R) func([]T) []R {
	return func(xs []T) []R {

		result := []R{}

		for _, x := range xs {
			result = append(result, callback(x)...)
		}

		return result
	}
}

// See FlatMap but callback receives index of element.
func FlatMapWithIndex[T any, R any](callback func(T, int) []R) func([]T) []R {
	return func(xs []T) []R {

		result := []R{}

		for i, x := range xs {
			result = append(result, callback(x, i)...)
		}

		return result
	}
}

// Like FlatMap but callback receives index of element and the whole array.
func FlatMapWithSlice[T any, R any](callback func(T, int, []T) []R) func([]T) []R {
	return func(xs []T) []R {

		result := []R{}

		for i, x := range xs {
			result = append(result, callback(x, i, xs)...)
		}

		return result
	}
}

func FlatMapWithError[T any, R any](callback func(T) ([]R, error)) func([]T) ([]R, error) {
	return func(xs []T) ([]R, error) {
		result := []R{}

		for _, x := range xs {
			targetValues, err := callback(x)
			if err != nil {
				return make([]R, 0), err
			}
			result = append(result, targetValues...)
		}

		return result, nil
	}
}

func FlatMapWithErrorAndTransform[T any, R any, Z any](callback func(T) ([]R, error),
	transform func(T, R) Z) func([]T) ([]Z, error) {
	return func(xs []T) ([]Z, error) {
		result := []Z{}

		for _, x := range xs {
			targetValues, err := callback(x)
			transformedValues := Map(func(targetValue R) Z {
				return transform(x, targetValue)
			})(targetValues)
			if err != nil {
				return make([]Z, 0), err
			}
			result = append(result, transformedValues...)
		}

		return result, nil
	}
}

func FlatMapWithErrorParallel[T any, R any](callback func(T) ([]R, error)) func([]T) ([]R, error) {
	return func(xs []T) ([]R, error) {
		// Result channel for collecting results.
		type result struct {
			values []R
			err    error
		}

		results := make(chan result, len(xs))
		var wg sync.WaitGroup

		// Start goroutines for parallel execution.
		for _, x := range xs {
			wg.Add(1)
			go func(x T) {
				defer wg.Done()
				res, err := callback(x)
				results <- result{values: res, err: err}
			}(x)
		}

		// Close the results channel when all goroutines complete.
		go func() {
			wg.Wait()
			close(results)
		}()

		// Collect results.
		finalResults := []R{}
		for r := range results {
			if r.err != nil {
				return nil, r.err
			}
			finalResults = append(finalResults, r.values...)
		}

		return finalResults, nil
	}
}
