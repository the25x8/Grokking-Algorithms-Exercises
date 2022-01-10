.PHONY: run-tests
run-tests:
	yarn test
	go test -v ./...
