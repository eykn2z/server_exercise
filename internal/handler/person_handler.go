package handler

import (
	"encoding/json"
	"net/http"

	faker "github.com/eykn2z/go_faker"
	"github.com/labstack/echo"
)

func getPersons() []faker.Person {
	return faker.GetPersons(20)
}

//net/httpのhanlderはhandlerの中でメソッド分岐する
func GetPersons(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		persons := getPersons()
		json.NewEncoder(w).Encode(persons)
	}
}

//Echoはhandlerの書き方が違う
func GetEchoPersons(c echo.Context) error {
	persons := getPersons()
	data, err := json.Marshal(persons)
	if err != nil {
		return err
	}
	c.String(http.StatusOK, string(data))
	return nil
}
