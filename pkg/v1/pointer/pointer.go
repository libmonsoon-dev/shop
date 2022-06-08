package pointer

// GetValue returns value by ptr if ptr is not nil, zero value
func GetValue[T any](ptr *T) (val T) {
	if ptr != nil {
		val = *ptr
	}

	return
}
