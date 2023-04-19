package common

import (
	"reflect"
	"sort"
)

// SortMapByKey SortMapByKey[K, V]
//
//	@Description: sorts the map by key.
//	@param m map[K]V
//	@return []K
func SortMapByKey[K ~int, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

// Max Max[Comparable[T]]
//
//	@Description: returns the max value in the slice.
//	@param o []T
//	@return T
func Max[T Comparable[T]](o ...T) T {
	if len(o) == 0 {
		var zero T
		return zero
	}
	maxIndex := 0
	for i := len(o) - 1; i > 0; i-- {
		if o[i].CompareTo(o[maxIndex]) > 0 {
			maxIndex = i
		}
	}
	return o[maxIndex]
}

// SliceMax SliceMax[Comparable[T]]
//
//	@Description: returns the max value in the slice.
//	@param o []T
//	@return T
func SliceMax[T Comparable[T]](o []T) T {
	return Max(o...)
}

// Min Min[Comparable[T]]
//
//	@Description:
//	@param o
//	@return T
func Min[T Comparable[T]](o []T) T {
	if len(o) == 0 {
		var zero T
		return zero
	}
	maxIndex := 0
	for i := len(o) - 1; i > 0; i-- {
		if o[i].CompareTo(o[maxIndex]) < 0 {
			maxIndex = i
		}
	}
	return o[maxIndex]
}

// SliceMin SliceMin[Comparable[T]]
//
//	@Description: returns the min value in the slice.
//	@param o []T
//	@return T
func SliceMin[T Comparable[T]](o []T) T {
	return Min(o)
}

// IndexOf IndexOf[Comparable[T]]
//
//	@Description: returns the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element.
//	@param o T
//	@param t []T
//	@return int
func IndexOf[T Comparable[T]](o T, t []T) int {
	if len(t) == 0 {
		return -1
	}
	for i := len(t) - 1; i >= 0; i-- {
		if t[i].CompareTo(o) == 0 {
			return i
		}
	}
	return -1
}

// MaxNum MaxNum[T Num]
//
//	@Description: returns the max value in the slice.
//	@param ns []T
//	@return T
func MaxNum[T Num](ns []T) T {
	if len(ns) == 0 {
		var zero T
		return zero
	}
	maxIndex := 0
	for i := len(ns) - 1; i > 0; i-- {
		if ns[i]-ns[maxIndex] > 0 {
			maxIndex = i
		}
	}
	return ns[maxIndex]
}

// MinNum MinNum[T Num]
//
//	@Description: returns the min value in the slice.
//	@param ns []T
//	@return T
func MinNum[T Num](ns []T) T {
	if len(ns) == 0 {
		var zero T
		return zero
	}
	minIndex := 0
	for i := len(ns) - 1; i > 0; i-- {
		if ns[i]-ns[minIndex] < 0 {
			minIndex = i
		}
	}
	return ns[minIndex]
}

// Contain
//
//	@Description: checks if the target contains the obj.
//	@param target slice array map
//	@param obj interface{}
//	@return bool
func Contain(target interface{}, obj interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}
	return false
}

// SliceContain SliceContain[T comparable]
//
//	@Description: checks if the target slice contains the obj.
//	@param target []T
//	@param obj T
//	@return bool
func SliceContain[T comparable](target []T, obj T) bool {
	for i := 0; i < len(target); i++ {
		if target[i] == obj {
			return true
		}
	}
	return false
}

// MapContainKey MapContainKey[T1 comparable, T2 any]
//
//	@Description: checks if the target map keys contains the obj.
//	@param target map[T1]T2
//	@param obj T1
//	@return bool
func MapContainKey[T1 comparable, T2 any](target map[T1]T2, obj T1) bool {
	_, ok := target[obj]
	return ok
}

// MapContainValue MapContainValue[T1, T2 comparable]
//
//	@Description: checks if the target map values contains the obj.
//	@param target map[T1]T2
//	@param obj T2
//	@return bool
func MapContainValue[T1, T2 comparable](target map[T1]T2, obj T2) bool {
	for _, v := range target {
		if v == obj {
			return true
		}
	}
	return false
}

// RemoveIndex RemoveIndex[T comparable]
//
//	@Description: removes the element at the specified position in this list.
//	@param obj []T
//	@param indices ...int
//	@return []T the new slice
func RemoveIndex[T comparable](obj []T, indices ...int) []T {
	objT := make([]T, 0, len(obj)-len(indices))
	mapI := map[int]struct{}{}
	for _, i := range indices {
		mapI[i] = struct{}{}
	}
	for i := 0; i < len(obj); i++ {
		if _, ok := mapI[i]; !ok {
			objT = append(objT, obj[i])
		}
	}
	return objT
}

// RemoveSafe RemoveSafe[T comparable]
//
//	@Description: removes the first occurrence of the specified element from this list, if it is present.
//	@param target []T
//	@param obj T
//	@return []T the new slice
//	@return bool true if the element was removed
func RemoveSafe[T comparable](target []T, obj T) ([]T, bool) {
	targetT := make([]T, len(target))
	i := 0
	for ; i < len(target); i++ {
		if target[i] == obj {
			break
		}
	}
	if i == len(target) {
		return target, false
	}
	copy(targetT, target[:i])
	copy(targetT[i:], target[i+1:])
	return targetT, true
}

// Remove Remove[T comparable]
//
//	@Description: removes the first occurrence of the specified element from this list, if it is present.
//	@param target []T
//	@param obj T
//	@return []T
//	@return error ErrNotFound
func Remove[T comparable](target []T, obj T) ([]T, error) {
	for i := 0; i < len(target); i++ {
		if target[i] == obj {
			target = append(target[:i], target[i+1:]...)
			return target, nil
		}
	}
	return target, ErrNotFound
}

// SliceEqual SliceEqual[T comparable]
//
//	@Description: checks if the two slices are equal.
//	@param obj1 []T
//	@param obj2 []T
//	@return bool true if the two slices are equal
func SliceEqual[T comparable](obj1 []T, obj2 []T) bool {
	if len(obj1) != len(obj2) {
		return false
	}
	for i := 0; i < len(obj1); i++ {
		if obj1[i] != obj2[i] {
			return false
		}
	}
	return true
}

// MapEqual MapEqual[T1, T2 comparable]
//
//	@Description: checks if the two maps are equal.
//	@param obj1 map[T1]T2
//	@param obj2 map[T1]T2
//	@return bool true if the two maps are equal
func MapEqual[T1, T2 comparable](obj1 map[T1]T2, obj2 map[T1]T2) bool {
	if len(obj1) != len(obj2) {
		return false
	}
	for k, v := range obj1 {
		val, ok := obj2[k]
		if !ok || val != v {
			return false
		}
	}
	return true
}

// MapSlice MapSlice[T1, T2 any]
//
//	@Description: returns a slice consisting of the results of applying the given function to the elements of this slice.
//	@param obj []T1
//	@param fn  func(T1) T2
//	@param T2] []T2
//	@return []T2 the new slice
func MapSlice[T1, T2 any](obj []T1, fn MapFunc[T1, T2]) []T2 {
	objT := make([]T2, len(obj))
	for i := 0; i < len(obj); i++ {
		objT[i] = fn(obj[i])
	}
	return objT
}
