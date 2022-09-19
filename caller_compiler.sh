solc --abi --bin contracts/callee.sol -o build &&
abigen --bin=./build/Callee.bin --abi=./build/Callee.abi --pkg=Callee --out=pkg/Callee.go

solc  --overwrite  --abi --bin contracts/caller.sol -o build &&
abigen --bin=./build/Caller.bin --abi=./build/Caller.abi --pkg=Caller --out=pkg/Caller.go


solc  --overwrite  --abi --bin contracts/Counter.sol -o build &&
abigen --bin=./build/Counter.bin --abi=./build/Counter.abi --pkg=Counter --out=pkg/Counter.go