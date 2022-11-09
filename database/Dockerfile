FROM mysql
ENV MYSQL_ROOT_PASSWORD hippomysqlpassword
COPY recreate.sql /docker-entrypoint-initdb.d
EXPOSE 3306