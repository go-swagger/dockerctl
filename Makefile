# download and replace the latest swagger spec for docker engine
# not used yet. The latest swagger causes some problems in cli generate
sync-swagger:
	curl https://raw.githubusercontent.com/docker/engine/master/api/swagger.yaml -o api/swagger.yaml

build-container:
	docker build -t go-swagger:dev -f dev/Dockerfile .

generate:
	swagger generate cli --target=. --spec=api/swagger.yaml

clean-generate:
	rm -rf cli client models cmd

build:
	go build -o cmd/cli/dockerctl cmd/cli/main.go

clean:
	rm -f cmd/cli/dockerctl