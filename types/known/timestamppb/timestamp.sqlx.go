package timestamppb

import (
	"database/sql/driver"
	"errors"
	"time"
)

// Scan implements the Scanner interface.
func (x *Timestamp) Scan(src interface{}) error {
	switch t := src.(type) {
	case time.Time:
		*x = *New(t)
	default:
		return errors.New("not a Time")
	}
	return nil
}

// Value implements the driver Valuer interface.
func (x *Timestamp) Value() (driver.Value, error) {
	if !x.IsValid() {
		return nil, nil
	}
	return x.AsTime(), nil
}
