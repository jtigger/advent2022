#!/usr/bin/env bash

set -o errexit -o nounset -o pipefail
set -o xtrace

(
  cd day08
  go build tree-scan.go

  ./get-input.sh
  cat puzzle-input.txt | ./tree-scan
)
