package domain

import (
	"bytes"
	"testing"
	"time"
)

func TestIdentity_Success(t *testing.T) {
	name := "My Name"
	pubKey := []byte("this is a pubKey")
	pk := []byte("this is a pk")
	createdOn := time.Now().UTC()
	ins, err := NewBuilder().Create().WithName(name).WithPublic(pubKey).WithPrivate(pk).CreatedOn(createdOn).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins.Name() != name {
		t.Errorf("the name was expected to be '%s', '%s' returned", name, ins.Name())
		return
	}

	if !ins.CreatedOn().Equal(createdOn) {
		t.Errorf("the creation time was expected to be '%s', '%s' returned", createdOn.String(), ins.CreatedOn().String())
		return
	}

	if bytes.Compare(pubKey, ins.Public()) != 0 {
		t.Errorf("the publicKey is invalid")
		return
	}

	if bytes.Compare(pk, ins.Private()) != 0 {
		t.Errorf("the pk is invalid")
		return
	}

}

func TestIdentity_withoutName_ReturnsError(t *testing.T) {
	pubKey := []byte("this is a pubKey")
	pk := []byte("this is a pk")
	createdOn := time.Now().UTC()
	_, err := NewBuilder().Create().WithPublic(pubKey).WithPrivate(pk).CreatedOn(createdOn).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

}

func TestIdentity_withEmptyPyubKey_ReturnsError(t *testing.T) {
	name := "My Name"
	pubKey := []byte{}
	pk := []byte("this is a pk")
	createdOn := time.Now().UTC()
	_, err := NewBuilder().Create().WithName(name).WithPublic(pubKey).WithPrivate(pk).CreatedOn(createdOn).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

}

func TestIdentity_withoutPubKey_ReturnsError(t *testing.T) {
	name := "My Name"
	pk := []byte("this is a pk")
	createdOn := time.Now().UTC()
	_, err := NewBuilder().Create().WithName(name).WithPrivate(pk).CreatedOn(createdOn).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

}

func TestIdentity_withEmptyPK_ReturnsError(t *testing.T) {
	name := "My Name"
	pubKey := []byte("this is a pubKey")
	pk := []byte{}
	createdOn := time.Now().UTC()
	_, err := NewBuilder().Create().WithName(name).WithPublic(pubKey).WithPrivate(pk).CreatedOn(createdOn).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

}

func TestIdentity_withoutPK_ReturnsError(t *testing.T) {
	name := "My Name"
	pubKey := []byte("this is a pubKey")
	createdOn := time.Now().UTC()
	_, err := NewBuilder().Create().WithName(name).WithPublic(pubKey).CreatedOn(createdOn).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

}

func TestIdentity_withoutCreationTime_ReturnsError(t *testing.T) {
	name := "My Name"
	pubKey := []byte("this is a pubKey")
	pk := []byte("this is a pk")
	_, err := NewBuilder().Create().WithName(name).WithPublic(pubKey).WithPrivate(pk).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

}
