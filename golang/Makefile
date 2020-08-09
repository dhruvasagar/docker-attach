.PHONY: build run stop logs attach debug

default: build

build:
	docker build . -t docker-attach-golang

run:
	docker run -d --rm -p 1234:1234 --name docker-attach-golang -it docker-attach-golang

stop:
	docker stop docker-attach-golang

logs:
	docker logs -f docker-attach-golang

attach:
	docker attach docker-attach-golang

debug:
	curl localhost:1234
