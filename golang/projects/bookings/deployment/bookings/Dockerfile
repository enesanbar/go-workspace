#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN go get github.com/gobuffalo/pop/...
RUN go build -o /go/bin/app -v ./cmd/web/*.go

#final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/bin/app /app/bookings
COPY --from=builder /go/bin/soda /usr/local/bin/soda
COPY database.compose.yml /app/database.yml
COPY email-templates /app/email-templates
COPY migrations /app/migrations
COPY static /app/static
COPY templates /app/templates
ENTRYPOINT ["/app/bookings"]
LABEL Name=bookings Version=0.0.1
EXPOSE 8080
