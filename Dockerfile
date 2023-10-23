FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

ENV DB_HOSTS="cassandra"
ENV DB_KEYSPACE=inventory
ENV DB_NAME=books
ENV DB_PORT=9042
ENV DB_USER=cassandra
ENV DB_PASSWORD=cassandra
EXPOSE 8080


COPY . .
CMD [ "go", "run", "book-tracker.go" ]