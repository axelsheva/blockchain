package repositories

import (
	"log"

	"bitbucket.org/axelsheva/blockchain/models"
)

var (
	Blocks IBlocksRepository
)

func init() {
	Blocks = &BlocksRepository{
		indexes: make(map[string]*models.Block),
	}
}

type IBlocksRepository interface {
	IsExists(ID string) bool
	Push(block *models.Block) *models.Block
	Pop() *models.Block
	GetLast() *models.Block
}

type BlocksRepository struct {
	blocks  []*models.Block
	indexes map[string]*models.Block
}

func (r *BlocksRepository) Push(block *models.Block) *models.Block {
	log.Printf("[Repository][Block][Push] Block ID: %s", block.ID)

	r.blocks = append(r.blocks, block)
	r.indexes[block.ID] = block

	return block
}

func (r *BlocksRepository) Pop() *models.Block {
	if len(r.blocks) > 1 {
		lastBlock := r.blocks[len(r.blocks)-1]
		r.blocks = r.blocks[:len(r.blocks)-1]
		delete(r.indexes, lastBlock.ID)

		return lastBlock
	}
	return nil
}

func (r *BlocksRepository) IsExists(ID string) bool {
	_, ok := r.indexes[ID]

	return ok
}

func (r *BlocksRepository) GetLast() *models.Block {
	if len(r.blocks) > 0 {
		return r.blocks[len(r.blocks)-1]
	}
	return nil
}
