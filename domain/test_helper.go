package domain

import "time"

// NewIdentityForTests creates a new identity for tests
func NewIdentityForTests(name string) Identity {
	pubKey := []byte("this is a pubKey")
	pk := []byte("this is a pk")
	createdOn := time.Now().UTC()
	ins, err := NewBuilder().Create().
		WithName(name).
		WithPublic(pubKey).
		WithPrivate(pk).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
