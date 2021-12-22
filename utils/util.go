package utils

import (
	"fmt"
	"time"
)

func GetNowFormatTodayTime () string {
	now := time.Now()
	dataStr := fmt.Sprintf("%02d-%02d-%02d", now.Year(), now.Month(), now.Day())

	return dataStr
}