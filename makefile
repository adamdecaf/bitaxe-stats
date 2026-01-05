PWD := $(shell pwd)

.PHONY: check
check:
ifeq ($(OS),Windows_NT)
	go test ./...
else
	@wget -O lint-project.sh https://raw.githubusercontent.com/moov-io/infra/master/go/lint-project.sh
	@chmod +x ./lint-project.sh
	COVER_THRESHOLD=0.0 ./lint-project.sh
endif

.PHONY: generate generate-client generate-cleanup
generate: generate-client generate-cleanup

generate-client:
	docker run --rm -v $(CURDIR):/local openapitools/openapi-generator-cli generate \
	  -i https://raw.githubusercontent.com/bitaxeorg/ESP-Miner/refs/heads/master/main/http_server/openapi.yaml \
	  -g go \
	  -o /local/internal/bitaxeclient

generate-cleanup:
	find internal/bitaxeclient/ -mindepth 1 -type d | xargs rm -rf
	find internal/bitaxeclient/ -type f -not -name "*.go" | xargs rm -rf

.PHONY: clean
clean:
	@rm -rf ./bin/ ./tmp/ coverage.txt misspell* staticcheck lint-project.sh

.PHONY: cover-test cover-web
cover-test:
	go test -coverprofile=cover.out ./...
cover-web:
	go tool cover -html=cover.out
