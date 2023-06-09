FROM golang:1.18.10-alpine3.17 as compiler

WORKDIR /app
COPY . .
RUN go mod download

COPY *.go ./

RUN go build -o /ssr-service

EXPOSE 8080

CMD [ "/ssr-service" ]
