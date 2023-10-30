FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

EXPOSE 8080


COPY . .
CMD [ "go", "run", "book-tracker.go" ]