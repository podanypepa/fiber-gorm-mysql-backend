# fiber-gorm-mysql-backend
Golang backend with Fiber, Gorm and MySQL

Rushed in a few hours for a colleague.

## Run on local with MySQL connection config

If You have MySQL server somewhere MySQL in internet or on Your machine.

```bash
go build -o backend
DBHOST=localhost DBPORT=3306 DBUSER=root DBPASSWORD=password DBNAME=todo backend
```

## Run in Docker from docker-compose.yaml

Run locally together `backend` and `MySQL`.

```bash
docker-compose build
docker-compose up -d
```

## Test

```bash
curl http://localhost:8080
# {"Status":"ok","TS":"2022-04-20T09:49:07.665761043Z"}
```
