GOPATH ?= $(shell go env GOPATH)
IMG_PREFIX ?= caas4/{{cookiecutter.project_slug}}
TAG ?= latest
REPO_URL ?= $(shell git remote get-url origin)
BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)
COMMIT_REF ?= $(shell git rev-parse --verify HEAD)
LOW_VERSION_GIT_REPO_URL ?= $(shell git ls-remote --get-url origin)

apps = '{{cookiecutter.app_slug}}'

{{cookiecutter.project_slug.upper()}}_DB_NAME ?= "/tmp/{{cookiecutter.project_slug}}.sqlite"
{{cookiecutter.project_slug.upper()}}_DB_HOST ?= ""
{{cookiecutter.project_slug.upper()}}_DB_USERNAME ?= ""
{{cookiecutter.project_slug.upper()}}_DB_PASSWORD ?= ""
{{cookiecutter.project_slug.upper()}}_DB_TYPE ?= "sqlite3"
{{cookiecutter.project_slug.upper()}}_DB_PORT ?= ""

VERSION ?= $(shell git rev-parse --short HEAD)-$(shell date -u '+%Y%m%d%I%M%S')
BUILDX ?= false

.PHONY: build
build:
	gofmt -s -w test pkg cmd tools internal
	for app in $(apps) ;\
	do \
		CGO_ENABLED=1 go build -o dist/$$app -a -ldflags "-w -s -X {{cookiecutter.project_slug}}/pkg/server/version.Version=${VERSION}" ./cmd/$$app;\
	done

.PHONY: swagger
swagger:
	go run tools/doc-gen/main.go --output=swagger-ui/swagger.json

.PHONY: swagger-server
swagger-server:
	go run swagger-ui/swagger.go

.PHONY: fmt
fmt:
	gofmt -s -w pkg cmd tools internal test

.PHONY: image
image:
ifeq ($(BUILDX), false)
	for app in $(apps) ;\
	do \
		docker build -f build/$$app/Dockerfile --build-arg REPO_URL=$(REPO_URL) --build-arg BRANCH=$(BRANCH) --build-arg COMMIT_REF=$(COMMIT_REF) --build-arg VERSION=$(VERSION) -t ${IMG_PREFIX}-$$app:${TAG} . ;\
	done
else
	for app in $(apps) ;\
	do \
		docker buildx build -f build/$$app/Dockerfile --build-arg REPO_URL=$(REPO_URL) --build-arg BRANCH=$(BRANCH) --build-arg COMMIT_REF=$(COMMIT_REF) --build-arg VERSION=$(VERSION) --platform linux/amd64,linux/arm64 --push -t ${IMG_PREFIX}-$$app:${TAG} . ;\
	done
endif

.PHONY: test
test:
	@if [ $({{cookiecutter.project_slug.upper()}}_DB_TYPE) = sqlite3 ]; \
	then \
		rm -rf $({{cookiecutter.project_slug.upper()}}_DB_NAME); \
		CGO_ENABLED=1 go test ./test/controller -args db-name=$({{cookiecutter.project_slug.upper()}}_DB_NAME); \
	else \
		CGO_ENABLED=1 go test ./test/controller -args db-name=$({{cookiecutter.project_slug.upper()}}_DB_NAME) db-host=$({{cookiecutter.project_slug.upper()}}_DB_HOST) db-password=$({{cookiecutter.project_slug.upper()}}_DB_PASSWORD) db-username=$({{cookiecutter.project_slug.upper()}}_DB_USERNAME) db-type=$({{cookiecutter.project_slug.upper()}}_DB_TYPE) db-port=$({{cookiecutter.project_slug.upper()}}_DB_PORT); \
	fi

.PHONY: full-test
full-test:
	rm -rf $({{cookiecutter.project_slug.upper()}}_DB_NAME)
	CGO_ENABLED=1 go test ./test/controller -args db-name=$({{cookiecutter.project_slug.upper()}}_DB_NAME)
	CGO_ENABLED=1 go test ./test/conditional -args db-name=$({{cookiecutter.project_slug.upper()}}_DB_NAME) ssh-host=172.16.30.34
