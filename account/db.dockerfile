FROM postgres:10.3

COPY up.sql /docker-entrypoint-init-db.d/1.sql

CMD [ "postgres" ]