#!/usr/bin/env bash
docker run --rm -it -p 15432:5432 \
    -v $(pwd)/init.sql:/docker-entrypoint-initdb.d/init.sql \
    -v $(pwd)/pg_hba.conf:/usr/share/postgresql/10/pg_hba.conf \
    -v $(pwd)/postgresql.conf:/etc/postgresql/postgresql.conf \
    debezium/postgres:10.0 -c 'config_file=/etc/postgresql/postgresql.conf'