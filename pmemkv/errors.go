package pmemkv

import "errors"

var (
	ErrNotFound = errors.New("Not Found")
	ErrFail     = errors.New("Fail")
	ErrIllegal  = errors.New("Illegal error number")
)

func pmemkvError(err int8) error {
	switch err {
	case -1:
		return ErrFail
	case 0:
		return ErrNotFound
	case 1:
		return nil
	default:
		return ErrIllegal
	}
}
