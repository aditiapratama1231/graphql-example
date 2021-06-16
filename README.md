# Prerequisite

1. Install [gqlgen]("https://gqlgen.com/getting-started/") for generating graphql schema

2. Install [go-migrate]("https://github.com/golang-migrate/migrate") for running migration

# gqlgen

gqlgen is a golang library to generate .graphql into golang code, so we only create schema or domain with .graphql files.

if you want to add or update graph/schema/schema.graphqls, you have to run below command.

```
go run github.com/99designs/gqlgen generate
```

# Migration

Run below command to run migration

```
migrate -path migration -database "mysql://user:password@tcp(host:port)/dbname?query"
```

To create a new migration file

```
migrate create -ext sql -dir migration name
```

# How To Run

Run below command to run app server

```
go run cmd/product/main.go
```

To access GraphQL Playground

```
http://localhost:8080
```