# LavandeApp

## Pre-Requisite
1. Install golang, at least go1.17
2. Install docker or docker-supported containter tools

## Installation
1. Copy `env.sample` into `.env`
2. Initiate go mod vendor, run `go mod tidy && go mod vendor -v`

## How To Run Backend Service
1. run `make docker-start`
2. run `make run-backend`
