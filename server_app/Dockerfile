# Copyright 2025 Dasmlab
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.



# STAGE 1 - Build the server_manager binary
FROM golang:latest AS builder
WORKDIR /workspace

# Build Arch ARG
ARG ARCH

# Allow overload of goproxy (--build_arg $(go env GOPROXY))
ARG goproxy=https://proxy.golang.org
ENV GOPROXY=$goproxy

# Install SWAG CMD
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy the Go Mod files
COPY main-app/go.mod go.mod
COPY main-app/go.sum go.sum

# Cache our dependancies
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download


# Copy sources (main-app and sub dirs)
COPY main-app/. ./

# Build the swagger files
RUN swag init --generalInfo main.go --output docs

# Build ARGS
ARG package=.
ARG ldflags

# Dont Force Rebuild of up-to-date pkgs
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} \
    go build -ldflags "${ldflags} -extldflags '-static'" \
    -o mcp-go-server-manager ${package}

## STAGE 2 - Create Production Image
#FROM gcr.io/distroless/static:nonroot
FROM ubuntu:latest
WORKDIR /
COPY --from=builder /workspace/mcp-go-server-manager .
# Use uid of nonroot user (65532) because kubernetes expects numeric user when applying pod security policies
USER 65532
ENTRYPOINT ["/mcp-go-server-manager"]
