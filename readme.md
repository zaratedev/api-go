# API de Registro y Login en Go

Esta es una API desarrollada en Go para permitir el registro de usuarios y autenticación mediante el login.

## Instalación

1. Asegúrate de tener Go instalado en tu sistema. Puedes descargarlo desde https://golang.org/
2. Clona este repositorio en tu máquina local:

```bash
git clone https://github.com/zaratedev/api-go.git

cd api-go

go mod tidy

```


## Uso

Para ejecutar la API, simplemente ejecuta el siguiente comando:

```
go run main.go
```


## Endpoints

Registro de usuarios
URL: `api/users`
Metodo: `POST`

Ejemplo:
```
curl --location 'http://127.0.0.1:8080/api/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "zaratedev",
    "email": "zaratedev@gmail.com",
    "phone": "7471499895",
    "password": "passworD1$"
}'
```

Login
URL: `api/login`
Metodo: `POST`

Ejemplo:
```
curl --location 'http://127.0.0.1:8080/api/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "zaratedev@gmail.com",
    "password": "passworD1$"
}'
```
