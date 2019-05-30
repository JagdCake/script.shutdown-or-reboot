.PHONY: all
.PHONY: test
.PHONY: build

.PHONY: vet
.PHONY: clean

all: test build

test:
	go test ./**/

build:
	go build

vet:
	go vet

clean:
	go clean
