package data_type

import (
	"capstone/config"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type DateTime struct {
	origin   time.Time
	parseErr error
}

func (dt DateTime) layout() string {
	return time.RFC3339
}

func (dt DateTime) databaseLayout() string {
	return "2006-01-02 15:04:05.999999"
}

func (dt DateTime) IsoLayout() string {
	return "YYYY-MM-DDThh:mm:ssZ"
}

func (dt DateTime) Time() time.Time {
	return dt.origin
}

func (dt DateTime) IsZero() bool {
	return dt.Time().IsZero()
}

func (dt DateTime) Format(layout string) string {
	return dt.Time().Format(layout)
}

func (dt DateTime) FormatUTC(layout string) string {
	return dt.Time().UTC().Format(layout)
}

func (dt *DateTime) parse(layout string, s string) {
	dt.origin, dt.parseErr = time.Time{}, nil
	if s == "" {
		return
	}

	t, err := time.Parse(layout, s)
	dt.parseErr = err
	if err == nil {
		*dt = NewDateTime(t)
	}
}

// implement Stringer interface
func (dt DateTime) String() string {
	s := "0000-00-00T00:00:00Z"

	if !dt.IsZero() {
		s = dt.Format(dt.layout())
	}

	return s
}

// implement MarshalJSON interface
func (dt DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(dt.String())
}

// implement UnmarshalJSON interface
func (dt *DateTime) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	dt.parse(dt.layout(), s)

	return nil
}

// implement UnmarshalText interface
func (dt *DateTime) UnmarshalText(b []byte) error {
	dt.parse(dt.layout(), string(b))
	return nil
}

// implement database Scanner interface
func (dt *DateTime) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		*dt = NewDateTime(v)
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, dt)
	}

	return nil
}

// implement database Valuer inteface
func (dt DateTime) Value() (driver.Value, error) {
	s := "0000-00-00 00:00:00.000000"

	if !dt.IsZero() {
		s = dt.FormatUTC(dt.databaseLayout())
	}

	return s, nil
}

func NewDateTime(t time.Time) DateTime {
	return DateTime{origin: t.In(config.GetTimeLocation())}
}
