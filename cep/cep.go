package main

import (
	"log"
	"regexp"
)

func main() {
	cepvalidation := regexp.MustCompile(`^\d{5}[-\s]?\d{3}$`)

	var (
		example1 = "65000 000"
		example2 = "65000000"
		example3 = "65000-000"
	)

	if !cepvalidation.MatchString(example1) {
		log.Println("Exemplo 1 - Invalido")
	}

	if !cepvalidation.MatchString(example2) {
		log.Println("Exemplo 2 - Invalido")
	}

	if !cepvalidation.MatchString(example3) {
		log.Println("Exemplo 3 - Invalido")
	}
}
