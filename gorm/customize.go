package gorm

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/spf13/cast"
)

type (
	TimeString string
	JsonString []string
	JsonInt64  []int64
	JsonInt    []int
)

func (c TimeString) Value() (driver.Value, error) {
	return string(c), nil
}

func (c *TimeString) Scan(value any) error {
	if value == nil {
		return nil
	}
	switch s := value.(type) {
	case time.Time:
		*c = TimeString(s.Format("2006-01-02 15:04:05"))
		return nil
	}
	return nil
}

func (c JsonString) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *JsonString) Scan(value any) error {
	return json.Unmarshal(value.([]byte), c)
}

func (c JsonInt64) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return cast.ToInt64(string(b)), err
}

func (c *JsonInt64) Scan(value any) error {
	return json.Unmarshal(value.([]byte), c)
}

func (c JsonInt) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return cast.ToInt(string(b)), err
}

func (c *JsonInt) Scan(value any) error {
	return json.Unmarshal(value.([]byte), c)
}
