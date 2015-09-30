SOURCE=./...

all: 
	godep test

godep:
	go get github.com/tools/godep

savedeps:
	godep save $(SOURCE)

loaddeps:
	godep restore

install: build
	go install -v $(SOURCE)

build:
	go build -v $(SOURCE)

integration-loud:
	godep go test -v -tags=integration -timeout 30m $(SOURCE)

integration: 
	godep go test -tags=integration -timeout 30m $(SOURCE)

test-loud:
	godep go test -v $(SOURCE)

test:
	godep go test $(SOURCE)