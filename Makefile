compile: 
	protoc api/v1/*.proto \
		--go_out=. \
		--go-grpc_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--proto_path=.
test:
	go test -race -v ./...

clean:
	rm -rf dist
	go clean -modcache && go mod tidy

GITBRANCH=$(shell git branch --show-current)
release:
	@if [ $(GITBRANCH) == "main" ]; then \
		goreleaser check; \
		goreleaser release --rm-dist; \
	else \
		echo "you can only release on the main branch"; \
		exit 1; \
	fi 
mock-release:
	goreleaser check
	goreleaser release --snapshot --rm-dist
