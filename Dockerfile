# Build stage
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN sqlc generate

RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Install migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.2/migrate.linux-amd64.tar.gz | tar xvz
RUN mv migrate /usr/bin/migrate

# Final stage
FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/.env ./.env 
COPY --from=builder /app/app .
COPY --from=builder /app/db/migrations ./db/migrations
COPY --from=builder /usr/bin/migrate /usr/bin/migrate
COPY entrypoint.sh .

# Install PostgreSQL client tools
RUN apt-get update && apt-get install -y postgresql-client && rm -rf /var/lib/apt/lists/*

RUN chmod +x entrypoint.sh

# Use environment variables for configuration
ENV DATABASE_URL=postgresql://root:secret@localhost:5432/quiz?sslmode=disable

ENTRYPOINT ["./entrypoint.sh"]
