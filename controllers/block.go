package controllers

import (
	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/repositories"
)

type BlockController struct{}

var (
	Block BlockController
)

func (c *BlockController) OnReceive(block *models.Block) error {

	repositories.Blocks.Push(block)
	for _, transaction := range block.Transactions {
		sender := repositories.Accounts.Get(transaction.SenderPublicKey)
		transaction.ApplyUnconfirmed(sender)
	}

	return nil
}

func (c *BlockController) Generate(block *models.Block) (*models.Block, error) {
	return nil, nil
}
