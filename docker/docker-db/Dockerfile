FROM postgres

ENV POSTGRES_DB database
ENV POSTGRES_USER webstartuser
ENV POSTGRES_PASSWORD webstartpassword

COPY schema.sql /docker-entrypoint-initdb.d
EXPOSE 5432
