package repositories

import (
	"log"

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
	log.Printf("[Repository][Peer][Add] Added peer: %s, heigth: %d, broadhash: %s", peer.IP, peer.Height, peer.Broadhash)

	r.peers[peer.IP] = peer
}

func (r *PeerRepository) GetByIP(IP string) *models.Peer {
	return r.peers[IP]
}
