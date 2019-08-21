start: build
	@./app

build:
	@go build -a -tags netgo -ldflags '-w' -o app .

docker: dockerBuild
	@docker run --rm --name eiei -d eiei

dockerBuild:
	@docker build -t eiei .

.PHONY: build start docker dockerBuild
