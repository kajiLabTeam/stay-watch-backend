FROM mysql:8.4

ADD ./my.cnf /etc/mysql/conf.d/my.cnf
ADD ./hosts /etc/hosts

COPY ./script/init.sh /docker-entrypoint-initdb.d/init.sh

RUN ln -sf  /usr/share/zoneinfo/Asia/Tokyo /etc/localtime &&\
  chmod 644 /etc/mysql/conf.d/my.cnf

EXPOSE 33060
EXPOSE 3306
EXPOSE 22

CMD [ "mysqld" ]

