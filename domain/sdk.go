package domain

import (
	"time"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an identity builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPrivate(private []byte) Builder
	WithPublic(public []byte) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Identity, error)
}

// Identity represents an identity
type Identity interface {
	Name() string
	Private() []byte
	Public() []byte
	CreatedOn() time.Time
}

// RepositoryBuilder represents a repository builder
type RepositoryBuilder interface {
	Create() RepositoryBuilder
	WithContext(context uint) RepositoryBuilder
	Now() (Repository, error)
}

// Repository represents the identity repository
type Repository interface {
	List() ([]string, error)
	ListDeleted() ([]string, error)
	Retrieve(name string, password []byte) (Identity, error)
}

// ServiceBuilder represents a service builder
type ServiceBuilder interface {
	Create() ServiceBuilder
	WithContext(context uint) ServiceBuilder
	WithKind(kind uint) ServiceBuilder
	Now() (Service, error)
}

// Service represents the identity service
type Service interface {
	Insert(identity Identity, password []byte) error
	Update(name string, updated Identity, originalPassword []byte, newPassword []byte) error
	Delete(identity Identity, password []byte) error
}
