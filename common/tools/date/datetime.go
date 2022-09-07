package date

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type DateTime time.Time

var _ driver.ValueConverter = Now()
var _ driver.Value = Now()

func Now() DateTime {
	return DateTime(time.Now())
}

func (t DateTime) String() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}

func NowStr() string {
	return Now().String()
}
func (t DateTime) ConvertValue(v any) (driver.Value, error) {
	switch v.(type) {
	case time.Time:
		return v, nil
	}
	return time.Time(t), nil
}

func (t DateTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if time.Time(t).UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return time.Time(t), nil
}

func (t *DateTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(data), time.Local)
	*t = DateTime(now)
	return
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.String())), nil
}

//func StrToTime(str string) (*time.Time, error) {
//	t, err := time.Parse("2006-01-02 15:04:05", str)
//	if err != nil {
//		return nil, err
//	}
//	return &t, nil
//}
