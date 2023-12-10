.PHONY: builds routing basiauth ssl clean

BASE_IMAGE ?= local-nginx

builds:
	@docker build -t $(BASE_IMAGE):base -f ./nginx/nginx.base.Dockerfile .
	@docker build -t $(BASE_IMAGE):basicauth -f ./nginx/nginx.basicauth.Dockerfile .
	@docker build -t $(BASE_IMAGE):ssl -f ./nginx/nginx.ssl.Dockerfile .

routing:
	@docker compose -f ./docker/docker-compose.routing.yml up -d

basicauth:
	@docker compose -f ./docker/docker-compose.basic-auth.yml up -d

ssl:
	@docker compose -f ./docker/docker-compose.routing-ssl.yml up -d

clean:
	@docker compose -f ./docker/docker-compose.routing.yml down
	@docker compose -f ./docker/docker-compose.basicauth.yml down
	@docker compose -f ./docker/docker-compose.ssl.yml down

