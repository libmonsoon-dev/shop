package sqlrepo

func NewFactory[R any](fn func(db DB) R) Factory[R] {
	return Factory[R]{fn: fn}
}

type Factory[R any] struct {
	fn func(db DB) R
}

func (f Factory[R]) Create(db DB) R {
	return f.fn(db)
}
