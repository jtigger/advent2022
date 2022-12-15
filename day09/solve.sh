#!/usr/bin/env bash

set -o errexit -o nounset -o pipefail
set -o xtrace

(
  cd day09
  go build trackrope.go

  ./get-input.sh
  cat puzzle-input.txt | ./trackrope
)
