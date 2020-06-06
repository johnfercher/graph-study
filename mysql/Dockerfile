FROM mysql:8.0

COPY ./sql-scripts/create_tables.sql /docker-entrypoint-initdb.d/
COPY ./sql-scripts/insert_data.sql /docker-entrypoint-initdb.d/