FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/service

FROM scratch
COPY --from=builder /app/service /service
EXPOSE 8080
ENTRYPOINT ["/service"]