FROM golang:1.24-alpine AS builder
WORKDIR /app

ARG APP_ENV
ARG APP_VERSION
ARG BUILD_TIME
ARG GIT_COMMIT

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build \
  -ldflags "\
    -X 'main.AppEnv=${APP_ENV}' \
    -X 'main.Version=${APP_VERSION}' \
    -X 'main.BuildTime=${BUILD_TIME}' \
    -X 'main.GitCommit=${GIT_COMMIT}'" \
  -o server .

FROM alpine:latest
WORKDIR /app

ENV APP_ENV=${APP_ENV}

COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]
