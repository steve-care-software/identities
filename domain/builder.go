package domain

import (
	"errors"
	"time"
)

type builder struct {
	name       string
	private    []byte
	public     []byte
	pCreatedOn *time.Time
}

func createBuilder() Builder {
	out := builder{
		name:       "",
		private:    nil,
		public:     nil,
		pCreatedOn: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithPrivate adds a pk to the builder
func (app *builder) WithPrivate(private []byte) Builder {
	app.private = private
	return app
}

// WithPublic adds a pubKey to the builder
func (app *builder) WithPublic(public []byte) Builder {
	app.public = public
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.pCreatedOn = &createdOn
	return app
}

// Now builds a new Identity instance
func (app *builder) Now() (Identity, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Identity instance")
	}

	if app.private != nil && len(app.private) <= 0 {
		app.private = nil
	}

	if app.private == nil {
		return nil, errors.New("the privateKey is mandatory in order to build an Identity instance")
	}

	if app.public != nil && len(app.public) <= 0 {
		app.public = nil
	}

	if app.public == nil {
		return nil, errors.New("the publicKey is mandatory in order to build an Identity instance")
	}

	if app.pCreatedOn == nil {
		return nil, errors.New("the creationTime is mandatory in order to build an Identity instance")
	}

	return createIdentity(app.name, app.private, app.public, *app.pCreatedOn), nil
}
