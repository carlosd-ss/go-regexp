package main

import (
	"errors"
	"log"
	"regexp"
)

func CpfValidate(s string) error {
	cpfvalidation := regexp.MustCompile(`^(?:\d{3}[.-]?){3}\d{2}$`)

	if !cpfvalidation.MatchString(s) {
		return errors.New("CPF invalido")
	}

	return nil
}

func main() {
	s := "00000000000"

	err := CpfValidate(s)
	if err != nil {
		log.Println("Padrão de CPF Invalido")
	}
}
