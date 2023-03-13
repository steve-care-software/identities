package databases

import (
	"errors"

	database_application "github.com/steve-care-software/databases/applications"
	identities "github.com/steve-care-software/identities/domain"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type identityRepositoryBuilder struct {
	hashAdapter hash.Adapter
	database    database_application.Application
	builder     identities.Builder
	pContext    *uint
}

func createIdentityRepositoryBuilder(
	hashAdapter hash.Adapter,
	database database_application.Application,
	builder identities.Builder,
) identities.RepositoryBuilder {
	out := identityRepositoryBuilder{
		hashAdapter: hashAdapter,
		database:    database,
		builder:     builder,
		pContext:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *identityRepositoryBuilder) Create() identities.RepositoryBuilder {
	return createIdentityRepositoryBuilder(app.hashAdapter, app.database, app.builder)
}

// WithContext adds a context to the builder
func (app *identityRepositoryBuilder) WithContext(context uint) identities.RepositoryBuilder {
	app.pContext = &context
	return app
}

// Now builds a new Repository instance
func (app *identityRepositoryBuilder) Now() (identities.Repository, error) {
	if app.pContext == nil {
		return nil, errors.New("the context is mandatory in order to build a Repository instance")
	}

	return createIdentityRepository(
		app.hashAdapter,
		app.database,
		app.builder,
		*app.pContext,
	), nil
}
