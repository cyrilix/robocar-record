.PHONY: test docker

DOCKER_IMG = cyrilix/robocar-record

test:
	go test -mod vendor ./cmd/rc-record ./part

docker:
	export DOCKEC_CLI_EXPERIMENTAL=enabled
	docker buildx build . --platform linux/arm/7,linux/arm64,linux/amd64 -t ${DOCKER_IMG} --push

