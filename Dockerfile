FROM golang:1.14

WORKDIR /go/src/app
COPY . .
EXPOSE  8081
RUN go run main.go
CMD ["curl","http://localhost:8081/hi"]
