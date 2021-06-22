FROM postgres:13

ENV POSTGRES_PASSWORD ewg8z09dsd98zb
USER postgres

COPY ./*.sh /docker-entrypoint-initdb.d/
COPY ./init.sql /docker-entrypoint-initdb.d/initial
