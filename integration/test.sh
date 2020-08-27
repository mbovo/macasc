#!/bin/bash

for i in $(find . -type f -name "*.yml"); do
  outfile=$(basename $i .yml)
  ../bin/yacasc verify $i > /tmp/out
  diff ${outfile}.out /tmp/out
  if [ $? -eq 0 ]; then
    echo "$i  OK"
  else
    echo "$i  FAIL"
    exit 1
  fi
done

