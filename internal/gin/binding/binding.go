package binding

import (
	"github.com/gin-gonic/gin/binding"
)

// These implement the Binding interface and can be used to bind the data
// present in the request to struct instances.
var (
	Form          = formBinding{}
	FormMultipart = formMultipartBinding{}
	Ignored       = ignoredBinding{}
)

func validate(obj interface{}) error {
	if binding.Validator == nil {
		return nil
	}
	return binding.Validator.ValidateStruct(obj)
}

// Default returns the appropriate Binding instance based on the HTTP method
// and the content type.
func Default(method, contentType string) binding.Binding {
	b := binding.Default(method, contentType)

	switch b {
	case binding.JSON:
		return b
	case binding.Form:
		return Form
	case binding.FormMultipart:
		return FormMultipart
	default:
		return Ignored
	}
}
