package unixTime

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type UnixTime int64

func (ut *UnixTime) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		*ut = UnixTime(v.Unix())
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
	return nil
}

func (ut UnixTime) Value() (driver.Value, error) {
	t := time.Unix(int64(ut), 0)
	return t, nil
}

func (ut UnixTime) TimeX() time.Time {
	return time.Unix(int64(ut), 0)
}

func (ut UnixTime) GetUpdf(v time.Time) UnixTime {
	return UnixTime(v.Unix())
}
