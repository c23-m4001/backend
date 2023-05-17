package translation

import (
	"fmt"
	"regexp"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var (
	bracketPlaceholderRe = regexp.MustCompile(`\[.*\]`)
)

type translation struct {
	tag             string
	translation     string
	override        bool
	customRegisFunc validator.RegisterTranslationsFunc
	customTransFunc validator.TranslationFunc
}

func registerTranslation(v *validator.Validate, trans ut.Translator, t translation) (err error) {
	if t.customTransFunc != nil && t.customRegisFunc != nil {
		err = v.RegisterTranslation(t.tag, trans, t.customRegisFunc, t.customTransFunc)
	} else if t.customTransFunc != nil && t.customRegisFunc == nil {
		err = v.RegisterTranslation(t.tag, trans, registrationFunc(t.tag, t.translation, t.override), t.customTransFunc)
	} else if t.customTransFunc == nil && t.customRegisFunc != nil {
		err = v.RegisterTranslation(t.tag, trans, t.customRegisFunc, translateFunc)
	} else {
		err = v.RegisterTranslation(t.tag, trans, registrationFunc(t.tag, t.translation, t.override), translateFunc)
	}

	return err
}

func registrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) (err error) {
		if err = ut.Add(tag, translation, override); err != nil {
			return
		}

		return
	}
}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), extractField(fe))
	if err != nil {
		fmt.Printf("warning: error translating FieldError: %#v", fe)
		return fe.(error).Error()
	}

	return t
}

func translateFuncFieldComparison(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), extractField(fe), extractComparedField(fe))
	if err != nil {
		fmt.Printf("warning: error translating FieldError: %#v", fe)
		return fe.(error).Error()
	}

	return t
}

func translateFuncValueComparison(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), extractField(fe), fe.Param())
	if err != nil {
		fmt.Printf("warning: error translating FieldError: %#v", fe)
		return fe.(error).Error()
	}

	return t
}

func extractField(fe validator.FieldError) string {
	return bracketPlaceholderRe.ReplaceAllString(fe.Field(), "")
}

func extractDatetimeParam(fe validator.FieldError) string {
	param := fe.Param()
	switch param {
	case "2006-01-02":
		param = "YYYY-MM-DD"
	}

	return param
}

func extractStructLevelValidationField(fe validator.FieldError) string {
	chunkStructNs := strings.Split(fe.StructNamespace(), ".")
	// StructNamespace will return "very.long.struct.namespace.fieldname."
	fieldName := chunkStructNs[len(chunkStructNs)-2]
	return bracketPlaceholderRe.ReplaceAllString(fieldName, "")
}

func extractComparedField(fe validator.FieldError) string {
	namespace := fe.Param()
	namespaceChunks := strings.Split(namespace, ".")
	comparedField := namespaceChunks[len(namespaceChunks)-1]

	return bracketPlaceholderRe.ReplaceAllString(comparedField, "")
}
