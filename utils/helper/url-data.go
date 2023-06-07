package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type customerParam struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type Response struct {
	Data []customerParam `json:"data"`
}

func ValidateRequestCreateOrUpdateCustomer(firstName string, lastName string, email string) ([]customerParam, error) {
	url := "https://reqres.in/api/users?page=2"

	// Make the GET request
	response, err := http.Get(url)
	if err != nil {
		return []customerParam{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(response.Body)

	// Read the response body
	body, err := io.ReadAll(response.Body)

	var forbiddenCustomer Response
	err = json.Unmarshal(body, &forbiddenCustomer)
	if err != nil {
		return []customerParam{}, err
	}

	err = ValidateCreateCustomer(firstName, lastName, email, forbiddenCustomer.Data)
	return nil, err
}

func checkFirstNameLastNameEmail(firstName string, lastName string, email string, forbiddenCustomer customerParam) error {
	if firstName == forbiddenCustomer.FirstName && lastName == forbiddenCustomer.LastName && email == forbiddenCustomer.Email {
		return errors.New("customer cannot be added (already exists in other service)")
	}

	return nil
}

func ValidateCreateCustomer(firstName string, lastName string, email string, forbiddenCustomer []customerParam) error {
	for _, element := range forbiddenCustomer {
		err := checkFirstNameLastNameEmail(firstName, lastName, email, element)
		if err != nil {
			return err
		}
	}

	return nil
}
