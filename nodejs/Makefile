.PHONY: build run stop logs attach debug

default: build

build:
	docker build . -t docker-attach-nodejs

run:
	docker run -d --rm -p 1234:1234 --name docker-attach-nodejs -it docker-attach-nodejs

stop:
	docker stop docker-attach-nodejs

logs:
	docker logs -f docker-attach-nodejs

attach:
	docker attach docker-attach-nodejs

debug:
	curl localhost:1234
