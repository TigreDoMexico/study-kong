# Estudo Kong
Repo para estudo das funcionalidades do Kong.

## Sobre o Kong
[Kong](https://konghq.com/) é um API Gateway que permite unificar e gerenciar as rotas de diferentes APIs em um só endpoint, facilitando o acesso, segurança, manipulação de dados e controle de tráfego.

## Arquitetura do Estudo

Esta solução é composta por 3 APIs:

- [API Usuário 1 - em .NET](./api-usuario1)
- [API Usuário 2 - em Node](./api-usuario2)
- [API Usuário 3 - em Go](./api-usuario3)

Cada uma destas APIs rodam de forma independente no Docker e possuem rotas diferentes, simulando um ambiente de microsserviços.

A ideia é ter o Kong como Gateway dessas três APIs e estudar a manipulação de dados entre elas.

## Como rodar e testar

Basta rodar o docker compose e depois fazer as seguintes requisições (todas estão no arquivo routes.http):

- http://localhost:8000/api1/usuario (GET)
- http://localhost:8000/api2/api/usuario (GET)
- http://localhost:8000/api3/usuario (GET)