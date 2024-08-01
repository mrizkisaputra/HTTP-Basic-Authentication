package repositories

import (
	. "basic_authentication/src/main/mrizkisaputra/entity"
	"context"
)

type ContactRepository interface {
	GetContacts(ctx context.Context) ([]Contact, error)
	GetContactById(ctx context.Context, id string) (Contact, error)
}
