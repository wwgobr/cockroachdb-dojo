package main

import "log"
import "github.com/wwgobr/cockroachdb-dojo/parte1/conectar"

const nome = "Augusto"

func main() {
	db := conectar.Conectar()

	_, err := db.Exec("DELETE FROM bichinhos WHERE nome=$1", nome)

	if err != nil {
		log.Fatal(err)
	}

}
