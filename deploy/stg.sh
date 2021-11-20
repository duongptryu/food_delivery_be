#!/usr/bin/env bash

APP_NAME=food-delivery

docker load -i ${APP_NAME}.tar
docker rm -f ${APP_NAME}

docker run -d --name ${APP_NAME} \
  --net my-net \
  -e VIRTUAL_HOST="send2kindle.me" \
  -e LETSENCRYPT_HOST="send2kindle.me" \
  -e LETSENCRYPT_EMAIL="duongptryu@gmail.com" \
  -e DBConnectionStr="root:root@tcp(localhost:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local" \
  -e S3BucketName="..." \
  -e Region="..." \
  -e S3API="..." \
  -e S3SeretKey="..." \
  -e S3Domain="..." \
  -e SYSTEM_SECRET="..." \
  -p 8080:8080 \
  ${APP_NAME}


docker run -d --name food-delivery \
  --net my-net \
  -e VIRTUAL_HOST="send2kindle.me" \
  -e LETSENCRYPT_HOST="send2kindle.me" \
  -e LETSENCRYPT_EMAIL="duongptryu@gmail.com" \
  -e DBConnectionStr="root:root@tcp(localhost:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local" \
  -p 8080:8080 \
  food-delivery