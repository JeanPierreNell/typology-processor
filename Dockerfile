FROM golang:1.20

WORKDIR /app

# COPY go.mod go.sum db-lib proto structs ./

ADD . /app

RUN go mod download

# COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping
CMD ["/docker-gs-ping"]