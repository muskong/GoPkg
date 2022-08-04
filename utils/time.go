package utils

import (
	"database/sql"
	"time"
)

func NullStringToFormat(dt sql.NullString, formater string) string {
	if !dt.Valid {
		return ""
	}

	local, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return ""
	}
	t, err := time.ParseInLocation(time.RFC3339, dt.String, local)
	if err != nil {
		return ""
	}
	return t.Format(formater)
}

func StringToFormat(dt string, formater string) string {
	local, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return ""
	}
	t, err := time.ParseInLocation(time.RFC3339, dt, local)
	if err != nil {
		return ""
	}
	return t.Format(formater)
}
