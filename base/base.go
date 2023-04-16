package base

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](a, b T) T {
	return MinWithCmp(a, b, CmpOrdered[T])
}

func Max[T constraints.Ordered](a, b T) T {
	return MaxWithCmp(a, b, CmpOrdered[T])
}

func MinWithCmp[T any](a, b T, cmp Comparator[T]) T {
	if cmp(a, b) < 0 {
		return a
	}
	return b
}

func MaxWithCmp[T any](a, b T, cmp Comparator[T]) T {
	if cmp(a, b) > 0 {
		return a
	}
	return b
}
