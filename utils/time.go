package utils

import (
	"database/sql"
	"time"
)

func TimeToFormat(dt sql.NullString, formater string) string {
	if !dt.Valid {
		return ""
	}

	local, _ := time.LoadLocation("Asia/Shanghai")
	t, err := time.ParseInLocation(formater, dt.String, local)
	if err != nil {
		return ""
	}
	return t.Format(formater)
}
