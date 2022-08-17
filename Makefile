.PHONY: list
list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'

build: clean deps
	env go build -o bin/gotask2 main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deps:
	go mod tidy

lint: deps
	golangci-lint run -c ./.golangci.yml

test: deps
	go test ./pkg/...

run: build
	./bin/gotask2

