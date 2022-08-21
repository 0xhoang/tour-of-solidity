mkdir -p build/caller &&
mkdir -p build/caller/deploy &&

solc --overwrite --abi contracts/caller/Caller.sol --bin contracts/caller/Caller.sol -o build/caller &&
solc --overwrite --abi contracts/caller/Callee.sol --bin contracts/caller/Callee.sol -o build/caller &&

abigen --abi=build/caller/Caller.abi --pkg=caller --type=caller --out=build/caller/Caller.go &&
abigen --abi=build/caller/Caller.abi --pkg=deploy --type=caller --out=build/caller/deploy/DeployCaller.go --bin=build/caller/Caller.bin &&

abigen --abi=build/caller/Callee.abi --pkg=caller --type=callee --out=build/caller/Callee.go &&
abigen --abi=build/caller/Callee.abi --pkg=deploy --type=callee --out=build/caller/deploy/DeployCallee.go --bin=build/caller/Callee.bin
