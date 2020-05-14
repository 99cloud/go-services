GOPATH ?= $(shell go env GOPATH)
IMG ?= caas4/app_lower_46ea591951824d8e9376b0f98fe4d48a
TAG ?= latest
REPO_URL ?= $(shell git remote get-url origin)
BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)
COMMIT_REF ?= $(shell git rev-parse --verify HEAD)
LOW_VERSION_GIT_REPO_URL ?= $(shell git ls-remote --get-url origin)

apps = 'APP_46ea591951824d8e9376b0f98fe4d48a'

APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_NAME ?= "/tmp/APP_46ea591951824d8e9376b0f98fe4d48adb"
APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_HOST ?= ""
APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_USERNAME ?= ""
APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_PASSWORD ?= ""
APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_TYPE ?= "sqlite3"
APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_PORT ?= ""

VERSION ?= $(shell git rev-parse --short HEAD)-$(shell date -u '+%Y%m%d%I%M%S')

.PHONY: build
build:
	gofmt -s -w test pkg cmd tools internal
	for app in $(apps) ;\
	do \
		CGO_ENABLED=1 go build -o dist/$$app -a -ldflags "-w -s -X PROJECT_46ea591951824d8e9376b0f98fe4d48a/pkg/server/version.Version=${VERSION}" ./cmd/$$app;\
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
	docker build --build-arg REPO_URL=$(REPO_URL) --build-arg BRANCH=$(BRANCH) --build-arg COMMIT_REF=$(COMMIT_REF) --build-arg VERSION=$(VERSION) -t ${IMG}:${TAG} .

#issue: https://github.com/gliderlabs/docker-alpine/issues/307
.PHONY: jenkins-cicd-image
jenkins-cicd-image:
	docker build --network host --build-arg REPO_URL=$(LOW_VERSION_GIT_REPO_URL) --build-arg BRANCH=$(BRANCH) --build-arg COMMIT_REF=$(COMMIT_REF) --build-arg VERSION=$(VERSION) -t ${IMG}:${TAG} .

.PHONY: push-to-dockerhub
push-to-dockerhub:
	docker push ${IMG}:${TAG}

.PHONY: test
test:
	@if [ $(APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_TYPE) = sqlite3 ]; \
	then \
		rm -rf $(APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_NAME); \
		CGO_ENABLED=1 go test ./test/controller -args db-name=$(APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_NAME); \
	else \
		CGO_ENABLED=1 go test ./test/controller -args db-name=$(APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_NAME) db-host=$(APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_HOST) db-password=$(APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_PASSWORD) db-username=$(APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_USERNAME) db-type=$(APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_TYPE) db-port=$(APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_PORT); \
	fi

.PHONY: full-test
full-test:
	rm -rf $(APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_NAME)
	CGO_ENABLED=1 go test ./test/controller -args db-name=$(APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_NAME)
	CGO_ENABLED=1 go test ./test/conditional -args db-name=$(APP_UPPER_46ea591951824d8e9376b0f98fe4d48a_DB_NAME) ssh-host=172.16.30.34
