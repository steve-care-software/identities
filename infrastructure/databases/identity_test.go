package databases

import (
	"os"
	"testing"

	"github.com/steve-care-software/databases/infrastructure/files"
	identities "github.com/steve-care-software/identities/domain"
)

const kindForTests = 1
const kindNameForTests = 2

func TestIdentity_repositoryAndService_Success(t *testing.T) {
	dirPath := "./test_files"
	dstExtension := "destination"
	bckExtension := "backup"
	readChunkSize := uint(1000000)
	defer func() {
		os.RemoveAll(dirPath)
	}()

	fileName := "my_file.db"
	dbApp := files.NewApplication(dirPath, dstExtension, bckExtension, readChunkSize)
	err := dbApp.New(fileName)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pContext, err := dbApp.Open(fileName)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	defer dbApp.Close(*pContext)

	// create repository:
	repositoryBuilder := NewIdentityRepositoryBuilder(dbApp)
	repository, err := repositoryBuilder.Create().WithContext(*pContext).WithKind(kindForTests).WithNameKind(kindNameForTests).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// create service:
	service, err := NewIdentityServiceBuilder(dbApp, repositoryBuilder).Create().WithContext(*pContext).WithKind(kindForTests).WithNameKind(kindNameForTests).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	//creates the identity:
	name := "Roger"
	idIns := identities.NewIdentityForTests(name)
	pass := []byte("12345")

	// insert the identity:
	err = service.Insert(idIns, pass)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// commit:
	err = dbApp.Commit(*pContext)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// retrieve the identity:
	retIdentity, err := repository.Retrieve(name, pass)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retIdentity.Name() != name {
		t.Errorf("the returned identity is invalid")
		return
	}

	// delete the identity:
	err = service.Delete(name, pass)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// commit:
	err = dbApp.Commit(*pContext)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// re-insert the identity:
	err = service.Insert(idIns, pass)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// commit:
	err = dbApp.Commit(*pContext)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pubKey := []byte("this is another pubkey update")
	pk := []byte("this is another privatekey update")
	createdOn := idIns.CreatedOn()
	updatedIns, err := identities.NewBuilder().Create().WithName(name).WithPrivate(pubKey).WithPublic(pk).CreatedOn(createdOn).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// update the identity:
	newPass := []byte("45678")
	err = service.Update(name, updatedIns, pass, newPass)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// commit:
	err = dbApp.Commit(*pContext)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// retrieve the list:
	list, err := repository.List()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(list) != 1 {
		t.Errorf("%d identities were expected, %d returned", 1, len(list))
		return
	}

	if list[0] != name {
		t.Errorf("the element was expected to be '%s', '%s' returned", name, list[0])
		return
	}

}

func TestIdentity_application_Success(t *testing.T) {
	dirPath := "./test_files"
	dstExtension := "destination"
	bckExtension := "backup"
	readChunkSize := uint(1000000)
	defer func() {
		os.RemoveAll(dirPath)
	}()

	fileName := "my_file.db"
	dbApp := files.NewApplication(dirPath, dstExtension, bckExtension, readChunkSize)

	// create application:
	repositoryBuilder := NewIdentityRepositoryBuilder(dbApp)
	application := NewIdentityApplication(
		repositoryBuilder,
		NewIdentityServiceBuilder(dbApp, repositoryBuilder),
		dbApp,
		fileName,
		0,
		1,
	)

	//creates the identity:
	name := "Roger"
	idIns := identities.NewIdentityForTests(name)
	pass := []byte("12345")

	// insert the identity:
	err := application.Insert(idIns, pass)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// retrieve the identity:
	retIdentity, err := application.Retrieve(name, pass)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retIdentity.Name() != name {
		t.Errorf("the returned identity is invalid")
		return
	}

	// delete the identity:
	err = application.Delete(name, pass)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// re-insert the identity:
	err = application.Insert(idIns, pass)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pubKey := []byte("this is another pubkey update")
	pk := []byte("this is another privatekey update")
	createdOn := idIns.CreatedOn()
	updatedIns, err := identities.NewBuilder().Create().WithName(name).WithPrivate(pubKey).WithPublic(pk).CreatedOn(createdOn).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// update the identity:
	newPass := []byte("45678")
	err = application.Update(name, updatedIns, pass, newPass)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// retrieve the list:
	list, err := application.List()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(list) != 1 {
		t.Errorf("%d identities were expected, %d returned", 1, len(list))
		return
	}

	if list[0] != name {
		t.Errorf("the element was expected to be '%s', '%s' returned", name, list[0])
		return
	}

}
