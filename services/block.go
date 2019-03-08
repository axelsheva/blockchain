package services

import "bitbucket.org/axelsheva/blockchain/models"

var (
	Block IBlockService
)

func init() {
	Block = &BlockService{}
}

type IBlockService interface {
	SetLastBlock(block *models.Block)
	GetLastBlock() *models.Block
}

type BlockService struct {
	lastBlock *models.Block
}

func (bs *BlockService) SetLastBlock(block *models.Block) {
	bs.lastBlock = block
}

func (bs *BlockService) GetLastBlock() *models.Block {
	return bs.lastBlock
}
