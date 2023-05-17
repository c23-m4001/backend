package validator

import "regexp"

var (
	alphanumSymbolRegex = regexp.MustCompile(`[^\s]+`)
	splitParamsRegex    = regexp.MustCompile(`'[^']*'|\S+`)
)
