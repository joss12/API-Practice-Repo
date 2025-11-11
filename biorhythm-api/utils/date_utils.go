package utils

import (
	"fmt"
	"time"
)

// ParseDate parses date string in YYYY-MM-DD format
func ParseDate(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, fmt.Errorf("date string cannot be empty")
	}

	// Parse the date using the RFC3339 date format
	parsedTime, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format, expected YYYY-MM-DD, got: %s", dateStr)
	}

	return parsedTime, nil
}

// FormatDate formats time to YYYY-MM-DD string
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// DaysBetween calculates days between two dates
func DaysBetween(start, end time.Time) int {
	return int(end.Sub(start).Hours() / 24)
}

// IsValidDate checks if date string is in valid format
func IsValidDate(dateStr string) bool {
	_, err := ParseDate(dateStr)
	return err == nil
}

// AddDays adds specified number of days to a date
func AddDays(date time.Time, days int) time.Time {
	return date.AddDate(0, 0, days)
}

// ValidateDateRange ensures target date is not before birth date
func ValidateDateRange(birthDate, targetDate time.Time) error {
	if targetDate.Before(birthDate) {
		return fmt.Errorf("target date (%s) cannot be before birth date (%s)",
			FormatDate(targetDate), FormatDate(birthDate))
	}
	return nil
}

// GetCurrentDate returns current date in YYYY-MM-DD format
func GetCurrentDate() string {
	return time.Now().Format("2006-01-02")
}

// ParseDateWithValidation parses and validates date format
func ParseDateWithValidation(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, fmt.Errorf("date cannot be empty")
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format '%s', expected format: YYYY-MM-DD (e.g., 1990-05-15)", dateStr)
	}

	// Check if date is reasonable (not too far in past/future)
	now := time.Now()
	if date.After(now.AddDate(100, 0, 0)) {
		return time.Time{}, fmt.Errorf("date cannot be more than 100 years in the future")
	}

	if date.Before(time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)) {
		return time.Time{}, fmt.Errorf("date cannot be before year 1900")
	}

	return date, nil
}
