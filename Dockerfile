# Build stage
FROM golang:1.22.5-alpine3.20 as builder
WORKDIR /app
COPY . .

RUN go build -o main main.go

# Run stage
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY db/migration ./db/migration

RUN chmod 0770 /app/start.sh

EXPOSE 8080 9090
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]