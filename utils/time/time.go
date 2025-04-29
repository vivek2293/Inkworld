package timeutils

import "time"

var defaultLocation *time.Location

func InitLocation(location string) error {
	var err error
	defaultLocation, err = time.LoadLocation(location)
	return err
}

// GetCurrentLocalTime returns the current local time based on the initialized location.
func GetCurrentLocalTime() time.Time {
	if defaultLocation == nil {
		return time.Now()
	}

	return time.Now().In(defaultLocation)
}

// GetCurrentUTCTime returns the current UTC time.
func GetCurrentUTCTime() time.Time {
	return time.Now().UTC()
}
