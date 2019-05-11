package utils

import "time"

func CaltulateFirstSlot(timestamp time.Time, slotInverval uint8, activeDevegates int) int64 {
	slot := (timestamp.Unix() - int64(slotInverval)) / int64(slotInverval)
	diff := slot % int64(activeDevegates)
	return slot - diff*int64(slotInverval)
}
