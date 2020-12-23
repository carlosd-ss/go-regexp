package main

import (
	"log"
	"regexp"
)

func main() {
	cnpjvalidation := regexp.MustCompile(`^\d{2}[.-]?(?:\d{3}[.-]?){2}\/?\d{4}[.-]?\d{2}$`)

	var (
		example1 = "00.000.000/0000-00"
		example2 = "00000000000000"
		example3 = "00-000-000-0000-00"
	)

	if !cnpjvalidation.MatchString(example1) {
		log.Println("Exemplo 1 - Invalido")
	}

	if !cnpjvalidation.MatchString(example2) {
		log.Println("Exemplo 2 - Invalido")
	}

	if !cnpjvalidation.MatchString(example3) {
		log.Println("Exemplo 3 - Invalido")
	}
}
