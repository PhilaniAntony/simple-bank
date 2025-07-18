# ---------- Build stage ----------
FROM golang:1.24.4-alpine3.22 AS builder

WORKDIR /app

# Install build tools and copy source code
RUN apk add --no-cache wget tar

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY app.env .

# Build the Go binary
RUN go build -o main main.go

# Download and extract golang-migrate binary
RUN wget https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz && \
  tar -xvzf migrate.linux-amd64.tar.gz && \
  chmod +x migrate

# ---------- Final stage ----------
FROM alpine:3.22

WORKDIR /app

# Copy necessary files from builder
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY --from=builder /app/start.sh .
COPY --from=builder /app/wait-for.sh .
COPY --from=builder /app/app.env .
COPY --from=builder /app/db/migration ./db/migration

# Make scripts executable
RUN chmod +x /app/start.sh /app/wait-for.sh

EXPOSE 8080

# Default entrypoint for api service
ENTRYPOINT ["/app/start.sh"]