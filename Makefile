GOARCH?=amd64
GOOS?=linux
APP?=apigateway
PROJECT?=github.com/seagullbird/headr-apigateway
COMMIT?=$(shell git rev-parse --short HEAD)
PORT?=:8689


clean:
	rm -f ${APP}

build: clean
	GOARCH=${GOARCH} GOOS=${GOOS} go build \
	-ldflags "-s -w -X ${PROJECT}/config.PORT=${PORT}" \
	-o ${APP}

container: build
	docker build -t apigateway:${COMMIT} .

minikube: container
	cat k8s/deployment.yaml | gsed -E "s/\{\{(\s*)\.Commit(\s*)\}\}/$(COMMIT)/g" > tmp.yaml
	kubectl apply -f tmp.yaml
	rm -f tmp.yaml
