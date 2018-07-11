# Parte 1

Primeiro devemos inicializar o banco. Para isso vamos seguir as instruções listadas em: https://www.cockroachlabs.com/docs/stable/build-a-go-app-with-cockroachdb.html


```
cockroach start \
--insecure \
--store=hello-1 \
--host=localhost
```

`cockroach user set usuario --insecure`

`cockroach sql --insecure -e 'CREATE DATABASE banco'`

`cockroach sql --insecure -e 'GRANT ALL ON DATABASE banco TO usuario'`

Antes de usar as operações CRUD, rode o init para criar uma tabela no banco de dados.

Não é roubar se quiser consultar documentação e exemplos, como:
- https://github.com/cockroachdb/examples-go
- https://www.cockroachlabs.com/docs/stable/
- https://www.cockroachlabs.com/docs/stable/build-a-go-app-with-cockroachdb.html