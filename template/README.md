# APP_UPPER_46ea591951824d8e9376b0f98fe4d48a Services

## 1. Arch

## 2. Quick Start

### 2.1 Build binary

```console
$ make
gofmt -s -w pkg cmd tools internal
for app in 'APP_46ea591951824d8e9376b0f98fe4d48a' ;\
	do \
		CGO_ENABLED=1 go build -o dist/$app -a -ldflags "-w -s -X PROJECT_46ea591951824d8e9376b0f98fe4d48a/pkg/server/version.Version=2d6676f-20200312043644" ./cmd/$app;\
	done
```

### 2.2 Build image

```bash
$ make image
```

### 2.3 Launch

```bash
# Sqlite：/root/test.db can not exist
./APP_46ea591951824d8e9376b0f98fe4d48a --db-name=/root/test.db

# 或者 Docker
docker run -d --name docker-APP_46ea591951824d8e9376b0f98fe4d48a --rm -p 8080:8080 caas4/APP_46ea591951824d8e9376b0f98fe4d48a /app/APP_46ea591951824d8e9376b0f98fe4d48a --db-name=test.db

# PG
docker stop APP_46ea591951824d8e9376b0f98fe4d48a-pg
docker run -d --rm --name APP_46ea591951824d8e9376b0f98fe4d48a-pg -p 5432:5432 -e "POSTGRES_USER=APP_46ea591951824d8e9376b0f98fe4d48atest" -e "POSTGRES_DB=APP_46ea591951824d8e9376b0f98fe4d48atest" -e "POSTGRES_PASSWORD=APP_46ea591951824d8e9376b0f98fe4d48atest" caas4/postgres:12.2
./APP_46ea591951824d8e9376b0f98fe4d48a --db-host="localhost" --db-name="APP_46ea591951824d8e9376b0f98fe4d48atest" --db-password="APP_46ea591951824d8e9376b0f98fe4d48atest" --db-port="5432" --db-type="postgres" --db-username="APP_46ea591951824d8e9376b0f98fe4d48atest" --insecure-port=8080 --loglevel=debug

# MySQL
docker stop APP_46ea591951824d8e9376b0f98fe4d48a-mysql
docker run -d --rm --name APP_46ea591951824d8e9376b0f98fe4d48a-mysql -p:3306:3306 -e "MYSQL_ROOT_PASSWORD=APP_46ea591951824d8e9376b0f98fe4d48atest" mysql
docker exec -it APP_46ea591951824d8e9376b0f98fe4d48a-mysql mysql -p
APP_46ea591951824d8e9376b0f98fe4d48atest
create database APP_46ea591951824d8e9376b0f98fe4d48atest;
exit
./APP_46ea591951824d8e9376b0f98fe4d48a --db-host="localhost" --db-name="APP_46ea591951824d8e9376b0f98fe4d48atest" --db-password="APP_46ea591951824d8e9376b0f98fe4d48atest" --db-port="3306" --db-type="mysql" --db-username="root" --insecure-port=8080 --loglevel=debug
```

### 2.4 Smoke test

```bash
# SQLite
make test

# PG
docker stop APP_46ea591951824d8e9376b0f98fe4d48a-pg
docker run -d --rm --name APP_46ea591951824d8e9376b0f98fe4d48a-pg -p 5432:5432 -e "POSTGRES_USER=APP_46ea591951824d8e9376b0f98fe4d48atest" -e "POSTGRES_DB=APP_46ea591951824d8e9376b0f98fe4d48atest" -e "POSTGRES_PASSWORD=APP_46ea591951824d8e9376b0f98fe4d48atest" caas4/postgres:12.2
make test APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_NAME="APP_46ea591951824d8e9376b0f98fe4d48atest" APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_HOST="localhost" APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_USERNAME="APP_46ea591951824d8e9376b0f98fe4d48atest" APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_PASSWORD="APP_46ea591951824d8e9376b0f98fe4d48atest" APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_TYPE="postgres" APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_PORT="5432"

# MySQL
# Attension !!! MySQL json data has a limitation 65535 bytes
docker stop APP_46ea591951824d8e9376b0f98fe4d48a-mysql
docker run -d --rm --name APP_46ea591951824d8e9376b0f98fe4d48a-mysql -p:3306:3306 -e "MYSQL_ROOT_PASSWORD=APP_46ea591951824d8e9376b0f98fe4d48atest" mysql
docker exec -it APP_46ea591951824d8e9376b0f98fe4d48a-mysql mysql -p
APP_46ea591951824d8e9376b0f98fe4d48atest
create database APP_46ea591951824d8e9376b0f98fe4d48atest;
exit
make test APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_NAME="APP_46ea591951824d8e9376b0f98fe4d48atest" APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_HOST="localhost" APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_USERNAME="root" APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_PASSWORD="APP_46ea591951824d8e9376b0f98fe4d48atest" APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_TYPE="mysql" APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_PORT="3306"
```

