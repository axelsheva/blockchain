package services

import (
	"errors"
	"sort"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/repositories"
	"bitbucket.org/axelsheva/blockchain/utils"
)

var (
	Block IBlockService
)

func init() {
	Block = &BlockService{}
}

type IBlockService interface {
	SetLastBlock(block *models.Block)
	GetLastBlock() *models.Block
	ApplyGenesisBlock(block *models.Block) error
	Process(block *models.Block) error
}

type BlockService struct {
	lastBlock *models.Block
}

func (s *BlockService) SetLastBlock(block *models.Block) {
	s.lastBlock = block
}

func (s *BlockService) GetLastBlock() *models.Block {
	return s.lastBlock
}

func (s *BlockService) ApplyGenesisBlock(block *models.Block) error {
	for _, transaction := range block.Transactions {
		if repositories.Accounts.Get(transaction.SenderPublicKey) != nil {
			continue
		}

		account := models.Account{
			PublicKey: transaction.SenderPublicKey,
		}

		repositories.Accounts.Add(&account)
	}

	sort.Sort(utils.BlockSort(block.Transactions))

	return s.Process(block)
}

func (s *BlockService) Process(block *models.Block) error {
	if repositories.Blocks.IsExists(block.ID) {
		return errors.New("Block is already exists")
	}

	for _, transaction := range block.Transactions {
		sender := repositories.Accounts.Get(transaction.SenderPublicKey)

		transaction.ApplyUnconfirmed(sender)
	}

	repositories.Blocks.Push(block)

	return nil
}
