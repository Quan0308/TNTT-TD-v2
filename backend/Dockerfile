FROM golang:1.20 as builder

WORKDIR /app

COPY main-api/go.mod main-api/go.sum ./
RUN go mod download

COPY main-api/ .

RUN go build -o main-api ./cmd/main.go

FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=builder /app/main-api .

EXPOSE 8080

CMD ["./main-api"]
