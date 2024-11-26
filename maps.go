package fp

func EvaluateMap[K comparable, V any](m map[K](func() V)) map[K]V {
	res := make(map[K]V)

	for k, f := range m {
		res[k] = f()
	}

	return res
}

func EvaluateMapWithError[K comparable, V any](m map[K](func() (V, error))) (map[K]V, error) {
	res := make(map[K]V)

	for k, f := range m {
		v, err := f()
		if err != nil {
			return nil, err
		}

		res[k] = v
	}

	return res, nil
}
