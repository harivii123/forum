## multi-stage build
FROM golang:1.27-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -tags sqlite_fts5 -o forum .

FROM alpine:3.22

LABEL org.opencontainers.image.title="Literary Lions Forum"
LABEL org.opencontainers.image.description="Forum for book worms who want to share their feelings and thoughts about their books they have read."
LABEL org.opencontainers.image.authors="Hari, Hang, Ville, Vivian"
LABEL org.opencontainers.image.version="1.0"

WORKDIR /app

COPY --from=builder /app/forum .

EXPOSE 8080

CMD ["./forum", "-seed"]

## docker container prune
## docker image prune
## docker builder prune
## docker system prune

## docker logs forum-app
## docker ps -a
## docker images