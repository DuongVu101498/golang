FROM golang:1.14

WORKDIR /go/src/app
COPY . .
EXPOSE  8081

CMD ["app"]
