#!/bin/zsh

echo "Migrate UP"

migrate -path=./cmd/migrate/migrations -database="postgresql://gsarbaj:!Genryh38312290966@localhost:5432/social?sslmode=disable" up