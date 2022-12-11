#!/usr/bin/env bash

set -o errexit -o nounset -o pipefail
set -o xtrace

(
  cd day07
  go build filesystem.go

  ./get-input.sh
  cat puzzle-input.txt | ./filesystem
)
