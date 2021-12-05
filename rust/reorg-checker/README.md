# Reorg Checker

## config

slack notification

## How to run

```bash
./reorg-checker --chain BSC --network testnet
```

## What is it

This program monitors the network whether the network has reorganization or not.
This program checks the network using two methods:

1. If the best block number is reducecd, it means there was a network oreorganization.
2. If the new block's parent hash is not the previous block's hash, it means the network has a reorganization.
