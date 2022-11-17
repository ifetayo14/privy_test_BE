FROM golang:1.19-alpine
MAINTAINER github.com/ifetayo14

WORKDIR /app

#DOWNLOAD MODUlES
COPY . .
RUN go mod download && go mod verify

RUN go build -o /privy_cake_store

CMD ["/privy_cake_store"]
