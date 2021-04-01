# download and replace the latest swagger spec for docker engine
sync-swagger:
	curl https://raw.githubusercontent.com/docker/engine/master/api/swagger.yaml -o api/swagger.yaml