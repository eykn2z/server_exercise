package handler

import (
	"encoding/json"
	"github.com/aknfujii/faker"
	"net/http"
)

func GetPersonsHandler(w http.ResponseWriter, r *http.Request){
	persons := faker.GetPersons(20)
	json.NewEncoder(w).Encode(persons)
}
