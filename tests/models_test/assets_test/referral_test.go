package assets_test

import (
	"testing"

	"bitbucket.org/axelsheva/blockchain/models/assets"
)

func TestReferralCalculateFee(t *testing.T) {
	asset := assets.Referral{}

	got := asset.CalculateFee()
	want := int64(0)

	if got != want {
		t.Errorf("Fee calculation is incorrect. Actual %d, Expected %d", got, want)
	}
}

func TestReferralGetAmount(t *testing.T) {
	asset := assets.Referral{}

	got := asset.GetAmount()
	want := int64(0)

	if got != want {
		t.Errorf("Amount is incorrect. Actual %d, Expected %d", got, want)
	}
}
