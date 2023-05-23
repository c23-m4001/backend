package data_type

//go:generate go run capstone/tool/stringer -linecomment -type=CategoryLogoType -output=category_enum_gen.go -swagoutput=../tool/swag/enum_gen/category_enum_gen.go -custom
type CategoryLogoType int // @name CategoryLogoTypeEnum

const (
	CategoryLogoTypeFoodBeverage    CategoryLogoType = iota + 1 // FOOD_BEVERAGE
	CategoryLogoTypeTransportation                              // TRANSPORTATION
	CategoryLogoTypeRental                                      // RENTAL
	CategoryLogoTypeWaterBill                                   // WATER_BILL
	CategoryLogoTypePhoneBill                                   // PHONE_BILL
	CategoryLogoTypeElectricityBill                             // ELECTRICITY_BILL
	CategoryLogoTypeEducation                                   // EDUCATION
	CategoryLogoTypePets                                        // PETS
	CategoryLogoTypeFitness                                     // FITNESS
	CategoryLogoTypeGames                                       // GAMES
	CategoryLogoTypeOther                                       // OTHER
)
