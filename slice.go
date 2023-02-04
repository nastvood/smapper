package smapper

import "unsafe"

// SliceByStructs slice by slice of structs.
func SliceByStructs[T, E any](s []T, offset uintptr) []E {
	if len(s) == 0 {
		return nil
	}

	size := unsafe.Sizeof(s[0])
	if size < offset {
		panic("wrong offset")
	}

	res := make([]E, 0, len(s))

	var pI unsafe.Pointer
	for _, el := range s {
		if pI == nil {
			//nolint:gosec
			pI = unsafe.Add(unsafe.Pointer(&el), offset)
		}

		res = append(res, *(*E)(pI))
	}

	return res
}

// SliceByStructsP slice by slice of pointers of struct.
func SliceByStructsP[T, E any](s []*T, offset uintptr) []E {
	if len(s) == 0 {
		return nil
	}

	res := make([]E, 0, len(s))

	var pEl, pI unsafe.Pointer
	for _, el := range s {
		pEl = unsafe.Pointer(el)
		if pEl == nil {
			continue
		}

		pI = unsafe.Add(unsafe.Pointer((*T)(pEl)), offset)

		res = append(res, *(*E)(pI))
	}

	return res
}
