FROM golang:1.21

WORKDIR /app

COPY src/go.mod src/go.sum ./

RUN go mod download && go mod verify

EXPOSE 8080

ENV DB_HOSTS="book-tracker-database"
ENV DB_KEYSPACE="inventory"
ENV DB_NAME="books"
ENV DB_PORT=9042
ENV DB_USER="cassandra"
ENV DB_PASSWORD="cassandra"
ENV CLIENT_ID="book-tracker-client"
ENV BOOK_ADMIN_ROLE="book-admin"
ENV PROVIDER_URL="http://keycloak:8080/realms/myrealm"


COPY src ./
CMD [ "go", "run", "book-tracker.go" ]