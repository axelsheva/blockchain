package utils

import "time"

func CaltulateFirstSlot(timestamp time.Time, slotInverval uint8, activeDevegates uint16) uint64 {
	slot := (timestamp.Unix() - int64(slotInverval)) / int64(slotInverval)
	diff := slot % int64(activeDevegates)
	return uint64(slot - diff*int64(slotInverval))
}
