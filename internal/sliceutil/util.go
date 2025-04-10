package sliceutil

import "slices"

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

func Find[T any](arr []T, callback func(value T) bool) (value T, ok bool) {
	index := slices.IndexFunc(arr, callback)
	if index == -1 {
		return
	}
	return arr[index], true
}
