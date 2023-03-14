package applications

import (
	identities "github.com/steve-care-software/identities/domain"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// OnSignFn represents an onSign fn
type OnSignFn func(hash hash.Hash, pk []byte) ([]byte, error)

// OnVerifySignatureFn represents an onVerifySignature fn
type OnVerifySignatureFn func(pubKey []byte, signature []byte, hash hash.Hash) (bool, error)

// Application represents the chain application
type Application interface {
	List() ([]string, error)
	Insert(identity identities.Identity, password []byte) error
	Update(name string, updated identities.Identity, originalPassword []byte, newPassword []byte) error
	Retrieve(name string, password []byte) (identities.Identity, error)
	Delete(name string, password []byte) error
	Sign(hash hash.Hash, identity identities.Identity) ([]byte, error)
	VerifySignature(signature []byte, hash hash.Hash) (bool, error)
}
