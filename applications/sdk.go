package applications

import identities "github.com/steve-care-software/identities/domain"

// GenerateKeyPairFn is a func to generate key pair
type GenerateKeyPairFn func() ([]byte, []byte)

// Application represents the chain application
type Application interface {
	List() ([]string, error)
	Insert(identity identities.Identity, password []byte) error
	Update(name string, updated identities.Identity, originalPassword []byte, newPassword []byte) error
	Retrieve(name string, password []byte) (identities.Identity, error)
}
