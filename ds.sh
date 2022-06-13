#!/bin/bash

docker-compose up -d vol_mysql
sleep 120
docker-compose up -d vol_golang 
