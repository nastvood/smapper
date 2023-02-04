package smapper

func ref[T any](t T) *T {
	return &t
}
