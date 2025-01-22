# Migrations

## Pre-requisites
Migrations are done using [go-migrate](https://github.com/golang-migrate/migrate/) package.
Migrations are done in code in `core/infrastructure/postgres/posgres.go` file.

## How to run create migration
To create migration run:
```bash
migrate create -ext sql -dir ./documents/migrations -seq <migration_name>
```

## How to run migrations
In case if you want to run migrations manually, run:
```bash
migrate -source ./documents/migrations -database <postgresql connection string> up
```

## Potential issues
If you see error `Dirty database version 1. Fix and force version.` run:
```bash
migrate -source ./documents/migrations -database <postgresql connection string> force 1
```
Replace `1` with the number you see in the error message. And then run `migrate up` from the previous step again.

## Configurations
- Please ensure that the correct migration path is specified in the .env file of the services which interacts with the postgresql data.