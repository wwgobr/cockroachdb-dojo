package conectar

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Vamos copiar esta função nos quatro arquivos que vamos criar.
func Conectar() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://usuario@localhost:26257/banco?sslmode=disable")
	if err != nil {
		log.Fatal("Erro conectando no banco: ", err)
	}
	return db
}
