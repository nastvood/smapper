package smapper

import (
	"unsafe"
)

func MapBySlice[T any, I comparable](s []T, offset uintptr) map[I]T {
	if len(s) == 0 {
		return nil
	}

	size := unsafe.Sizeof(s[0])
	if size < offset {
		panic("wrong offset")
	}

	m := make(map[I]T)

	var pI unsafe.Pointer
	for _, el := range s {
		if pI == nil {
			//nolint:gosec
			pI = unsafe.Add(unsafe.Pointer(&el), offset)
		}

		m[*(*I)(pI)] = el
	}

	return m
}

func MapBySliceP[T any, I comparable](s []*T, offset uintptr) map[I]*T {
	if len(s) == 0 {
		return nil
	}

	m := make(map[I]*T)

	var pEl, pI unsafe.Pointer
	for _, el := range s {
		pEl = unsafe.Pointer(el)
		if pEl == nil {
			continue
		}

		pI = unsafe.Add(unsafe.Pointer((*T)(pEl)), offset)

		m[*(*I)(pI)] = el
	}

	return m
}

func SetBySlice[T any, I comparable](s []T, offset uintptr) map[I]struct{} {
	if len(s) == 0 {
		return nil
	}

	size := unsafe.Sizeof(s[0])
	if size < offset {
		panic("wrong offset")
	}

	m := make(map[I]struct{})

	var pI unsafe.Pointer
	for _, el := range s {
		if pI == nil {
			//nolint:gosec
			pI = unsafe.Add(unsafe.Pointer(&el), offset)
		}

		m[*(*I)(pI)] = struct{}{}
	}

	return m
}

func SetBySliceP[T any, I comparable](s []*T, offset uintptr) map[I]struct{} {
	if len(s) == 0 {
		return nil
	}

	m := make(map[I]struct{})

	var pEl, pI unsafe.Pointer
	for _, el := range s {
		pEl = unsafe.Pointer(el)
		if pEl == nil {
			continue
		}

		pI = unsafe.Add(unsafe.Pointer((*T)(pEl)), offset)

		k := *(*I)(pI)

		m[k] = struct{}{}
	}

	return m
}
