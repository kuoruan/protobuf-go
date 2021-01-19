package durationpb

import (
	"database/sql/driver"
	"errors"
	"time"

	"github.com/golang/protobuf/ptypes"
)

func (x *Duration) Scan(value interface{}) error {
	var dd time.Duration
	switch d := value.(type) {
	case int64:
		dd = time.Duration(d)
	case nil:
		dd = time.Duration(0)
	default:
		return errors.New("incompatible type for google.protobuf.Duration")
	}

	*x = *ptypes.DurationProto(dd)
	return nil
}

func (x *Duration) Value() (driver.Value, error) {
	d, err := ptypes.Duration(x)
	if err != nil {
		return nil, err
	}

	return int64(d), nil
}
