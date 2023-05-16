package data_type

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type NullDateTime struct {
	dateTime DateTime
	isValid  bool
}

func (dt NullDateTime) layout() string {
	return dt.dateTime.layout()
}

func (dt NullDateTime) get() *DateTime {
	var dateTime *DateTime

	if dt.isValid {
		dateTime = new(DateTime)
		*dateTime = dt.dateTime
	}

	return dateTime
}

func (dt *NullDateTime) set(dateTime *DateTime) {
	dt.dateTime, dt.isValid = DateTime{}, false
	if dateTime != nil {
		dt.dateTime, dt.isValid = *dateTime, true
	}
}

func (dt *NullDateTime) parse(s *string) {
	if s == nil {
		dt.set(nil)
		return
	}
	dateTime := DateTime{}
	dateTime.parse(dt.layout(), *s)
	dt.set(&dateTime)
}

func (dt NullDateTime) IsNil() bool {
	return dt.get() == nil
}

func (dt NullDateTime) DateTime() DateTime {
	dateTime := dt.get()
	if dateTime == nil {
		return DateTime{}
	}
	return *dateTime
}

func (dt NullDateTime) DateTimeP() *DateTime {
	return dt.get()
}

func (dt NullDateTime) Add(duration time.Duration) NullDateTime {
	dateTime := dt.get()
	if dateTime != nil {
		*dateTime = dateTime.Add(duration)
	}

	return NewNullDateTime(dateTime)
}

// implement Stringer interface
func (dt NullDateTime) String() string {
	dateTime := dt.get()
	if dateTime == nil {
		return `<nil>`
	}
	return dateTime.String()
}

// implement MarshalJSON interface
func (dt NullDateTime) MarshalJSON() ([]byte, error) {
	dateTime := dt.get()
	if dateTime == nil {
		return []byte("null"), nil
	}
	return json.Marshal(dateTime.String())
}

// implement UnmarshalJSON interface
func (dt *NullDateTime) UnmarshalJSON(b []byte) error {
	var s *string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	dt.parse(s)
	return nil
}

// implement UnmarshalText interface
func (dt *NullDateTime) UnmarshalText(b []byte) error {
	s := string(b)
	dt.parse(&s)
	return nil
}

// implement database Scanner interface
func (dt *NullDateTime) Scan(value interface{}) error {
	if value == nil {
		dt.set(nil)
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		dateTime := NewDateTime(v)
		dt.set(&dateTime)
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, dt)
	}

	return nil
}

// implement database Valuer interface
func (dt NullDateTime) Value() (driver.Value, error) {
	dateTime := dt.get()
	if dateTime == nil {
		return nil, nil
	}

	return dateTime.Value()
}

func NewNullDateTime(v *DateTime) NullDateTime {
	dateTime := NullDateTime{}
	dateTime.set(v)

	return dateTime
}
