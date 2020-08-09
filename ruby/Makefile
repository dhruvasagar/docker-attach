.PHONY: build run stop logs attach debug

default: build

build:
	docker build . -t docker-attach-ruby

run:
	docker run -d --rm -p 1234:1234 --name docker-attach-ruby -it docker-attach-ruby

stop:
	docker stop docker-attach-ruby

logs:
	docker logs -f docker-attach-ruby

attach:
	docker attach docker-attach-ruby

debug:
	curl localhost:1234
