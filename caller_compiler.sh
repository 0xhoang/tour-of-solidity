mkdir -p build/caller && \
mkdir -p build/caller/deploy && \
solc --overwrite  --abi contracts/caller/Caller.sol -o build/caller && \
solc --overwrite  --bin contracts/caller/Caller.sol -o build/caller && \
abigen --abi=build/caller/Caller.abi --pkg=caller --out=build/caller/Caller.go && \
abigen --abi=build/caller/Caller.abi --pkg=caller --out=build/caller/deploy/DeployCaller.go --bin=build/caller/Caller.bin