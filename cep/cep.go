package main

import (
	"errors"
	"log"
	"regexp"
)

func CepValidate(s string) error {
	cepvalidation := regexp.MustCompile(`^\d{5}[-\s]?\d{3}$`)

	if !cepvalidation.MatchString(s) {
		return errors.New("Cep invalido")
	}

	return nil
}

func main() {
	s := "65000 0000"

	err := CepValidate(s)
	if err != nil {
		log.Println("Padrão de CEP Invalido")
	}
}
