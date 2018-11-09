help:
	@echo "Makefile help"

magic-cube-go:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

docker-image: magic-cube-go
	docker build -t mingz2013/magic-cube-go .


commit-docker:docker-image
	docker login
	docker push mingz2013/magic-cube-go

run:
	docker run -d --link redis-mq:redis-mq --name magic-cube-go -it mingz2013/magic-cube-go:latest


.PYONY: help, commit-docker, docker-image, magic-cube-go, run

