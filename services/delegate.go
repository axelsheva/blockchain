package services

import (
	"sort"
	"time"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/repositories"
	"bitbucket.org/axelsheva/blockchain/utils"
	"github.com/jamesruan/sodium"
)

var (
	Delegate IDelegateService
)

func init() {
	Delegate = &DelegateService{
		active: 3,
	}
}

type IDelegateService interface {
	Forge(timestamp time.Time)
	GetActive() []*models.Account
}

type DelegateService struct {
	keyPair *sodium.SignKP
	active  int
}

func (s *DelegateService) Forge(timestamp time.Time) {
	Block.Generate(s.keyPair, timestamp)
}

func (s *DelegateService) GetActive() []*models.Account {
	delegates := repositories.Delegates.GetAll()
	sort.Sort(sort.Reverse(utils.ByVotes(delegates)))
	active := make([]*models.Account, 0, s.active)
	for _, delegate := range delegates {
		active = append(active, delegate)
		if len(active) >= s.active {
			break
		}
	}
	return active
}
