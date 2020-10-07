VERSION    = v0.2
GO_LDFLAGS = -ldflags=' -X "main.Version=$(VERSION)"'

histidine:
	go build $(GO_LDFLAGS)

.PHONY: build
build: darwin linux windows

test:
	go test
	printf "2.4\n2.45\n0.1\n1.0\n2.5\n4\n" | go run . -f s
	printf "0.1m\n1.0h\n2.5s\n4m8s\n" | go run . -f d

.PHONY: darwin linux windows
darwin linux windows:
	mkdir -p build/$@
	GOOS=$@ go build $(GO_LDFLAGS) -o build/$@/histidine
	tar -c -f histidine-$@.tgz -z -C build/$@ histidine


clean:
	-rm -rf build histidine-*.go histidine
