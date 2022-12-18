#!/usr/bin/env bash

set -o errexit -o nounset -o pipefail
set -o xtrace

(
  cd day10
  go build cpu.go

  ./get-input.sh
  cat puzzle-input.txt | ./cpu
)
