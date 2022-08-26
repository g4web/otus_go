# golang-migrate

https://pkg.go.dev/github.com/golang-migrate/migrate/v4/cmd/migrate#section-readme

## Installation

```
$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
$ apt-get update
$ apt-get install -y migrate
```


## Using
```
migrate create -ext sql -dir migrations some_table
migrate -source file://../migrations -database postgres://user:password@host:5432/db_name?sslmode=disable up
migrate -source file://../migrations -database postgres://user:password@host:5432/db_name?sslmode=disable down
```