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

## Docker
```bash
docker build -t backend .
docker run -e database_config__host='localhost' -e database_config__port='5432' -e database_config__username='postgres' -e database_config__password='1' -e database_config__database='test' -p 9000:9000 backend
```