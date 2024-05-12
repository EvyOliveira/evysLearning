FROM golang:1.22.2 AS BUILDER

WORKDIR /evys-learning

COPY go.mod ./
COPY go.sum ./
COPY . .

RUN go build -o evys-learning .

FROM scratch 
COPY --from=builder ./evys-learning /evys-learning
CMD ["./evys-learning"]