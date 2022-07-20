package sqlx

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

func Sql(filed, table, where, sortField, sort string, limit, offset int64) string {
	sql := `SELECT %s
		FROM %s
		WHERE %s
		ORDER BY %s %s
		LIMIT %d
		OFFSET %d`

	if where == "" {
		where = "1=1"
	}

	return fmt.Sprintf(sql, filed, table, where, sortField, sort, limit, offset)
}

func SqlDeleted(filed, table, where, sortField, sort string, limit, offset int64) string {
	where = "deleted_at IS NULL AND " + where
	return Sql(filed, table, where, sortField, sort, limit, offset)
}

type JsonString []string

func (c JsonString) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *JsonString) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
