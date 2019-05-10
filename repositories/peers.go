package repositories

import (
	"bitbucket.org/axelsheva/blockchain/models"
)

type IPeerRepository interface {
	Add(peer *models.Peer)
	GetByIP(IP string) *models.Peer
}

type PeerRepository struct {
	peers map[string]*models.Peer
}

var Peers IPeerRepository

func init() {
	Peers = &PeerRepository{
		peers: make(map[string]*models.Peer),
	}
}

func (r *PeerRepository) Add(peer *models.Peer) {
	r.peers[peer.IP] = peer
}

func (r *PeerRepository) GetByIP(IP string) *models.Peer {
	return r.peers[IP]
}
