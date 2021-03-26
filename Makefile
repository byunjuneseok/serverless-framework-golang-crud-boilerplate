.PHONY: clean

functions := $(shell find src/lambda -name \*main.go  | awk -F'/' '{print $$2"/"$$3"/"$$4}')

build:
	@for function in $(functions) ; do \
  		echo $$function ; \
		GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/$$function src/$$function/main.go ; \
	done
	chmod -R 755 bin/

clean:
	rm -rf ./bin

dv: clean build
	sls deploy --verbose
