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
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/htmxpoc /
COPY --from=builder /src/config.yaml /
COPY --from=builder /src/internal/db/migrations /internal/db/migrations

LABEL Name=gopoc Version=0.0.1
EXPOSE 8080
ENTRYPOINT ["./htmxpoc"]
