BINARY_NAME=robotcli
gover=1.18.1

build: # create an executable under bin folder.
	CGO_ENABLED=0 go build -o ./bin/$(BINARY_NAME) ./cmd

dockerbuild: # use this to build the binaries for linux(amd64), windows(amd64), macos(amd64) and macos(arm64)
	DOCKER_BUILDKIT=1 docker build --build-arg GO_VER=${gover} --build-arg BIN_NAME=${BINARY_NAME} --output type=tar,dest=./bin/build.tar .

test: # run unit tests.
	go test ./... -v

cover: # run test and show coverage information
	go test -coverprofile=coverage.out ./cmd
	go tool cover -html=coverage.out

run: # run the robotcli
	CGO_ENABLED=0 go run ./cmd

clean: # clean build artifacts
	go clean
	rm ./bin/*