package main

import (
	"errors"
	"log"
	"regexp"
)

func emailValidate(s string) error {
	emailvalidation := regexp.MustCompile(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)

	if !emailvalidation.MatchString(s) {
		return errors.New("Email invalido")
	}

	return nil
}

func main() {
	s := "gopher@gmail.com"

	err := emailValidate(s)
	if err != nil {
		log.Println("Padrão de Email Invalido")
	}
}
