package handler

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"

	"github.com/aknfujii/faker"
)

func GetPersons(w http.ResponseWriter, r *http.Request) {
	persons := faker.GetPersons(20)
	json.NewEncoder(w).Encode(persons)
}

func GetEchoPersons(c echo.Context) error {
	persons := faker.GetPersons(20)
	data, err := json.Marshal(persons)
	if err != nil {
		return err
	}
	c.String(http.StatusOK, string(data))
	return nil
}
