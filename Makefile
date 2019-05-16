
GO_COMPILER_OPTS = -a -tags netgo -ldflags '-w -extldflags "-static"'

build:
	go build

install:
	install -m755 keyto /usr/bin/keyto

all: fmt win lin linarm

fmt:
	gofmt -w *.go */*.go

dep:
	go get -u github.com/eyedeekay/httptunnel/keyto

win: win32 win64

win64:
	GOOS=windows GOARCH=amd64 go build \
		$(GO_COMPILER_OPTS) \
		-buildmode=exe \
		-o ./keyto.exe
	@echo "built"

win32:
	GOOS=windows GOARCH=386 go build \
		$(GO_COMPILER_OPTS) \
		-buildmode=exe \
		-o ./keyto.exe
	@echo "built"

lin: lin64 lin32

lin64:
	GOOS=linux GOARCH=amd64 go build \
		$(GO_COMPILER_OPTS) \
		-o ./keyto-64
	@echo "built"

lin32:
	GOOS=linux GOARCH=386 go build \
		$(GO_COMPILER_OPTS) \
		-o ./keyto-32
	@echo "built"

linarm: linarm32 linarm64

linarm64:
	GOOS=linux GOARCH=arm64 go build \
		$(GO_COMPILER_OPTS) \
		-o ./keyto-arm64
	@echo "built"

linarm32:
	GOOS=linux GOARCH=arm go build \
		$(GO_COMPILER_OPTS) \
		-o ./keyto-arm32
	@echo "built"

vet:
	go vet ./*.go
	go vet ./convert/*.go

clean:
	rm -f keyto-* *.exe *.log
