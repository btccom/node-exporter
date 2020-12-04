package peers

import "github.com/btcsuite/btcd/wire"

const (
	// MainNet represents the main bitcoin network.
	BitcoinMainNet wire.BitcoinNet = 0xd9b4bef9

	// TestNet represents the regression test network.
	BitcoinTestNet wire.BitcoinNet = 0xdab5bffa

	// TestNet3 represents the test network (version 3).
	BitcoinTestNet3 wire.BitcoinNet = 0x0709110b

	// SimNet represents the simulation test network.
	BitcoinSimNet wire.BitcoinNet = 0x12141c16
)

const (
	// MainNet represents the main bitcoin network.
	BitcoinCashMainNet wire.BitcoinNet = 0xe8f3e1e3

	// TestNet represents the regression test network.
	BitcoinCashTestNet wire.BitcoinNet = 0xfabfb5da

	// TestNet3 represents the test network (version 3).
	BitcoinCashTestNet3 wire.BitcoinNet = 0xf4f3e5f4

	// SimNet represents the simulation test network.
	BitcoinCashSimNet wire.BitcoinNet = 0x12141c16
)
