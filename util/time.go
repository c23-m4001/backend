package util

import (
	"capstone/data_type"
	"time"
)

func CurrentDateTime() data_type.DateTime {
	return data_type.NewDateTime(time.Now())
}

func CurrentNullDateTime() data_type.NullDateTime {
	return CurrentDateTime().NullDateTime()
}

func ParseDateTime(s string) data_type.DateTime {
	dateTime := data_type.DateTime{}
	if err := dateTime.UnmarshalJSON([]byte(s)); err != nil {
		return data_type.DateTime{}
	}
	return dateTime
}

func ParseNullDateTime(s string) data_type.NullDateTime {
	dateTime := data_type.NullDateTime{}
	if err := dateTime.UnmarshalJSON([]byte(s)); err != nil {
		return data_type.NullDateTime{}
	}
	return dateTime
}
