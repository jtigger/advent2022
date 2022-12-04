#!/usr/bin/env bash
set -o errexit -o nounset -o pipefail
#set -o xtrace

(
  cd day1
  ./get-input.sh
  ./most-calories.sh
)

