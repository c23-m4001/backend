package translation

import (
	"capstone/data_type"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/locales"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// RegisterEnTranslations registers a set of english translations
// for all built in tag's/custom tag's in validator; you may add your own as desired.
func RegisterEnTranslations(v *validator.Validate, trans ut.Translator) (err error) {
	translations := []translation{
		{
			tag:         "required",
			translation: "{0} is a required field",
			override:    false,
		},
		{
			tag:         "required_if",
			translation: "{0} is a required field",
			override:    false,
		},
		{
			tag:         "required_with",
			translation: "{0} is a required field",
			override:    false,
		},
		{
			tag:         "required_without",
			translation: "{0} is a required field",
			override:    false,
		},
		{
			tag:         "hostname_port",
			translation: "{0} format must be hostname:port",
			override:    false,
		},
		{
			tag: "len",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("len-string", "{0} must be {1} in length", false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-string-character", "{0} character", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-string-character", "{0} characters", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("len-number", "{0} must be equal to {1}", false); err != nil {
					return
				}

				if err = ut.Add("len-items", "{0} must contain {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-items-item", "{0} items", locales.PluralRuleOther, false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string

				var digits uint64
				var kind reflect.Kind

				if idx := strings.Index(fe.Param(), "."); idx != -1 {
					digits = uint64(len(fe.Param()[idx+1:]))
				}

				f64, err := strconv.ParseFloat(fe.Param(), 64)
				if err != nil {
					goto END
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:
					var c string

					c, err = ut.C("len-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("len-string", extractField(fe), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					c, err = ut.C("len-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("len-items", extractField(fe), c)

				default:
					t, err = ut.T("len-number", extractField(fe), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "min",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("min-string", "{0} must be at least {1} in length", false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-string-character", "{0} character", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-string-character", "{0} characters", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("min-number", "{0} must be {1} or greater", false); err != nil {
					return
				}

				if err = ut.Add("min-items", "{0} must contain at least {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-items-item", "{0} items", locales.PluralRuleOther, false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string

				var digits uint64
				var kind reflect.Kind

				if idx := strings.Index(fe.Param(), "."); idx != -1 {
					digits = uint64(len(fe.Param()[idx+1:]))
				}

				f64, err := strconv.ParseFloat(fe.Param(), 64)
				if err != nil {
					goto END
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:
					var c string

					c, err = ut.C("min-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("min-string", extractField(fe), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					c, err = ut.C("min-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("min-items", extractField(fe), c)

				default:
					t, err = ut.T("min-number", extractField(fe), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "max",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("max-string", "{0} must be a maximum of {1} in length", false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-string-character", "{0} character", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-string-character", "{0} characters", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("max-number", "{0} must be {1} or less", false); err != nil {
					return
				}

				if err = ut.Add("max-items", "{0} must contain at maximum {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-items-item", "{0} items", locales.PluralRuleOther, false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string

				var digits uint64
				var kind reflect.Kind

				if idx := strings.Index(fe.Param(), "."); idx != -1 {
					digits = uint64(len(fe.Param()[idx+1:]))
				}

				f64, err := strconv.ParseFloat(fe.Param(), 64)
				if err != nil {
					goto END
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:
					var c string

					c, err = ut.C("max-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("max-string", extractField(fe), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					c, err = ut.C("max-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("max-items", extractField(fe), c)

				default:
					t, err = ut.T("max-number", extractField(fe), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:             "eq",
			translation:     "{0} is not equal to {1}",
			override:        false,
			customTransFunc: translateFuncValueComparison,
		},
		{
			tag:             "ne",
			translation:     "{0} should not be equal to {1}",
			override:        false,
			customTransFunc: translateFuncValueComparison,
		},
		{
			tag: "lt",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("lt-string", "{0} must be less than {1} in length", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-string-character", "{0} character", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-string-character", "{0} characters", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-number", "{0} must be less than {1}", false); err != nil {
					return
				}

				if err = ut.Add("lt-items", "{0} must contain less than {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-items-item", "{0} items", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-datetime", "{0} must be less than the current Date & Time", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				fn := func() (err error) {
					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lt-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lt-string", extractField(fe), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lt-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lt-items", extractField(fe), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
						goto END
					}

					t, err = ut.T("lt-datetime", extractField(fe))

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("lt-number", extractField(fe), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "lte",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("lte-string", "{0} must be at maximum {1} in length", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-string-character", "{0} character", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-string-character", "{0} characters", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-number", "{0} must be {1} or less", false); err != nil {
					return
				}

				if err = ut.Add("lte-items", "{0} must contain at maximum {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-items-item", "{0} items", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-datetime", "{0} must be less than or equal to the current Date & Time", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				fn := func() (err error) {
					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lte-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lte-string", extractField(fe), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lte-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lte-items", extractField(fe), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
						goto END
					}

					t, err = ut.T("lte-datetime", extractField(fe))

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("lte-number", extractField(fe), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "gt",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("gt-string", "{0} must be greater than {1} in length", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-string-character", "{0} character", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-string-character", "{0} characters", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-number", "{0} must be greater than {1}", false); err != nil {
					return
				}

				if err = ut.Add("gt-items", "{0} must contain more than {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-items-item", "{0} items", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-datetime", "{0} must be greater than the current Date & Time", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				fn := func() (err error) {
					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gt-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gt-string", extractField(fe), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gt-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gt-items", extractField(fe), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
						goto END
					}

					t, err = ut.T("gt-datetime", extractField(fe))

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("gt-number", extractField(fe), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "gte",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("gte-string", "{0} must be at least {1} in length", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-string-character", "{0} character", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-string-character", "{0} characters", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-number", "{0} must be {1} or greater", false); err != nil {
					return
				}

				if err = ut.Add("gte-items", "{0} must contain at least {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-items-item", "{0} items", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-datetime", "{0} must be greater than or equal to the current Date & Time", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				fn := func() (err error) {
					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gte-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gte-string", extractField(fe), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gte-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gte-items", extractField(fe), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
						goto END
					}

					t, err = ut.T("gte-datetime", extractField(fe))

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("gte-number", extractField(fe), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:             "eqfield",
			translation:     "{0} must be equal to {1}",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:             "eqcsfield",
			translation:     "{0} must be equal to {1}",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:             "necsfield",
			translation:     "{0} cannot be equal to {1}",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:             "gtcsfield",
			translation:     "{0} must be greater than {1}",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:             "gtecsfield",
			translation:     "{0} must be greater than or equal to {1}",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:             "ltcsfield",
			translation:     "{0} must be less than {1}",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:             "ltecsfield",
			translation:     "{0} must be less than or equal to {1}",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:             "nefield",
			translation:     "{0} cannot be equal to {1}",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:             "gtfield",
			translation:     "{0} must be greater than {1}",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:             "gtefield",
			translation:     "{0} must be greater than or equal to {1}",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:             "ltfield",
			translation:     "{0} must be less than {1}",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:             "ltefield",
			translation:     "{0} must be less than or equal to {1}",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:         "alpha",
			translation: "{0} can only contain alphabetic characters",
			override:    false,
		},
		{
			tag:         "alphanum",
			translation: "{0} can only contain alphanumeric characters",
			override:    false,
		},
		{
			tag:         "numeric",
			translation: "{0} must be a valid numeric value",
			override:    false,
		},
		{
			tag:         "number",
			translation: "{0} must be a valid number",
			override:    false,
		},
		{
			tag:         "hexadecimal",
			translation: "{0} must be a valid hexadecimal",
			override:    false,
		},
		{
			tag:         "hexcolor",
			translation: "{0} must be a valid HEX color",
			override:    false,
		},
		{
			tag:         "rgb",
			translation: "{0} must be a valid RGB color",
			override:    false,
		},
		{
			tag:         "rgba",
			translation: "{0} must be a valid RGBA color",
			override:    false,
		},
		{
			tag:         "hsl",
			translation: "{0} must be a valid HSL color",
			override:    false,
		},
		{
			tag:         "hsla",
			translation: "{0} must be a valid HSLA color",
			override:    false,
		},
		{
			tag:         "e164",
			translation: "{0} must be a valid format phone number +62 (XXX) XXXX-XXXX",
			override:    false,
		},
		{
			tag:         "email",
			translation: "{0} must be a valid email address",
			override:    false,
		},
		{
			tag:         "url",
			translation: "{0} must be a valid URL",
			override:    false,
		},
		{
			tag:         "uri",
			translation: "{0} must be a valid URI",
			override:    false,
		},
		{
			tag:         "base64",
			translation: "{0} must be a valid Base64 string",
			override:    false,
		},
		{
			tag:             "contains",
			translation:     "{0} must contain the text '{1}'",
			override:        false,
			customTransFunc: translateFuncValueComparison,
		},
		{
			tag:             "containsany",
			translation:     "{0} must contain at least one of the following characters '{1}'",
			override:        false,
			customTransFunc: translateFuncValueComparison,
		},
		{
			tag:             "excludes",
			translation:     "{0} cannot contain the text '{1}'",
			override:        false,
			customTransFunc: translateFuncValueComparison,
		},
		{
			tag:             "excludesall",
			translation:     "{0} cannot contain any of the following characters '{1}'",
			override:        false,
			customTransFunc: translateFuncValueComparison,
		},
		{
			tag:             "excludesrune",
			translation:     "{0} cannot contain the following '{1}'",
			override:        false,
			customTransFunc: translateFuncValueComparison,
		},
		{
			tag:         "isbn",
			translation: "{0} must be a valid ISBN number",
			override:    false,
		},
		{
			tag:         "isbn10",
			translation: "{0} must be a valid ISBN-10 number",
			override:    false,
		},
		{
			tag:         "isbn13",
			translation: "{0} must be a valid ISBN-13 number",
			override:    false,
		},
		{
			tag:         "uuid",
			translation: "{0} must be a valid UUID",
			override:    false,
		},
		{
			tag:         "uuid3",
			translation: "{0} must be a valid version 3 UUID",
			override:    false,
		},
		{
			tag:         "uuid4",
			translation: "{0} must be a valid version 4 UUID",
			override:    false,
		},
		{
			tag:         "uuid5",
			translation: "{0} must be a valid version 5 UUID",
			override:    false,
		},
		{
			tag:         "ulid",
			translation: "{0} must be a valid ULID",
			override:    false,
		},
		{
			tag:         "ascii",
			translation: "{0} must contain only ascii characters",
			override:    false,
		},
		{
			tag:         "printascii",
			translation: "{0} must contain only printable ascii characters",
			override:    false,
		},
		{
			tag:         "multibyte",
			translation: "{0} must contain multibyte characters",
			override:    false,
		},
		{
			tag:         "datauri",
			translation: "{0} must contain a valid Data URI",
			override:    false,
		},
		{
			tag:         "latitude",
			translation: "{0} must contain valid latitude coordinates",
			override:    false,
		},
		{
			tag:         "longitude",
			translation: "{0} must contain a valid longitude coordinates",
			override:    false,
		},
		{
			tag:         "ssn",
			translation: "{0} must be a valid SSN number",
			override:    false,
		},
		{
			tag:         "ipv4",
			translation: "{0} must be a valid IPv4 address",
			override:    false,
		},
		{
			tag:         "ipv6",
			translation: "{0} must be a valid IPv6 address",
			override:    false,
		},
		{
			tag:         "ip",
			translation: "{0} must be a valid IP address",
			override:    false,
		},
		{
			tag:         "cidr",
			translation: "{0} must contain a valid CIDR notation",
			override:    false,
		},
		{
			tag:         "cidrv4",
			translation: "{0} must contain a valid CIDR notation for an IPv4 address",
			override:    false,
		},
		{
			tag:         "cidrv6",
			translation: "{0} must contain a valid CIDR notation for an IPv6 address",
			override:    false,
		},
		{
			tag:         "tcp_addr",
			translation: "{0} must be a valid TCP address",
			override:    false,
		},
		{
			tag:         "tcp4_addr",
			translation: "{0} must be a valid IPv4 TCP address",
			override:    false,
		},
		{
			tag:         "tcp6_addr",
			translation: "{0} must be a valid IPv6 TCP address",
			override:    false,
		},
		{
			tag:         "udp_addr",
			translation: "{0} must be a valid UDP address",
			override:    false,
		},
		{
			tag:         "udp4_addr",
			translation: "{0} must be a valid IPv4 UDP address",
			override:    false,
		},
		{
			tag:         "udp6_addr",
			translation: "{0} must be a valid IPv6 UDP address",
			override:    false,
		},
		{
			tag:         "ip_addr",
			translation: "{0} must be a resolvable IP address",
			override:    false,
		},
		{
			tag:         "ip4_addr",
			translation: "{0} must be a resolvable IPv4 address",
			override:    false,
		},
		{
			tag:         "ip6_addr",
			translation: "{0} must be a resolvable IPv6 address",
			override:    false,
		},
		{
			tag:         "unix_addr",
			translation: "{0} must be a resolvable UNIX address",
			override:    false,
		},
		{
			tag:         "mac",
			translation: "{0} must contain a valid MAC address",
			override:    false,
		},
		{
			tag:         "unique",
			translation: "{0} must contain unique values",
			override:    false,
		},
		{
			tag:         "iscolor",
			translation: "{0} must be a valid color",
			override:    false,
		},
		{
			tag:             "oneof",
			translation:     "{0} must be one of [{1}]",
			override:        false,
			customTransFunc: translateFuncValueComparison,
		},
		{
			tag:         "json",
			translation: "{0} must be a valid json string",
			override:    false,
		},
		{
			tag:         "jwt",
			translation: "{0} must be a valid jwt string",
			override:    false,
		},
		{
			tag:         "lowercase",
			translation: "{0} must be a lowercase string",
			override:    false,
		},
		{
			tag:         "uppercase",
			translation: "{0} must be an uppercase string",
			override:    false,
		},
		{
			tag:         "datetime",
			translation: "{0} does not match the {1} format",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), extractField(fe), extractDatetimeParam(fe))
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "boolean",
			translation: "{0} must be a valid boolean value",
			override:    false,
		},
		{
			tag:         "required_if_stringer",
			translation: "{0} is a required field",
			override:    false,
		},
		{
			tag:         "alphanum_symbol",
			translation: "{0} can only contain alphanumeric and symbol characters",
			override:    false,
		},
		{
			tag:         "not_empty",
			translation: "{0} must not empty",
			override:    false,
		},
		{
			tag:         "data_type_enum",
			translation: "{0} must be one of {1}",
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var t string
				var err error

				if v, ok := fe.Value().(data_type.EnumValidator); ok {
					t, err = ut.T(fe.Tag(), extractField(fe), fmt.Sprintf("[%s]", v.GetValidValuesString()))
				} else {
					err = fmt.Errorf("tag '%s' cannot be used on non-data_type.EnumValidator interface", fe.Tag())
				}

				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "data_type_date",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("data-type-date-required", "{0} is a required field", false); err != nil {
					return
				}

				if err = ut.Add("data-type-date-format-invalid", "{0} does not match the {1} format", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var t string
				var err error

				if v, ok := fe.Value().(data_type.Date); ok {
					if v.HasParseErr() {
						t, err = ut.T("data-type-date-format-invalid", extractStructLevelValidationField(fe), v.IsoLayout())
					} else {
						t, err = ut.T("data-type-date-required", extractStructLevelValidationField(fe))
					}
				} else {
					err = fmt.Errorf("tag '%s' cannot be used on non-data_type.Date struct", fe.Tag())
				}

				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "data_type_null_date",
			translation: "{0} does not match the {1} format",
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var t string
				var err error

				if v, ok := fe.Value().(data_type.NullDate); ok {
					t, err = ut.T(fe.Tag(), extractStructLevelValidationField(fe), v.IsoLayout())
				} else {
					err = fmt.Errorf("tag '%s' cannot be used on non-data_type.NullDate struct", fe.Tag())
				}

				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "data_type_date_time",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("data-type-date-time-required", "{0} is a required field", false); err != nil {
					return
				}

				if err = ut.Add("data-type-date-time-format-invalid", "{0} does not match the {1} format", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var t string
				var err error

				if v, ok := fe.Value().(data_type.DateTime); ok {
					if v.HasParseErr() {
						t, err = ut.T("data-type-date-time-format-invalid", extractStructLevelValidationField(fe), v.IsoLayout())
					} else {
						t, err = ut.T("data-type-date-time-required", extractStructLevelValidationField(fe))
					}
				} else {
					err = fmt.Errorf("tag '%s' cannot be used on non-data_type.DateTime struct", fe.Tag())
				}

				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "data_type_null_date_time",
			translation: "{0} does not match the {1} format",
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var t string
				var err error

				if v, ok := fe.Value().(data_type.NullDateTime); ok {
					t, err = ut.T(fe.Tag(), extractStructLevelValidationField(fe), v.IsoLayout())
				} else {
					err = fmt.Errorf("tag '%s' cannot be used on non-data_type.NullDateTime struct", fe.Tag())
				}

				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
	}

	for _, t := range translations {
		err = registerTranslation(v, trans, t)

		if err != nil {
			return
		}
	}

	return
}
