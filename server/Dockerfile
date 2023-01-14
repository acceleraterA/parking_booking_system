FROM golang:1.16-alpine as builder
ENV MONGO_DB_USERNAME=admin \
    MONGO_DB_PWD=password
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /book-parking
## Deploy
FROM gcr.io/distroless/static-debian11

WORKDIR /

COPY --from=builder /app/book-parking /book-parking

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/book-parking"]