package service

import (
	"errors"
	"time"
)

func parseDOB(dob string) (time.Time, error) {
	parsedDOB, err := time.Parse(
		"2006-01-02",
		dob,
	)

	if err != nil {
		return time.Time{}, errors.New(
			"invalid date format, use YYYY-MM-DD",
		)
	}

	if parsedDOB.After(time.Now()) {
		return time.Time{}, errors.New(
			"date of birth cannot be in the future",
		)
	}

	return parsedDOB, nil
}
