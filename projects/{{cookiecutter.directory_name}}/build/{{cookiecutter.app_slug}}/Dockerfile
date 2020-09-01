FROM golang:1.13-alpine3.10 AS builder
ARG VERSION
RUN apk update
RUN apk add gcc musl-dev \
    && mkdir /app
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=1 GOOS=linux go build -o {{cookiecutter.app_slug}} -a -ldflags "-w -s -X {{cookiecutter.project_slug}}/pkg/server/version.Version=${VERSION}" ./cmd/{{cookiecutter.app_slug}}

FROM alpine:3.10
ARG REPO_URL
ARG BRANCH
ARG COMMIT_REF
LABEL repo-url=${REPO_URL}
LABEL branch=${BRANCH}
LABEL commit-ref=${COMMIT_REF}
RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    && update-ca-certificates 2>/dev/null || true \
    && mkdir /app
WORKDIR /app
COPY --from=builder /app/{{cookiecutter.app_slug}} /app/{{cookiecutter.app_slug}}
EXPOSE 8080
VOLUME [ "/app/data" ]
