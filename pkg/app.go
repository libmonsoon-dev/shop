package shop

// TODO: move to golib
type Property[T any] struct {
	val T
}

func (p *Property[T]) Set(val T) {
	p.val = val
}

func (p *Property[T]) Get() T {
	return p.val
}

//var _ json.Marshaler = Property[any] //TODO
//var _ json.Unmarshaler = Property[any] //TODO
//var _ driver.Valuer = Property[any] //TODO
