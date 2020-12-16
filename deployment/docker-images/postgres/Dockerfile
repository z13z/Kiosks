FROM postgres:latest

COPY init_tables.sql /docker-entrypoint-initdb.d/

ENV POSTGRES_DB=kiosks