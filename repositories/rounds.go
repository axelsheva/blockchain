package repositories

import (
	"log"

	"bitbucket.org/axelsheva/blockchain/models"
)

var (
	Rounds IRoundsRepository
)

func init() {
	Rounds = &RoundsRepository{}
}

type IRoundsRepository interface {
	Push(round *models.Round)
	Pop() *models.Round
}

type RoundsRepository struct {
	rounds []*models.Round
}

func (r *RoundsRepository) Push(round *models.Round) {
	log.Printf("[Repository][Round][Push] %+v", round.Slots)

	r.rounds = append(r.rounds, round)
}

func (r *RoundsRepository) Pop() *models.Round {
	if len(r.rounds) > 1 {
		lastRound := r.rounds[len(r.rounds)-1]
		r.rounds = r.rounds[:len(r.rounds)-1]

		return lastRound
	}
	return nil
}
