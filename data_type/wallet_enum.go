package data_type

//go:generate go run capstone/tool/stringer -linecomment -type=WalletLogoType -output=wallet_enum_gen.go -swagoutput=../tool/swag/enum_gen/wallet_enum_gen.go -custom
type WalletLogoType int // @name WalletLogoTypeEnum

const (
	WalletLogoTypeCash       WalletLogoType = iota + 1 // CASH
	WalletLogoTypeBank                                 // BANK
	WalletLogoTypeCreditCard                           // CREDIT_CARD
	WalletLogoTypeLoan                                 // LOAN
	WalletLogoTypeInsurance                            // INSURANCE
	WalletLogoTypeInvestment                           // INVESTMENT
	WalletLogoTypeMortgage                             // MORTGAGE
	WalletLogoTypeBonus                                // BONUS
	WalletLogoTypeOther                                // OTHER
)
