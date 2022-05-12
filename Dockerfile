ARG GO_VER=1.18
FROM golang:$GO_VER AS build-stage
RUN mkdir /app
COPY . /app
WORKDIR /app
ARG BIN_NAME=robotcli
RUN for os in "darwin" "linux" "windows"; do for arch in "arm64" "amd64"; do GOOS=$os GOARCH=$arch CGO_ENABLED=0 go build -o ./bin/${BIN_NAME}_${os}_${arch} ./cmd ; done ; done

FROM scratch AS export-stage
COPY --from=build-stage /app/bin /