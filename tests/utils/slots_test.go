package utils_test

import (
	"testing"
	"time"

	"bitbucket.org/axelsheva/blockchain/utils"
)

func TestCaltulateFirstSlot(t *testing.T) {
	timestamp, err := time.Parse(time.RFC3339, "2019-01-01T00:00:05.000Z")
	if err != nil {
		panic(err)
	}

	firstSlot := utils.CaltulateFirstSlot(timestamp, 10, 3)
	expectedFirstSlot := int64(154630059)
	if firstSlot != expectedFirstSlot {
		t.Errorf("Invalid first slot calculation. Expected %d, actual: %d", expectedFirstSlot, firstSlot)
	}
}
