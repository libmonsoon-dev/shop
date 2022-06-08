package unitofwork

import (
	"github.com/libmonsoon-dev/shop/pkg/v1/pointer"
	"github.com/libmonsoon-dev/shop/pkg/v1/slice"
)

func MergeErrors(errors ...*error) error {
	panic("impl me")
	// TODO: use outer liberr pkg function
	// return liberr.MergeErrors(ptrErrors2Errors(errors))
}

//TOOD:
//func ptrErrors2Errors(ptrs []*error) []error {
//	return slice.Map(ptrs, pointer.GetValue)
//}
func ptrErrors2Errors(ptrs []*error) []error {
	return slice.Map(ptrs, pointer.GetValue[error])
}
