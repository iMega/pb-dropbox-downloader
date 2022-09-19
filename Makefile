REPO = github.com/imega/pb-dropbox-downloader
CWD = /go/src/$(REPO)
GO_IMG = golang:alpine
SDK_IMG = 5keeve/pocketbook-go-sdk:6.3.0-b288-v1

test: lint unit

lint:
	@docker run --rm -t -v $(CURDIR):$(CWD) -w $(CWD) \
		golangci/golangci-lint golangci-lint run

unit:
	@docker run --rm -v $(CURDIR):$(CWD) \
		-w $(CWD) \
		-e GOFLAGS=-mod=mod \
		-e CGO_ENABLED=0 \
		$(GO_IMG) \
		sh -c 'go test -v `go list ./... | grep -v tests`'

build:
	@docker run --rm -v $(CURDIR):/app $(SDK_IMG) \
		build -v -tags=UI -ldflags="-s -w" -o d7xsync.app .

genlogo:
	@docker run --rm -v $(CURDIR):/data -w /data --entrypoint=sh $(SDK_IMG) \
		-c "go run logogen/main.go && pbres -c pbsdk/logo.c logo.bmp && rm logo.bmp"
