FROM ubuntu:18.04


ENV DEBIAN_FRONTEND 'noninteractive'

RUN apt-get -y update && apt-get install redis-server

EXPOSE 6379

RUN apt-get -y upgrade
# RUN add-apt-repository universe
RUN apt install -y gcc git wget

RUN wget https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz
RUN tar -xvf go1.12.7.linux-amd64.tar.gz
RUN mv go /usr/local

ENV GOROOT /usr/local/go
ENV GOPATH /opt/go
ENV PATH $GOROOT/bin:$GOPATH/bin:/usr/local/go/bin:$PATH

WORKDIR $GOPATH/backend
COPY . .

EXPOSE 8888


USER root

CMD systemctl enable redis-server.service && go run cmd/main.go