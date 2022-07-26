package gorm

import (
	"database/sql/driver"
	"encoding/binary"
	"encoding/json"
)

type (
	JsonString []string
	JsonInt64  []int64
	JsonInt    []int
)

func (c JsonString) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *JsonString) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

func (c JsonInt64) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return int64(binary.BigEndian.Uint64(b)), err
}

func (c *JsonInt64) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

func (c JsonInt) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return int(binary.BigEndian.Uint64(b)), err
}

func (c *JsonInt) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
