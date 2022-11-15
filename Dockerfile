FROM golang:alpine AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app
COPY go.mod .
COPY go.sum .

COPY . .
RUN go build -o main .
WORKDIR /dist
RUN cp /app/main .
FROM scratch
COPY --from=builder /dist/main /
EXPOSE 8080

ENTRYPOINT ["/main"]