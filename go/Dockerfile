FROM golang:alpine

WORKDIR /app
EXPOSE 80
EXPOSE 443

COPY ./docker_rsa /root/.ssh/docker_rsa

RUN apk --update add tzdata&&\
  apk add --no-cache openssh &&\
  cp  /usr/share/zoneinfo/Asia/Tokyo /etc/localtime&&\
  apk del tzdata&&\
  rm -rf /var/cache/apk/*

CMD ["go","run","main.go"]



