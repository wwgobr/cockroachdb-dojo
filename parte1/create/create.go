package main

import "log"
import "github.com/ellenkorbes/cockroachdb-dojo/parte1/conectar"

const nome = "Marcelinho"
const idade = 5
const especie = true // true é gato, false é cachorro.

func main() {
	db := conectar.Conectar()

	_, err := db.Exec("INSERT INTO banco.bichinhos (nome, idade, especie) VALUES ($1, $2, $3)", nome, idade, especie)

	if err != nil {
		log.Fatal(err)
	}

}
