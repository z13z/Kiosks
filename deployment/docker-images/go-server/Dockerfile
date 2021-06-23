FROM ubuntu:latest

WORKDIR /go/src/app
COPY . .

RUN apt update
RUN apt install -y wget
RUN apt install -y tar
RUN apt install -y curl
RUN apt install -y sudo
RUN apt install -y build-essential
RUN apt install -y openssh-client
RUN apt install -y rsync
RUN wget -c https://dl.google.com/go/go1.16.1.linux-amd64.tar.gz -O - | tar -xz -C /usr/local
ENV PATH="/usr/local/go/bin:/root/go/bin:${PATH}"
RUN go get -d -v github.com/z13z/Kiosks@v1.0.0
RUN go install -v github.com/z13z/Kiosks/master-server@v1.0.0

CMD ["master-server"]