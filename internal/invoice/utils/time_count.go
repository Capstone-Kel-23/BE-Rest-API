package utils

import (
	"strconv"
	"time"
)

func TimeDueCount(dateStart time.Time, value string) time.Time {
	var result time.Time
	if value == "7" {
		result = dateStart.Add(time.Hour * 24 * 7)
	} else if value == "15" {
		result = dateStart.Add(time.Hour * 24 * 15)
	} else if value == "30" {
		result = dateStart.Add(time.Hour * 24 * 30)
	} else {
		result, _ = time.Parse("2006-01-02", value)
	}
	return result
}

func ExcelSerialDateToTime(serial string) string {
	var excelEpoch = time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
	days, _ := strconv.Atoi(serial)
	convertedTime := excelEpoch.Add(time.Second * time.Duration(days*86400)).Format("2006-01-02")
	return convertedTime
}
