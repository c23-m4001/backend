// Code generated by "go run clinic-pilot/tool/stringer -linecomment -type=CategoryLogoType -output=category_enum_gen.go -swagoutput=../tool/swag/enum_gen/category_enum_gen.go -custom"; DO NOT EDIT.

package data_type

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CategoryLogoTypeFoodBeverage-1]
	_ = x[CategoryLogoTypeTransportation-2]
	_ = x[CategoryLogoTypeRental-3]
	_ = x[CategoryLogoTypeWaterBill-4]
	_ = x[CategoryLogoTypePhoneBill-5]
	_ = x[CategoryLogoTypeElectricityBill-6]
	_ = x[CategoryLogoTypeEducation-7]
	_ = x[CategoryLogoTypePets-8]
	_ = x[CategoryLogoTypeFitness-9]
	_ = x[CategoryLogoTypeGames-10]
	_ = x[CategoryLogoTypeOther-11]
}

const _CategoryLogoType_nameReadable = "FOOD_BEVERAGE, TRANSPORTATION, RENTAL, WATER_BILL, PHONE_BILL, ELECTRICITY_BILL, EDUCATION, PETS, FITNESS, GAMES, OTHER"

const _CategoryLogoType_name = "FOOD_BEVERAGETRANSPORTATIONRENTALWATER_BILLPHONE_BILLELECTRICITY_BILLEDUCATIONPETSFITNESSGAMESOTHER"

var _CategoryLogoType_index = [...]uint8{0, 13, 27, 33, 43, 53, 69, 78, 82, 89, 94, 99}

func (i *CategoryLogoType) determine(s string) {
	switch s {
	case "FOOD_BEVERAGE":
		*i = CategoryLogoTypeFoodBeverage
	case "TRANSPORTATION":
		*i = CategoryLogoTypeTransportation
	case "RENTAL":
		*i = CategoryLogoTypeRental
	case "WATER_BILL":
		*i = CategoryLogoTypeWaterBill
	case "PHONE_BILL":
		*i = CategoryLogoTypePhoneBill
	case "ELECTRICITY_BILL":
		*i = CategoryLogoTypeElectricityBill
	case "EDUCATION":
		*i = CategoryLogoTypeEducation
	case "PETS":
		*i = CategoryLogoTypePets
	case "FITNESS":
		*i = CategoryLogoTypeFitness
	case "GAMES":
		*i = CategoryLogoTypeGames
	case "OTHER":
		*i = CategoryLogoTypeOther
	default:
		*i = 0
	}
}

func (i CategoryLogoType) IsValid() bool {
	if i == 0 {
		return false
	}

	return true
}

func (i CategoryLogoType) GetValidValuesString() string {
	return _CategoryLogoType_nameReadable
}

func (i CategoryLogoType) String() string {
	i -= 1
	if i < 0 || i >= CategoryLogoType(len(_CategoryLogoType_index)-1) {
		return "CategoryLogoType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}

	return _CategoryLogoType_name[_CategoryLogoType_index[i]:_CategoryLogoType_index[i+1]]
}

func (i CategoryLogoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *CategoryLogoType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	i.determine(s)

	return nil
}

func (i *CategoryLogoType) UnmarshalText(b []byte) error {
	i.determine(string(b))

	return nil
}

func (i *CategoryLogoType) Scan(value interface{}) error {
	switch s := value.(type) {
	case string:
		i.determine(s)
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, i)
	}

	return nil
}

func (i CategoryLogoType) Value() (driver.Value, error) {
	return i.String(), nil
}

func CategoryLogoTypeP(v CategoryLogoType) *CategoryLogoType {
	return &v
}

func ListCategoryLogoType() []CategoryLogoType {
	return []CategoryLogoType{
		CategoryLogoTypeFoodBeverage,
		CategoryLogoTypeTransportation,
		CategoryLogoTypeRental,
		CategoryLogoTypeWaterBill,
		CategoryLogoTypePhoneBill,
		CategoryLogoTypeElectricityBill,
		CategoryLogoTypeEducation,
		CategoryLogoTypePets,
		CategoryLogoTypeFitness,
		CategoryLogoTypeGames,
		CategoryLogoTypeOther,
	}
}

func ListCategoryLogoTypeString() []string {
	return []string{
		CategoryLogoTypeFoodBeverage.String(),
		CategoryLogoTypeTransportation.String(),
		CategoryLogoTypeRental.String(),
		CategoryLogoTypeWaterBill.String(),
		CategoryLogoTypePhoneBill.String(),
		CategoryLogoTypeElectricityBill.String(),
		CategoryLogoTypeEducation.String(),
		CategoryLogoTypePets.String(),
		CategoryLogoTypeFitness.String(),
		CategoryLogoTypeGames.String(),
		CategoryLogoTypeOther.String(),
	}
}
