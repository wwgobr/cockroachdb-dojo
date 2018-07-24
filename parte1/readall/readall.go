package main

import "fmt"
import "log"
import "github.com/wwgobr/cockroachdb-dojo/parte1/conectar"

const pesquisa = "Marcelinho"

func main() {
	db := conectar.Conectar()

	resultado, err := db.Query("SELECT * FROM bichinhos")
	if err != nil {
		log.Fatal(err)
	}
	defer resultado.Close()
	for resultado.Next() {
		var nome string
		var idade int
		var especie bool
		if err := resultado.Scan(&nome, &idade, &especie); err != nil {
			log.Fatal(err)
		}
		var especieNome string
		if especie {
			especieNome = "gatinho"
		} else {
			especieNome = "cachorrinho"
		}
		fmt.Printf("%v %v %v\n", nome, idade, especieNome)
	}

}
