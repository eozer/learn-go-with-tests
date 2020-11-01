# See the base Dockerfile from:
# https://github.com/microsoft/vscode-dev-containers/blob/v0.137.0/containers/go/.devcontainer/base.Dockerfile
ARG GOLANGVERSION="1"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${GOLANGVERSION}

# [Optional] Uncomment this section to install additional OS packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends vim

##
# https://taskfile.dev/#/installation
RUN curl -sL https://taskfile.dev/install.sh | sh

# [Optional] Uncomment the next line to use go get to install anything else you need
# RUN go get -x <your-dependency-or-tool>
