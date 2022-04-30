FROM golang:1.16

WORKDIR /app
COPY . .
COPY go.mod ./
RUN apt-get update


RUN go mod download
RUN go get -t
RUN go build
CMD [ "./service_customer" ]
