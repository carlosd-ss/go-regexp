package main

import (
	"errors"
	"log"
	"regexp"
)

func phoneValidate(s string) error {
	phonevalidation := regexp.MustCompile(`^(?:\+?55\s?)?(?:\(?\d{2}\)?[\s-]?)?\d{4,5}[-\s]?\d{4}$`)

	if !phonevalidation.MatchString(s) {
		return errors.New("Cep invalido")
	}

	return nil
}

func main() {
	s := "11 98888-8888"

	err := phoneValidate(s)
	if err != nil {
		log.Println("Padrão de telefone invalido")
	}
}
