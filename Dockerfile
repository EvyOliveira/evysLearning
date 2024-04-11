FROM golang:1.22.2

WORKDIR /evys-learning

COPY go.mod ./
COPY main.go ./

RUN go build -o /server

EXPOSE 8080

CMD ["/server"]