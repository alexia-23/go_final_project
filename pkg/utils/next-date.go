package utils

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func handleDailyPeriod(now time.Time, dstart string, interval int) (string, error) {
	t, err := time.Parse("20060102", dstart)
	if err != nil {
		return "", err
	}
	diff := now.Sub(t)
	days := int(diff.Hours() / 24)
	r := days % interval
	delay := interval - r
	next := now.AddDate(0, 0, delay)

	return next.Format("20060102"), nil
}

func handleYearPeriod(now time.Time, dstart string) (string, error) {
	t, err := time.Parse("20060102", dstart)
	if err != nil {
		return "", err
	}
	next := time.Date(now.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	if !next.After(now) {
		next = time.Date(now.Year()+1, t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	}
	return next.Format("20060102"), nil
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	entries := strings.Split(repeat, " ")
	period := entries[0]
	if period == "d" {
		if len(entries) < 2 {
			return "", errors.New("invalid period")
		}
		interval, err := strconv.Atoi(entries[1])
		if err != nil {
			return "", errors.New("invalid interval")
		}
		return handleDailyPeriod(now, dstart, interval)
	}
	if period == "y" {
		return handleYearPeriod(now, dstart)
	}
	return "", errors.New("invalid period")
}
