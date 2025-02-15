#!/bin/bash

export $(grep -v '^#' env/api/.env | xargs)
go run ./api/main.go