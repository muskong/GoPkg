package gorm

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/muskong/GoPkg/idworker"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type (
	NullString    string
	TimeString    string
	JsonAny       []any
	JsonMapString map[string]string
	JsonString    []string
	JsonInt64     []int64
	JsonInt       []int

	Model struct {
		gorm.Model
		ID        int        `json:"ID,omitempty" db:"id" gorm:"primarykey"`
		Uuid      string     `json:"Uuid,omitempty" db:"uuid"`
		CreatedAt TimeString `json:"CreatedAt,omitempty" db:"created_at"`
		UpdatedAt TimeString `json:"UpdatedAt,omitempty" db:"updated_at"`
		DeletedAt NullString `json:"DeletedAt,omitempty" db:"deleted_at" gorm:"index"`
	}
)

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	m.Uuid = idworker.StringNanoid(30)
	return
}

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

func (c JsonAny) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return b, err
}

func (c *JsonAny) Scan(value any) error {
	return json.Unmarshal(value.([]byte), c)
}

func (c JsonMapString) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *JsonMapString) Scan(value any) error {
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
