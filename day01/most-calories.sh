#!/usr/bin/env bash

set -o errexit -o nounset -o pipefail
# set -o xtrace

function sum_groups() {
  IFS='' 
  while read item ; do
    if [[ "${item}" == "" ]]; then
      echo ${elf_total}
      elf_total=0
    else 
      elf_total=$((${elf_total} + ${item}))
    fi
  done
  if [[ ${elf_total} -ne 0 ]]; then
     echo ${elf_total}
  fi
}

input=$(cat puzzle-input.txt)

# sum up how many calories each elf has
elf_total=0
elves=$(echo "${input}" | sum_groups)

# which elf has the most calories?
echo -n "The most: "
echo "${elves[*]}" | sort --numeric --reverse | head -n 1

echo -n "Sum of the top three: "
echo "${elves[*]}" | sort --numeric --reverse | head -n 3 | sum_groups

