# drone-github-comment

Drone CI - Plugin to add comments to GitHub Issues/PRs

[![Build Status](https://img.shields.io/drone/build/xoxys/drone-github-comment?logo=drone)](https://cloud.drone.io/xoxys/drone-github-comment)
[![Docker Hub](https://img.shields.io/badge/dockerhub-latest-blue.svg?logo=docker&logoColor=white)](https://hub.docker.com/r/xoxys/drone-github-comment)
[![Quay.io](https://img.shields.io/badge/quay-latest-blue.svg?logo=docker&logoColor=white)](https://quay.io/repository/thegeeklab/drone-github-comment)
[![Source: GitHub](https://img.shields.io/badge/source-github-blue.svg?logo=github&logoColor=white)](https://github.com/xoxys/drone-github-comment)
[![License: MIT](https://img.shields.io/github/license/xoxys/drone-github-comment)]([LICENSE](https://github.com/xoxys/drone-github-comment/blob/master/LICENSE))

Drone plugin to add comments to GitHub Issues/PR's.

## Build

Build the binary with the following command:

```Shell
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

go build -v -a -tags netgo -o release/drone-github-comment
```

Build the Docker image with the following command:

```Shell
docker build \
  --label org.label-schema.build-date=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
  --label org.label-schema.vcs-ref=$(git rev-parse --short HEAD) \
  --file docker/Dockerfile.amd64 --tag xoxys/drone-github-comment .
```

## Usage

```Shell
docker run --rm \
  -e DRONE_BUILD_EVENT=pull \
  -e DRONE_REPO_OWNER=octocat \
  -e DRONE_REPO_NAME=foo \
  -e DRONE_PULL_REQUEST=1
  -e PLUGIN_API_KEY=abc123 \
  -e PLUGIN_MESSAGE="Demo comment" \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  xoxys/drone-github-comment
```
