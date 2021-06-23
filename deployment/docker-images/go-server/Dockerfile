FROM golang:1.16

WORKDIR /go/src/app
COPY . .

RUN go get -d -v github.com/z13z/Kiosks
RUN go install -v github.com/z13z/Kiosks/master-server@latest

CMD ["master-server"]