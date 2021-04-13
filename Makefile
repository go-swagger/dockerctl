# download and replace the latest swagger spec for docker engine
# not used yet. The latest swagger causes some problems in cli generate
sync-swagger:
	curl https://raw.githubusercontent.com/docker/engine/master/api/swagger.yaml -o api/swagger.yaml

build-container:
	docker build -t dockerctl:dev -f dev/Dockerfile .

generate:
	swagger generate cli --target=. --spec=api/swagger.yaml --cli-app-name dockerctl

generate-completion:
	cmd/dockerctl/dockerctl completion bash > cmd/completion/dockerctl.bash-completion.sh

clean-generate:
	rm -rf cli client models cmd/dockerctl

build:
	CGO_ENABLED=0 go build -o cmd/dockerctl/dockerctl cmd/dockerctl/main.go

clean:
	rm -f cmd/dockerctl/dockerctl

dind-start:
	docker run --privileged --name dind -d \
		--network docker --network-alias docker \
		-v ${PWD}:/repo -w /repo \
		dockerctl:dev

dind-stop:
	docker container stop dind
	docker container rm dind

shell:
	docker exec -it dind bash

expose-docker:
	socat TCP-LISTEN:12345,bind=127.0.0.1,reuseaddr,fork,range=127.0.0.0/8 UNIX-CLIENT:/var/run/docker.sock &
