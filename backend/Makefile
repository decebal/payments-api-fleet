build-deps:
	go mod download

test-local:
	ginkgo ./test/...
	go test -tags=integration

coverage:
	courtney ./src/...
