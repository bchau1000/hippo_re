FROM mysql
ENV MYSQL_ROOT_PASSWORD hippodbpassword
COPY recreate.sql /docker-entrypoint-initdb.d
EXPOSE 3306