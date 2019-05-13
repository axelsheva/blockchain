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

func InitDelegate(forgeSecret string, activeCount uint16) {
	keyPair := utils.GenerateKeyPair(forgeSecret)
	publicKey := hex.EncodeToString(keyPair.PublicKey.Bytes)

	log.Printf("[Service][Delegate][Init] public key: %s", publicKey)

	Delegate = &DelegateService{
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
		diff := forgeTime.Sub(time.Now()).Nanoseconds() / int64(time.Millisecond)

		if (diff + 1000) > 0 {
			log.Printf("[Service][Delegate][Forge] Delegate will be forge in slot %d, after %d ms", slot, diff)
		} else {
			log.Printf("[Service][Delegate][Forge] Skip slot: %d, time: %d ms", slot, diff)
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
