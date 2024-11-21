FROM golang:1.23.3

WORKDIR /go/src/app
COPY src/map.go .

RUN go build map.go
RUN chmod +x map && mv map /go/bin/app

CMD ["app"]
