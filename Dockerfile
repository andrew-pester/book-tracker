FROM golang:1.21

WORKDIR /app

COPY src/go.mod src/go.sum ./

RUN go mod download && go mod verify

EXPOSE 8080


COPY src ./
CMD [ "go", "run", "book-tracker.go" ]