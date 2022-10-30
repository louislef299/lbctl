RELEASE_TYPE ?= patch 
# patch, minor, major
SEMVER_IMAGE ?= alpine/semver
# https://hub.docker.com/r/alpine/semver

CURRENT_VERSION := $(shell git describe --tags `git rev-list --tags --max-count=1`)
PRETTY_VERSION := $(shell go run main.go --version)
GO_VERSION := $(lastword $(subst :, ,$(PRETTY_VERSION)))

ifndef CURRENT_VERSION
  CURRENT_VERSION := v0.0.0
endif

NEXT_VERSION := v$(shell docker run --rm $(SEMVER_IMAGE) semver -c -i $(RELEASE_TYPE) $(CURRENT_VERSION))

BRANCH = $(shell git rev-parse --abbrev-ref HEAD)

current-version:
	@echo $(CURRENT_VERSION)

next-version:
	@echo $(NEXT_VERSION)

check-version:
	if [ "$(GO_VERSION)" != "$(NEXT_VERSION)" ]; then \
		echo "Versions don't match up(GO_VERSION=$(GO_VERSION) and NEXT_VERSION=$(NEXT_VERSION)), update the clctl in version.go to match the next version" ;\
		exit 1 ;\
	fi

test:
	go test -race -v ./...

clctl: main.go go.mod go.sum
	go build

release: check-version ~/.config/goreleaser/gitea_token
	if [ "$(BRANCH)" = "main" ];then \
	  git tag $(NEXT_VERSION) ;\
	  goreleaser check ;\
	  goreleaser release --rm-dist ;\
	  git push --tags ;\
	fi

test-release: check-version
	goreleaser check
	goreleaser release --snapshot --rm-dist

clean:
	rm -rf clctl dist

clean-go:
	go clean -modcache
	go get -u && go mod tidy
