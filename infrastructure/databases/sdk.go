package databases

import (
	database_application "github.com/steve-care-software/databases/applications"
	application_identity "github.com/steve-care-software/identities/applications"
	identities "github.com/steve-care-software/identities/domain"
	"github.com/steve-care-software/libs/cryptography/hash"
)

const identityList = "identity:list"
const identityListDeleted = "identity:list:deleted"

// NewIdentityApplication creates a new identity application
func NewIdentityApplication(
	repositoryBuilder identities.RepositoryBuilder,
	serviceBuilder identities.ServiceBuilder,
	dbName string,
	database database_application.Application,
) application_identity.Application {
	return createIdentityApplication(database, repositoryBuilder, serviceBuilder, dbName)
}

// NewIdentityServiceBuilder creates a new identity service builder
func NewIdentityServiceBuilder(
	repositoryBuilder identities.RepositoryBuilder,
	database database_application.Application,
) identities.ServiceBuilder {
	hashAdapter := hash.NewAdapter()
	return createIdentityServiceBuilder(hashAdapter, database, repositoryBuilder)
}

// NewIdentityRepositoryBuilder creates a new identy repository builder
func NewIdentityRepositoryBuilder(
	database database_application.Application,
) identities.RepositoryBuilder {
	hashAdapter := hash.NewAdapter()
	builder := identities.NewBuilder()
	return createIdentityRepositoryBuilder(hashAdapter, database, builder)
}
