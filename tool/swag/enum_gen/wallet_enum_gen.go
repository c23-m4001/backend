// Code generated by "go run clinic-pilot/tool/stringer -linecomment -type=WalletLogoType -output=wallet_enum_gen.go -swagoutput=../tool/swag/enum_gen/wallet_enum_gen.go -custom"; DO NOT EDIT.

package enum_gen

import (
	"github.com/go-openapi/spec"
)

func init() {
	PostSwaggerDefinitions["WalletLogoTypeEnum"] = spec.Schema{
		SchemaProps: spec.SchemaProps{
			Type: []string{"string"},
			Enum: []interface{}{
				"CASH",
				"BANK",
				"CREDIT_CARD",
				"LOAN",
				"INSURANCE",
				"INVESTMENT",
				"MORTGAGE",
				"BONUS",
				"OTHER",
			},
		},
	}
}
