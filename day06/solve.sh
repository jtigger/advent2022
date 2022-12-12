#!/usr/bin/env bash

set -o errexit -o nounset -o pipefail
set -o xtrace

(
  cd day06
  go build cmd.go

  ./get-input.sh
  cat puzzle-input.txt | ./cmd
)
