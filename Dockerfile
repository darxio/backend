FROM ubuntu:18.04


RUN apt-get -y update
RUN apt-get -y upgrade
RUN apt install -y gcc git wget

RUN wget https://dl.google.com/go/go1.13.linux-amd64.tar.gz
RUN tar -xvf go1.13.linux-amd64.tar.gz
RUN mv go /usr/local

ENV GOROOT /usr/local/go
ENV GOPATH /opt/go
ENV PATH $GOROOT/bin:$GOPATH/bin:/usr/local/go/bin:$PATH

WORKDIR $GOPATH/backend
COPY . .

EXPOSE 8888


RUN apt install -y redis-server

EXPOSE 6379


USER root

CMD systemctl enable redis-server.service && go run cmd/main.go