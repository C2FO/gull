SOURCE := ./source/...

build: installdeps	
	go install -ldflags "-X github.com/c2fo/gull/source/lib/common.ApplicationVersion=`cat ./VERSION.txt`" -v $(SOURCE)

installdeps: 
	glide install

integration-loud: build
	go test -v -tags=integration -timeout 30m $(SOURCE)

integration: build
	go test -tags=integration -timeout 30m $(SOURCE)

test-loud: build
	go test -v $(SOURCE)

test: build
	go test $(SOURCE)