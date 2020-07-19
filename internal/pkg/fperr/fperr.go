package fperr

import "errors"

const (
	NoContent = iota
)

var (
	fpErrNoContent = errors.New("no content")
)

var fpErrMap = map[int]error{
	NoContent: fpErrNoContent,
}

func New(code int) error {
	if v, ok := fpErrMap[code]; ok {
		return v
	}

	panic(errors.New("no such fpErr"))
}
