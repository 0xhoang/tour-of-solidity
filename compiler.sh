#!/bin/bash

if [[ $1 == "callee" ]]; then
  solc --abi contracts/callee.sol -o build &&
    abigen --abi=./pkg/Callee.abi --pkg=Callee --out=pkg/Callee.go
elif [[ $1 == "caller" ]]; then
  solc --abi contracts/caller.sol -o pkg &&
    abigen --abi=./pkg/Caller.abi --pkg=Caller --out=pkg/Caller.go
elif [[ $1 == "counter" ]]; then
  solc --abi contracts/Counter.sol -o pkg &&
    abigen --abi=./pkg/Counter.abi --pkg=Counter --out=pkg/Counter.go
elif [[ $1 == "simplestorage" ]]; then
  solc --abi contracts/SimpleStorage.sol -o pkg &&
    abigen --abi=./pkg/SimpleStorage.abi --pkg=SimpleStorage --out=pkg/SimpleStorage.go
else
  echo Hey that\'s is wrong params.
fi
