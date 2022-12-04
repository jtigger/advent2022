#!/usr/bin/env bash

set -o errexit -o nounset -o pipefail
set -o xtrace

(
  cd day02
  go build score.go

  ./get-input.sh
  cat puzzle-input.txt | ./score
)
