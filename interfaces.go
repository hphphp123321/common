package common

type MapFunc[T1, T2 any] func(T1) T2

type Comparable[T any] interface {
	CompareTo(o T) int
}
