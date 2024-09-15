# Load environment variables from the .env file
ENV_FILE := .env
include $(ENV_FILE)

export $(shell sed 's/=.*//' $(ENV_FILE))

PROJECT_ID := $(GCP_PROJECT_ID)
REGION := $(GCP_REGION)
REPOSITORY_ID := $(GCP_REPOSITORY_ID)
SERVICE := $(GCP_SERVICE)
DB_PASSWORD := $(DB_PASSWORD)
DB_USER := $(DB_USER)

IMAGE_SHA := $(shell git rev-parse HEAD)
IMAGE_NAME := $(REGION)-docker.pkg.dev/$(PROJECT_ID)/$(REPOSITORY_ID)/$(SERVICE)

.PHONY: run cloud-sql-proxy

gen-migrate:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrate:
	migrate -path db/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@localhost:5432/eisa?sslmode=disable" up

migrate-down:
	migrate -path db/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@localhost:5432/eisa?sslmode=disable" down

seed:
	go run cmd/seed/main.go

# Authenticate with Google Cloud
gcloud-auth:
	gcloud auth login
	gcloud auth configure-docker $(REGION)-docker.pkg.dev

# Build and push Docker images
docker-build:
	docker build -t $(IMAGE_NAME):$(IMAGE_SHA) -t $(IMAGE_NAME):latest --platform=linux/amd64 -f ./docker/prod/Dockerfile ./

docker-push:
	docker push $(IMAGE_NAME):$(IMAGE_SHA)
	docker push $(IMAGE_NAME):latest

# Deploy to Cloud Run
deploy-cloudrun:
	gcloud run deploy $(SERVICE) \
		--image $(IMAGE_NAME):$(IMAGE_SHA) \
		--region $(REGION) \
		--platform managed

# Full deployment process
deploy: docker-build docker-push deploy-cloudrun


# gcloud auth loginを実行ずみであること
# gcloud set project PROJECT_IDを実行ずみであること
setup-proxy:
	curl -o cloud-sql-proxy https://storage.googleapis.com/cloud-sql-connectors/cloud-sql-proxy/v2.1.2/cloud-sql-proxy.darwin.amd64
	chmod 744 cloud-sql-proxy
	gcloud components install cloud_sql_proxy --quiet
	rm cloud-sql-proxy

cloud-sql-proxy:
	cloud_sql_proxy yanbaru-eisa:asia-northeast1:yanbaru-eisa-eisa-db-prod=tcp:0.0.0.0:5432 --credential-file=key.json
