FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

ENV DB_HOST="cassandra"
ENV DB_PORT=9042
ENV DB_USER=cassandra
ENV DB_PASSWORD=cassandra
ENV DB_NAME=inventory
EXPOSE 8080


COPY . .
CMD [ "go", "run", "book-tracker.go" ]