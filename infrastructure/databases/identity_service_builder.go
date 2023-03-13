package databases

import (
	"errors"

	database_application "github.com/steve-care-software/databases/applications"
	identities "github.com/steve-care-software/identities/domain"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type identityServiceBuilder struct {
	hashAdapter       hash.Adapter
	database          database_application.Application
	repositoryBuilder identities.RepositoryBuilder
	pContext          *uint
	pKind             *uint
	pNameKind         *uint
}

func createIdentityServiceBuilder(
	hashAdapter hash.Adapter,
	database database_application.Application,
	repositoryBuilder identities.RepositoryBuilder,
) identities.ServiceBuilder {
	out := identityServiceBuilder{
		hashAdapter:       hashAdapter,
		database:          database,
		repositoryBuilder: repositoryBuilder,
		pContext:          nil,
		pKind:             nil,
		pNameKind:         nil,
	}

	return &out
}

// Create initializes the builder
func (app *identityServiceBuilder) Create() identities.ServiceBuilder {
	return createIdentityServiceBuilder(app.hashAdapter, app.database, app.repositoryBuilder)
}

// WithContext adds a context to the builder
func (app *identityServiceBuilder) WithContext(context uint) identities.ServiceBuilder {
	app.pContext = &context
	return app
}

// WithKind adds a kind to the builder
func (app *identityServiceBuilder) WithKind(kind uint) identities.ServiceBuilder {
	app.pKind = &kind
	return app
}

// WithNameKind adds a nameKind to the builder
func (app *identityServiceBuilder) WithNameKind(nameKind uint) identities.ServiceBuilder {
	app.pNameKind = &nameKind
	return app
}

// Now builds a new Service instance
func (app *identityServiceBuilder) Now() (identities.Service, error) {
	if app.pContext == nil {
		return nil, errors.New("the context is mandatory in order to build a Service instance")
	}

	if app.pKind == nil {
		return nil, errors.New("the kind is mandatory in order to build a Service instance")
	}

	if app.pNameKind == nil {
		return nil, errors.New("the name kind is mandatory in order to build a Repository instance")
	}

	repository, err := app.repositoryBuilder.Create().WithContext(*app.pContext).WithKind(*app.pKind).WithNameKind(*app.pNameKind).Now()
	if err != nil {
		return nil, err
	}

	return createIdentityService(
		app.hashAdapter,
		app.database,
		repository,
		*app.pContext,
		*app.pKind,
		*app.pNameKind,
	), nil
}
