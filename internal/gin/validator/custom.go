package validator

import (
	"capstone/data_type"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	// customValidators is the custom map of ValidationFunc
	// you can add, remove or even replace items to suite your needs,
	// or even disregard and use your own map if so desired.
	customValidators = map[string]validator.Func{
		requiredIfStringerTag: requiredIfStringer,
		"alphanum_symbol":     isAlphanumSymbol,
		"not_empty":           isNotEmpty,
		"data_type_enum":      isValidDataTypeEnum,
	}
)

var (
	oneofValsCache       = map[string][]string{}
	oneofValsCacheRWLock = sync.RWMutex{}
)

func parseOneOfParam2(s string) []string {
	oneofValsCacheRWLock.RLock()
	vals, ok := oneofValsCache[s]
	oneofValsCacheRWLock.RUnlock()
	if !ok {
		oneofValsCacheRWLock.Lock()
		vals = splitParamsRegex.FindAllString(s, -1)
		for i := 0; i < len(vals); i++ {
			vals[i] = strings.Replace(vals[i], "'", "", -1)
		}
		oneofValsCache[s] = vals
		oneofValsCacheRWLock.Unlock()
	}
	return vals
}

// requireCheckStringerFieldValue is a func for check field value that implement stringer
func requireCheckStringerFieldValue(fl validator.FieldLevel, param string, value string, defaultNotFoundValue bool) bool {
	field, _, _, found := fl.GetStructFieldOKAdvanced2(fl.Parent(), param)
	if !found {
		return defaultNotFoundValue
	}

	if isTypeFmtStringer(field) {
		v := field.Interface().(fmt.Stringer)
		return v.String() == value
	}

	panic(fmt.Sprintf("field %s doesn't implement fmt.Stringer interface", param))
}

// requiredIfStringer is the validation function
// The field under validation must be present and not empty only if all the other specified fields are equal to the value following with the specified field.
func requiredIfStringer(fl validator.FieldLevel) bool {
	params := parseOneOfParam2(fl.Param())
	if len(params)%2 != 0 {
		panic(fmt.Sprintf("Bad param number for required_if_stringer %s", fl.FieldName()))
	}
	for i := 0; i < len(params); i += 2 {
		if !requireCheckStringerFieldValue(fl, params[i], params[i+1], false) {
			return true
		}
	}

	field := fl.Field()
	switch field.Kind() {
	case reflect.Slice, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func:
		return !field.IsNil()
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

func isAlphanumSymbol(fl validator.FieldLevel) bool {
	return alphanumSymbolRegex.MatchString(fl.Field().String())
}

func isNotEmpty(fl validator.FieldLevel) bool {
	field := fl.Field()
	switch field.Kind() {
	case reflect.Slice:
		isEmpty := field.Len() == 0

		return !isEmpty

	case reflect.String:
		val := strings.TrimSpace(field.String())
		isEmpty := len(val) == 0

		return !isEmpty

	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

func isValidDataTypeEnum(fl validator.FieldLevel) bool {
	field := fl.Field()
	if v, ok := field.Interface().(data_type.EnumValidator); ok {
		return v.IsValid()
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}
