package main

import (
	"errors"
	"log"
	"regexp"
)

func emailValidate(s string) error {
	emailvalidation := regexp.MustCompile(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)

	var err error

	emailerr := errors.New("Email invalido")

	if !emailvalidation.MatchString(s) {
		err = emailerr
	}

	return err
}

func main() {
	s := "gopher@gmail.com"

	err := emailValidate(s)
	if err != nil {
		log.Println("Padrão de Email Invalido")
	}
}
