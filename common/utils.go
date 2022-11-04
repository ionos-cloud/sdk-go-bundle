package common

import (
	"unicode/utf8"
)

// ToValue - returns the value of the bool pointer passed in
func ToValue[T any](ptr *T) T {
	return *ptr
}

func ToValueDefault[T any](ptr *T) T {
	var defaultVal T
	if ptr == nil {
		return defaultVal
	}
	return *ptr
}

func SliceToValueDefault[T any](ptrSlice *[]T) []T {
	return append([]T{}, *ptrSlice...)
}

type Nullable[T any] struct {
	value *T
	isSet bool
}

func (v Nullable[T]) Get() *T {
	return v.value
}

func (v *Nullable[T]) Set(val *T) {
	v.value = val
	v.isSet = true
}

func (v Nullable[T]) IsSet() bool {
	return v.isSet
}

func (v *Nullable[T]) Unset() {
	v.value = nil
	v.isSet = false
}

func Strlen(s string) int {
	return utf8.RuneCountInString(s)
}
