#!/bin/bash

AWS_REGION=$1
ECR_REGISTRY=$2
CONTAINER_NAME=$3
ECR_REPOSITORY=$4
IMAGE_TAG=$5

aws ecr get-login-password --region "$AWS_REGION" | \
    docker login --username AWS --password-stdin $ECR_REGISTRY

if docker ps -a --format '{{.Names}}' | grep -q "$CONTAINER_NAME"; then
  docker stop $CONTAINER_NAME
  docker rm $CONTAINER_NAME
fi

docker pull $ECR_REGISTRY/$ECR_REPOSITORY:$IMAGE_TAG
docker run -d -p 80:3000 --name $CONTAINER_NAME \
    $ECR_REGISTRY/$ECR_REPOSITORY:$IMAGE_TAG
