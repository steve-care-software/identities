package objects

import "time"

// Identity represents an identity
type Identity struct {
	Name      string    `json:"name"`
	Private   []byte    `json:"pk"`
	Public    []byte    `json:"pubkey"`
	CreatedOn time.Time `json:"created_on"`
}
