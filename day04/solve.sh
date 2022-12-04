#!/usr/bin/env bash

set -o errexit -o nounset -o pipefail
set -o xtrace

(
  cd day04
  go build organize-cleanup.go

  ./get-input.sh
  cat puzzle-input.txt | ./organize-cleanup
)
