package assets_test

import (
	"testing"

	"bitbucket.org/axelsheva/blockchain/models/assets"
)

func TestBaseCalculateFee(t *testing.T) {
	asset := assets.Base{}

	got := asset.CalculateFee()
	want := int64(0)

	if got != want {
		t.Errorf("Fee calculation is incorrect. Actual %d, Expected %d", got, want)
	}
}

func TestBaseGetAmount(t *testing.T) {
	asset := assets.Base{}

	got := asset.GetAmount()
	want := int64(0)

	if got != want {
		t.Errorf("Amount is incorrect. Actual %d, Expected %d", got, want)
	}
}
