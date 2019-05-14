package services

import (
	"encoding/hex"
	"log"
	"sort"
	"time"

	"bitbucket.org/axelsheva/blockchain/configs"
	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/repositories"
	"bitbucket.org/axelsheva/blockchain/utils"
	"github.com/jamesruan/sodium"
)

var (
	Delegate IDelegateService
)

func NewDelegate(forgeSecret string, activeCount uint16) IDelegateService {
	keyPair := utils.GenerateKeyPair(forgeSecret)
	publicKey := hex.EncodeToString(keyPair.PublicKey.Bytes)

	log.Printf("[Service][Delegate][NewDelegate] public key: %s", publicKey)

	return &DelegateService{
		keyPair:     keyPair,
		publicKey:   models.PublicKey(publicKey),
		activeCount: activeCount,
	}
}

type IDelegateService interface {
	Forge(round *models.Round)
	GetActive() []*models.Account
	GetPublicKey() models.PublicKey
}

type DelegateService struct {
	publicKey   models.PublicKey
	keyPair     *sodium.SignKP
	activeCount uint16
}

func (s *DelegateService) Forge(round *models.Round) {
	if slot, ok := round.Slots[s.publicKey]; ok {
		forgeTime := time.Unix(slot*int64(configs.Const.SlotInterval), 0)
		diff := forgeTime.Sub(time.Now())

		if (diff.Seconds() + 1) > 0 {
			log.Printf("[Service][Delegate][Forge] Delegate will be forge in slot %d, after %s", slot, diff.String())

			go func() {
				time.Sleep(diff)

				Block.Generate(s.keyPair, forgeTime)
			}()
		} else {
			log.Printf("[Service][Delegate][Forge] Skip slot: %d, time: %s", slot, diff.String())
		}
	}

	// diff := timestamp.Sub(time.Now())
	// log.Printf("[Service][Delegate][Forge] %d", diff.Seconds())

	// Block.Generate(s.keyPair, timestamp)
}

func (s *DelegateService) GetActive() []*models.Account {
	delegates := repositories.Delegates.GetAll()
	sort.Sort(sort.Reverse(utils.ByVotes(delegates)))
	active := make([]*models.Account, 0, s.activeCount)
	for _, delegate := range delegates {
		active = append(active, delegate)
		if len(active) >= int(s.activeCount) {
			break
		}
	}
	return active
}

func (s *DelegateService) GetPublicKey() models.PublicKey {
	return s.publicKey
}
