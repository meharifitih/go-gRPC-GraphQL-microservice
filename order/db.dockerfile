FROM postgres:16-alpine

COPY up.sql /docker-entrypoint-init-db.d/1.sql

CMD ["postgres"]