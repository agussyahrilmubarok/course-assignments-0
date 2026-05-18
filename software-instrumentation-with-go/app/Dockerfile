FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main cmd/course/main.go


FROM alpine:3.23
WORKDIR /app
COPY --from=builder /app/main course
COPY --from=builder /app/configs ./configs

EXPOSE 8080
CMD ["./course"]