package timestamppb

import (
	"database/sql/driver"
	"errors"
	"time"

	"github.com/golang/protobuf/ptypes"
)

// Scan implements the Scanner interface.
func (x *Timestamp) Scan(value interface{}) error {
	var tt time.Time

	switch t := value.(type) {
	case time.Time:
		tt = t
	case string:
		timePoint, err := time.Parse(time.RFC3339, t)
		if err != nil {
			return err
		}

		tt = timePoint
	default:
		return errors.New("incompatible type for google.protobuf.Timestamp")
	}

	ts, err := ptypes.TimestampProto(tt)
	if err != nil {
		return err
	}

	*x = Timestamp{
		Seconds: ts.GetSeconds(),
		Nanos:   ts.GetNanos(),
	}
	return nil
}

// Value implements the driver Valuer interface.
func (x *Timestamp) Value() (driver.Value, error) {
	return ptypes.Timestamp(x)
}
