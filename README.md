# Go Project Structure Example

[![Build Status](https://travis-ci.org/koungkub/go-structure.svg?branch=master)](https://travis-ci.org/koungkub/go-structure)

using [echo](https://echo.labstack.com/) for web framework

this project not following [golang-project-layout](https://github.com/golang-standards/project-layout)

## How to use this project

- edit env in `config/env.yml`

- edit import path

- run !!

## Run !!

**start normal**

`make start`

**start with docker**

`make docker`

## Folder Explain

`main.go` main program

`src/connection` for 3rd party connection

`src/controller` controller in `MVC` model

`src/middleware` custom middleware

`src/route` setup middleware and routing

`src/service` business logic

`src/utils` for utility
