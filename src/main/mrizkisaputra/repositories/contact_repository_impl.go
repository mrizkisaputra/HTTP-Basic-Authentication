package repositories

import (
	. "basic_authentication/src/main/mrizkisaputra/entity"
	"basic_authentication/src/main/mrizkisaputra/utils"
	"context"
	"database/sql"
)

type contactRepositoryImpl struct {
	DB *sql.DB
}

func NewContactRepository(db *sql.DB) ContactRepository {
	contactRepository := new(contactRepositoryImpl)
	contactRepository.DB = db
	return contactRepository
}

func (repository contactRepositoryImpl) GetContactById(ctx context.Context, id string) (Contact, error) {
	conn := repository.DB
	defer conn.Close()
	rows, err := conn.QueryContext(ctx, utils.QueryByIdContact, id)
	defer rows.Close()
	if err != nil {
		return Contact{}, err
	}
	var contact Contact = Contact{}
	if rows.Next() {
		err := rows.Scan(&contact.Id, &contact.Name, &contact.PhoneNumber, &contact.CreatedAt)
		if err != nil {
			return Contact{}, err
		}
	}
	return contact, nil
}

func (repository contactRepositoryImpl) GetContacts(ctx context.Context) ([]Contact, error) {
	conn := repository.DB
	defer conn.Close()
	rows, err := conn.QueryContext(ctx, utils.QueryContacts)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var contacts []Contact
	for rows.Next() {
		var contact = Contact{}
		err := rows.Scan(&contact.Id, &contact.Name, &contact.PhoneNumber, &contact.CreatedAt)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}
