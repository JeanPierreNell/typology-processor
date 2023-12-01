FROM golang:1.20

WORKDIR /app
ADD . /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-go-tp

ENV REDIS_DB=0
ENV REDIS_AUTH="exampleAuth"
ENV REDIS_SERVERS='[{"host":"127.0.0.1", "port":6379}, {"host":"127.0.0.1", "port":6380}]'
ENV REDIS_IS_CLUSTER=false

# ArangoDB
ENV DATABASE_NAME=Configuration
ENV DATABASE_URL=tcp://0.0.0.0:8529
ENV DATABASE_USER=root
ENV DATABASE_PASSWORD=''
ENV COLLECTION_NAME=typologyExpression
ENV DATABASE_CERT_PATH=/usr/local/share/ca-certificates/ca-certificates.crt
ENV CACHE_ENABLED=false
ENV CACHE_TTL=0

# NATS
ENV PRODUCER_STREAM=
ENV CONSUMER_STREAM=
ENV SERVER_URL=0.0.0.0:4222
ENV STARTUP_TYPE=nats
ENV CMS_PRODUCER=CMS

CMD ["/docker-go-tp"]