package helpers

func MapArray[TInput any, TOutput any](values []TInput, handler func(value TInput) TOutput) []TOutput {
	var result []TOutput

	for _, v := range values {
		result = append(result, handler(v))
	}

	return result
}
