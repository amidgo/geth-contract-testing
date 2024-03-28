compile-contract:
	docker run --rm -v ${PWD}/contract:/sources ethereum/solc:0.8.0 --abi --bin /sources/contract.sol -o /sources --overwrite 
	docker run --rm -v ${PWD}/contract:/sources ethereum/client-go:alltools-v1.10.26 abigen --abi /sources/Owner.abi --bin /sources/Owner.bin --pkg contract --out /sources/Contract.go
	sudo chmod 777 contract/Contract.go