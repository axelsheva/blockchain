package utils_test

import (
	"reflect"
	"sort"
	"testing"
	"time"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/models/types"
	"bitbucket.org/axelsheva/blockchain/utils"
)

func TestBlockSort(t *testing.T) {
	transactions := []*models.Transaction{
		{Type: types.SendType, CreatedAt: time.Unix(40, 0), ID: "b"},
		{Type: types.SendType, CreatedAt: time.Unix(40, 0), ID: "a"},
		{Type: types.SendType, CreatedAt: time.Unix(10, 0), ID: "h"},
		{Type: types.ReferralType, CreatedAt: time.Unix(50, 0), ID: "z"},
	}

	expected := []*models.Transaction{
		transactions[3],
		transactions[2],
		transactions[1],
		transactions[0],
	}

	sort.Sort(utils.BlockSort(transactions))

	isEqual := reflect.DeepEqual(transactions, expected)
	if !isEqual {
		t.Error("Transactions sorting is incorrect")
	}
}

func TestDelegatesSort(t *testing.T) {
	delegates := []*models.Account{
		{
			PublicKey: "1",
			Delegate: &models.Delegate{
				Votes: 10,
			},
		},
		{
			PublicKey: "2",
			Delegate: &models.Delegate{
				Votes: 20,
			},
		},
		{
			PublicKey: "3",
			Delegate: &models.Delegate{
				Votes: 40,
			},
		},
		{
			PublicKey: "4",
			Delegate: &models.Delegate{
				Votes: 30,
			},
		},
	}

	expected := []*models.Account{
		delegates[2],
		delegates[3],
		delegates[1],
		delegates[0],
	}

	sort.Sort(sort.Reverse(utils.ByVotes(delegates)))

	isEqual := reflect.DeepEqual(delegates, expected)
	if !isEqual {
		t.Error("Delegates sorting is incorrect")
	}
}
