package repositories

import (
	"bitbucket.org/axelsheva/blockchain/models"
)

var (
	Blocks IBlocksRepository
)

func init() {
	Blocks = &BlocksRepository{}
}

type IBlocksRepository interface {
	AddOne(block models.Block) *models.Block
	RemoveLast()
}

type BlocksRepository struct {
	blocks []*models.Block
}

func (br *BlocksRepository) AddOne(block models.Block) *models.Block {
	br.blocks = append(br.blocks, &block)
	return &block
}

func (br *BlocksRepository) RemoveLast() {
	if len(br.blocks) > 0 {
		br.blocks = br.blocks[:len(br.blocks)-1]
	}
}
