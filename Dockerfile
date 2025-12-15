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
    -X 'osmity-web-backend/internal/buildinfo.AppEnvLd=${APP_ENV}' \
    -X 'osmity-web-backend/internal/buildinfo.VersionLd=${APP_VERSION}' \
    -X 'osmity-web-backend/internal/buildinfo.BuildTimeLd=${BUILD_TIME}' \
    -X 'osmity-web-backend/internal/buildinfo.CommitSHALd=${GIT_COMMIT}'" \
  -o server ./cmd/server

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]
