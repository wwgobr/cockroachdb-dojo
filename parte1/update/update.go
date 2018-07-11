package main

import "log"
import "github.com/ellenkorbes/cockroachdb-dojo/parte1/conectar"

const nome = "Marcelinho"
const idade = 6
const especie = false // true é gato, false é cachorro.

func main() {
	db := conectar.Conectar()

	_, err := db.Exec("UPDATE bichinhos SET idade=$1, especie=$2 WHERE nome=$3", idade, especie, nome)

	if err != nil {
		log.Fatal(err)
	}

}
