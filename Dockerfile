FROM golang:1.15.7-buster

WORKDIR /build

COPY go.mod .
COPY go.sum .795
RUN go mod download

COPY . .

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

EXPOSE 8010
CMD ["/dist/main"]