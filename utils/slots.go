package utils

import "time"

func CalculateFirstSlot(timestamp time.Time, slotInverval uint8, activeDevegates uint16) int64 {
	slot := CalculateSlot(timestamp, slotInverval)
	return slot / int64(activeDevegates) * int64(activeDevegates)
}

func CalculateSlot(timestamp time.Time, slotInverval uint8) int64 {
	return timestamp.Unix() / int64(slotInverval)
}
