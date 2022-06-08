package pointer

type Pointer[T any] interface {
	*T
}

// GetValue returns value by ptr if ptr is not nil, zero value
// GetValue returns value by ptr if Pointer[T] is non-nil, otherwise zero value of T
// TODO:
func GetValue[T any, Ptr Pointer[T]](ptr Ptr) (val T) {
	// func GetValue[T any](ptr *T) (val T) {
	if ptr != nil {
		val = *ptr
	}

	return
}
