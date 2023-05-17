package binding

import (
	"errors"
	"net/http"
)

var (
	// Only accept json, form and multipart form
	ErrIgnoredBinding = errors.New("cannot bind request to struct (ignored binding method)")
)

type ignoredBinding struct{}

func (ignoredBinding) Name() string {
	return "ignored"
}

func (ignoredBinding) Bind(req *http.Request, obj interface{}) error {
	return ErrIgnoredBinding
}
