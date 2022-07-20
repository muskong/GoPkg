package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func Interface2String(inter interface{}) string {
	switch inter.(type) {

	case string:
		return fmt.Sprintf("%s", inter.(string))
	case int:
		return fmt.Sprintf("%d", inter.(int))
	case int64:
		return fmt.Sprintf("%d", inter.(int64))
	case float64:
		return fmt.Sprintf("%0.f", inter.(float64))
	case json.Number:
		return inter.(json.Number).String()
	}

	return ""
}

func StrToFloat64(str string) float64 {
	if len(str) <= 0 {
		return 0
	}

	floatValue, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return floatValue
}

func StrToInt(str string) int {
	if len(str) <= 0 {
		return 0
	}

	intValue, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return intValue
}

func StrToInt64(str string) int64 {
	if len(str) <= 0 {
		return 0
	}

	intValue, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}

	return intValue
}

/**
 * 金额处理，10000分 == 100.00元
 */
func Percentile(money string) string {
	return fmt.Sprintf("%.2f", StrToFloat64(money)/100)
}

/**
 * 金额处理，10000分 == 100.00元
 */
func PercentileString(money float64) string {
	return fmt.Sprintf("%.2f", money/100)
}

/**
 * 金额处理，10000分 == 100.00元
 */
func PercentileFloat(money float64) float64 {
	return StrToFloat64(fmt.Sprintf("%.2f", money/100))
}

/**
 * 数字千分位处理，10000 == 10,000
 */
func NumberFormat(str string) string {
	length := len(str)
	if length < 4 {
		return str
	}
	arr := strings.Split(str, ".")
	length1 := len(arr[0])
	if length1 < 4 {
		return str
	}
	count := (length1 - 1) / 3
	for i := 0; i < count; i++ {
		arr[0] = arr[0][:length1-(i+1)*3] + "," + arr[0][length1-(i+1)*3:]
	}
	return strings.Join(arr, ".")
}
