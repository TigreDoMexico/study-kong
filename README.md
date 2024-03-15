# Estudo Kong
Repo para estudo das funcionalidades do Kong

## Arquitetura Inicial
Temos 3 APIs que rodam em containers no Docker. A ideia é ter um Kong como Gateway dessas três APIs e estudar como que os dados são trafegados entre elas

## Como rodar e testar
Basta rodar o docker compose e depois fazer as seguintes requisições (todas estão no arquivo routes.http):

- http://localhost:8000/api1/usuario (GET)
- http://localhost:8000/api2/api/usuario (GET)
- http://localhost:8000/api3/usuario (GET)