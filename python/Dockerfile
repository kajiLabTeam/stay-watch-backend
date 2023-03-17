FROM python:3.9.6
#タイムゾーンを日本に設定
ENV TZ=Asia/Tokyo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apt-get update && \
  apt-get install -y fonts-ipafont-gothic

WORKDIR /app
COPY ./app ./
RUN pip3 install --no-cache-dir -r requirements.txt
