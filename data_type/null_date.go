package data_type

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type NullDate struct {
	date    Date
	isValid bool
}

func (dt NullDate) IsoLayout() string {
	return dt.date.IsoLayout()
}

func (dt NullDate) get() *Date {
	var date *Date

	if dt.isValid {
		date = new(Date)
		*date = dt.date
	}

	return date
}

func (dt *NullDate) set(v *Date) {
	date, isValid := Date{}, false
	if v != nil {
		date, isValid = *v, true
	}

	dt.date, dt.isValid = date, isValid
}

func (dt *NullDate) parse(s *string) {
	if s == nil {
		dt.set(nil)
		return
	}

	date := Date{}
	date.parse(*s)
	dt.set(&date)
}

func (dt NullDate) Date() Date {
	date := dt.get()
	if date == nil {
		return Date{}
	}

	return *date
}

func (dt NullDate) DateP() *Date {
	return dt.get()
}

func (dt NullDate) DateTimeEndOfDay() DateTime {
	date := dt.get()
	if date == nil {
		return DateTime{}
	}

	return date.DateTimeEndOfDay()
}

func (dt NullDate) DateTimeStartOfDay() DateTime {
	date := dt.get()
	if date == nil {
		return DateTime{}
	}

	return date.DateTimeStartOfDay()
}

func (dt NullDate) NullDateTimeEndOfDay() NullDateTime {
	return dt.DateTimeEndOfDay().NullDateTime()
}

func (dt NullDate) NullDateTimeStartOfDay() NullDateTime {
	return dt.DateTimeStartOfDay().NullDateTime()
}

func (dt NullDate) AddDay(d int) NullDate {
	date := dt.get()
	if date != nil {
		*date = date.AddDay(d)
	}

	return date.NullDate()
}

func (dt NullDate) IsEqual(u NullDate) bool {
	value, operand := dt.DateP(), u.DateP()
	if (value == nil && operand != nil) || (operand == nil && value != nil) {
		return false
	}

	return (value == nil && operand == nil) || (value.IsEqual(*operand))
}

func (dt NullDate) IsGreaterThan(u NullDate) bool {
	isGreaterThan := false

	value, operand := dt.DateP(), u.DateP()
	if value != nil && operand != nil {
		isGreaterThan = value.IsGreaterThan(*operand)
	}

	return isGreaterThan
}

func (dt NullDate) IsGreaterThanOrEqual(u NullDate) bool {
	return dt.IsGreaterThan(u) || dt.IsEqual(u)
}

func (dt NullDate) IsLessThan(u NullDate) bool {
	isLessThan := false

	value, operand := dt.DateP(), u.DateP()
	if value != nil && operand != nil {
		isLessThan = value.IsLessThan(*operand)
	}

	return isLessThan
}

func (dt NullDate) IsLessThanOrEqual(u NullDate) bool {
	return dt.IsLessThan(u) || dt.IsEqual(u)
}

func (dt NullDate) String() string {
	date := dt.get()
	if date == nil {
		return `<nil>`
	}

	return date.String()
}

func (dt NullDate) MarshalJSON() ([]byte, error) {
	date := dt.get()
	if date == nil {
		return []byte(`null`), nil
	}

	return json.Marshal(date.String())
}

func (dt *NullDate) UnmarshalJSON(b []byte) error {
	var s *string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	dt.parse(s)

	return nil
}

func (dt *NullDate) UnmarshalText(b []byte) error {
	s := string(b)
	dt.parse(&s)

	return nil
}

func (dt *NullDate) Scan(value interface{}) error {
	if value == nil {
		dt.set(nil)
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		year, month, day := v.Date()
		date := newDate(year, month, day)
		dt.set(&date)
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, dt)
	}

	return nil
}

func (dt NullDate) Value() (driver.Value, error) {
	date := dt.get()
	if date == nil {
		return nil, nil
	}

	return date.Value()
}

func NewNullDate(v *Date) NullDate {
	date := NullDate{}
	date.set(v)

	return date
}

func NullDateValidationFn(sl validator.StructLevel) {
	nullDate := sl.Current().Interface().(NullDate)
	date := nullDate.DateP()
	if date != nil && !date.IsValid() {
		sl.ReportError(nullDate, "", "", "data_type_null_date", "")
	}
}
