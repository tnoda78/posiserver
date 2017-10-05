NAME := "posiserver"
VERSION := "0.4.0"

RELEASE_DIR := releases

check-variables:
	echo "NAME: ${NAME}"
	echo "VERSION: ${VERSION}"

setup:
	go get github.com/Masterminds/glide
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports
	go get github.com/Songmu/make2help/cmd/make2help

test: deps
	go test $$(glide novendor)

deps: setup
	glide install

update: setup
	glide update

lint:
	go vet $$(glide novendor)
	for pkg in $$(glide novendor -x); do \
		golint -set_exit_status $$pkg || exit $?; \
	done

fmt:
	goimports -w $$(glide nv -x)

build:
	go build -ldflags "$(LDFLAGS)" -o ${RELEASE_DIR}/${GOOS}_${GOARCH}/${NAME}${SUFFIX} cmd/posiserver/main.go

build-all: build-darwin-amd64 build-linux-amd64 build-linux-386 build-windows-amd64 build-windows-386

build-darwin-amd64:
	@$(MAKE) build GOOS=darwin GOARCH=amd64

build-linux-amd64:
	@$(MAKE) build GOOS=linux GOARCH=amd64

build-linux-386:
	@$(MAKE) build GOOS=linux GOARCH=386

build-windows-amd64:
	@$(MAKE) build GOOS=windows GOARCH=amd64 SUFFIX=.exe

build-windows-386:
	@$(MAKE) build GOOS=windows GOARCH=386 SUFFIX=.exe
