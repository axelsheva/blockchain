package utils_test

import (
	"testing"
	"time"

	"bitbucket.org/axelsheva/blockchain/utils"
)

type firstSlotTest struct {
	Time              int64
	ExpectedFirstSlot uint64
}

func TestCalculateFirstSlot(t *testing.T) {
	activeDelegatesCount := uint16(3)
	slotInterval := uint8(10)

	tests := []firstSlotTest{
		{
			Time:              0,
			ExpectedFirstSlot: 0,
		},
		{
			Time:              10,
			ExpectedFirstSlot: 0,
		},
		{
			Time:              90,
			ExpectedFirstSlot: 9,
		},
		{
			Time:              100,
			ExpectedFirstSlot: 9,
		},
		{
			Time:              110,
			ExpectedFirstSlot: 9,
		},
		{
			Time:              120,
			ExpectedFirstSlot: 12,
		},
	}

	for _, test := range tests {
		firstSlot := utils.CalculateFirstSlot(time.Unix(test.Time, 0), slotInterval, activeDelegatesCount)

		if firstSlot != test.ExpectedFirstSlot {
			t.Errorf("Invalid first slot calculation. Expected %d, actual: %d", test.ExpectedFirstSlot, firstSlot)
		}
	}
}

func TestCalculateSlot(t *testing.T) {
	timestamp, err := time.Parse(time.RFC3339, "2019-01-01T00:00:05.000Z")
	if err != nil {
		panic(err)
	}

	slot := utils.CalculateSlot(timestamp, 10)
	expectedSlot := uint64(154630080)
	if slot != expectedSlot {
		t.Errorf("Invalid slot calculation. Expected %d, actual: %d", expectedSlot, slot)
	}
}

func TestCalculateSlot2(t *testing.T) {
	slotInterval := uint8(10)
	count := 100
	for index := 0; index < count; index++ {
		slot := utils.CalculateSlot(time.Unix(int64(index*int(slotInterval)), 0), slotInterval)

		if slot != uint64(index) {
			t.Errorf("Invalid slot calculation. Expected %d, actual: %d", index, slot)
		}
	}
}
