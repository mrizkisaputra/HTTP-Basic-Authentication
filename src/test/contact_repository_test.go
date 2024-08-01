package test

import (
	"basic_authentication/src/config"
	"basic_authentication/src/main/mrizkisaputra/repositories"
	"context"
	"fmt"
	"testing"
)

var repository = repositories.NewContactRepository(config.GetConnectionDB())

func TestGetContacts(t *testing.T) {
	contacts, err := repository.GetContacts(context.Background())
	if err != nil {
		fmt.Errorf("terjadi error :: %s", err.Error())
	}
	fmt.Println(len(contacts))
}

func TestGetContactById(t *testing.T) {
	contact, err := repository.GetContactById(context.Background(), "001")
	if err != nil {
		fmt.Errorf("error do :: %w", err)
	}
	fmt.Println(contact)
}
