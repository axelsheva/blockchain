package repositories_test

import (
	"testing"

	"bitbucket.org/axelsheva/blockchain/models"

	"bitbucket.org/axelsheva/blockchain/repositories"
)

func TestGetLast(t *testing.T) {
	lastBlock := repositories.Blocks.GetLast()

	if lastBlock != nil {
		t.Errorf("Incorrect last block. Actual %v, Expected nil", lastBlock)
	}

	block := models.Block{}
	repositories.Blocks.Push(&block)

	lastBlock = repositories.Blocks.GetLast()
	if lastBlock != &block {
		t.Errorf("Incorrect last block. Actual %v, Expected %v", lastBlock, block)
	}
}
