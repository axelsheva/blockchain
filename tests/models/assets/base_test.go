package assets_test

import (
	"testing"

	"bitbucket.org/axelsheva/blockchain/models/assets"
)

func TestCalculateFee(t *testing.T) {
	base := assets.Base{}

	got := base.CalculateFee()
	want := uint64(0)

	if got != want {
		t.Errorf("Fee calculation is incorrect. Actual %d, Expected %d", got, want)
	}
}

func TestGetAmount(t *testing.T) {
	base := assets.Base{}

	got := base.GetAmount()
	want := uint64(0)

	if got != want {
		t.Errorf("Amount is incorrect. Actual %d, Expected %d", got, want)
	}
}
