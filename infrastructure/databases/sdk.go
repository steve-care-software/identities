package databases

import (
	database_application "github.com/steve-care-software/databases/applications"
	application_identity "github.com/steve-care-software/identities/applications"
	identities "github.com/steve-care-software/identities/domain"
	"github.com/steve-care-software/libs/cryptography/hash"
	"go.dedis.ch/kyber/v3/group/edwards25519"
)

const identityList = "identity:list"
const identityListDeleted = "identity:list:deleted"

var curve = edwards25519.NewBlakeSHA256Ed25519()

// NewIdentityApplication creates a new identity application
func NewIdentityApplication(
	repositoryBuilder identities.RepositoryBuilder,
	serviceBuilder identities.ServiceBuilder,
	database database_application.Application,
	dbName string,
	kind uint,
	nameKind uint,
) application_identity.Application {
	return createIdentityApplication(
		database,
		repositoryBuilder,
		serviceBuilder,
		dbName,
		kind,
		nameKind,
	)
}

// NewIdentityServiceBuilder creates a new identity service builder
func NewIdentityServiceBuilder(
	database database_application.Application,
	repositoryBuilder identities.RepositoryBuilder,
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
