package util

import (
	"time"
)

func GetLastMonth(now time.Time) (startDate, endDate time.Time) {
	// take now, go back 1 month
	// get first day to get "startDate"
	// take now, get first day, go back one day to get "endDate"
	lastMonth := now.AddDate(0, -1, 0)
	startDate = time.Date(lastMonth.Year(), lastMonth.Month(), 1, 0, 0, 0, 0, lastMonth.Location())
	thisMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	endDate = thisMonth.AddDate(0, 0, -1)
	return
}
