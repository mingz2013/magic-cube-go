help:
	@echo "Makefile help"

clean:
	rm magic-cube-go


magic-cube-go: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

docker-image: magic-cube-go
	docker build -t mingz2013/magic-cube-go .


commit-docker: docker-image
	docker login
	docker push mingz2013/magic-cube-go


remove-container:
	docker rm magic-cube-go


run: remove-container
	docker run -d --link redis-mq:redis-mq --name magic-cube-go -it mingz2013/magic-cube-go:latest

logs:
	docker logs magic-cube-go

.PYONY: help, commit-docker, docker-image, magic-cube-go, run, remove-container, logs

