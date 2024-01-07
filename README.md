## Go Rest DDD V2

### Run App With Docker
```sh
    git clone git@github.com:geekbim/Sagara.git
    cd Sagara
    docker-compose up
```

### Run App Without Docker
```sh
    git clone git@github.com:geekbim/Sagara.git
    cd Sagara
    cp .env.example .env
    go mod tidy
    sh run-service.sh
```

### Run Migration
```sh
    sh run-migration.sh
```
### Generate Mock
```sh
    mockery --all --case underscore
```

### Run Test
```sh
    sh run-test.sh
```

### Docs Swagger
```sh
    http://localhost:8080/docs/index.html
```

### Docs postman
```sh
    https://documenter.getpostman.com/view/1850032/Uz5CLJ9Q
```
