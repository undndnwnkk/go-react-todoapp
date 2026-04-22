FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/todoapp ./cmd/go-todoapp

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
COPY --from=builder /bin/todoapp /bin/todoapp
CMD ["/bin/todoapp"]