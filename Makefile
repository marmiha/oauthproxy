.DEFAULT_GOAL: help
.PHONY: help docker/build docker/build/% docker/run docker/run/%

DOCKER_TAG := oauthproxy
DOCKER_VERSION := latest
PORT := 8081

help:
	@echo "------------------------------------------------------------------------"
	@echo "Container commands:"
	@echo " \033[36mdocker/build\033[0m		# Builds \033[36m${DOCKER_TAG}:${DOCKER_VERSION}\033[0m image."
	@echo " \033[36mdocker/build/%\033[0m		# Builds \033[36m${DOCKER_TAG}:%\033[0m image."
	@echo
	@echo " \033[36mdocker/run\033[0m		# Runs \033[36m${DOCKER_TAG}:${DOCKER_VERSION}\033[0m image."
	@echo " \033[36mdocker/run/%\033[0m		# Runs \033[36m${DOCKER_TAG}:%\033[0m image."
	@echo "------------------------------------------------------------------------"

docker/build:
	@make docker/build/${DOCKER_VERSION}

docker/build/%:
	@echo "\033[36m[Building]\033[0m ${DOCKER_TAG}:\033[36m${*}\033[0m"
	@docker build -t ${DOCKER_TAG}:$* -f build/Dockerfile .

docker/run:
	@make docker/run/${DOCKER_VERSION}

docker/run/%:
	@echo "\033[36m[Launching]\033[0m ${DOCKER_TAG}:\033[36m${*}\033[0m on port \033[36m${PORT}\033[0m"
	@docker run -p ${PORT}:8081 ${DOCKER_TAG}:$*