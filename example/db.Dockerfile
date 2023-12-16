#copying the migration file and exporting to the docker entrypont
FROM mysql:8.0.23
COPY ./database/*sql /docker-entrypoint-initdb.d/
