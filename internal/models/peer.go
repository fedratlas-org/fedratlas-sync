package models

type Peer struct {
	ID        string `json:"id"`
	BaseURL   string `json:"base_url"`
	PublicKey string `json:"public_key"`
	LastSeen  int64  `json:"last_seen"`
}
