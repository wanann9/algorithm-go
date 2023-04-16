package base

import "golang.org/x/exp/constraints"

type Comparator[T any] func(T, T) int

func CmpOrdered[T constraints.Ordered](a, b T) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func CmpComparable[T comparable](a, b T) int {
	return Bool2Int(a != b)
}

func SliceCmp[T any](cmp Comparator[T]) Comparator[[]T] {
	return func(a, b []T) int {
		m, n := len(a), len(b)
		for i := 0; i < m && i < n; i++ {
			if v := cmp(a[i], b[i]); v != 0 {
				return v
			}
		}
		return m - n
	}
}

func ReverseCmp[T any](cmp Comparator[T]) Comparator[T] {
	return func(a, b T) int { return -cmp(a, b) }
}
