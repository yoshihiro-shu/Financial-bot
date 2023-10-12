package etcd

import "errors"

var (
	ErrCtxCanceled         = errors.New("ctx is canceled by another routine")
	ErrCtxDeadlineExceeded = errors.New("ctx is attached with a deadline is exceeded")
	ErrEmptyKey            = errors.New("client-side error")
	ErrDefault             = errors.New("bad cluster endpoints, which are not etcd servers")
)
