package main

//go:generate solc --optimize --abi ./src/hodl.sol -o build --overwrite
//go:generate solc --optimize --bin ./src/hodl.sol -o build --overwrite
//go:generate abigen --abi=./build/Hodl.abi --bin=./build/Hodl.bin --pkg=api --out=./api/hodl.go
