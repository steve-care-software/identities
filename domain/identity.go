package domain

import "time"

type identity struct {
	name      string
	private   []byte
	public    []byte
	createdOn time.Time
}

func createIdentity(
	name string,
	private []byte,
	public []byte,
	createdOn time.Time,
) Identity {
	out := identity{
		name:      name,
		private:   private,
		public:    public,
		createdOn: createdOn,
	}

	return &out
}

// Name returns the name
func (obj *identity) Name() string {
	return obj.name
}

// Private returns the private key
func (obj *identity) Private() []byte {
	return obj.private
}

// Public returns the public key
func (obj *identity) Public() []byte {
	return obj.public
}

// CreatedOn returns the creation time
func (obj *identity) CreatedOn() time.Time {
	return obj.createdOn
}
