package assets_test

import (
	"testing"

	"bitbucket.org/axelsheva/blockchain/models/assets"
)

func TestDelegateCalculateFee(t *testing.T) {
	asset := assets.Delegate{}

	got := asset.CalculateFee()
	want := uint64(1000000000)

	if got != want {
		t.Errorf("Fee calculation is incorrect. Actual %d, Expected %d", got, want)
	}
}

func TestDelegateGetAmount(t *testing.T) {
	asset := assets.Delegate{}

	got := asset.GetAmount()
	want := uint64(0)

	if got != want {
		t.Errorf("Amount is incorrect. Actual %d, Expected %d", got, want)
	}
}
