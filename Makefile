.PHONY: build 

all: build

build:
	@docker build -t bigmoonmoon/sport:v1 -f Dockerfile .
	@docker push bigmoonmoon/sport:v1

