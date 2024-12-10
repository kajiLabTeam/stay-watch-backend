FROM mysql:8.4

COPY ./my.cnf /etc/mysql/conf.d/my.cnf
COPY ./hosts /etc/hosts
COPY ./script/init.sh /init.sh
COPY ./docker_rsa.pub /root/.ssh/authorized_keys

RUN ln -sf  /usr/share/zoneinfo/Asia/Tokyo /etc/localtime &&\
  chmod 644 /etc/mysql/conf.d/my.cnf &&\
  microdnf update &&\
  microdnf install -y openssh-server &&\
  mkdir -p /run/sshd &&\
  ssh-keygen -A &&\
  chmod 744 /init.sh

EXPOSE 33060
EXPOSE 3306
EXPOSE 22

CMD [ "./init.sh" ]

