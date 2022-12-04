#!/usr/bin/env bash
set -o errexit -o nounset -o pipefail
#set -o xtrace

if [[ ! -v AOC_SESSION_TOKEN ]]; then
  echo "Need: environment variable AOC_SESSION_TOKEN to be set with Advent of Code website token."
  echo "  (this token is used to fetch the puzzle input)"
  echo "  (In Chrome, token can be found in the Developer Tools; Application > Cookies > https://adventofcode.com > session)"
  echo "  e.g."
  echo '    export AOC_SESSION_TOKEN="...128-character-session-token..."'
  exit 1
fi

input=$(
  curl 'https://adventofcode.com/2022/day/4/input' \
    -H "cookie: session=${AOC_SESSION_TOKEN}" \
  2>curl-stderr.txt
)

echo "${input}" >puzzle-input.txt
