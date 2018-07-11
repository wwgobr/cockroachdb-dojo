package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type bichinho struct {
	nome    string
	idade   int
	especie bool
}

func Conectar() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://usuario@localhost:26257/banco?sslmode=disable")
	if err != nil {
		log.Fatal("Erro conectando no banco: ", err)
	}
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS bichinhos (nome TEXT, idade INT, especie BOOL)"); err != nil {
		log.Fatal(err)
	}
	return db
}

func create(db *sql.DB, item bichinho) {
	_, err := db.Exec("INSERT INTO banco.bichinhos (nome, idade, especie) VALUES ($1, $2, $3)", item.nome, item.idade, item.especie)
	if err != nil {
		log.Fatal(err)
	}
}

func delete(db *sql.DB, nome string) {
	_, err := db.Exec("DELETE FROM bichinhos WHERE nome=$1", nome)
	if err != nil {
		log.Fatal(err)
	}
}

func read(db *sql.DB, nome string) {
	resultado, err := db.Query("SELECT * FROM bichinhos WHERE nome=$1", nome)
	if err != nil {
		log.Fatal(err)
	}
	defer resultado.Close()
	for resultado.Next() {
		var item bichinho
		if err := resultado.Scan(&item.nome, &item.idade, &item.especie); err != nil {
			log.Fatal(err)
		}
		var especieNome string
		if item.especie {
			especieNome = "gatinho"
		} else {
			especieNome = "cachorrinho"
		}
		fmt.Printf("%v %v %v\n", item.nome, item.idade, especieNome)
	}
}

func update(db *sql.DB, item bichinho) {
	_, err := db.Exec("UPDATE bichinhos SET idade=$1, especie=$2 WHERE nome=$3", item.idade, item.especie, item.nome)
	if err != nil {
		log.Fatal(err)
	}
}

func readall(db *sql.DB) {
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

func argsToBichinho(args []string) bichinho {
	idade, err := strconv.Atoi(args[3])
	if err != nil {
		panic("strconv error")
	}
	especie, err := strconv.ParseBool(args[4])
	if err != nil {
		panic("strconv error")
	}
	return bichinho{args[2], idade, especie}
}

func main() {
	db := Conectar()
	if len(os.Args) < 2 {
		fmt.Println("Favor utilizar create, read, update, delete, ou readall.")
		return
	}

	switch os.Args[1] {
	case "create":
		if len(os.Args) < 5 {
			fmt.Println("Favor utilizar o formato: comando create nome idade especie")
			return
		}
		create(db, argsToBichinho(os.Args))
	case "read":
		if len(os.Args) < 3 {
			fmt.Println("Favor utilizar o formato: comando read nome")
			return
		}
		read(db, os.Args[2])
	case "update":
		if len(os.Args) < 5 {
			fmt.Println("Favor utilizar o formato: comando update nome idade especie")
			return
		}
		update(db, argsToBichinho(os.Args))
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Favor utilizar o formato: comando delete nome")
			return
		}
		delete(db, os.Args[2])
	case "readall":
		readall(db)
	default:
		fmt.Println("Favor utilizar create, read, update, delete, ou readall.")
	}
}
