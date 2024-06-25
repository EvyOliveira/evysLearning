FROM golang:1.22.2 AS BUILDER

RUN apt-get update && apt-get install -y busybox

WORKDIR /evys-learning

COPY . .

RUN busybox sh -c 'apt-get update && apt-get install -y libpq-dev'

RUN go build -o evys-learning .

WORKDIR /evys-learning

FROM scratch 
COPY --from=builder ./evys-learning /evys-learning
CMD ["./evys-learning"]

ENTRYPOINT ["/main"]

ENV PGSSLMODE=disable
ENV PGDATA=/var/lib/postgresql/data