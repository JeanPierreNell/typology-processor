FROM golang:1.20

WORKDIR /app
ADD . /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-go-tp
CMD ["/docker-go-tp"]