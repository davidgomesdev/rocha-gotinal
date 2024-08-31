#!/usr/bin/env bash

# Colors

START=$(tput setaf 4)
SUCCESS=$(tput setaf 2)
WARNING=$(tput setaf 3)
ERROR=$(tput setaf 1)
INFO=$(tput setaf 6)
RESET=$(tput sgr0)

# Constants

REPO_URL="https://github.com/DavidGomesDev/rocha-gotinal"
BINARY_PATH="/tmp/rocha-gotinal"

arch=$(uname -m | xargs echo -n)

if [[ "$arch" == 'aarch64' ]]; then
  arch='arm64'
fi

wget "$REPO_URL/releases/latest/download/rocha-gotinal_linux_$arch" -o /tmp/rocha-gotinal-wget.log -O "$BINARY_PATH" || { cat /tmp/rocha-gotinal-wget.log ; exit 1; }

chmod +x "$BINARY_PATH" || exit 1

"$BINARY_PATH" || (echo 'Executable failed.' && exit 1)
