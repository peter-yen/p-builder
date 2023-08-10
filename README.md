# p-builder


## Description

🚀 This is a tool that can help you generate database table and column into struct code

❗️Currently only supports mysql and postgre


## Getting started

```
go install github.com/peter-yen/p-builder
```

## Usage
driver: mysql, postgres (default: postgres)

dir: database connection string

folder: generate model folder path 
(default: ./model)
(✅ If the folder does not exist, it will help you create a folder)
```
p-builder -driver postgres -dir postgresql://peter:123456@localhost:5432/tmpl?sslmode=disable -folder /path/to/model
```