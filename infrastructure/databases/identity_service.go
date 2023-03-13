package databases

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	database_application "github.com/steve-care-software/databases/applications"
	identities "github.com/steve-care-software/identities/domain"
	"github.com/steve-care-software/identities/infrastructure/objects"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type identityService struct {
	hashAdapter hash.Adapter
	database    database_application.Application
	repository  identities.Repository
	context     uint
	kind        uint
}

func createIdentityService(
	hashAdapter hash.Adapter,
	database database_application.Application,
	repository identities.Repository,
	context uint,
	kind uint,
) identities.Service {
	out := identityService{
		hashAdapter: hashAdapter,
		database:    database,
		repository:  repository,
		context:     context,
		kind:        kind,
	}

	return &out
}

// Insert inserts an identity
func (app *identityService) Insert(identity identities.Identity, password []byte) error {
	pHash, err := app.hashAdapter.FromBytes([]byte(identity.Name()))
	if err != nil {
		return err
	}

	ins := objects.Identity{
		Name:      identity.Name(),
		Private:   identity.Private(),
		Public:    identity.Public(),
		CreatedOn: identity.CreatedOn(),
	}

	js, err := json.Marshal(ins)
	if err != nil {
		return err
	}

	cipher, err := app.encrypt(password, js)
	if err != nil {
		return err
	}

	return app.database.Write(
		app.context,
		app.kind,
		*pHash,
		cipher,
	)
}

// Update updates an identity
func (app *identityService) Update(name string, updated identities.Identity, originalPassword []byte, newPassword []byte) error {
	retIdentity, err := app.repository.Retrieve(name, originalPassword)
	if err != nil {
		return err
	}

	if !updated.CreatedOn().Equal(retIdentity.CreatedOn()) {
		str := fmt.Sprintf("the identity's creation time (original: %s, updated: %s) was expected to NOT change", retIdentity.CreatedOn().String(), updated.CreatedOn().String())
		return errors.New(str)
	}

	defer app.Delete(retIdentity, originalPassword)
	return app.Insert(updated, newPassword)
}

// Delete deletes an identity
func (app *identityService) Delete(identity identities.Identity, password []byte) error {
	pHash, err := app.hashAdapter.FromBytes([]byte(identity.Name()))
	if err != nil {
		return err
	}

	list, err := app.repository.ListDeleted()
	if err != nil {
		return err
	}

	list = append(list, identity.Name())
	js, err := json.Marshal(list)
	if err != nil {
		return err
	}

	return app.database.Write(
		app.context,
		app.kind,
		*pHash,
		js,
	)
}

func (app *identityService) encrypt(password []byte, message []byte) ([]byte, error) {
	pHash, err := app.hashAdapter.FromBytes(password)
	if err != nil {
		return nil, err
	}

	block, blockErr := aes.NewCipher(*pHash)
	if blockErr != nil {
		return nil, blockErr
	}

	cipherBytes := make([]byte, aes.BlockSize+len(message))
	iv := cipherBytes[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherBytes[aes.BlockSize:], message)

	return cipherBytes, nil
}
