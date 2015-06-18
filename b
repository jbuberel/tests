#!/bin/bash

for i in {0..9}
do
    go test -bench=. | tail >>  $1
done
