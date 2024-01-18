.PHONY: run/docker

APP = virtualbookstore

run/docker:
	 @DOCKER_BUILDKIT=1 docker build -t $(APP) . && \
		docker run --rm --name $(APP) $(APP)
