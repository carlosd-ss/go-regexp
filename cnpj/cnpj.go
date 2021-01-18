package main

import (
	"errors"
	"log"
	"regexp"
)

func CnpjValidate(s string) error {
	cnpjvalidation := regexp.MustCompile(`^\d{2}[.-]?(?:\d{3}[.-]?){2}\/?\d{4}[.-]?\d{2}$`)

	if !cnpjvalidation.MatchString(s) {
		return errors.New("CNPJ invalido")
	}

	return nil
}

func main() {
	s := "00000000000000"

	err := CnpjValidate(s)
	if err != nil {
		log.Println("Padrão de CNPJ Invalido")
	}
}
