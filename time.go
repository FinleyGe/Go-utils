package utility

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

// 自定义的时间
type Time struct {
	time.Time
}

// JSON 序列化
func (t Time) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

// JSON 反序列化
func (t *Time) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	var err error
	str := string(b)
	timeStr := strings.Trim(str, "\"")
	tl, err := time.Parse("2006-01-02 15:04:05", timeStr)
	// fmt.Println(timeStr)
	// fmt.Println(tl)
	t.Time = tl
	return err
}

func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *Time) Scan(v interface{}) error {
	// fmt.Println(v)
	value, ok := v.(time.Time)
	if ok {
		*t = Time{Time: value}
		return nil
	}
	log.Fatalln("Time Error", v)
	return errors.New("Time Error")
}
