package assets_test

import (
	"testing"

	"bitbucket.org/axelsheva/blockchain/models/assets"
)

func TestSendCalculateFee(t *testing.T) {
	send := assets.Send{
		Amount: 100000000,
	}

	got := send.CalculateFee()
	want := int64(10000)

	if got != want {
		t.Errorf("Send fee calculation is incorrect. Actual %d, Expected %d", got, want)
	}
}
