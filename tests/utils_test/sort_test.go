package utils_test

import (
	"reflect"
	"sort"
	"testing"
	"time"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/utils"
)

func TestBlockSort(t *testing.T) {
	transactions := []*models.Transaction{
		{Type: models.SendType, CreatedAt: time.Unix(40, 0), ID: "b"},
		{Type: models.SendType, CreatedAt: time.Unix(40, 0), ID: "a"},
		{Type: models.SendType, CreatedAt: time.Unix(10, 0), ID: "h"},
		{Type: models.ReferralType, CreatedAt: time.Unix(50, 0), ID: "z"},
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
