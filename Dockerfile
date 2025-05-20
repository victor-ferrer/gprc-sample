FROM golang:1.24.3

WORKDIR /app

COPY . .
COPY go.mod go.sum ./

COPY infrastructure/db/migrations/ ./migrations/


RUN CGO_ENABLED=0 GOOS=linux go build -o gprc-sample ./cmd/gprc-sample


# Specify the default command
EXPOSE 8080
CMD ["/app/gprc-sample"]