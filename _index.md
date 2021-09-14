---
title: drone-github-comment
---

[![Build Status](https://img.shields.io/drone/build/thegeeklab/drone-github-comment?logo=drone&server=https%3A%2F%2Fdrone.thegeeklab.de)](https://drone.thegeeklab.de/thegeeklab/drone-github-comment)
[![Docker Hub](https://img.shields.io/badge/dockerhub-latest-blue.svg?logo=docker&logoColor=white)](https://hub.docker.com/r/thegeeklab/drone-github-comment)
[![Quay.io](https://img.shields.io/badge/quay-latest-blue.svg?logo=docker&logoColor=white)](https://quay.io/repository/thegeeklab/drone-github-comment)
[![GitHub contributors](https://img.shields.io/github/contributors/thegeeklab/drone-github-comment)](https://github.com/thegeeklab/drone-github-comment/graphs/contributors)
[![Source: GitHub](https://img.shields.io/badge/source-github-blue.svg?logo=github&logoColor=white)](https://github.com/thegeeklab/drone-github-comment)
[![License: MIT](https://img.shields.io/github/license/thegeeklab/drone-github-comment)](https://github.com/thegeeklab/drone-github-comment/blob/main/LICENSE)

Drone plugin to add comments to GitHub Issues and Pull Requests.

<!-- prettier-ignore-start -->
<!-- spellchecker-disable -->
{{< toc >}}
<!-- spellchecker-enable -->
<!-- prettier-ignore-end -->

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
docker build --file docker/Dockerfile.amd64 --tag thegeeklab/drone-github-comment .
```

## Usage

```Shell
docker run --rm \
  -e DRONE_BUILD_EVENT=pull_request \
  -e DRONE_REPO_OWNER=octocat \
  -e DRONE_REPO_NAME=foo \
  -e DRONE_PULL_REQUEST=1
  -e PLUGIN_API_KEY=abc123 \
  -e PLUGIN_MESSAGE="Demo comment" \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  thegeeklab/drone-github-comment
```

### Parameters

api_key
: sets api key to access github api

base_url
: sets api url; need to be changed for gh enterprise (default `https://api.github.com`)

key
: sets unique key to assign to comment

message
: sets file or string with comment message

update
: enables update of an existing comment that matches the key

skip_missing
: skips comment creation if the given message file does not exist (default `false`)
