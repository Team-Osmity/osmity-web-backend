FROM golang:1.24-alpine AS builder
WORKDIR /app

ARG APP_ENV
ARG BACK_VERSION
ARG BACK_BUILD_TIME
ARG BACK_COMMIT_SHA

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build \
  -ldflags "\
    -X 'osmity-web-backend/internal/buildinfo.AppEnvLd=${APP_ENV}' \
    -X 'osmity-web-backend/internal/buildinfo.VersionLd=${BACK_VERSION}' \
    -X 'osmity-web-backend/internal/buildinfo.BuildTimeLd=${BACK_BUILD_TIME}' \
    -X 'osmity-web-backend/internal/buildinfo.CommitSHALd=${BACK_COMMIT_SHA}'" \
  -o server ./cmd/server

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]