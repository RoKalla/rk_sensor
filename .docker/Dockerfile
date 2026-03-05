FROM golang:1.25.3-trixie AS builder

WORKDIR /src

# COPY go.mod go.sum ./
# RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main ./cmd/sensor

FROM gcr.io/distroless/static-debian13:non-root

COPY --from=builder /src/main .

ENTRYPOINT [ "main" ]