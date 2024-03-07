# Bun migrations example

### Create migration tables:
```
go run main.go -c ../config/config.json db init
```

### To create a SQL migration:
```
go run main.go -c ../config/config.json db create_sql sql_migration_name
```

### To run migrations:
```
go run main.go -c ../config/config.json db migrate
```

### To rollback migrations:
```
go run main.go -c ../config/config.json db rollback
```

### To view status of migrations:
```
go run main.go -c ../config/config.json db status
```

### To create a Go migration:
```
go run main.go -c ../config/config.json db create_go go_migration_name
```