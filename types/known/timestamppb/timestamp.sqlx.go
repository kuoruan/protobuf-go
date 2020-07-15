package timestamppb

import (
	"database/sql/driver"
	"errors"
	"time"
)

// Scan implements the Scanner interface.
func (x *Timestamp) Scan(value interface{}) error {
	switch t := value.(type) {
	case time.Time:
		*x = *New(t)
	case string:
		timePoint, err := time.Parse(time.RFC3339, t)
		if err != nil {
			return err
		}

		*x = *New(timePoint)
	default:
		return errors.New("incompatible type for Timestamp")
	}
	return nil
}

// Value implements the driver Valuer interface.
func (x *Timestamp) Value() (driver.Value, error) {
	if x == nil || !x.IsValid() {
		return time.Unix(0, 0), nil
	}
	return x.AsTime(), nil
}
