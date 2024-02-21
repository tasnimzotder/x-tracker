#!/bin/bash

FLAG=$1

# check if .env.local file exist and add env variables
if [ -f .env.local ]; then
    while IFS= read -r line; do
        export "$line"
    done < .env.local
fi

if [ -f backend/app.env ]; then
    while IFS= read -r line; do
        export "$line"
    done < backend/app.env
fi

# set env variables
export ENVIRONMENT=staging

# login to aws
#aws ecr get-login-password --region $AWS_REGION | \
#   docker login --username AWS --password-stdin $ECR_REGISTRY

if  [ $FLAG = "up" ]; then
  # run docker compose with env variables
  AWS_REGION=$AWS_REGION \
    ECR_REGISTRY=$ECR_REGISTRY \
    ECR_REPOSITORY_BE=$ECR_REPOSITORY_BE \
    ECR_REPOSITORY_FE=$ECR_REPOSITORY_FE \
    IMAGE_TAG=$IMAGE_TAG \
    DB_VOLUME_HOST=$DB_VOLUME_HOST \
    DB_PASSWORD=$DB_PASSWORD \
    XT_DB_SOURCE=$XT_DB_SOURCE \
    GIN_MODE=$GIN_MODE \
    SERVER_ADDRESS=$SERVER_ADDRESS \
    NEXT_PUBLIC_API_URL=$NEXT_PUBLIC_API_URL \
    NEXT_PUBLIC_MAPBOX_TOKEN=$NEXT_PUBLIC_MAPBOX_TOKEN \
    ENVIRONMENT=$ENVIRONMENT \
      docker compose up -d --build
#      docker compose up -d --pull always --build
elif [ $FLAG = "down" ]; then
    docker compose down
fi