package applications

import identities "github.com/steve-care-software/identities/domain"

// Application represents the chain application
type Application interface {
	List() ([]string, error)
	Insert(identity identities.Identity, password []byte) error
	Update(name string, updated identities.Identity, originalPassword []byte, newPassword []byte) error
	Retrieve(name string, password []byte) (identities.Identity, error)
	Delete(name string, password []byte) error
}
