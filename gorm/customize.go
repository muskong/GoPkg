package gorm

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/spf13/cast"
)

type (
	NullString string
	TimeString string
	JsonMap    []map[string]any
	JsonString []string
	JsonInt64  []int64
	JsonInt    []int
)

func (c NullString) Value() (driver.Value, error) {
	if len(c) > 0 {
		return string(c), nil
	}
	return nil, nil
}

func (c *NullString) Scan(value any) error {
	if value == nil {
		return nil
	}
	switch s := value.(type) {
	case time.Time:
		*c = NullString(s.Format("2006-01-02 15:04:05"))
		return nil
	}
	return nil
}

func (c TimeString) Value() (driver.Value, error) {
	if len(c) > 0 {
		return string(c), nil
	}
	return time.Now().Format("2006-01-02 15:04:05"), nil
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

func (c JsonMap) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *JsonMap) Scan(value any) error {
	return json.Unmarshal(value.([]byte), c)
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
