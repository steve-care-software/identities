package databases

import (
	database_application "github.com/steve-care-software/databases/applications"
	application_identity "github.com/steve-care-software/identities/applications"
	identities "github.com/steve-care-software/identities/domain"
)

type identityApplication struct {
	database          database_application.Application
	repositoryBuilder identities.RepositoryBuilder
	serviceBuilder    identities.ServiceBuilder
	dbName            string
}

func createIdentityApplication(
	database database_application.Application,
	repositoryBuilder identities.RepositoryBuilder,
	serviceBuilder identities.ServiceBuilder,
	dbName string,
) application_identity.Application {
	out := identityApplication{
		database:          database,
		repositoryBuilder: repositoryBuilder,
		serviceBuilder:    serviceBuilder,
		dbName:            dbName,
	}

	return &out
}

// List returns the list of identities
func (app *identityApplication) List() ([]string, error) {
	repository, pContext, err := app.openRepository()
	if err != nil {
		return nil, err
	}

	defer app.database.Close(*pContext)
	return repository.List()
}

// Insert inserts an identity
func (app *identityApplication) Insert(identity identities.Identity, password []byte) error {
	service, pContext, err := app.openService()
	if err != nil {
		return err
	}

	err = service.Insert(identity, password)
	if err != nil {
		return app.cancelService(*pContext)
	}

	return app.saveService(*pContext)
}

// Update updates an identity
func (app *identityApplication) Update(name string, updated identities.Identity, originalPassword []byte, newPassword []byte) error {
	service, pContext, err := app.openService()
	if err != nil {
		return err
	}

	err = service.Update(name, updated, originalPassword, newPassword)
	if err != nil {
		return app.cancelService(*pContext)
	}

	return app.saveService(*pContext)
}

// Retrieve retrieves an identity
func (app *identityApplication) Retrieve(name string, password []byte) (identities.Identity, error) {
	repository, pContext, err := app.openRepository()
	if err != nil {
		return nil, err
	}

	defer app.database.Close(*pContext)
	return repository.Retrieve(name, password)
}

func (app *identityApplication) cancelService(context uint) error {
	defer app.database.Close(context)
	return app.database.Cancel(context)
}

func (app *identityApplication) saveService(context uint) error {
	defer app.database.Close(context)
	return app.database.Commit(context)
}

func (app *identityApplication) openService() (identities.Service, *uint, error) {
	pContext, err := app.database.Open(app.dbName)
	if err != nil {
		return nil, nil, err
	}

	service, err := app.serviceBuilder.Create().WithContext(*pContext).Now()
	if err != nil {
		return nil, nil, err
	}

	return service, pContext, nil
}

func (app *identityApplication) openRepository() (identities.Repository, *uint, error) {
	pContext, err := app.database.Open(app.dbName)
	if err != nil {
		return nil, nil, err
	}

	repository, err := app.repositoryBuilder.Create().WithContext(*pContext).Now()
	if err != nil {
		return nil, nil, err
	}

	return repository, pContext, nil
}
