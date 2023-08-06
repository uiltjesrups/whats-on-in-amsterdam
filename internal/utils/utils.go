package utils

import "time"

func CurrentDate() time.Time {
	currentTime := time.Now()
	currentDate := currentTime.UTC().Truncate(24 * time.Hour)
	return currentDate
}

func AddOneDay(date time.Time) time.Time {
	oneDay := 24 * time.Hour
	newDate := date.Add(oneDay)
	return newDate
}
