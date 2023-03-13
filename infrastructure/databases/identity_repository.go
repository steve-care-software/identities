package databases

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"errors"

	database_application "github.com/steve-care-software/databases/applications"
	identities "github.com/steve-care-software/identities/domain"
	"github.com/steve-care-software/identities/infrastructure/objects"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type identityRepository struct {
	hashAdapter hash.Adapter
	database    database_application.Application
	builder     identities.Builder
	context     uint
	kind        uint
	nameKind    uint
}

func createIdentityRepository(
	hashAdapter hash.Adapter,
	database database_application.Application,
	builder identities.Builder,
	context uint,
	kind uint,
	nameKind uint,
) identities.Repository {
	out := identityRepository{
		hashAdapter: hashAdapter,
		database:    database,
		builder:     builder,
		context:     context,
		kind:        kind,
		nameKind:    nameKind,
	}

	return &out
}

// List returns the list of identity names
func (app *identityRepository) List() ([]string, error) {
	contentKeys, err := app.database.ContentKeysByKind(app.context, app.nameKind)
	if err != nil {
		return nil, err
	}

	hashes := []hash.Hash{}
	contentKeysList := contentKeys.List()
	for _, oneContentKey := range contentKeysList {
		hashes = append(hashes, oneContentKey.Hash())
	}

	values, err := app.database.ReadAllByHashes(app.context, app.nameKind, hashes)
	if err != nil {
		return nil, err
	}

	output := []string{}
	for _, oneValue := range values {
		output = append(output, string(oneValue))
	}

	return output, nil
}

// Retrieve retrieves an identity by name and password
func (app *identityRepository) Retrieve(name string, password []byte) (identities.Identity, error) {
	pHash, err := app.hashAdapter.FromBytes([]byte(name))
	if err != nil {
		return nil, err
	}

	cipher, err := app.database.ReadByHash(app.context, app.kind, *pHash)
	if err != nil {
		return nil, err
	}

	js, err := app.decrypt(password, cipher)
	if err != nil {
		return nil, err
	}

	ins := new(objects.Identity)
	err = json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithName(ins.Name).
		WithPrivate(ins.Private).
		WithPublic(ins.Public).
		CreatedOn(ins.CreatedOn).
		Now()
}

func (app *identityRepository) decrypt(password []byte, cipherBytes []byte) ([]byte, error) {
	hasher := curve.Hash()
	hasher.Write(password)
	hashedPass := hasher.Sum(nil)
	block, blockErr := aes.NewCipher(hashedPass)
	if blockErr != nil {
		return nil, blockErr
	}

	if len(cipherBytes) < aes.BlockSize {
		return nil, errors.New("the encrypted text cannot be decoded using this password: ciphertext block size is too short")
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	iv := cipherBytes[:aes.BlockSize]
	cipherBytes = cipherBytes[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherBytes, cipherBytes)

	// returns the decoded message:
	return cipherBytes, nil
}
