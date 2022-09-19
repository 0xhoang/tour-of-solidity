#!/bin/bash

if [[ $1 == "callee" ]]; then
  solc --abi --bin contracts/callee.sol -o build &&
    abigen --bin=./build/Callee.bin --abi=./build/Callee.abi --pkg=Callee --out=pkg/Callee.go
elif [[ $1 == "caller" ]]; then
  solc --overwrite --abi --bin contracts/caller.sol -o build &&
    abigen --bin=./build/Caller.bin --abi=./build/Caller.abi --pkg=Caller --out=pkg/Caller.go
elif [[ $1 == "counter" ]]; then
  solc --overwrite --abi --bin contracts/Counter.sol -o build &&
    abigen --bin=./build/Counter.bin --abi=./build/Counter.abi --pkg=Counter --out=pkg/Counter.go
elif [[ $1 == "simplestorage" ]]; then
  solc --overwrite --abi --bin contracts/SimpleStorage.sol -o build &&
    abigen --bin=./build/SimpleStorage.bin --abi=./build/SimpleStorage.abi --pkg=SimpleStorage --out=pkg/SimpleStorage.go
else
  echo Hey that\'s a large number.
fi
