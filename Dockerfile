FROM golang:1.18-alpine as base

FROM base as builder

WORKDIR /app

COPY . .

RUN go build -o main main.go


FROM alpine

WORKDIR /app

COPY --from=builder /app/main .
COPY config.yaml .

EXPOSE 8080

CMD [ "/app/main" ]