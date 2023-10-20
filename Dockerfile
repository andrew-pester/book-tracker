ARG IMAGE=golang
ARG VERSION=1.21

FROM IMAGE:VERSION

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify


COPY . .