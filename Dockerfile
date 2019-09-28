FROM ubuntu:19.04

ENV PGSQLVER 11
ENV DEBIAN_FRONTEND 'noninteractive'
RUN echo 'Europe/Moscow' > '/etc/timezone'

RUN apt-get -y update
RUN apt install -y gcc git wget
RUN apt install -y postgresql-$PGSQLVER

RUN wget https://dl.google.com/go/go1.13.linux-amd64.tar.gz
RUN tar -xvf go1.13.linux-amd64.tar.gz
RUN mv go /usr/local

ENV GOROOT /usr/local/go
ENV GOPATH /opt/go
ENV PATH $GOROOT/bin:$GOPATH/bin:/usr/local/go/bin:$PATH

WORKDIR $GOPATH/backend
COPY . .

USER postgres

RUN /etc/init.d/postgresql start &&\
    psql --echo-all --command "CREATE USER ksu WITH SUPERUSER PASSWORD 'pswd';" &&\
    createdb -O ksu backend_db &&\
    psql --dbname=backend_db --echo-all --command 'CREATE EXTENSION IF NOT EXISTS citext;' &&\
    psql backend_db -f $GOPATH/backend/internal/database/scripts/scheme.sql &&\
    /etc/init.d/postgresql stop

RUN echo "host all  all    0.0.0.0/0  md5" >> /etc/postgresql/$PGSQLVER/main/pg_hba.conf
RUN echo "listen_addresses='*'" >> /etc/postgresql/$PGSQLVER/main/postgresql.conf
RUN echo "synchronous_commit = off" >> /etc/postgresql/$PGSQLVER/main/postgresql.conf
RUN echo "shared_buffers = 128MB" >> /etc/postgresql/$PGSQLVER/main/postgresql.conf

USER root

CMD service postgresql start && go run $GOPATH/backend/cmd/main.go