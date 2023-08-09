# p-builder


## Description

❗️Currently only supports mysql and postgre


## Getting started

```
go get -u github.com/my-owner-projects1/p-builder
```

## Usage
driver: mysql, postgres

dir: database connection string

folder: generate model folder path
```
p-builder -driver postgres -dir postgresql://peter:123456@localhost:5432/tmpl?sslmode=disable -folder /path/to/model
```