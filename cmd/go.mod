module github.com/breez/breez/cmd

go 1.13

require (
	github.com/breez/breez v0.0.0-20191222101539-d9f37e24beee
	github.com/deadsy/go-cli v0.0.0-20191117003156-1fbe7fd20d78
	github.com/golang/protobuf v1.3.1
)

replace (
	github.com/breez/breez => ../
	github.com/btcsuite/btcd v0.20.0-beta => github.com/btcsuite/btcd v0.20.1-beta
	github.com/btcsuite/btcwallet v0.10.0 => github.com/breez/btcwallet v0.10.1-0.20191121081139-3f579e0a038c
	github.com/btcsuite/btcwallet/walletdb v1.1.0 => github.com/breez/btcwallet/walletdb v1.1.1-0.20191121081139-3f579e0a038c
	github.com/btcsuite/btcwallet/wtxmgr v1.0.0 => github.com/breez/btcwallet/wtxmgr v1.0.1-0.20191121081139-3f579e0a038c

	github.com/lightninglabs/neutrino => github.com/breez/neutrino v0.10.1-0.20191121084819-28462a8edb3a
	github.com/lightningnetwork/lnd => github.com/breez/lnd v0.8.0-beta.0.20191212101524-70e19d1e5d35
)