package main

import (
	"log"
	"regexp"
)

func main() {

	emailvalidation := regexp.MustCompile(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)

	var (
		example1 = "gopher@gmail.com"
		example2 = "gopher-go@outlook.com"
		example3 = "gopher-br@hotmail.com.br"
	)

	if !emailvalidation.MatchString(example1) {
		log.Println("Exemplo 1 - Invalido")
	}

	if !emailvalidation.MatchString(example2) {
		log.Println("Exemplo 2 - Invalido")
	}

	if !emailvalidation.MatchString(example3) {
		log.Println("Exemplo 3 - Invalido")
	}
}
