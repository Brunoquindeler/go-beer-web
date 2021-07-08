FROM golang:1.16.0-stretch

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
ENV CGO_ENABLED=0

RUN go get github.com/mattn/go-sqlite3

CMD ["tail", "-f", "/dev/null"]