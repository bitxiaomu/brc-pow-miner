# BRC PoW Miner

**This is an interest based beta version project, feel free to clone and use.**

## Why Brc-Pow?

The inscription of the BRC20 protocol causes a lot of network congestion during fair launch, which has a negative impact on the Bitcoin ecosystem. After all, the Bitcoin network is a shared ecosystem, so on this basis, adding pre PoW brings more fairness and greenery.

## How to use

prepare

1. git tool
2. install go dev env
3. sync full bitcoin node / online version inscription tool
4. a wallet address with satoshis

download code

```bash
# download code
git clone https://github.com/bitxiaomu/brc-pow-miner.git
cd brc-pow-miner
```

configure etc/config.yaml

```bash
# in folder brc-pow-miner

cd etc
cp config.example.yaml config.yaml
```

**TODO**: explain how to configure the file

build

```bash
# in folder brc-pow-miner

go --version
go build .
```

run

```bash
# in folder brc-pow-miner

chmod +x ./brc-pow-miner
./brc-pow-miner
```
