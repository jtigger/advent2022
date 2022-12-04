#!/usr/bin/env bash

set -o errexit -o nounset -o pipefail
#set -o xtrace

input=$(cat puzzle-input.txt)

# sum up how many calories each elf has
elf_total=0
elves=$(
IFS='' echo "${input}" | while read item ; do
  if [[ "${item}" == "" ]]; then
    echo ${elf_total}
    elf_total=0
  else 
    elf_total=$((${elf_total} + ${item}))
  fi
done
)

# which elf has the most calories?
echo "${elves[*]}" | sort --numeric --reverse | head -n 1

