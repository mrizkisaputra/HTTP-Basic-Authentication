package controllers

import (
	"basic_authentication/src/config"
	. "basic_authentication/src/main/mrizkisaputra/repositories"
	"context"
	"encoding/json"
	"net/http"
)

var repository = NewContactRepository(config.GetConnectionDB())

func GetContactHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, "Invalid request method", http.StatusMethodNotAllowed)
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if id := request.URL.Query().Get("id"); id != "" {
		contactById, err := repository.GetContactById(context.Background(), id)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		responseJSON(contactById, writer)
		return
	}

	contacts, err := repository.GetContacts(context.Background())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	responseJSON(contacts, writer)

}

func responseJSON(data any, writer http.ResponseWriter) {
	res, err := json.Marshal(data)
	if err != nil {
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(res)
}
