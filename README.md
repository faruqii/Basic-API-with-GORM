# Go Book API

### Requirement

- Docker

- Run docker compose
```bash
$ docker-compose up 
```

- Dependencies
```bash
go get github.com/gofiber/fiber/v2
go get github.com/google/uuid
go get gorm.io/driver/postgres
go get gorm.io/gorm
```

- Run
```bash
$ go run main.go
```
- Database Migration
```bash
$ go run migration.go
```
