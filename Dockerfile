# Copyright © Observerly Ltd.

# //////////////////////////////////////////////////////////////////////////////////////////////////////////////////// #

FROM golang:1.23-bookworm AS development

# //////////////////////////////////////////////////////////////////////////////////////////////////////////////////// #

# Install necessary packages: gcc, make, libc-dev, bash, curl, and openssh-client
RUN apt-get update && apt-get install -y --no-install-recommends \
    bash \
    curl \
    gcc \
    git \
    libc-dev \
    make \
    openssh-client \
    unzip \
    wget \
    zsh \
    && rm -rf /var/lib/apt/lists/*

# Set the working directory to /usr/src/app
WORKDIR /usr/src/app

# Ensure staticcheck is executable and in the PATH
RUN go install honnef.co/go/tools/cmd/staticcheck@latest

# Ensure go-critic is executable and in the PATH
RUN go install github.com/go-critic/go-critic/cmd/gocritic@latest

# Add Go binaries to PATH
ENV PATH="$PATH:$(go env GOPATH)/bin"

# Install Oh My Zsh non-interactively
RUN sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" "" --unattended

# Install zsh-in-docker non-interactively
RUN sh -c "$(wget -O- https://github.com/deluan/zsh-in-docker/releases/download/v1.1.5/zsh-in-docker.sh)" -- \
    -t https://github.com/denysdovhan/spaceship-prompt \
    -a 'SPACESHIP_PROMPT_ADD_NEWLINE="false"' \
    -a 'SPACESHIP_PROMPT_SEPARATE_LINE="false"' \
    -p git \
    -p ssh-agent \
    -p https://github.com/zsh-users/zsh-autosuggestions \
    -p https://github.com/zsh-users/zsh-completions

# Copy application code
COPY . /usr/src/app/

# //////////////////////////////////////////////////////////////////////////////////////////////////////////////////// #

# Set the default shell to zsh
SHELL ["/bin/zsh", "-c"]

# //////////////////////////////////////////////////////////////////////////////////////////////////////////////////// #