package sliceutil

func Map[T any, U any](arr []T, callback func(value T) U) []U {
	var res []U
	for _, item := range arr {
		res = append(res, callback(item))
	}
	return res
}

func FlatMap[T any, U any](arr []T, callback func(value T) []U) []U {
	var res []U
	for _, item := range arr {
		res = append(res, callback(item)...)
	}
	return res
}
