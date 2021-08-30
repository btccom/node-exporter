package peers

import (
	"fmt"
	"github.com/btcsuite/btcd/addrmgr"
	"github.com/btcsuite/btcd/wire"
	"github.com/sirupsen/logrus"
	"net"
	"time"
)

type Result struct {
	Status          string
	Address         string
	Peers           string
	UserAgent       string
	ErrMsg          string
	CoinType        string
	Timestamp       string
	HeightChangedAt string
	Height          int64
}

type NodePeer struct {
	conn           net.Conn
	Address        string
	ConnectTimeout time.Duration
	Coin           string
	Peers          []string
}

func NewNodePeer(address string, coin string) *NodePeer {
	b := NodePeer{Address: address, Coin: coin, ConnectTimeout: 100 * time.Second}
	return &b
}

// Connect
func (b *NodePeer) Connect() error {
	if b.conn != nil {
		return fmt.Errorf("peer already connected, can't connect again")
	}
	conn, err := net.DialTimeout("tcp", b.Address, b.ConnectTimeout)
	if err != nil {
		return err
	}
	b.conn = conn
	return nil
}

// DisConnect
func (b *NodePeer) DisConnect() {
	if b.conn == nil {
		return
	}
	if err := b.conn.Close(); err != nil {
		logrus.Error(err)
	}
}

func (b *NodePeer) Handshake() (result Result, err error) {
	if b.conn == nil {
		err = fmt.Errorf("peer is not connected, can't handshake")
		return
	}
	nonce, err := GetRandomUint64()
	if err != nil {
		return
	}
	msgVersion := wire.NewMsgVersion(wire.NewNetAddress(b.conn.LocalAddr().(*net.TCPAddr), 0),
		wire.NewNetAddress(b.conn.RemoteAddr().(*net.TCPAddr), 0),
		nonce, 0)
	msgVersion.UserAgent = "ua"
	msgVersion.DisableRelayTx = true
	if err = b.WriteMessage(msgVersion); err != nil {
		return
	}
	// read the version response.
	msg, _, err := b.ReadMessage()
	if err != nil {
		return
	}
	vmsg, ok := msg.(*wire.MsgVersion)
	if !ok {
		err = fmt.Errorf("did not receive version message: %T", vmsg)
		return
	}
	// send ver ack.
	if err = b.WriteMessage(wire.NewMsgVerAck()); err != nil {
		return
	}
	result.Height = int64(vmsg.LastBlock)
	result.Address = b.Address
	result.UserAgent = vmsg.UserAgent
	result.Timestamp = vmsg.Timestamp.UTC().Format("2006-01-02 15:04:05")
	return
}

// SendGetAddr send getaddr command
func (b *NodePeer) SendGetAddr() error {
	err := b.WriteMessage(wire.NewMsgGetAddr())
	return err
}

func (b *NodePeer) GetAddrResponse() ([]string, error) {
	var firstReceived = -1
	var tolerateMessages = 3
	var otherMessages []string
	var addresses []string
	for {
		if len(otherMessages) > tolerateMessages {
			return addresses, nil
		}
		// read message in loop
		msg, _, err := b.ReadMessage()
		if err != nil {
			otherMessages = append(otherMessages, err.Error())
			logrus.Warningf("[%s] Failed to read message: %v", b.Address, err)
			continue
		}

		switch tmsg := msg.(type) {
		case *wire.MsgAddr:
			for _, addrList := range tmsg.AddrList {
				// node age
				addresses = append(addresses, addrmgr.NetAddressKey(addrList))
			}
			if firstReceived == -1 {
				firstReceived = len(tmsg.AddrList)
			} else if firstReceived > len(tmsg.AddrList) || firstReceived == 0 {
				// probably done.
				return addresses, nil
			}
		default:
			otherMessages = append(otherMessages, tmsg.Command())
		}
	}
}

func (b *NodePeer) WriteMessage(msg wire.Message) error {
	return wire.WriteMessage(b.conn, msg, wire.ProtocolVersion, b.GetBitcoinNet())
}

func (b *NodePeer) ReadMessage() (wire.Message, []byte, error) {
	return wire.ReadMessage(b.conn, wire.ProtocolVersion, b.GetBitcoinNet())
}

func (b *NodePeer) GetBitcoinNet() wire.BitcoinNet {
	if b.Coin == "bch" {
		return BitcoinCashMainNet
	}
	return BitcoinMainNet
}
