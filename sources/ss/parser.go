package ss

import (
	"github.com/btccom/node-exporter/sources/ss/btc"
	"github.com/btccom/node-exporter/sources/ss/ckb"
	"github.com/btccom/node-exporter/sources/ss/dcr"
)

func ParseHeight(coin string, resp interface{}) (height int64, err error) {
	nResp := resp.(NotifyRes)
	if coin == "dcr" {
		return dcr.ParseHeight(nResp.CoinbaseTX1)
	}
	if coin == "ckb" {
		return ckb.ParseHeight(nResp.Height)
	}
	if coin == "eth" || coin == "etc" {
		return ckb.ParseHeight(nResp.Height)
	}
	if coin == "beam" {
		return ckb.ParseHeight(nResp.Height)
	}
	if coin == "grin" {
		return ckb.ParseHeight(nResp.Height)
	}
	return btc.ParseHeight(nResp.CoinbaseTX1)
}

func ParsePrevHash(coin string, resp interface{}) (hash string) {
	nResp := resp.(NotifyRes)
	if coin == "ckb" {
		return ckb.ParsePrevHash(nResp.ParentHash)
	}
	return btc.ParsePrevHash(nResp.Hash)
}
