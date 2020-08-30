#!/bin/bash

function checkStdOut() {
  diff -q "$1" "$2" | grep -o differ
  if [ $? -eq 0 ]; then
    return 0
  else
    return 1
  fi
}

function checkStdErr(){
  if [ -e "$1" ]; then
    diff -q "$1" "$2" | grep -o differ
    if [ $? -eq 0 ]; then
      return 0
    else
      return 1
    fi
  fi
  return 0
}

for i in $(find . -type f -name "*.yml"); do
  outfile=$(basename $i .yml)
  ../bin/yacasc verify $i > /tmp/out 2> /tmp/err

  (checkStdOut "${outfile}.out" "/tmp/out")
  if [ $? -eq 0 ]; then
    b=$(checkStdErr "${outfile}.err" "/tmp/err")
    if [ $? -eq 0 ]; then
      echo "$i  OK"
      continue
    fi
  fi
  echo "$i  FAIL"
  exit 1
done

