# ebanx-bank

This project is a test assigned by ebanx technical team.


## How to run it
Depedencies: docker/docker-compose

- [Introduction](#introduction)
- [Challenge](#challenge)
- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
## Introduction

Banking like project to showcase knowledge.
The project consists in a API with 3 routes.

/reset
/event
/balance?account_id=?

## Features

## Challenge: 

--
## Reset state before starting tests

POST /reset

200 OK


--
## Get balance for non-existing account

GET /balance?account_id=1234

404 0


--
## Create account with initial balance

POST /event {"type":"deposit", "destination":"100", "amount":10}

201 {"destination": {"id":"100", "balance":10}}


--
## Deposit into existing account

POST /event {"type":"deposit", "destination":"100", "amount":10}

201 {"destination": {"id":"100", "balance":20}}


--
## Get balance for existing account

GET /balance?account_id=100

200 20

--
## Withdraw from non-existing account

POST /event {"type":"withdraw", "origin":"200", "amount":10}

404 0

--
## Withdraw from existing account

POST /event {"type":"withdraw", "origin":"100", "amount":5}

201 {"origin": {"id":"100", "balance":15}}

--
## Transfer from existing account

POST /event {"type":"transfer", "origin":"100", "amount":15, "destination":"300"}

201 {"origin": {"id":"100", "balance":0}, "destination": {"id":"300", "balance":15}}

--
## Transfer from non-existing account

POST /event {"type":"transfer", "origin":"200", "amount":15, "destination":"300"}

404 0

## Instalation

Install docker
run docker-compose up --build
url = localhost:8080/

## Usage

## Event Route Request Examples
POST /event
{
    "type":"deposit",
    "destination":"1",
    "amount": 100
}

{
    "type":"withdraw",
    "origin":"1",
    "amount": 100
}

{
    "type":"transfer",
    "destination":"1",
    "origin":"2",
    "amount": 15
}

## Balance Route Request Example

GET /balance?account_id=1 
{
    "type":"deposit",
    "destination":"1",
    "amount": 100
}

## Reset Route Usage
POST /reset

## Testing

go test -v -c ./...