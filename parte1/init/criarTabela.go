package main

import "log"
import "github.com/wwgobr/cockroachdb-dojo/parte1/conectar"
import "database/sql"

// Esta função cria uma tabela no banco de dados. Precisamos executá-la somente uma vez.
func criarTabela(db *sql.DB) {
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS bichinhos (nome TEXT, idade INT, especie BOOL)"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	db := conectar.Conectar()
	criarTabela(db)
}
