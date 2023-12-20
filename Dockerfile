#build stage
# TODO: use tagged version 
FROM golang:alpine AS builder
WORKDIR /src

COPY . .
RUN go get -d -v ./... && mkdir -p /go/bin
RUN go build -o /go/bin/htmxpoc -v ./main.go

#final stage
# TODO: use tagged version 
FROM alpine:latest 

WORKDIR /app

RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/htmxpoc /app
COPY --from=builder /src/config.yaml /app
COPY --from=builder /src/internal/db/migrations /app/internal/db/migrations
COPY --from=builder /src/ui/static /app/ui/static

LABEL Name=gopoc Version=0.0.1
EXPOSE 8080
ENTRYPOINT ["./htmxpoc"]
