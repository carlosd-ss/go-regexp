package main

import (
	"log"
	"regexp"
)

func main() {
	cpfvalidation := regexp.MustCompile(`^(?:\d{3}[.-]?){3}\d{2}$`)

	var (
		example1 = "000.000.000.00"
		example2 = "00000000000"
		example3 = "000-000-000-00"
	)

	if !cpfvalidation.MatchString(example1) {
		log.Println("Exemplo 1 - Invalido")
	}

	if !cpfvalidation.MatchString(example2) {
		log.Println("Exemplo 2 - Invalido")
	}

	if !cpfvalidation.MatchString(example3) {
		log.Println("Exemplo 3 - Invalido")
	}
}
