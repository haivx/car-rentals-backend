FROM golang:1.18-alpine as base

FROM base as builder

WORKDIR /app

COPY . .

RUN go build -o main main.go

RUN apk add curl

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz



FROM alpine

WORKDIR /app

COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY --from=builder /app/main .
COPY config.yaml .
COPY ./script/start.sh .
COPY ./script/wait-for.sh .
COPY db/migration ./migration

EXPOSE 8080

CMD [ "/app/main" ]

ENTRYPOINT [ "/app/start.sh" ]