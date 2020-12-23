package main

import (
	"log"
	"regexp"
)

func main() {
	telefonevalidation := regexp.MustCompile(`^(?:\+?55\s?)?(?:\(?\d{2}\)?[\s-]?)?\d{4,5}[-\s]?\d{4}$`)

	var (
		example1 = "+55 11 98888-8888"
		example2 = "11 98888-8888"
		example3 = "988888888"
	)

	if !telefonevalidation.MatchString(example1) {
		log.Println("Exemplo 1 - Invalido")
	}

	if !telefonevalidation.MatchString(example2) {
		log.Println("Exemplo 2 - Invalido")
	}

	if !telefonevalidation.MatchString(example3) {
		log.Println("Exemplo 3 - Invalido")
	}
}
