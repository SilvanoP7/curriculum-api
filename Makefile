.PHONY: image image-build image-push

# Versioning information.
IMAGE_TAG=1.0.0

# Container registry information.
CR_NAME = cr

# Use Bash shell for better syntax and additional features.
SHELL := /bin/bash

CONTAINER_REGISTRY = ${CR_NAME}.sac23.io
DB_IMAGE_NAME = ${CONTAINER_REGISTRY}/curriculum-api/curriculumdb
API_IMAGE_NAME=${CONTAINER_REGISTRY}/curriculum-api/api

db-image-build:
	cd postgres-docker && docker build -t ${DB_IMAGE_NAME}\:${IMAGE_TAG} .
	docker push ${DB_IMAGE_NAME}\:${IMAGE_TAG}

api-image-build:
	docker build -t ${API_IMAGE_NAME}\:${IMAGE_TAG} .
	docker push ${API_IMAGE_NAME}\:${IMAGE_TAG}

