FROM golang:1.24.3 as builder

WORKDIR /app

COPY . .
COPY go.mod go.sum ./

RUN CGO_ENABLED=0 GOOS=linux go build -o gprc-sample ./cmd/gprc-sample


FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/gprc-sample .

EXPOSE 8080
CMD ["/app/gprc-sample"]