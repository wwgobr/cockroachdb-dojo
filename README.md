# Dojo: CockroachDB

Olá!

Neste projeto vamos criar uma API para utilizar funções CRUD num banco de dados CockroachDB.

Este guia está organizado em partes. A idéia é você ler a descrição de cada parte e tentar implementar, e só olhar a resposta caso não consiga fazer sozinha.

- Parte 1: Operações CRUD
- Parte 2: Linha de comando
- Parte 3: Web API

Antes de começar é necessário instalar a linguagem Go e o CockroachDB em sua máquina. Siga os links:

- https://golang.org/dl/
- https://www.cockroachlabs.com/docs/stable/install-cockroachdb.html

## Parte 1: Operações CRUD

Para utilizar um banco de dados é necessário que possamos fazer quatro operações: criar entradas, lê-las, atualizá-las, e apagá-las. A isso dá-se o acrônimo CRUD, que vem do nome destas operações em inglês (create, read, update, delete).

O CockroachDB é um banco de dados que pode ser utilizado através de SQL (Sequel Query Language), que é uma linguagem "universal" de bancos de dados. Mais especificamente, para que nosso programa em Go possa interagir com o CockroachDB vamos precisar de um driver de Postgres. O mais utilizado é o `pq`.

Ou seja, nesta primeira parte nossa missão é:

- Descobrir como conectar um programa em Go com o CockroachDB através do driver `pq`. Siga a documentação em: https://godoc.org/github.com/lib/pq
- Descobrir como inicializar (criar o banco, criar uma tabela) e fazer as operações CRUD através de queries SQL. Para documentação sobre SQL, veja as entradas Create Database, Create Table, Insert Query, Select Query, Update Query, e Delete Query em: https://www.tutorialspoint.com/postgresql/postgresql_insert_query.htm
- Criar quatro programas separados, que podem ser `create.go`, `read.go`, `update.go` e `delete.go`, onde o primeiro cria uma entrada no banco, o segundo lê, etc.

O schema que vamos seguir neste exemplo vai ter três campos: nome (string), idade (int), gatinho ou cachorrinho (bool).

## Parte 2: Linha de comando

Não faz muito sentido um programa que sempre cria a mesma entrada no banco, certo? Então agora vamos tornar as coisas mais interativas.

Sua missão é aprender sobre argumentos via linha de comando, e utilizá-los para que um usuário possa fazer, por exemplo:

`$ meu-programa-em-go create marcelinho 5 gatinho`

`$ meu-programa-em-go delete marcelinho`

Lembrete #1: Não queremos mais quatro programas separados. Queremos um programa só que consiga escolher de maneira interativa que operação o usuário quer fazer. Utilize funções ou métodos para organizar seu código e não fazer uma bagunça!

Lembrete #2: Usuários são imprevisíveis. O que acontece alguem digitar letras no campo idade, por exemplo?

Esta página contém tudo o que você precisa saber sobre argumentos via linha de comando: https://gobyexample.com/command-line-arguments

## Parte 3: Web API

Agora queremos que nosso banco seja acessível através da web, utilizando uma API. 

Assim um usuário pode fazer operações no banco remotamente. Por exemplo, acessando as URLs:

`http://meu-app-em-go.com/create/marcelinho,5,gatinho`

`http://meu-app-em-go.com/delete/marcelinho`

Parece complicado, mas a maior parte do trabalho já está feita.

- Descubra como utilizar as funções NewServeMux, HandleFunc, e ListenAndServe do package http: https://golang.org/pkg/net/http/
- Re-escreva seu programa de forma que as diferentes funções CRUD sejam chamadas não através dos argumentos da linha de comando, mas através do ServeMux.

Para testar, use o browser!
