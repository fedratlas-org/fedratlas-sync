package peer

import "fedratlas-sync/internal/models"

var peers = make(map[string]models.Peer)

func AddPeer(p models.Peer) {
	peers[p.ID] = p
}

func GetPeers() map[string]models.Peer {
	return peers
}
