# Backend Boilerplate Golang

## Technical
- BUN
- GIN
- SWAGGER


## Building Services
To build a service you can specify the `main.go` file in the src folder and build as normal.

```bash
cd src/
go build -o ../../bin/backend main.go 
```

## Generate swagger
```bash
cd src
swag init --parseDependency --parseInternal
```