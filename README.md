# LavandeApp

## Pre-Requisite
1. Install golang, at least go1.17
2. Install docker or docker-supported containter tools
3. Install ruby 3.0.0 & rails (>=7.0.0)
    - Recommended: Install RVM https://rvm.io/rvm/install

## Installation
1. Copy `env.sample` into `.env`
2. Initiate go mod vendor, run `go mod tidy && go mod vendor -v`

## How To Run Backend Service
1. Run `make docker-start`
2. Run `make run-backend`

## How To Run Frontend Service
1. Navigate to `lavande_fe`
2. Run `bundle install`
2. Run `rails s`
