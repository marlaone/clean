#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE ROLE speedtest_user PASSWORD 'yd339WQgS7rskS9Q' SUPERUSER LOGIN;
    CREATE DATABASE speedtest OWNER speedtest_user;
EOSQL

psql -v ON_ERROR_STOP=1 --username "speedtest_user" --dbname "speedtest" < /docker-entrypoint-initdb.d/initial