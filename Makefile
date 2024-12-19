GOOS=linux

.PHONY: build
build:
	go build -o ./bin/app ./daemon/.

.PHONY: docker-build
docker-build:
	docker build -t alexbenncr.azurecr.io/aks-imex/aks-imex-daemon:latest --file ./daemon/Dockerfile .

.PHONY: docker-push
docker-push:
	docker push alexbenncr.azurecr.io/aks-imex/aks-imex-daemon
