package satellite

import "context"

type DB interface {
	Users() Users
}

type Users interface {
	GetByCredentials(ctx context.Context, password, email string) (*User, error)
	Get(ctx context.Context, id uuid.UUID) (*User, error)
	Insert(ctx context.Context, user *User) error
	Disable(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, user *User) error
}

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string

	PasswordHash []byte
}
