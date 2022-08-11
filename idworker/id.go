package idworker

import (
	"fmt"
	"math/rand"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func Nanoid(size int) string {
	return gonanoid.MustGenerate("123456789abcdefghijklmnpqrstuvwxyzABCDEFGHIJKLMNPQRSTUVWXYZ", size)
}

func NumberNanoid(size int) string {
	return gonanoid.MustGenerate("123456789", size)
}

func StringNanoid(size int) string {
	return gonanoid.MustGenerate("abcdefghijklmnpqrstuvwxyzABCDEFGHIJKLMNPQRSTUVWXYZ", size)
}

/**
 * 订单编号`
 */
func IdWorker(uuidServer string) string {
	t := time.Now()
	return fmt.Sprintf("%s%d%d",
		uuidServer,
		t.Unix(),
		t.Nanosecond()+rand.Intn(999))
}

/**
 * 唯一编号
 */
func NumberWorker() string {
	nowTime := time.Now()
	return fmt.Sprintf("%s%s%s%s%s",
		nowTime.Format("06"),               // 2位年份
		PadLeft(nowTime.YearDay(), 3, "0"), // 3位天数
		PadLeft(nowTime.Hour(), 2, "0"),    // 2位小时
		PadLeft(nowTime.Minute(), 2, "0"),  // 2位分钟
		// PadLeft(nowTime.Nanosecond(), 2, "0"),  // 2位分钟
		PadLeft(rand.Intn(999), 3, "0"),
		// PadLeft(nowTime.Second(), 2, "0"),  // 2位秒
		// nowTime.Nanosecond(),               //那一秒内的纳秒
	)
}

func IdInt64(uuidServer int64) (id int64) {

	now := time.Now()
	var base int64 = 1000000000000

	year := base * int64(now.Year()%100)
	days := base / 1000 * int64(now.YearDay())
	// seconds := int64(now.Hour()*60*60+now.Minute()*60+now.Second())
	nano := int64(now.Nanosecond()+now.Second()) / 1000

	// 设备号
	machine := base / 1000000 * uuidServer

	id = year + days + machine + nano

	return id
}
