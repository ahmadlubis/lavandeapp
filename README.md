# LavandeApp

## Pre-Requisite
1. Install golang, at least go1.17
2. Install docker or docker-supported containter tools
3. Install ruby, at least 3.0.0
3. Install rails, at least 7.0.0

## Installation
1. Copy `env.sample` into `.env`
2. Initiate go mod vendor, run `go mod tidy && go mod vendor -v`

## How To Run Backend Service
1. run `make docker-start`
2. run `make run-backend`

## How To Run Frontend Service
1. Navigate to `lavande_fe`
2. Run `rails s`
