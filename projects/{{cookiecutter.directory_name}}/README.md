# {{cookiecutter.app_slug.upper()}} Services

## 1. Arch

## 2. Quick Start

### 2.1 Build binary

```console
$ make
gofmt -s -w pkg cmd tools internal
for app in '{{cookiecutter.app_slug}}' ;\
	do \
		CGO_ENABLED=1 go build -o dist/$app -a -ldflags "-w -s -X {{cookiecutter.project_slug}}/pkg/server/version.Version=2d6676f-20200312043644" ./cmd/$app;\
	done
```

### 2.2 Build image

```bash
$ make image
```

### 2.3 Launch

```bash
# Sqlite：/root/test.db can not exist
./{{cookiecutter.app_slug}} --db-name=/root/test.db

# 或者 Docker
docker run -d --name docker-{{cookiecutter.app_slug}} --rm -p 8080:8080 caas4/{{cookiecutter.app_slug}} /app/{{cookiecutter.app_slug}} --db-name=test.db

# PG
docker stop {{cookiecutter.app_slug}}-pg
docker run -d --rm --name {{cookiecutter.app_slug}}-pg -p 5432:5432 -e "POSTGRES_USER={{cookiecutter.app_slug}}test" -e "POSTGRES_DB={{cookiecutter.app_slug}}test" -e "POSTGRES_PASSWORD={{cookiecutter.app_slug}}test" caas4/postgres:12.2
./{{cookiecutter.app_slug}} --db-host="localhost" --db-name="{{cookiecutter.app_slug}}test" --db-password="{{cookiecutter.app_slug}}test" --db-port="5432" --db-type="postgres" --db-username="{{cookiecutter.app_slug}}test" --insecure-port=8080 --loglevel=debug

# MySQL
docker stop {{cookiecutter.app_slug}}-mysql
docker run -d --rm --name {{cookiecutter.app_slug}}-mysql -p:3306:3306 -e "MYSQL_ROOT_PASSWORD={{cookiecutter.app_slug}}test" mysql
docker exec -it {{cookiecutter.app_slug}}-mysql mysql -p
{{cookiecutter.app_slug}}test
create database {{cookiecutter.app_slug}}test;
exit
./{{cookiecutter.app_slug}} --db-host="localhost" --db-name="{{cookiecutter.app_slug}}test" --db-password="{{cookiecutter.app_slug}}test" --db-port="3306" --db-type="mysql" --db-username="root" --insecure-port=8080 --loglevel=debug
```

### 2.4 Smoke test

```bash
# SQLite
make test

# PG
docker stop {{cookiecutter.app_slug}}-pg
docker run -d --rm --name {{cookiecutter.app_slug}}-pg -p 5432:5432 -e "POSTGRES_USER={{cookiecutter.app_slug}}test" -e "POSTGRES_DB={{cookiecutter.app_slug}}test" -e "POSTGRES_PASSWORD={{cookiecutter.app_slug}}test" caas4/postgres:12.2
make test {{cookiecutter.app_slug.upper()}}_DB_NAME="{{cookiecutter.app_slug}}test" {{cookiecutter.app_slug.upper()}}_DB_HOST="localhost" {{cookiecutter.app_slug.upper()}}_DB_USERNAME="{{cookiecutter.app_slug}}test" {{cookiecutter.app_slug.upper()}}_DB_PASSWORD="{{cookiecutter.app_slug}}test" {{cookiecutter.app_slug.upper()}}_DB_TYPE="postgres" {{cookiecutter.app_slug.upper()}}_DB_PORT="5432"

# MySQL
# Attension !!! MySQL json data has a limitation 65535 bytes
docker stop {{cookiecutter.app_slug}}-mysql
docker run -d --rm --name {{cookiecutter.app_slug}}-mysql -p:3306:3306 -e "MYSQL_ROOT_PASSWORD={{cookiecutter.app_slug}}test" mysql
docker exec -it {{cookiecutter.app_slug}}-mysql mysql -p
{{cookiecutter.app_slug}}test
create database {{cookiecutter.app_slug}}test;
exit
make test {{cookiecutter.app_slug.upper()}}_DB_NAME="{{cookiecutter.app_slug}}test" {{cookiecutter.app_slug.upper()}}_DB_HOST="localhost" {{cookiecutter.app_slug.upper()}}_DB_USERNAME="root" {{cookiecutter.app_slug.upper()}}_DB_PASSWORD="{{cookiecutter.app_slug}}test" {{cookiecutter.app_slug.upper()}}_DB_TYPE="mysql" {{cookiecutter.app_slug.upper()}}_DB_PORT="3306"
```

