package models_test

import (
	"testing"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/models/assets"
	"bitbucket.org/axelsheva/blockchain/repositories"
)

func TestSendApplyUnconfirmed(t *testing.T) {
	sender := models.Account{
		Balance:   100,
		PublicKey: "1",
	}
	recipient := models.Account{
		PublicKey: "2",
	}

	repositories.Accounts.Add(&recipient)

	transaction := models.Transaction{
		Fee: 1,
		Asset: &assets.Send{
			RecipientPublicKey: "2",
			Amount:             1,
		},
	}

	transaction.ApplyUnconfirmed(&sender)

	expectedSender := models.Account{
		Balance:   98,
		PublicKey: "1",
	}
	if sender.Balance != expectedSender.Balance {
		t.Errorf("Invalid sender balance. Actual %d, Expected %d", sender.Balance, expectedSender.Balance)
	}

	expectedRecipient := models.Account{
		Balance:   1,
		PublicKey: "2",
	}
	if recipient.Balance != expectedRecipient.Balance {
		t.Errorf("Invalid recipient balance. Actual %d, Expected %d", recipient.Balance, expectedRecipient.Balance)
	}
}
