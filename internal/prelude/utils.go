package prelude

func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

func Ptr[T any](v T) *T {
	return &v
}
