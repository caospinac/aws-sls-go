.PHONY: build clean deploy remove deploy-function

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/greet handler/greet/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose

remove: clean
	sls remove --verbose

deploy-function: clean build
	sls deploy function -f ${name}
